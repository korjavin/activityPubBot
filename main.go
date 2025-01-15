package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

func main() {
	// Get the access token from the environment variable
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN environment variable is not set")
	}

	// Parse command line arguments
	if len(os.Args) < 3 {
		log.Fatal("Usage: activityPubBot SERVER TEXT [IMAGE_PATH]")
	}
	server := os.Args[1]
	text := os.Args[2]

	// Create a new Mastodon client
	client := mastodon.NewClient(&mastodon.Config{
		Server:      server,
		AccessToken: token,
	})

	var mediaID mastodon.ID
	if len(os.Args) == 4 {
		imagePath := os.Args[3]
		file, err := os.Open(imagePath)
		if err != nil {
			log.Fatalf("Failed to open image file: %v", err)
		}
		defer file.Close()

		media, err := client.UploadMedia(context.Background(), imagePath)
		if err != nil {
			log.Fatalf("Failed to upload media: %v", err)
		}
		mediaID = media.ID
	}

	// Post the status
	status := &mastodon.Toot{
		Status: text,
	}
	if mediaID != "" {
		status.MediaIDs = []mastodon.ID{mediaID}
	}
	_, err := client.PostStatus(context.Background(), status)
	if err != nil {
		log.Fatalf("Failed to post status: %v", err)
	}

	fmt.Println("Status posted successfully!")
}
