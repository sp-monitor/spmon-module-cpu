package main

import (
	"fmt"
	"runtime"
)


func main() {
	os := runtime.GOOS
	if os == "linux" {
		linux()
	} else if os == "windows" {
		windows()
	} else {
		fmt.Println("OS not supported")
	}
}

func linux() {
	fmt.Println("Linux")
}

func windows() {
	fmt.Println("Windows")
}



