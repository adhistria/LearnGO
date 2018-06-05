package main

import (
	"fmt"
	"strings"
	//"os"
	//"log"
	//"io/ioutil"
	//"strconv"
	"sort"
)

func main(){
	name := "Hello Adhidd"
	fmt.Println(strings.Contains(name,"ad"))
	fmt.Println(strings.Index(name,"Ad"))
	fmt.Println(strings.Count(name,"a"))
	fmt.Println(strings.Replace(name,"d","h",1))


	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString,","))

	listOfLetters := []string{"c","a","i"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters L",listOfLetters)

	listOfNums :=strings.Join([]string{"3","2","1"},",")
	fmt.Println(listOfNums)

//	file IO
}