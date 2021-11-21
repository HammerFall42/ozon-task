package main

import (
	"fmt"
	"os"

	"ozon-task/app"
)

func main() {
	mode := os.Args[1]
	if err := app.Run(mode[0]); err != nil {
		fmt.Print(err.Error())
	}
}