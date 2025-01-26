package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/fatih/color"
)

var (
	greenPrintf = color.New(color.FgGreen).PrintfFunc()
	redPrintf = color.New(color.FgRed).PrintfFunc()
	bluePrintf = color.New(color.FgBlue).PrintfFunc()
	yellowPrintf = color.New(color.FgHiYellow).PrintfFunc()
	greyPrintf = color.New(color.FgHiBlack).PrintfFunc()
	hiRedPrintf = color.New(color.FgHiRed).PrintfFunc()
)

//=====================================================
func main() {
	setTerminalTitle("Конвертор")
	for {
		clearScreen()
		inputNum, inputBase := inputUser()
		convertBase(inputNum, inputBase)
		// fmt.Println(strconv.FormatInt(int64(inputNum), inputBase))
		pause()
	}
}
//=====================================================

// перевод числа в указанную систему счисления
func convertBase(n, base int) {
	hexMap := map[int]string{10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F"}
	
	if base == 16 {
		yellowPrintf("\nВ шестнадцатеричной системе счисления используются\nчисла от 0 до 9 и буквы от A до F\n")
		yellowPrintf("10 = A, 11 = B, 12 = C, 13 = D, 14 = E, 15 = F\n")
	}

	yellowPrintf("\nВыполняем преобразование:\n")
	result := []string{}
	for n > 0 {
		remainder := n % base
		if remainder >= 10 && remainder <= 15 {
			yellowPrintf("%v / %v = %v остаток %v ≡ %v \n", n, base, n / base, n % base, hexMap[remainder])
			result = append(result, []string{hexMap[remainder]}...)
		} else {
			yellowPrintf("%v / %v = %v остаток %v\n", n, base, n / base, n % base)
			result = append(result, []string{strconv.Itoa(remainder)}...)
		}
		
		n = n / base             
	}
	yellowPrintf("\nПереписываем остатки в обратном порядке:\n")
	slices.Reverse(result)
	greenPrintf("Результат = ")
	hiRedPrintf("%s\n", strings.Join(result, ""))
}

// Ввод пользователя
func inputUser() (int, int) {
	newScanner := bufio.NewScanner(os.Stdin)
	greenPrintf("Введите число и систему счисления (например, 29 16): ")
	
	for newScanner.Scan() {
		input := strings.Split(strings.TrimSpace(newScanner.Text()), " ")
		if len(input) == 2 {
			numbersStr, baseStr := input[0], input[1]
			if checkNumbers(numbersStr) && checkNumbers(baseStr) {
				numbers, _ := strconv.Atoi(numbersStr)
				base, _ := strconv.Atoi(baseStr)
				if base >= 2 && base <= 16 {
					return numbers, base
				}
			}
		}
		bluePrintf("Некорректный ввод. Попробуйте снова.\n")
		greenPrintf("Введите число и систему счисления (например, 29 16): ")
	}

	redPrintf("Ошибка ввода\n")
	return -1, -1
}

// проверяем является ли строка числом
func checkNumbers(n string) bool {
	for _, r := range n {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Очистка экрана и перевод курсора верхний левый угол
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Пауза до нажатии клавиши
func pause() {
	greyPrintf("Продолжить Enter...⏎")
	bufio.NewScanner(os.Stdin).Scan()
}

// Устанавливаем заголовок терминала
func setTerminalTitle(title string) {
	// ANSI код для изменения заголовка терминала
	fmt.Printf("\033]0;%s\007", title)
}
