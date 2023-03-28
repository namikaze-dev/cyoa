package main_test

import (
	"strings"
	"testing"

	web "github.com/namikaze-dev/cyoa/cmd/web"
)

func TestNewStoryStore(t *testing.T) {
	store, err := web.NewStoryStore(strings.NewReader(sample_json))
	if err != nil {
		t.Errorf("unexpected error from web.NewStoryStore: %v", err)
	}

	assertEqual(t, store["intro"], web.Story{
		Title: "Title 1",
		Story: []string{
			"story line 1",
			"story line 2",
			"story line 3",
		},
		Options: []web.StoryOption{
			{
				Text: "text 1",
				Arc: "valley",
			},
		},
	})

	assertEqual(t, store["valley"], web.Story{
		Title: "Valley of Code",
		Story: []string{
			"story line 1",
			"story line 2",
		},
		Options: []web.StoryOption{
			{
				Text: "start",
				Arc: "intro",
			},
			{
				Text: "end",
				Arc: "home",
			},
		},
	})
}

var sample_json = `
{
    "intro": {
        "title": "Title 1",
        "story": [
            "story line 1",
            "story line 2",
            "story line 3"
        ],
        "options": [
            {
                "text": "text 1",
                "arc": "valley"
            }
        ]
    },
    "valley": {
        "title": "Valley of Code",
        "story": [
            "story line 1",
            "story line 2"
        ],
        "options": [
            {
                "text": "start",
                "arc": "intro"
            },
            {
                "text": "end",
                "arc": "home"
            }
        ]
    }
}`