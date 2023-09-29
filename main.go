package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func dealWithErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	os := runtime.GOOS
	if os == "linux" {
		linux()
	} else if os == "windows" {
		windows()
	} else {
		fmt.Println("OS not supported")
	}
	crossPlatform()
}

func linux() {
	fmt.Println("Linux")
}

func windows() {
	fmt.Println("Windows")
}

func crossPlatform() {
	for {
		percentage, err := cpu.Percent(0, true)
		dealWithErr(err)
		for idx, percent := range percentage {
			fmt.Printf("Core %d: %f\n", idx, percent)
		}
		time.Sleep(1 * time.Second)
	}
}
