package dockerclient

import(
	"github.com/docker/docker/client"
)

func CreateClient() (*client.Client, error) {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.40", nil, nil)
	return cli, err
}
