package main

import (
	"log"
	"os"

	"github.com/korjavin/activityPubBot/internal/poster/mastodon"
)

type Poster interface {
	Post(server, text, imagePath string) error
}

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
	var imagePath string
	if len(os.Args) == 4 {
		imagePath = os.Args[3]
	}

	// Create a new Mastodon poster
	poster := mastodon.NewPoster(token)

	// Post the status
	err := poster.Post(server, text, imagePath)
	if err != nil {
		log.Fatalf("Failed to post status: %v", err)
	}

	log.Println("Status posted successfully!")
}
