package main

import (
    "io"
    "log"
    "os"
)

func main() {

    src_g := "global_conf.json"
    dst_folder := "/opt/siliq/"


    info,err := os.Stat("/opt/siliq")
    if os.IsNotExist(err) {
//        log.Fatal("File does not exist.")
        if err := os.Mkdir("/opt/siliq", os.ModePerm); err != nil {
          log.Fatal(err)
	}
    }
    log.Println(info)


    fin, err := os.Open(src_g)
    if err != nil {
        log.Fatal(err)
    }
    defer fin.Close()

    fout, err := os.Create(dst_folder+src_g)
    if err != nil {
        log.Fatal(err)
    }
    defer fout.Close()

    _, err = io.Copy(fout, fin)

    if err != nil {
        log.Fatal(err)
    }
}
