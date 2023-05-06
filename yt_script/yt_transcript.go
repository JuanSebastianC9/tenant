package main

import (
    "fmt"
    "tensorflow/go"
    "io"
	"os"
	"strings"
    "net/url"

	"github.com/kkdai/youtube/v2"
)

func main() {

    transcribe()
}

func download(youtubeURL string) {
    //get videoID	
    videoID, err := getVideoIDFromURL(youtubeURL)
	
    if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Video ID:", videoID)
	}

    client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

func transcribe() {
    // Create a TensorFlow session.
    sess := tf.NewSession()

    // Load the speech recognition model.
    model := tf.LoadModel("model.pb")

    // Recognize speech.
    audio := tf.NewAudio("audio.mp4")
    transcript := model.Recognize(sess, audio)

    // Print the transcript.
    fmt.Println(transcript)
}

func getVideoIDFromURL(youtubeURL string) (string, error) {
    u, err := url.Parse(youtubeURL)
	
    //Check for error
    if err != nil {
		return "", err
	}
	
    //Ensure it is a youtube url
    if u.Host != "www.youtube.com" && u.Host != "youtube.com" {
		return "", fmt.Errorf("invalid youtube url")
	}
	
    if u.Path == "/watch" {
		values := u.Query()
		videoID := values.Get("v")
		return videoID, nil
	} 
    else if strings.HasPrefix(u.Path, "/embed/") {
		parts := strings.Split(u.Path, "/")
		if len(parts) > 2 {
			videoID := parts[2]
			return videoID, nil
		}
	}
	return "", fmt.Errorf("could not find video id in url")
}