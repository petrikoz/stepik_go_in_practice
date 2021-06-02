// https://stepik.org/lesson/526868/step/8?unit=519587

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...
	var abbr []rune
	for _, word := range strings.Fields(phrase) {
		for _, rune := range word {
			if unicode.IsLetter(rune) {
				abbr = append(abbr, unicode.ToUpper(rune))
			}
			break
		}
	}

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}

// ┌─────────────────────────────┐
// │ тесты для запуска в консоли │
// └─────────────────────────────┘

// echo "Today I learned" | go run step_8.go
// echo "Высшее учебное заведение" | go run step_8.go
// echo "Кот обладает талантом" | go run step_8.go
// echo "Ар 2 Ди #2" | go run step_8.go
// echo "Анна-Мария Волхонская" | go run step_8.go
