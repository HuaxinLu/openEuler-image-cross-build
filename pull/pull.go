package pull

import(
    "fmt"
    "log"
    "context"
    "encoding/json"
    "encoding/base64"
    "github.com/docker/docker/api/types"
	docker "hwbuild/dockerclient"
)

// pull image
func PullImage(image string) error { 
    // create docker client
	cli, err := docker.CreateClient()
	if err != nil {
        log.Fatal(err, " :unable to create docker client.")
        return err
    }

	// get username
	fmt.Print("Username: ")
	var userName string
	fmt.Scanln(&userName)
	// get password
	fmt.Print("Password: ")
    var password string
    fmt.Scanln(&password)

	// change string to json
    auth := types.AuthConfig{Username: userName, Password: password,}
    encodedJSON, err := json.Marshal(auth)
    if err != nil {
        log.Fatal(err, " :unable to encode auth.")
        return err
    }
    authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	// pull
	fmt.Printf("Start to pull...\n")
    pullResp, err := cli.ImagePull(context.Background(), 
		image,
        types.ImagePullOptions{RegistryAuth: authStr,})
	if err != nil {
		log.Fatal(err, " :unable to pull docker image")
        return err
	}
	
    err = docker.ResolvePull(pullResp, image)
    fmt.Printf("Pull finished...\n")
	return err
}
