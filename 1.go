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

	// карта для перевода римских в арабские

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

	// проверка на целочисленность, значение (1-10), b != 0 при делении
	if (a != "10" && len(a) > 1) || (a == "0") || (b == "0") && (!rim) {
		panic("только целые числа от 1 до 10")
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
	if (rim) && (A <= B) && (op == "-") {
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
		if (rim) && (A < B) {
			panic("паника")
		}
		C = A / B
	default:
		panic("Неправильный оператор!")
	}

	// карта для перевода арабских в римские
	arab_to_rome := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXX",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}

	if rim {
		fmt.Println(arab_to_rome[C])
	} else {
		fmt.Println(C)
	}

}
