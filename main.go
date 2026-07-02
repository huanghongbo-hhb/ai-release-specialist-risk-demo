package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, Greeting("zadig"))
	})

	log.Println("ai-release-specialist-risk-demo listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
