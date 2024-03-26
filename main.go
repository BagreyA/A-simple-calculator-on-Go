package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Простой калькулятор")

	for {
		var firstNumber, secondNumber float64
		var operator string

		fmt.Print("Введите первое число: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		firstNumber, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
			continue
		}

		fmt.Print("Выберите операцию (+, -, *, /) или введите 'q' для выхода: ")
		operator, _ = reader.ReadString('\n')
		operator = strings.TrimSpace(operator)
		if operator == "q" {
			fmt.Println("Программа завершена.")
			break
		} else if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
			continue
		}

		fmt.Print("Введите второе число: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		secondNumber, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
			continue
		}

		var result float64
		switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber == 0 {
				fmt.Println("Ошибка: деление на ноль невозможно.")
				continue
			}
			result = firstNumber / secondNumber
		}

		fmt.Printf("Результат: %v %s %v = %v\n", firstNumber, operator, secondNumber, result)
	}
}
