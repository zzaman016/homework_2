package main

import "fmt"

func main(){
	var a,b,c float64
	fmt.Print("Введите a,b,c: ")
	fmt.Scan(&a,&b,&c)
	if a>=b && a>=c{
		fmt.Println("Максимальное число :", a)
	} else if b>=a && b>=c{
		fmt.Println("Максимальное число: ", b)
	} else if c>=a && c>=b{
		fmt.Println("Максимальное число: ", c)
	}
	if a<=b && a<=c{
		fmt.Println("Минимальное число :", a)
	} else if b<=a && b<=c{
		fmt.Println("Минимальное число :", b)
	} else if c<=a && c<=b{
		fmt.Println("Минимальное число :", c)
	}
	if a>b && a<c{
		fmt.Println("Среднее число :", a)
	} else if b>a && b<c{
		fmt.Println("Среднее число :", b)
	} else if c>a && c<b{
		fmt.Println("Среднее число :", c)
	}
}
