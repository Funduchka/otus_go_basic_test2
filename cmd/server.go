package main // Имя текущего пакета

// Импорты других пакетов
import (
	"fmt"
	"time"
)

// func findUserBySNILS() (string, bool) {
// 	return "Moscow", false
// }

func dsc() int {
	return 13 + 20
}

func dsc2() (int, error) {
	return 13 + 20, nil
}

// func Example2(age int) {
// 	city := findUserBySNILS()
// 	if age <= 14 {
// 		fmt.Printf("Ребенок\n")
// 	} else if age == 13 && city == "Moscow" {
// 		fmt.Printf("Москвичка\n")
// 	} else {
// 		fmt.Printf("Человек %\n", age)
// 	}
// }

// func Example2(age int) {

// 	if myVar, err := myFuncCheck("fsdsad"); err != nil {
// 		fmt.Println("myFuncCheck", myVar, false)
// 	} else {
// 		fmt.Println("myFuncCheck", true)
// 	}
// }

// func Example2(age int) {
// 	myVar := "123"
// 	if myVar, err := myFuncCheck("ewqe"); err != nil {
// 		fmt.Println("myFuncCheck", myVar, false)
// 	} else {
// 		fmt.Println("myFuncCheck", myVar, true)
// 	}
// 	fmt.Println("myVar", myVar)
// }

// func Example2(age int, myBool bool) {
// 	myVar := "123"

// 	if age == 33 || !myBool {
// 		fmt.Println("Example2", myBool, myVar)
// 	} else if myVar, err := myFuncCheck("trtr"); err != nil {
// 		fmt.Println("myFuncCheck", myVar, err)
// 	} else {
// 		fmt.Println("myFuncCheck", myVar, true)
// 	}
// 	fmt.Println("myVar", myVar)
// }

// func Example2(age int, myBool bool, str string) {
// 	var myVar = "123"

// 	switch age {
// 	case 33, 25, 45, 46:
// 		fmt.Println(30)
// 		fmt.Println(25)
// 	case 14:
// 		fmt.Println(14)
// 	case 5:
// 		fmt.Println(5)
// 	default:
// 		fmt.Printf("Age = %d\n", age)
// 	}

// 	switch str {
// 	case "aaa", "bbb":
// 		fmt.Println("Not")
// 	case "ddd":
// 		fmt.Println("Yes")
// 	default:
// 		fmt.Printf("Str = %s\n", str)
// 	}
// 	fmt.Println("myVar", myVar)
// }

func Example2(age int, myBool bool, str string) {
	var myVar = "123"

	switch myAge, _ := dsc2(); myAge {
	case 25, 45:
		fmt.Println(30)
		fmt.Println(25)
	case dsc():
		fmt.Println(100)
		fallthrough
	case 13 + 20:
		fmt.Println(200)
		fallthrough
	default:
		fmt.Printf("Age = %d\n", age)
	}

	t := time.Now()
	switch {
	case t.Second() > 15:
		fmt.Printf("> 15 %v\n", t)
		fallthrough
	case t.UnixMicro() > 10000:
		fmt.Printf("> 10000 %v\n", t)
		fallthrough
	case !myBool:
		fmt.Printf("= myBool %v\n", t)
	default:
		fmt.Printf(">45 %v\n", t)
	}

	switch str {
	case "aaa", "bbb":
		fmt.Println("Not")
	case "ddd":
		fmt.Println("Yes")
	default:
		fmt.Printf("Str = %s\n", str)
	}
	fmt.Println("myVar", myVar)

	var s interface{}
	s = 12345.6789

	switch stype := s.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64")
	case float32:
		fmt.Println("float32")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", stype)
	}
}

// func myFuncCheck(v string) (string, error) {
// 	if len(v) > 0 {
// 		return "Default", nil
// 	}
// 	return "", fmt.Errorf("%s", "Here is error")
// }

// Функция main как точка входа
func main2() {
	Example2(33, true, "bbb")
	fmt.Println("Foo-1")
}
