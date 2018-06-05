package main

import(
	"fmt"
	"time"
	"strconv"
)

func count(id int){
	for i:=0 ; i< 10 ;i++{
		fmt.Println(id,":",i)
		time.Sleep(time.Millisecond *1000)
	}
}
var pizzaNum = 0
var pizzaName = ""
func makeDough(stringChan chan string){
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond *10)

}

func addSauce(stringChan chan string){
	pizza := <-stringChan

	fmt.Println("Add Sauce and Send", pizza,"for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond *10)

}

func addToppings(stringChan chan string){
	pizza := <- stringChan
	fmt.Println("Add Toppings to",pizza,"and ship")
}

func main(){

	stringChan := make(chan string)

	for i:=0 ; i<3 ;i++{
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond*500)
	}
	//for i:=0 ; i<10 ; i++{
	//	go count(i)
	//}
	//
	//time.Sleep(time.Millisecond *11000)
}