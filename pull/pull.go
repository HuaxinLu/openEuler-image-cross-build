package pull

import(
    "fmt"
    //"os"
    "io"
    "context"
	"strings"
    "encoding/json"
    "encoding/base64"
    "github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
)

//pull镜像
func PullImage(image string) error { 
    cli, err := client.NewClient("tcp://0.0.0.0:8088", "v1.40", nil, nil)
	if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }

    auth := types.AuthConfig{Username: "", Password: "",}
    encodedJSON, err := json.Marshal(auth)
    
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }

    authStr := base64.URLEncoding.EncodeToString(encodedJSON)
    pullResp, err := cli.ImagePull(context.Background(), 
		image,
        types.ImagePullOptions{RegistryAuth: authStr,})
    //io.Copy(os.Stdout, pullResp)
    d := json.NewDecoder(pullResp)

    type Event struct {
        Status         string `json:"status"`
        Error          string `json:"error"`
        Progress       string `json:"progress"`
        ProgressDetail struct {
            Current int `json:"current"`
            Total   int `json:"total"`
        } `json:"progressDetail"`
    }

    var event *Event
	fmt.Printf("Start to pull...\n")
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
        if strings.Contains(event.Status, fmt.Sprintf("Downloaded newer image for %s", image)) {
            fmt.Printf("Downloaded newer image for %s\n", image)
        }

        if strings.Contains(event.Status, fmt.Sprintf("Image is up to date for %s", image)) {
            fmt.Printf("Image is up to date for %s\n", image)
        }
    }
	return err
}
