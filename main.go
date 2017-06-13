package main

import (
	"container/list"
	"fmt"
	"github.com/jharshman/gosh/config"
	"github.com/jharshman/gosh/history"
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
	var conf config.Rc
	conf = config.Init(rcFile)

	// history use container/list instead of array / slice
	hList := list.New()
	history.Init(&hList)

}
