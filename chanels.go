package main

import ("fmt"
		"reflect"
		// "unicode"
)
func sum(res int, chanel chan int) int {
	res += 5
	chanel <- res
	return res
}

func main() {
	chanel := make(chan int)
	var digit int
	fmt.Print("Введите число: ")
	fmt.Scan(&digit)
	fmt.Println(reflect.TypeOf(digit))

	// if !unicode.IsInt(digit) {
	// 	fmt.Println("Введите число!")
	// }
	
	go sum(digit, chanel)
	result := <- chanel
	fmt.Println("Ваша скидка:", result)
}