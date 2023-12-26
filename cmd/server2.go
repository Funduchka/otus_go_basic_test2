package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		fmt.Printf("i = %d\n", i)
	}

	a := 1
	for a <= 3 {
		fmt.Printf("i = %d\n", a)
		a++
	}
	fmt.Println("Result", a)

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		} else if i > 3 {
			fmt.Printf("II = %d\n", i)
			break
		}
		fmt.Printf("i = %d\n", i)
	}

	for i := 1; i <= 6; i++ {
		fmt.Printf("i = %d\n", i)
		switch i {
		case 2:
			break
		default:
			fmt.Println("Ok")
		}
		fmt.Println(".")
	}

	letters := []string{"a", "b", "c"}
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}

	s := "My name"
	for i := 0; i < len(s); i++ {
		b := s[i]
		fmt.Println("b", b)
		// i строго последоваельно
		// b имеет тип byte, uint8
	}

	// руны
	for i, r := range s {
		fmt.Println("i, r", i, r)
		// i может перепрыгивать значения 1,2,4,6,9...
		// r - имеет тип rune, int32
	}

	for i, r := range s {
		fmt.Println("i, r", i, r, string(r))
		// i может перепрыгивать значения 1,2,4,6,9...
		// r - имеет тип rune, int32
	}

	// работа с каналами
	// func Example12() {
	// 	ch := make(chan string)
	// 	go pushToChannel(ch)
	// 	for val := range ch {
	// 		fmt.Println(val)
	// } }
	// func pushToChannel(ch chan<- string) {
	// 	ch <- "a"
	// 	ch <- "b"
	// 	ch <- "c"
	// close(ch) }

}
