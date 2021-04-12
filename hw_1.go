package main

import "fmt"

func main(){
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)
	if year % 4 != 0{
				fmt.Println("Обычный год")
	} else if year % 4 == 0 && year % 100 != 0{
				fmt.Print("Високосный год")
	} else if (year % 100 == 0 && year % 400 == 0) {
				fmt.Print("Високосный год")
	} else{
				fmt.Print("Обычный год")
	}
}

