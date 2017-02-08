package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gmemo"
	app.Usage = "print arguments" // ヘルプを表示する際に使用される
	app.Version = "1.0"           // ヘルプを表示する際に使用される
	//    app.Action = func(c *cli.Context) { // コマンド実行時の処理
	//        println("Number of arguments is", len(c.Args()))
	//        println("Fisrt argument is", c.Args()[0])
	//        println("Second argument is", c.Args()[1])
	//    }
	app.Commands = Commands
	app.Run(os.Args)
}
