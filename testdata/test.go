package main

import "fmt"

func main(){
	Println("pizza")
	_= Println("pizza")
	_,_=fmt.Println("pies")
}

func Println(a string)error{
	println(a)
	return nil
}
