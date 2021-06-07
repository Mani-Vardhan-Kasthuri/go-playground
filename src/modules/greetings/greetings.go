package greetings

import "fmt"

func Greet(name string) string {
	msg := fmt.Sprintf("Hello!! %v Welcome", name)
	return msg
}
