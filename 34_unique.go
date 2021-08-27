package main

import (
	"fmt"
)

func main() {
	// Задаем строку, которую будем проверять
	str := "heraps122"
	res := isUnique(str)
	// Проверяем вернувшийся результат и выводим ответ в консоль
	if !res {
		fmt.Println("Значения в строке не уникальны")
	} else {
		fmt.Println("Значения в строке уникальны")
	}

}

// Функция проверки уникальности значения в строке
func isUnique(s string) bool {
	// Выбираем первое значение для начала сравнения
	// Берем значение руны
	f := rune(s[0])

	// Сравниваем
	for i, r := range s {
		// Потому что первое значение уже объявлено
		if i == 0 {
			continue
		}
		// Если есть совпадение
		if r == f {
			return false
		}
		// Если нет, меняем значение для сравнения на следующее
		f = r
	}
	return true
}
