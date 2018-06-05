package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hellow Adhi")
	//randNum :=1
	//var randNum =1
	//var Num2 = 2
	//var x float64 = 1.123123123123123
	//var age = 22
	//var favNumber = 7
	//randNum = "Hello"
	//fmt.Print(randNum)

	//fmt.Println(age,favNumber)

	//var (
	//	varA = 2
	//	varB = 3
	//)
	//fmt.Print(varA,varB)

	//var name = "Adhi Satria"
	//fmt.Println(name," Muhammad")
	//fmt.Print(name,"\n \nMuhammad")

	var isOver bool = true
	fmt.Printf("%T \n",isOver)
//	%T buat ngeliat tipe data kalo %t buat ngeliat nilainya
	var pi = 3.1753456345
	fmt.Printf("%.3f \n",pi)
	fmt.Printf("%T \n",pi)
	fmt.Println(pi)
//
	var c = 200
	//desimal
	fmt.Printf("%d \n",100)
	//biner
	fmt.Printf("%b \n",c)
	//character (ASCII)
	fmt.Printf("%c \n",100)
	//hexa
	fmt.Printf("%x \n",17)
	//equation like calculator
	fmt.Printf("%e \n",pi)
//	logical operator
	fmt.Println("true and false", true && false)
	fmt.Println("true or false", true || false)
	fmt.Println("!true = ", !true )

//looping
	i:=1
	for i <=10{
		fmt.Println(i)
		i++
	}
	for j:=0;j<5;j++{
		fmt.Println(j)
	}
//	condition

	yourAge:=18

	if(yourAge>=18){
		fmt.Println("You Can Drive")
	}else{
		fmt.Println("You Can't Drive")
	}

//	switch statements
	switch yourAge {
	case 16 : fmt.Println("You Can Drive")
	case 18 : fmt.Println("You Can Vote")
	default : fmt.Println("You Can't Drive")

	}
//	array
	var favNums2[5] float64
	favNums2[0] = 1
	favNums2[1] = 11
	favNums2[2] = 111.11
	favNums2[3] = 1111.123
	favNums2[4] = 11111

	//fmt.Println(favNums2[4])
	favNums3 :=[5]float64 {1,2,3,4,5}

	for i, value := range favNums3{
		fmt.Println(value, i)
	}
	for _, value := range favNums3{
		fmt.Println(value)
	}


	numSlice2 :=[5]float64 {1,2,34,5,63}
	for i, value:=range numSlice2{
		fmt.Println(value,i)
	}

//	slice
	numSlice :=[]int{5,4,3,2,1}
	//numSlice3 := numSlice[1:4]
	//stop sebelum index2
	//numSlice3 := numSlice[:2]
	//mulai dari index ke 2
	numSlice3 := numSlice[2:]

	for _, value:= range numSlice3{
		fmt.Println(value)
	}

//	copy slice
// make integer array with size 5 and maximum size 10
	numSlice4 := make([]int, 5,10)
	fmt.Println(numSlice4)
	copy(numSlice4,numSlice)
	fmt.Println(numSlice4)
	fmt.Println(numSlice4[0])
	numSlice4 = append(numSlice4,0,-1,3,4,56,6,5,2,5,6)
	fmt.Println(numSlice4)
//	maps golang
//  make maps array with key
	presAge := make(map[string] int)
	presAge["Adhi"] = 7
	fmt.Println(presAge["Adhi"])
// 	fungsi
	listNums :=[]float64{1,2,3,4,5}
	fmt.Println("Sum ", addThemUp(listNums))

	num1,num2 :=next2Values(5)
	fmt.Println(num1,num2)

	fmt.Println(substracThem(1,2,3,4,5))
//	Closure
	num := 4

	doubleNum := func(n int) int {
		return n*2
	}
	fmt.Println(doubleNum(num))
	fmt.Println(doubleNum(num))

//	rekursig

	fmt.Println("Faktorial 3", factorial(3))

// 	defer eksekusi setelah fungsinya telah selesai semua
//	defer printTwo()
	printOne()

	// recover kalo error di catch
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(3,2))
	// panic
	demPanic()
//	Pointer
	x:=0
	changeVal(&x)
	fmt.Println("X =",x)
	fmt.Println("Memory address of x = ",&x)
	//fmt.Println(x)

	y := new(int)
	fmt.Println(*y)
	changeVal2(y)
	fmt.Println("y = ", *y)
	s:=30
	fmt.Println(s)

//	data type
	rect1 := Rectangle{0,50,10,10}
	fmt.Println(rect1.area())
}



//create new Type
type Rectangle struct{
	leftX float64
	topY float64
	height float64
	width float64
}
// data type
func (rect *Rectangle) area() float64{
	return rect.width * rect.height
}

//type Shape struct{
//	area() float64
//}
//pointer
func changeVal(x *int){
	*x = 2
}

//pointer 2
func changeVal2( y *int){
	*y = 100
}

func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()
	panic("Panic")
}

func safeDiv(num1,num2 int) int{
	//menangkal error wkkw
	defer func(){
		//fmt.Println(recover())
		recover()
	}()

	solution :=num1/num2
	return solution
}

func addThemUp(numbers []float64) float64{
	sum :=0.0

	for _,value := range numbers{
		sum+=value
	}
	return sum
}

func next2Values(number int)(int, int){
	return number+1, number+2
}

//function tanpa tau inputan apa
func substracThem(args ...int) int{
	finalValue:=0
	for _,value:=range args{
		finalValue-=value
	}
	return finalValue
}

func factorial(num int) int{
	if num ==0 {
		return 1
	}else{
		return num*factorial(num-1)
	}
}

func printOne(){fmt.Println(1)}
func printTwo(){fmt.Println(2)}

