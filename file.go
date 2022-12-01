package main

import (
	"fmt"
	"os"
	"io"
) 

func main () {
	// var n int
	filename := "test.txt"
	file , err := os.OpenFile(filename,os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	// fmt.Println(n)
	tmp := make([]byte ,1024)
	var val []byte

	content := "Hello GoCoder"
	_, err = io.WriteString(file, content)
	if err != nil{
		fmt.Println("Write file err =", err)
		return
	}
	file.Seek(0, os.SEEK_SET) //重置文件指针
	for {
		n,err := file.Read(tmp)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		if n == 0 {
			break
		}
		val = append(val,tmp[:n]...)
		fmt.Println("read once...",n)
	}
	fmt.Println(string(val))
	if err := file.Close() ; err != nil {
		fmt.Println("Close file err =", err)
		return
	}
	if err := file.Close() ; err != nil {
		fmt.Println("Close file err =", err)
		return
	}
	fmt.Println("Close file success")
}
