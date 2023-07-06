# mediaconv

Intuitive Media Converter CLI that calls into existing ffmpeg binary to
perform common video operations.

The CLI is written with cobra in go language. This allows you to install completions
for your favourite shell.

Example:

```txt
./mediaconv -h

mediaconv

Convert videos, audio, watermark your videos,
manipulate audio using this simple cli.

Note that you must have ffmpeg installed on your machine.

Usage:
  mediaconv [command]

Available Commands:
  addAudio       Add audio to a video. Trim at the shortest
  applyFilter    Apply video filter
  addSubtitles   Add subtitles to video from a file
  changeVolume   Adjust audio volume of audio or video
  completion     Generate the autocompletion script for the specified shell
  concatVideos   Concatenate multiple videos
  convertVideo   Convert one or more videos
  createGif      Create an animated GIF from a video
  cutVideo       A brief description of your command
  extractAudio   Extract audio from a video
  extractFrame   Extract video frame
  extractFreq    Generate a new audio file with frequencies within the specified range
  extractStereo  Seperate left and right stereo channels from input
  genImages      Generate Images from video
  help           Help about any command
  imageWatermark Apply a image watermark to video
  img2Video      Create a video from a folder of images
  metadata       Print media metadata
  resize         Resize Video
  rotateVideo    Rotate Video
  scaleImage     Scale image to specified width, height or aspect ratio
  textWatermark  Apply a text watermark to video
  trimSilence    Remove silent portions of audio from video/audio

Flags:
  -h, --help            help for mediaconv
  -i, --input string    Path to the input video
  -o, --output string   Path to the output video

Use "mediaconv [command] --help" for more information about a command.
```
