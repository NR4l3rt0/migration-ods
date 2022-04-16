package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {

	project := ""
	baseUrl := ""
	repo := ""
	url := fmt.Sprintf("%s/%s/%s", baseUrl, project, repo)

	user := ""
	token := ""

	fmt.Printf("Cloning repository: %q\n\n", repo)

	_, err := git.PlainClone(repo, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: user,
			Password: token,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	//CheckIfError(err)
}
