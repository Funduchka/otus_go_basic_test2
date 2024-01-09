// package main

// import "fmt"

// // Функция main как точка входа
// var a = 1 // <- уровень пакета

// func init() {
// 	fmt.Println("init", a)
// }

// func main() {
// 	fmt.Println("Functions lesson")
// 	fmt.Println("1: ", a)
// 	a := 2 // <-- уровень блока функции
// 	fmt.Println("2: ", a)

// 	{
// 		a := 3 // <-- уровень пустого блока
// 		fmt.Println("3: ", a)
// 	}
// 	fmt.Println("4: ", a) // <-- ???
// 	f()
// 	d(a)
// 	block()
// }

// func f() {
// 	fmt.Println("5: ", a) // <-- ???
// }

// func d(a int) {
// 	fmt.Println("d: ", a) // <-- ???
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

// Сколько раз объявляется переменная х?

// package main

// import "fmt"

// func f(x int) {
// 	fmt.Println("f", x)
// 	x = 100
// 	fmt.Println("100", x)
// 	for x := 0; x < 10; x++ {
// 		fmt.Println(x)
// 	}
// }

// var x int

// func main() {
// 	var x = 200
// 	fmt.Println("200", x)
// 	f(x)
// 	fmt.Println("200", x)
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

// package main

// import "fmt"

// func sum(nums ...int) {
// 	fmt.Print(nums, " ")
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }

// func main() {
// 	sum(5, 7)
// 	sum(3, 2, 1)
// 	nums := []int{1, 2, 3, 4}
// 	sum(nums...)
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	people := []string{"Alice", "Bob", "Dave"}
// 	sort.Slice(people, func(i, j int) bool {
// 		return len(people[i]) < len(people[j])
// 	})
// 	fmt.Println(people)
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func intSeq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func main() {
// 	nextInt := intSeq()
// 	fmt.Println(nextInt()) // <-- ?
// 	fmt.Println(nextInt()) // <-- ?
// 	fmt.Println(nextInt()) // <-- ?
// 	newInts := intSeq()
// 	fmt.Println(newInts()) // <-- ?
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	s := "hello"
// 	defer fmt.Println(">", s)
// 	defer func() {
// 		fmt.Println("_", s)
// 	}()

// 	s = "world"
// 	fmt.Println("!", s)
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

// РЕКУРСИЯ - возможнеость вызвать функцию саму в себе

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("!", factorial(6))
// }

// // n! = n×(n-1)! where n >0
// func factorial(num int) int {
// 	if num > 1 {
// 		return num * factorial(num-1)
// 	}
// 	return 1
// }

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------

package main

import "fmt"

type Book struct {
	pages int
}

func (b *Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

// "the same"
func Pages(b Book) int {
	return b.pages
}
func SetPages(b *Book, pages int) {
	b.pages = pages
}

func main() {
	bk := Book{
		pages: 34,
	}

	fmt.Println(bk.Pages())
	bk.SetPages(10)
	fmt.Println(bk.Pages())

	fmt.Println(Pages(bk))
	SetPages(&bk, 100)
	fmt.Println(bk.Pages())
}
