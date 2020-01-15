package push

import(
    "fmt"
    //"os"
    "io"
	//"strings"
    "context"
    "encoding/json"
    "encoding/base64"
    "github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
)

//push镜像
func PushImage(image string) error {
    cli, err := client.NewClient("tcp://0.0.0.0:8088", "v1.40", nil, nil)
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
    }

	auth := types.AuthConfig{Username: "huaxinlu", Password: "l95h08x27",}
	encodedJSON, err := json.Marshal(auth)
		
    if err != nil {
        fmt.Printf("%v\n", err)
        return err
	}
	
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	pushResp, err := cli.ImagePush(context.Background(),
		image,
		types.ImagePushOptions{RegistryAuth: authStr,})
	//io.Copy(os.Stdout, pushResp)
	
    d := json.NewDecoder(pushResp)                                                     
    
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
    fmt.Printf("Start to push...\n")                                                   
    for {
        if err := d.Decode(&event); err != nil {                                       
            if err == io.EOF {                                                         
                break                                                                  
            }
            panic(err)                                                                 
        }                                                                              
        
        if event.Error != "" {  
            fmt.Printf("Error: %s\n", event.Error)
        }
		//fmt.Println(event.Status)
    }

	fmt.Printf("Push finished...\n")

	return err
}
