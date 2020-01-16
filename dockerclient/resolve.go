package dockerclient

import (
	"fmt"
	"io"
	"encoding/json"
	"strings"
)

type EventPull struct {
	Status         string `json:"status"`
	Error          string `json:"error"`
	Progress       string `json:"progress"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
}

type EventPush struct {
	Status         string `json:"status"`
	Error          string `json:"error"`
	Progress       string `json:"progress"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
}

type EventBuild struct {
	Stream	string `json:"stream"`
	Error   string `json:"error"`
}

func ResolvePull(resp io.ReadCloser, image string) error {
	var event *EventPull
	d := json.NewDecoder(resp)
	for {
        if err := d.Decode(&event); err != nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }
        if event.Error != "" {
            fmt.Printf("Error:  %s\n", event.Error)
        }
    }
	if event != nil {
        if strings.Contains(event.Status, "Downloaded newer image") {
            fmt.Printf("Downloaded newer image for %s\n", image)
        }
        if strings.Contains(event.Status, "Image is up to date") {
            fmt.Printf("Image is up to date for %s\n", image)
        }
    }
    return nil
}

func ResolvePush(resp io.ReadCloser, image string) error {
    var event *EventPush
    d := json.NewDecoder(resp)
    for {
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
        }
		if event.Error != "" {
			fmt.Printf("Error:  %s\n", event.Error)
		}
	}
	return nil
}

func ResolveBuild(resp io.ReadCloser) error {
    var event *EventBuild
    d := json.NewDecoder(resp)
    for {
        if err := d.Decode(&event); err != nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }
		if event.Stream != "" && event.Stream != "\n" {
			fmt.Printf("Docker message: %s\n", event.Stream)
		}
		if event.Error != "" {
			fmt.Printf("Docker error: %s\n", event.Error)
		}
    }
    return nil
}

