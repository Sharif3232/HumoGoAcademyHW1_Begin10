package main

import (
	"fmt"
)

func main() {

	var a, b float64

	fmt.Println("Введите значение a и b через запятую:")
	fmt.Scan(&a, &b)

	fmt.Println("Сумма квадратов a и b равно:", (a*a)+(b*b))
	fmt.Println("Разность квадратов a и b:", (a*a)-(b*b))
	fmt.Println("Произведение квадратов a и b:", (a*a)*(b*b))
	fmt.Println("Частное квадратов a и b:", (a*a)/(b*b))

}
