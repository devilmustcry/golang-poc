package main

import (
	"fmt"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	err := ffmpeg_go.Input("./video-without-forcing.mp4").Output("./test.mp3").OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Print(err)
	}
}
