package main

import (
	"fmt"
	"math"
)

var (
	action    uint32
	action1   uint32
	question  uint32
	question1 uint32
	number1   uint64
	number2   uint64
)

func main() {

	for {
		fmt.Println("Привет, хотите начать?")
		fmt.Println("Выберите: 1 - да ; 2 - нет")
		fmt.Scan(&question)

		if question == 1 {
			for {
				fmt.Println("Выберите действие:\n0-(завершить); 1-(суммирование); 2-(вычитание); 3-(умножение); 4-(деление); 5-(следущие действия=>) ")
				fmt.Scan(&action)

				if action != 0 && action != 5 {
					first_act()
				} else if action == 5 {
					second_act()
				} else if action == 0 {
					fmt.Println("Завершение работы.\nДо свидания!")
					return
				}

				fmt.Println("Хотите продолжить?")
				fmt.Println("Выберите: 1 - да ; 2 - нет")
				fmt.Scan(&question1)

				if question1 == 2 {
					fmt.Println("До свидания!")
					return
				}
			}
		} else {
			fmt.Println("До свидания!")
			break
		}
	}
}

func first_act() {

	fmt.Println("Введите первое число:")
	fmt.Scan(&number1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&number2)

	if action == 1 {
		fmt.Println("Ответ:", number1+number2)
	} else if action == 2 {
		fmt.Println("Ответ:", number1-number2)
	} else if action == 3 {
		fmt.Println("Ответ:", number1*number2)
	} else if action == 4 {
		if number2 != 0 {
			fmt.Println("Ответ:", number1/number2)
		} else {
			fmt.Println("Деление на ноль невозможно!")
		}
	}
}

func second_act() {

	fmt.Println("Ввыберите действие:\n0-(вернутся назад); 1-(возведение в степень); 2-(вычислить корень)")
	fmt.Scan(&action1)
	if action1 == 1 {
		fmt.Println("Введите число, которое хотите возвести в x степень:")
		fmt.Scan(&number1)
		fmt.Println("Введите степень:")
		fmt.Scan(&number2)
		fmt.Println(math.Pow(float64(number1), float64(number2)))
	} else if action1 == 2 {
		fmt.Println("Введите число, из которого хотите извлечь корень:")
		fmt.Scan(&number1)
		fmt.Println(math.Sqrt(float64(number1)))
	}
}
