package build

import(
    "fmt"
    "bytes"
    "context"
	"log"
	"os"
	"io"
	"path/filepath"
	"errors"
	"archive/tar"
	"encoding/json"
    "github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
)

type BuildOpt struct{
    DockerfileName string
    ImageTag string
}

//build镜像
func BuildImage(filePath string, opt* BuildOpt)error {
	//新建docker客户
	cli, err := client.NewClient("tcp://0.0.0.0:8088", "v1.40", nil, nil)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }
	//合成dockerfile路径
	dockerfilePath := filepath.Join(filePath, opt.DockerfileName)
	fmt.Printf("The path of dockerfile is %s.\n", dockerfilePath)
	//修改dockerfile，生成用于跨平台构建的dockerfile
	dockerfileCrossName, err := ChangeFile(dockerfilePath)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }
	//dockerfileCrossName := "DockerfileCross"
	//打包文件，建立build上下文环境
    buf := new(bytes.Buffer)
    tw := tar.NewWriter(buf)
    defer tw.Close()
	//文件目录必须为文件夹
    sfileInfo, err := os.Stat(filePath)
    if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	if !sfileInfo.IsDir() {
		fmt.Println("The filepath must be folder.")
		return errors.New("The filepath must be folder")
	}
	//建立tar文件，注意不包含第一级目录，即直接打包文件
	tarCreate(filePath, tw)
	//构建镜像
	tarReader := bytes.NewReader(buf.Bytes())
    imageBuildResponse, err := cli.ImageBuild(
        context.Background(),
        tarReader,
        types.ImageBuildOptions{
            Context:    tarReader,
            Dockerfile: dockerfileCrossName,
            Tags:       []string{opt.ImageTag},
            Remove:     false})
    if err != nil {
        log.Fatal(err, " :unable to build docker image")
    }
	//处理构建中的消息
    //defer imageBuildResponse.Body.Close()
    //_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	buildResp := imageBuildResponse.Body
	d := json.NewDecoder(buildResp)
	
	type Event struct {
		Stream	string `json:"stream"`
		Error   string `json:"error"`
	}

	var event *Event
	fmt.Printf("Start to build...\n")
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
	
    return err
}
