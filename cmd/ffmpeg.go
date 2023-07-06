package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// ffmpeg -nostats -progress pipe:1
func ExecShell(command string) error {
	commandParts := strings.Split(command, " ")
	parts := []string{
		commandParts[0],
		"-y",
		"-stats",
		"-hide_banner",
		"-loglevel error",
	}

	parts = append(parts, commandParts[1:]...)

	// Rebuild command
	command = strings.Join(parts, " ")

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}

func ResizeVideo(input, output string, width, height int) error {
	// Construct the FFmpeg command
	command := fmt.Sprintf("ffmpeg -i %s -vf scale=%d:%d -c:v libx264 -preset medium -crf 23 -c:a copy %s", input, width, height, output)

	// Execute the FFmpeg command
	return ExecShell(command)
}

// Extract audio from a video with specified audio codec. default is to copy source codec
// startTimestamp and/or endTimestamp are optional.
// If codec is empty string will use source codec.
func ExtractAudioFromVideoWithTimestamps(videoFilename, outputFilename string, codec string,
	startTimestamp, endTimestamp string) error {
	if codec == "" {
		codec = "copy"
	}

	var command string
	if startTimestamp != "" && endTimestamp != "" {
		command = fmt.Sprintf("ffmpeg -i \"%s\" -ss %s -to %s -vn -acodec %s \"%s\"", videoFilename, startTimestamp, endTimestamp, codec, outputFilename)
	} else if startTimestamp != "" {
		command = fmt.Sprintf("ffmpeg -i \"%s\" -ss %s -vn -acodec %s \"%s\"", videoFilename, startTimestamp, codec, outputFilename)
	} else if endTimestamp != "" {
		command = fmt.Sprintf("ffmpeg -i \"%s\" -to %s -vn -acodec %s \"%s\"", videoFilename, endTimestamp, codec, outputFilename)
	} else {
		command = fmt.Sprintf("ffmpeg -i \"%s\" -vn -acodec %s \"%s\"", videoFilename, codec, outputFilename)
	}
	return ExecShell(command)
}

// Extract a video frame/image to jpg per timestamp.
func ExtractFrameAtTimestamps(inputFilename, outputDirectory string, timestamps []string) error {
	for i, timestamp := range timestamps {
		outputFilename := fmt.Sprintf("%s/frame%03d.jpg", outputDirectory, i+1)
		command := fmt.Sprintf("ffmpeg -ss %s -i %s -vframes 1 %s", timestamp, inputFilename, outputFilename)
		err := ExecShell(command)
		if err != nil {
			return fmt.Errorf("error extracting frame at timestamp %s: %s", timestamp, err)
		}
	}
	return nil
}

func RotateVideo(inputFilename, outputFilename string, angle int) error {
	transposeValue := 0

	switch angle {
	case 90:
		transposeValue = 1
	case 180:
		transposeValue = 2
	case 270:
		transposeValue = 3
	default:
		return fmt.Errorf("unsupported rotation angle")
	}

	transposeFilter := fmt.Sprintf("transpose=%d", transposeValue)

	command := fmt.Sprintf("ffmpeg -i %s -vf \"%s\" %s", inputFilename, transposeFilter, outputFilename)
	return ExecShell(command)
}

// Generate a new audio file containing only the frequencies within the specified range.
func ExtractAudioFrequencies(inputFilename, outputFilename string, startFrequency, endFrequency int) error {
	command := fmt.Sprintf("ffmpeg -i %s -af \"lowpass=%d,highpass=%d\" %s", inputFilename, endFrequency, startFrequency, outputFilename)
	return ExecShell(command)
}

/*
*
ExtractStereoChannels function takes the input audio file path and the output file paths for the left and right channels. It uses the pan audio filter with the 1c option to extract each channel separately. The c0=c0 parameter specifies that the left channel should be taken from the original audio's first channel (c0), and c0=c1 specifies that the right channel should be taken from the original audio's second channel (c1). The resulting audio files will contain the left and right channels of the stereo audio, respectively
*/
func ExtractStereoChannels(inputFilename, outputLeftFilename, outputRightFilename string) error {
	command := fmt.Sprintf("ffmpeg -i %s -filter_complex \"[0:a]pan=1c|c0=c0[left];[0:a]pan=1c|c0=c1[right]\" -map [left] %s -map [right] %s",
		inputFilename, outputLeftFilename, outputRightFilename)
	return ExecShell(command)
}

// Remove silent parts from a video or audio.
func TrimSilence(inputFilename, outputFilename string, silenceThreshold float64) error {
	command := fmt.Sprintf("ffmpeg -i \"%s\" -af silenceremove=1:%.2f dB -c:v copy \"%s\" -y", inputFilename, silenceThreshold, outputFilename)
	return ExecShell(command)
}

// Convert videos matching a glob pattern to output format with specified quality
// Codec is guessed from output format: e.g mp4, mkv, etc
// quality value should be between 1 and 31, where 1 represents the highest quality
func ConvertVideos(globPattern, outputFormat, outputDir string, quality int) error {
	globResult, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("error expanding glob pattern")
	}

	errList := make([]error, 0, len(globResult))

	// Create a channel to receive notifications when conversions are complete
	done := make(chan bool)

	if outputDir != "" {
		os.MkdirAll(outputDir, 0755)
		fmt.Println("Creating output dir: ", outputDir)
	}

	// Launch a goroutine for each video conversion
	for _, inputFilename := range globResult {
		go func(inputFilename string) {
			outputFilename := strings.TrimSuffix(inputFilename, filepath.Ext(inputFilename))

			if outputDir == "" {
				outputFilename += "." + outputFormat
			} else {
				outputFilename = filepath.Join(outputDir, filepath.Base(outputFilename)+"."+outputFormat)
			}

			command := fmt.Sprintf("ffmpeg -i \"%s\" -crf %d \"%s\"", inputFilename, quality, outputFilename)
			err := ExecShell(command)
			if err != nil {
				errList = append(errList, err)
			}

			// Notify the channel that the conversion is complete
			done <- true
		}(inputFilename)
	}

	// Wait for all conversions to complete
	for range globResult {
		<-done
	}

	if len(errList) == 0 {
		return nil
	}

	var errString string
	for _, e := range errList {
		errString += e.Error() + "\n"
	}
	return fmt.Errorf(errString)
}

func ConcatenateVideos(inputFilenames []string, outputFilename string, destWidth, destHeight int) error {
	// Prepare the FFmpeg command
	cmdArgs := []string{}

	// Add input files to the command
	for _, filename := range inputFilenames {
		cmdArgs = append(cmdArgs, "-i", filename)
	}

	// Set video filter to resize all videos to the destination width and height
	cmdArgs = append(cmdArgs, "-vf", fmt.Sprintf("scale=%d:%d", destWidth, destHeight))

	// Set the output file name
	cmdArgs = append(cmdArgs, "-c:v", "libx264", "-preset", "slow", "-crf", "22", outputFilename)

	// Generate the FFmpeg command
	commandArgs := append([]string{"ffmpeg"}, cmdArgs...)
	command := strings.Join(commandArgs, " ")
	return ExecShell(command)
}

// Function to extract a specific audio channel from a multi-channel audio file
func ExtractAudioChannel(inputFilename, outputFilename string, channelIndex int) error {
	command := fmt.Sprintf("ffmpeg -i %s -map 0:%d -c:a copy %s", inputFilename, channelIndex, outputFilename)
	return ExecShell(command)
}

// Function to convert a video file to a GIF
func ConvertToGif(inputFilename, outputFilename string, width, height, fps int, from, to string) error {
	command := fmt.Sprintf("ffmpeg -i %s -ss %s -to %s -vf \"fps=%d,scale=%d:%d:flags=lanczos\" -c:v gif %s", inputFilename, from, to, fps, width, height, outputFilename)
	return ExecShell(command)
}

// Function to extract a specific video segment based on duration
func ExtractVideoSegment(inputFilename, outputFilename, startTime, duration string) error {
	command := fmt.Sprintf("ffmpeg -i %s -ss %s -t %s -c copy %s", inputFilename, startTime, duration, outputFilename)
	return ExecShell(command)
}

// Function to add subtitles to a video
func AddSubtitles(inputFilename, outputFilename, subtitlesFilename string) error {
	command := fmt.Sprintf("ffmpeg -i %s -vf subtitles=%s %s", inputFilename, subtitlesFilename, outputFilename)
	return ExecShell(command)
}

// Function to extract frames from a video
func ExtractFrames(inputFilename, outputPattern string, frameRate int) error {
	command := fmt.Sprintf("ffmpeg -i %s -vf fps=%d %s", inputFilename, frameRate, outputPattern)
	return ExecShell(command)
}

// Function to merge an audio file with a video file
func MergeAudioWithVideo(audioFilename, videoFilename, outputFilename string) error {
	command := fmt.Sprintf("ffmpeg -i %s -i %s -c:v copy -c:a aac -map 0:v:0 -map 1:a:0 -shortest %s", videoFilename, audioFilename, outputFilename)
	return ExecShell(command)
}

// Function to apply a watermark image on a video with a specific size
func ApplyImageWatermark(inputFilename, outputFilename,
	watermarkImage string, position string, width, height int) error {
	command := fmt.Sprintf("ffmpeg -i %s -i %s -filter_complex \"[1]scale=%d:%d [overlay]; [0][overlay]overlay=%s\" -c:a copy %s",
		inputFilename, watermarkImage, width, height, position, outputFilename)
	return ExecShell(command)
}

// Function to apply a text watermark on a video with a specific size
func ApplyTextWatermark(inputFilename, outputFilename, text, position, bgColor, fgColor string, fontSize float32, fontPath string) error {
	filter := ""
	if fontPath != "" {
		filter = fmt.Sprintf("drawtext=fontfile='%s':", fontPath)
	} else {
		filter = "drawtext="
	}

	escapedText := escapeTextForFilter(text)
	command := fmt.Sprintf("ffmpeg -i %s -vf \"%s text='%s':fontsize=%f:box=1:boxcolor=%s@0.2:boxborderw=5:x=%s:y=%s:fontcolor=%s\" %s",
		inputFilename, filter, escapedText, fontSize,
		getColorString(fgColor), getPositionX(position),
		getPositionY(position), getColorString(bgColor),
		outputFilename)
	return ExecShell(command)
}

// Helper function to format the color string as valid FFmpeg color format
func getColorString(color string) string {
	// Check if the color string is in the format "0xRRGGBB" or "0xRRGGBBAA"
	if len(color) == 7 || len(color) == 9 {
		return fmt.Sprintf("0x%s", color[1:])
	}

	// If the color string is not in the expected format, return a default color
	return "black" // Default color, change it as per your requirement
}

// Helper function to escape special characters in the text for the FFmpeg filter
func escapeTextForFilter(text string) string {
	// Add escape characters for ':', '\', and '|'
	escapedText := strings.ReplaceAll(text, ":", "\\:")
	escapedText = strings.ReplaceAll(escapedText, "\\", "\\\\")
	escapedText = strings.ReplaceAll(escapedText, "|", "\\|")
	return escapedText
}

// Helper function to calculate the X position based on the specified position string
func getPositionX(position string) string {
	switch position {
	case "top-left", "middle-left", "bottom-left":
		return "5" // Adjust the value according to your needs
	case "top-right", "middle-right", "bottom-right":
		return "(w-text_w)-5" // Adjust the value according to your needs
	case "center":
		return "(w-text_w)/2"
	default: // Default to center position
		return "(w-text_w)/2"
	}
}

// Helper function to calculate the Y position based on the specified position string
func getPositionY(position string) string {
	switch position {
	case "top-left", "top-middle", "top-right":
		return "5" // Adjust the value according to your needs
	case "bottom-left", "bottom-middle", "bottom-right":
		return "(h-text_h)-5" // Adjust the value according to your needs
	case "center":
		return "(h-text_h)/2"
	default: // Default to center position
		return "(h-text_h)/2"
	}
}

// Function to change the volume of an audio file
func ChangeAudioVolume(inputFilename, outputFilename string, volumeFactor float32) error {
	command := fmt.Sprintf("ffmpeg -i %s -af \"volume=%.2f\" %s", inputFilename, volumeFactor, outputFilename)
	return ExecShell(command)
}

// Function to extract specific audio channels from a multi-channel audio file
func ExtractAudioChannels(inputFilename, outputFilename string, channelIndices []int) error {
	channelArgs := ""
	for i, index := range channelIndices {
		channelArg := fmt.Sprintf("pan=1c|c%d=c%d", i, index)
		channelArgs += channelArg
		if i != len(channelIndices)-1 {
			channelArgs += "|"
		}
	}
	command := fmt.Sprintf("ffmpeg -i %s -map 0 -af \"%s\" %s", inputFilename, channelArgs, outputFilename)
	return ExecShell(command)
}

// Function to apply a video filter to a video file
func ApplyVideoFilter(inputFilename, outputFilename, filterExpression string) error {
	command := fmt.Sprintf("ffmpeg -i %s -vf \"%s\" %s", inputFilename, filterExpression, outputFilename)
	return ExecShell(command)
}

// Function to convert a video file to a series of images
// To keep original dimensions of video. Pass width=0 and height=0
func VideoToImages(inputFilename, outputDirectory,
	imageFormat string, width, height int, from string, to string) error {
	scaleFilter := ""
	if width > 0 && height > 0 {
		scaleFilter = fmt.Sprintf("-vf \"scale=%d:%d\"", width, height)
	}

	timeFilter := ""
	if from != "" {
		timeFilter += fmt.Sprintf("-ss %s ", from)
	}

	if to != "" {
		timeFilter += fmt.Sprintf("-to %s ", to)
	}

	command := fmt.Sprintf("ffmpeg -i %s %s %s %s/%s", inputFilename, scaleFilter, timeFilter, outputDirectory, imageFormat)
	return ExecShell(command)
}

// Extract frame at given video timestamp.
// To use this function, you need to pass the input video file, output image filename,
// and the desired timestamp in the format "hh:mm:ss"
// (e.g., "00:01:23" to extract the frame at 1 minute and 23 seconds)
func ExtractFrame(inputFilename, outputFilename string, timestamp string) error {
	command := fmt.Sprintf("ffmpeg -i %s -ss %s -vframes 1 %s", inputFilename, timestamp, outputFilename)
	return ExecShell(command)
}

// Function to convert a series of images to a video file
// e.g ImagesToVideo("input_images/*.png", "output.mp4", 30, 640, 480, "slow")
func ImagesToVideo(inputFormat, outputFilename string, framerate, width, height int) error {
	scaleFilter := ""
	if width > 0 && height > 0 {
		scaleFilter += fmt.Sprintf(" -s %dx%d ", width, height)
	}
	command := fmt.Sprintf("ffmpeg -framerate %d -pattern_type glob -i '%s' -c:v libx264 -pix_fmt yuv420p %s %s",
		framerate, inputFormat, scaleFilter, outputFilename)
	return ExecShell(command)
}

// Scale image to given width and height.
func ScaleImage(inputFilename, outputFilename string, width, height int) error {
	command := fmt.Sprintf("ffmpeg -i %s -vf %s %s",
		inputFilename,
		fmt.Sprintf("scale=%d:%d", width, height),
		outputFilename,
	)

	return ExecShell(command)
}

// Scale image to a given aspect ratio.
func ScaleImageToAspectRation(inputFilename, outputFilename string, aspectRation float32) error {
	command := fmt.Sprintf("ffmpeg -i %s -vf \"scale=w=iw*%f:h=ih*%f\" %s",
		inputFilename, aspectRation, aspectRation, outputFilename)
	return ExecShell(command)
}

// Prints the metadata for a media file.
func PrintMediaMetadata(filename string) error {
	command := fmt.Sprintf("ffprobe -hide_banner -i %s", filename)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
