package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// Содержится ли подстрока в строке
		strings.Contains("test", "es"),
		// результат: true

		// Кол-во подстрок в строке
		strings.Count("test", "t"),
		// результат: 2

		// Начинается ли строка с префикса
		strings.HasPrefix("test", "te"),
		// результат: true

		// Заканчивается ли строка суффиксом
		strings.HasSuffix("test", "st"),
		// результат: true

		// Возвращает начальный индекс подстроки в строке, а при отсутствии вхождения возвращает -1
		strings.Index("test", "s"),
		// результат: 1

		// объединяет массив строк через символ
		strings.Join([]string{"hello", "world"}, "-"),
		// результат: "hello-world"

		// Повторяет строку n раз подряд
		strings.Repeat("a", 5),
		// результат: "aaaaa"

		// Функция Replace заменяет любое вхождение old в вашей строке на new
		// Если значение n равно -1, то будут заменены все вхождения.
		// Общий вид: func Replace(s, old, new string, n int) string
		// Пример:
		strings.Replace("blanotblanot", "not", "***", -1),
		// результат: "bla***bla***"

		// Разбивает строку согласно разделителю
		strings.Split("a-b-c-d-e", "-"),
		// результат: []string{"a","b","c","d","e"}

		// Возвращает строку c нижним регистром
		strings.ToLower("TEST"),
		// результат: "test"

		// Возвращает строку c верхним регистром
		strings.ToUpper("test"),
		// результат: "TEST"

		// Возвращает строку с вырезанным набором
		strings.Trim("tetstet", "te"),
		// результат: s
	)
}
