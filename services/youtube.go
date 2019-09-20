package services

import (
	"log"
	"net/http"
	"strings"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YoutubeSearcher struct {
	developKey string
	maxResult  int64
}

func NewYoutubeSearcher(developKey string) *YoutubeSearcher {
	return &YoutubeSearcher{developKey: developKey}
}

func (searcher *YoutubeSearcher) Search(query string) (string, string) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: searcher.developKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(query).
		MaxResults(1).
		Type("video").
		Order("relevance")
	response, err := call.Do()

	item := response.Items[0]

	return item.Id.VideoId, item.Snippet.Title
}

func (searcher *YoutubeSearcher) ValidateResult(songName, resultTitle string) bool {
	return strings.Contains(strings.ToLower(resultTitle), strings.ToLower(songName))
}
