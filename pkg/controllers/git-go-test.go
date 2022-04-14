package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	fmt.Println("git clone https://github.com/NR4l3rt0/migration-ods")

	// Clone the given repository to the given directory
	//Info("git clone https://github.com/NR4l3rt0/migration-ods")

	_, err := git.PlainClone("migration-ods", false, &git.CloneOptions{
		URL:      "https://github.com/NR4l3rt0/migration-ods",
		Progress: os.Stdout,
	})
	if err != nil {
		panic(err)
	}
	//CheckIfError(err)
}
