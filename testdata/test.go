package main

import "fmt"

func main(){
	PrintPie("pizza")
	_= PrintPizza("pizza")
//	_,_=fmt.Println("pies")
//	println("pizza")
	DoesNotReturnAnything()
}

func PrintPie(a string)error{
	//fmt.Println(a)
	return nil
}

func PrintPizza(a string)error{
	//fmt.Println(a)
	return nil
}

func DoesNotReturnAnything(){
}
