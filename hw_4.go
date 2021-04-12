package main

import "fmt"

func main(){
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	if num%2==0{
		fmt.Println("Число", num,  "является четным")
	} else {
		fmt.Println("Число", num, "является нечетным")
	}

}
