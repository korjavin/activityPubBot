package mastodon

import (
	"context"
	"os"

	"github.com/mattn/go-mastodon"
)

type Poster struct {
	client *mastodon.Client
}

func NewPoster(token string) *Poster {
	client := mastodon.NewClient(&mastodon.Config{
		AccessToken: token,
	})
	return &Poster{client: client}
}

func (p *Poster) Post(server, text, imagePath string) error {
	p.client.Config.Server = server

	var mediaID mastodon.ID
	if imagePath != "" {
		file, err := os.Open(imagePath)
		if err != nil {
			return err
		}
		defer file.Close()

		media, err := p.client.UploadMedia(context.Background(), imagePath)
		if err != nil {
			return err
		}
		mediaID = media.ID
	}

	status := &mastodon.Toot{
		Status: text,
	}
	if mediaID != "" {
		status.MediaIDs = []mastodon.ID{mediaID}
	}

	_, err := p.client.PostStatus(context.Background(), status)
	if err != nil {
		return err
	}

	return nil
}
