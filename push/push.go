package push

import(
	"fmt"
	"log"
	"context"
	"encoding/json"
	"encoding/base64"
	"github.com/docker/docker/api/types"
	docker "hwbuild/dockerclient"
)

//push镜像
func PushImage(image string) error {
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
	
	//push
	fmt.Printf("Start to push...\n")
	pushResp, err := cli.ImagePush(context.Background(),
		image,
		types.ImagePushOptions{RegistryAuth: authStr,})
	if err != nil {
		log.Fatal(err, " :unable to push docker image.")
		return err
	}
	
	err = docker.ResolvePush(pushResp, image)
	fmt.Printf("Push finished...\n")
	return err
}
