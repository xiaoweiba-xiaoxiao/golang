package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func testCli() {
	app := cli.NewApp()
	app.Name = "hello"
	app.Usage = "fight the lineless"
	app.Action = func(c *cli.Context) error {
		fmt.Println("hello friends")
		return nil
	}
	app.Run(os.Args)
}

func testCli2() {
	app := cli.NewApp()
	app.Name = "xiaoxiao"
	app.Usage = "fight he loneliness"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("hello %v\n", c.Args().Get(0))
		return nil
	}
	app.Run(os.Args)
}

func testCli3() {
	var language string
	var recusive bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang,l",
			Value:       "Enlish",
			Usage:       "language for the greeting",
			Destination: &language,
		},
		cli.BoolTFlag{
			Name:        "recusive,g",
			Usage:       "recusive for the greeting",
			Destination: &recusive,
		},
	}
	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			fmt.Printf("%v", c.Args())
			cmd = c.Args()[0]
			fmt.Printf("cmd is %v\n", cmd)
		}
		fmt.Printf("language is %v\n", language)
		fmt.Printf("recusive is %v\n", recusive)
		return nil
	}
	app.Run(os.Args)
}

func main() {
	// testCli()
	// testCli2()
	// testCli3()
	var language string
	var recusive bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang,l",
			Value:       "Enlish",
			Usage:       "language for the greeting",
			Destination: &language,
		},
		cli.BoolTFlag{
			Name:        "recusive,g",
			Usage:       "recusive for the greeting",
			Destination: &recusive,
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("language is %v\n", language)
		fmt.Printf("recusive is %v\n", recusive)
		return nil
	}
	app.Run(os.Args)
}
