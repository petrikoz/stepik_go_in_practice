// https://stepik.org/lesson/526867/step/3?unit=519586

package main

import (
	"fmt"
	"time"
)

func main() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	// выведите исходное значение
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	fmt.Println(s, "=", d.Minutes(), "min")
}

// ┌─────────────────────────────┐
// │ тесты для запуска в консоли │
// └─────────────────────────────┘

// echo "1h30m" | go run step_3.go
// echo "300s" | go run step_3.go
// echo "10m" | go run step_3.go
