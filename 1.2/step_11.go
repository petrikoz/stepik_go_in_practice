// https://stepik.org/lesson/526867/step/11?unit=519586

package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	lang := "Unknown"
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	}

	fmt.Println(lang)
}

// ┌─────────────────────────────┐
// │ тесты для запуска в консоли │
// └─────────────────────────────┘

// echo "en" | go run step_11.go
// echo "fr" | go run step_11.go
// echo "ru" | go run step_11.go
// echo "rus" | go run step_11.go
// echo "иначе" | go run step_11.go
