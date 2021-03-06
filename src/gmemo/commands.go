package main

import (
	"bufio"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
)

var Commands = []cli.Command{
	commandAdd,
	commandShow,
	commandDelete,
}

var commandAdd = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "add a memo to list",
	Action:  add,
}

const filePath = "./gmemo"

func add(c *cli.Context) {
	res := Exists(filePath)

	if res {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {

		}
		defer file.Close()
		file.Write(([]byte)(c.Args()[0]))
		file.Write(([]byte)("\n"))
	} else {
		file, err := os.Create(filePath)
		if err != nil {

		}
		defer file.Close()
		file.Write(([]byte)(c.Args()[0]))
		file.Write(([]byte)("\n"))
	}

	//	return nil
}

var commandShow = cli.Command{
	Name:    "show",
	Aliases: []string{"s"},
	Usage:   "show memo's",
	Action:  show,
}

func show(c *cli.Context) {
	println("your memo is ")
	file, err := os.Open(filePath)
	if err != nil {

	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		//println(sc.Text())
		fmt.Printf("%d. %s\n", i, sc.Text())
	}

	//	return nil
}

var commandDelete = cli.Command{
	Name:    "delete",
	Aliases: []string{"d"},
	Usage:   "delete a memo from list",
	Action:  _delete,
}

func _delete(c *cli.Context) {
	//	println("show memo's")
	file, err := os.Open(filePath)
	if err != nil {

	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	list := map[string]string{}
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			println("error")
			break
		}
		list[strconv.Itoa(i)] = sc.Text()
	}
	delete(list, c.Args()[0])

	file, err = os.Create(filePath)
	for _, v := range list {
		file.Write(([]byte)(v))
		file.Write(([]byte)("\n"))
	}
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
