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

	fmt.Println("Ввыберите действие:\n0-(вернутся назад); 1-(возведение в степень); 2-(вычислить корень); 3-(квадратное уранение)")
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
	} else if action1 == 3 {
		square_leveling()
	}
}

func square_leveling() {
	var (
		square_number           uint64
		square_number1          uint64
		square_number2          uint64
		square_action           uint64
		square_action1          uint64
		answer_x1               uint64
		answer_x2               uint64
		discriminant            uint64
		square_counter_question uint64
	)

	fmt.Println("Уровнение:\n a * x^2 _ b * x _ c = 0")

	fmt.Println("Введите число a:")
	fmt.Scan(&square_number)
	fmt.Println("Введите число b:")
	fmt.Scan(&square_number1)
	fmt.Println("Введите число c:")
	fmt.Scan(&square_number2)
	fmt.Println("Введите первый знак(между a и b):")
	fmt.Scan(&square_action)
	fmt.Println("Введите первый знак(между b и c):")
	fmt.Scan(&square_action1)

	fmt.Println("Ваше уравнение выглядит так:")

	fmt.Println(square_number, "* x^2", square_action, square_number1, "* x", square_action1, square_number2, "= 0")

	fmt.Println("Выберите: 1 - да ; 2 - нет")

	fmt.Scan(&square_counter_question)

	if square_counter_question == 1 {
		discriminant = uint64(math.Pow(float64(square_number1), 2)) - 4*square_number*square_number2

		if discriminant > 0 {
			fmt.Println("Дискриминант =", discriminant)
			answer_x1 = ((-square_number1) - uint64(math.Sqrt(float64(discriminant)))) / (2 * square_number)
			answer_x2 = ((-square_number1) + uint64(math.Sqrt(float64(discriminant)))) / (2 * square_number)
			fmt.Println("Ответ:")
			fmt.Println("x1 =", answer_x1)
			fmt.Println("x2 =", answer_x2)
		} else if discriminant == 0 {
			fmt.Println("Дискриминант =", discriminant)
			answer_x1 = (-square_number1) / (2 * square_number)
			fmt.Println("Ответ:")
			fmt.Println("x =", answer_x1)
		} else {
			fmt.Println("Так как дискриминант меньше 0, то корней нет!")
		}
	}

}