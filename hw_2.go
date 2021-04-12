package main

import "fmt"

func main(){
	var a,b,c float64
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("c: ")
	fmt.Scan(&c)
	if a+b<=c || a+c<=b || b+c<=a{
							fmt.Println("Треугольник не существует")
	} else if a!=b && b!=c && a!=c{
							fmt.Print("Разносторонний треугольник")
	} else if a==b && b==c && a==c{
							fmt.Print("Равносторонний треугольник")
	} else if a==b || b==c || a==c {
							fmt.Print("Равнобедренный треугольник")
	}


}
