package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/seadreamcode/sea/core"
)

func main() {
	fmt.Println("~")
	dir, _ := os.Getwd()
	err := core.Build(dir, path.Join(dir, "dest"))

	if err != nil {
		log.Fatal(err)
	}
}
