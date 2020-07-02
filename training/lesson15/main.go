package main

import (
	_ "flag"
	"fmt"
	"io/ioutil"
	"path"
)

func tree(dir string) {
	fmt.Printf("|----%s", path.Clean())
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if len(files) == 0 {
		return
	}
	for _, file := range files {
		fmt.Printf("%s\n", file.Name())
		if file.IsDir() {
			tree(path.Join(dir, file.Name()))
		}
	}
}

func main() {
	fmt.Println("|----/home/golang/test")
	tree("/home/golang/test")

}
