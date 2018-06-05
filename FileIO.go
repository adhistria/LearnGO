package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
)
func main(){
	file, err:=os.Create("coba.txt")

	if err!= nil{
		log.Fatal(err)
	}

	file.WriteString("This is random text")
	file.Close()

	stream,err:= ioutil.ReadFile("coba.txt")

	if(err!=nil){
		log.Fatal(err)
	}
	readString:= string(stream)
	fmt.Println(readString)
}