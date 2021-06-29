package main

import (
	"os"
	"fmt"
)

func main() {
	dir, err := os.Open(".")

	if err != nil {
		//Handle the error
		return
	}

	// Close the file
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		//Handle the error
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}