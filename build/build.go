package build

import(
	"fmt"
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"errors"
	"archive/tar"
	"github.com/docker/docker/api/types"
	docker "hwbuild/dockerclient"
)

type BuildOpt struct{
	DockerfileName string
	ImageTag string
}

//build image
func BuildImage(filePath string, opt* BuildOpt)error {
	// create docker client
	cli, err := docker.CreateClient()
	if err != nil {
		log.Fatal(err, " :unable to create docker client.")
		return err
	}
	// get whole Dockerfile path
	dockerfilePath := filepath.Join(filePath, opt.DockerfileName)
	//fmt.Printf("The path of dockerfile is %s.\n", dockerfilePath)
	//change dockerfile, add cross build model
	_, err = ChangeFile(dockerfilePath)
	if err != nil {
		log.Fatal(err, " :unable to change dockerfile.")
		return err
	}
	// tar the whole file for building, used for build context
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()
	// filePath must be a folder
	sfileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err, " :unable to get file status.")
		return err
	}
	if !sfileInfo.IsDir() {
		fmt.Println("error: the filepath must be folder.")
		return errors.New("The filepath must be folder")
	}
	tarCreate(filePath, tw)
	//build
	tarReader := bytes.NewReader(buf.Bytes())
	imageBuildResponse, err := cli.ImageBuild(
		context.Background(),
		tarReader,
		types.ImageBuildOptions{
			Context:    tarReader,
			Dockerfile: opt.DockerfileName+".cross",
			Tags:       []string{opt.ImageTag},
			Remove:     false})
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	buildResp := imageBuildResponse.Body
	err = docker.ResolveBuild(buildResp)
	fmt.Printf("Build finished...\n")	
	return err
}
