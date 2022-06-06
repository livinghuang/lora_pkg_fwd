package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func copyFile(src string, dstFolder string, perm fs.FileMode) error {

	targetFileName := dstFolder + src
	//check target file exist , if exist , remove
	info, err := os.Stat(targetFileName)
	if os.IsExist(err) {
		log.Fatal("File exist.")
		e := os.Remove(targetFileName)
		if e != nil {
			log.Fatal(e)
		}
	} else {
		log.Println("file created", info)
	}

	input, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(targetFileName, input, perm)
	if err != nil {
		fmt.Println("Error creating", dstFolder+src)
		fmt.Println(err)
	}
	return err

}

func main() {

	src_g := "global_conf.json"
	src_l := "local_conf.json"
	src_lpf := "lora_pkt_fwd"
	src_rls := "reset_lgw.sh"
	src_ses := "set_eui.sh"
	src_ss := "start.sh"
	dst_opt := "/opt/siliq/"
	dst_service := "/etc/systemd/system/"
	info, err := os.Stat(dst_opt)
	if os.IsNotExist(err) {
		log.Println("Target folder does not exist.")
		if err := os.Mkdir(dst_opt, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		log.Println("Target folder created", info)
	} else {
		log.Println("Target exist", info)
	}

	info1, err := os.Stat(dst_service)
	if os.IsNotExist(err) {
		log.Println("Target folder does not exist.")
		if err := os.Mkdir(dst_service, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		log.Println("Target folder created", info1)
	} else {
		log.Println("Target folder exist", info1)
	}
	copyFile(src_g, dst_opt, 0666)
	copyFile(src_l, dst_opt, 0666)
	copyFile(src_lpf, dst_opt, 0755)
	copyFile(src_rls, dst_opt, 0755)
	copyFile(src_ses, dst_opt, 0755)
	copyFile(src_ss, dst_opt, 0755)
	copyFile(src_lpf+".service", dst_service, 0755)
}
