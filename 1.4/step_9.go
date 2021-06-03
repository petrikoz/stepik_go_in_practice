// https://stepik.org/lesson/526869/step/9?unit=519588

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func shuffle(nums []int) {
	// Функция `Seed()` инициализирует генератор случайных чисел
	// здесь мы используем константу `42`, чтобы программу
	// можно было проверить тестами.
	// В реальных задачах не используйте константы!
	// Используйте, например, время в наносекундах:
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(42)
	// перетасуйте `nums` с помощью `rand.Shuffle()``
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

// ┌─────────────────────────────┐
// │ тесты для запуска в консоли │
// └─────────────────────────────┘

// echo "1 2 3 4 5 6" | go run step_9.go
