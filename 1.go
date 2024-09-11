package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var a, op, b string
	fmt.Scan(&a, &op, &b)

	var rim bool = false

	// проверка на целочисленность, значение (0-10), b != 0 при делении
	if (a != "10" && len(a) > 1) || (a == "0") || (b == "0"){
		panic("только целые числа от 1 до 10")
	}

	if (a[0] == 'I' || a[0] == 'V' || a[0] == 'X') && (b[0] == 'I' || b[0] == 'V' || b[0] == 'X') {
		rim = true

		switch a {
		case "I":
			a = "1"
		case "II":
			a = "2"
		case "III":
			a = "3"
		case "IV":
			a = "4"
		case "V":
			a = "5"
		case "VI":
			a = "6"
		case "VII":
			a = "7"
		case "VIII":
			a = "8"
		case "IX":
			a = "9"
		case "X":
			a = "10"
		}
		switch b {
		case "I":
			b = "1"
		case "II":
			b = "2"
		case "III":
			b = "3"
		case "IV":
			b = "4"
		case "V":
			b = "5"
		case "VI":
			b = "6"
		case "VII":
			b = "7"
		case "VIII":
			b = "8"
		case "IX":
			b = "9"
		case "X":
			b = "10"
		}
	}

	// преобразование строк в числа
	A, err := strconv.Atoi(a)

	if err != nil {
		log.Fatal(err)
	}

	B, err := strconv.Atoi(b)

	if err != nil {
		log.Fatal(err)
	}

	// римские числа не могут быть отрицательными
	if (rim == true) && (A <= B) && (op == "-") {
		panic("Римские числа не могут быть отрицательными!")
	}

	var C int

	switch op {
	case "+":
		C = A + B
	case "-":
		C = A - B
	case "*":
		C = A * B
	case "/":
		if B == 0 {
			panic("На ноль делить нельзя!")
		}
		C = A / B
	default:
		panic("Неправильный оператор!")
	}

	if rim == true {
		switch strconv.Itoa(C) {
		case "1":
			fmt.Println("I")
		case "2":
			fmt.Println("II")
		case "3":
			fmt.Println("III")
		case "4":
			fmt.Println("IV")
		case "5":
			fmt.Println("V")
		case "6":
			fmt.Println("VI")
		case "7":
			fmt.Println("VII")
		case "8":
			fmt.Println("VIII")
		case "9":
			fmt.Println("IX")
		case "10":
			fmt.Println("X")
		case "11":
			fmt.Println("XI")
		case "12":
			fmt.Println("XII")
		case "13":
			fmt.Println("XIII")
		case "14":
			fmt.Println("XIV")
		case "15":
			fmt.Println("XV")
		case "16":
			fmt.Println("XVI")
		case "17":
			fmt.Println("XVII")
		case "18":
			fmt.Println("XVIII")
		case "19":
			fmt.Println("XIX")
		case "20":
			fmt.Println("XX")
		default:
			panic("Паника!")
		}
	} else {
		fmt.Println(C)
	}

}
