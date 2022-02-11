package main

import "fmt"

func main() {
    // ваш код должен быть здесь
    for i := 0; i <= 10; i++ {
        var n int
        fmt.Scanln(&n)
        switch n {
            case 0:
            fmt.Println("Ноль")
            case 1:
            fmt.Println("Один")
            case 2:
            fmt.Println("Два")
            case 3:
            fmt.Println("Три")
            case 4:
            fmt.Println("Четыре")
            case 5:
            fmt.Println("Пять")
            case 6:
            fmt.Println("Шесть")
            case 7:
            fmt.Println("Семь")
            case 8:
            fmt.Println("Восемь")
            case 9:
            fmt.Println("Девять")
            case 10:
            fmt.Println("Десять")
        }
    }

}