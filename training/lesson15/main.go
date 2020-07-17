package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

//go 实现tree命令
func tree(dirname string, dep int) (err error) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return fmt.Errorf("Error: %s\n", err)
	}
	if dep == 1 {
		fmt.Printf("|----%s\n", filepath.Base(dirname))
	}
	sep := string(os.PathSeparator)
	for _, file := range files {
		fmt.Printf("|")
		for i := 0; i < dep; i++ {
			fmt.Printf("    |")
		}
		if file.IsDir() {
			fmt.Printf("----%s\n", file.Name())
			tree(dirname+sep+file.Name(), dep+1)
			continue
		}
		fmt.Printf("----%s\n", file.Name())
	}
	return
}

func main() {
	app := cli.NewApp()
	app.Name = "tree"
	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		return tree(dir, 1)
	}
	app.Run(os.Args)
}
