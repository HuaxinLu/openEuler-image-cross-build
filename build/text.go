package build

import(
    "os"
    "io"
    "bufio"
    //"fmt"
    "bytes"
)

func ChangeFile(filePath string) (string, error) {
    //打开原dockerfile
    df_ori, err := os.OpenFile(filePath, os.O_RDWR,0644)
    if err != nil {
        return "", err
    }
    defer df_ori.Close()

    //创建新的dockerfile
    df_new, err := os.OpenFile(filePath + "Cross", os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return "", err
    }
    defer df_new.Close()

    bfRd := bufio.NewReader(df_ori)
    bfWr := bufio.NewWriter(df_new)
	hasAdded := false 
    for {
        //按行读取原始dockerfile
        line, err := bfRd.ReadBytes('\n')
        if err != nil {
            if err == io.EOF {
                return "DockerfileCross", nil
            }
            return "DockerfileCross", err
        }
		//写入新的dockerfile
		_, err = bfWr.Write(line)
		if err != nil {
			return "DockerfileCross", err
		}	
        //按空格截取
		if hasAdded == false {
			fields := bytes.Fields(line)
			//for _, b := range fields {
			//	fmt.Printf("%q\n", b)
			//}
			//如果截取的命令包含FROM
			if "FROM" == string(fields[0]) {
				hasAdded = true
				bfWr.WriteString("COPY qemu-aarch64-static /usr/bin/qemu-aarch64-static\n")				
			}	
		}
		//刷新缓存
		if err = bfWr.Flush(); err != nil {
			return "DockerfileCross", err
		}
		
    }
   return "DockerfileCross", nil 
}
