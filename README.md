# Audio Whisper Processor

This Go program processes audio files in a directory using OpenAI's Whisper, converting speech to text into `.srt` subtitle files.  
It supports all audio and video formats that Whisper can handle and processes files sequentially, creating `.srt` files in the specified output directory.

## Usage example
`go run main.go --dir /path/to/files --output_dir /path/to/subtitles --lang Russian --model tiny`

### Optional args:
`--dir` where your files are(current directory by default)

`--output_dir` where your subtitles are(current directory by default)

`--lang` russian by default

`--model` medium by default, but you can choose: `tiny, tiny.en, base, base.en, small, small.en, medium, medium.en, large, large-v2`

## Features
- Skip already processed files (if `.srt` exists in output directory)
- Sequential processing (single thread)
- Automatically generates `.srt` files in the output directory
- Shows processed files with `Done: filename.srt` in green
- Supports specifying input directory, output directory, and language via command-line flags
- Optional GPU acceleration if Whisper supports CUDA

## Supported Formats
Whisper supports a wide range of audio and video formats, including:  
`.3gp, .aac, .aiff, .caf, .flac, .m4a, .mka, .mp3, .mp4, .mpeg, .mpg, .mov, .ogg, .opus, .wav, .webm`

## Requirements
- Go 1.20+ installed
- [Whisper CLI](https://github.com/openai/whisper) installed and available in PATH
- Optional: PyTorch with CUDA for GPU acceleration

Install Whisper:

`pip install git+https://github.com/openai/whisper.git`

or

```bash
python3 -m pip install --user pipx
python3 -m pipx ensurepath
pipx install git+https://github.com/openai/whisper.git
```