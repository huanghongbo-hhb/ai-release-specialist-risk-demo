package main

import "fmt"

func main() {
	fmt.Println(Greeting("zadig"))
}

func Greeting(name string) string {
	if name == "" {
		return "hello, anonymous"
	}
	return fmt.Sprintf("hello, %s", name)
}

func RiskyDiscount(price int, discount int) int {
	return price - discount
}

func IsSafeToRelease(scanBugs int, testPassRate float64) bool {
	return scanBugs <= 5 && testPassRate >= 90
}
