package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/jharshman/gosh/config"
	"github.com/jharshman/gosh/history"
	"os"
	"strings"
)

/*
  TODO:
  - load config files
  - read history file
  - initialize history
  - loop and process input
  - run cleanup if any
*/

func main() {

	// read .goshrc file
	conf := config.Init()

	// debug print
	fmt.Println(conf.HistorySize)
	fmt.Println(conf.HistoryFileSize)

	// history use container/list instead of array / slice
	hList := list.New()
	history.Init(&hList)

	// debug print history
	for e := hList.Front(); e.Next() != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// proof of concept write for protobuf
	history.WriteHistory(&hList)

	// main loop
	consoleReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("gosh> ")
		consoleInput, _ := consoleReader.ReadString('\n')
		println("you wrote " + consoleInput)

		if strings.Compare(consoleInput, "exit\n") == 0 {
			break
		}

	}

}
