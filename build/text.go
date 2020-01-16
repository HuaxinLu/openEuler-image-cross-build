package build

import(
    "os"
    "io"
    "bufio"
    //"fmt"
    "bytes"
)

func ChangeFile(filePath string) (string, error) {
    // open origin dockerfile
    df_ori, err := os.OpenFile(filePath, os.O_RDWR,0644)
    if err != nil {
        return "", err
    }
    defer df_ori.Close()

    // create new dockerfile
	filePathNew := filePath + ".cross"
    df_new, err := os.OpenFile(filePathNew, os.O_WRONLY|os.O_CREATE, 0644)
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
                return filePathNew, nil
            }
            return filePathNew, err
        }
		// write new dockerfile
		_, err = bfWr.Write(line)
		if err != nil {
			return filePathNew, err
		}	
        // find the FROM line
		if hasAdded == false {
			fields := bytes.Fields(line)
			//for _, b := range fields {
			//	fmt.Printf("%q\n", b)
			//}
			// if this is a FROM command
			if "FROM" == string(fields[0]) {
				hasAdded = true
				bfWr.WriteString("COPY qemu-aarch64-static /usr/bin/qemu-aarch64-static\n")				
			}	
		}
		if err = bfWr.Flush(); err != nil {
			return "DockerfileCross", err
		}
		
    }
   return filePathNew, nil 
}
