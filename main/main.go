package main

import (
	"fmt"
	"github.com/HammerFall42/ozon-task/app"
	"os"
)

func main() {
	mode := os.Args[1]
	if len(mode) == 1 {
		if err := app.Run(mode[0]); err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("unknown parameter")
	}
}