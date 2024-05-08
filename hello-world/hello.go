package main

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, " + name)
}

func main() {
	fmt.Println(Hello("Chris"))
}
