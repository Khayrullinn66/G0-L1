package main

import (
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

1.Не рекомендуется использовать глобальные переменные, таким образом
определение justString можно сделать и в функции(можно передать в функцию по указателю)
2.Также переменная v не понятно чего использована, после создание слайса
из нее остальные 1024-100=924 элемента являются неиспользуемыми, получается от
результата функции можно сразу взять первые 100 элементов.

*/

func createHuge(num int) string {
	str := []rune("12345678990qwertyuiopasdfghjkl")
	resultString := ""
	for i := 0; i < num; i++ {
		resultString += string(str[rand.Intn(len(str))])
	}
	return resultString
}
func SomeFunc() string {
	return createHuge(1 << 10)[:100]
}

func main() {
	fmt.Println(SomeFunc())
}
