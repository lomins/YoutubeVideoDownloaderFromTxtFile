package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {
	var videosIDs []string

	file, err := os.Open("videosIDs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		videosIDs = append(videosIDs, scanner.Text())
	}

	client := youtube.Client{}

	defer fmt.Scanln()
	for _, videoID := range videosIDs {

		video, err := client.GetVideo(videoID)
		if err != nil {
			panic(err)
		}

		formats := video.Formats.Quality("hd1080").WithAudioChannels()
		if formats == nil {
			formats = video.Formats.WithAudioChannels()
		}
		//formats := video.Formats.WithAudioChannels() // only get videos with audio

		stream, _, err := client.GetStream(video, &formats[0])

		if err != nil {
			panic(err)
		}

		videoTitle := video.Title
		videoTitle = symbolsReplacer(videoTitle)

		file, err = os.Create(videoTitle + ".mp4")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}

		fmt.Println("Succefull download: " + videoTitle)
	}
	fmt.Println("Нажмите Enter для выхода...")
}

func symbolsReplacer(videoTitle string) string {
	videoTitle = strings.Replace(videoTitle, "?", "", -1)
	videoTitle = strings.Replace(videoTitle, "!", "", -1)
	videoTitle = strings.Replace(videoTitle, "\\", "", -1)
	videoTitle = strings.Replace(videoTitle, "/", "", -1)
	videoTitle = strings.Replace(videoTitle, ":", "", -1)
	videoTitle = strings.Replace(videoTitle, "*", "", -1)
	videoTitle = strings.Replace(videoTitle, "\"", "", -1)
	videoTitle = strings.Replace(videoTitle, "<", "", -1)
	videoTitle = strings.Replace(videoTitle, ">", "", -1)
	videoTitle = strings.Replace(videoTitle, "|", "", -1)

	return videoTitle
}
