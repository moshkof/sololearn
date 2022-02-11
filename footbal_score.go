package main

import "fmt"

func main() {
	var c int
	var res int 
	var result string
	fmt.Scan(&result)
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	results = append(results, result)
	for _, v := range results{
		if v == "w" {
			c = 3
		} else if v == "d" {
			c = 1
		} else if v == "l" {
			c = 0
		} else {
			fmt.Println("Введите: w, d или l")
			break
		}
		res += c
		// fmt.Println(c) 
	}
	fmt.Println("Итог: ", res)
	
}