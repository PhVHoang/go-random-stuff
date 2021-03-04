package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PhVHoang/go-random-stuff/protocol_buffers/sample"
	"github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const dpPath = "mydb.pb"

func add(text string) error {
	task := &sample.Task{
		Text: text,
		Done: false,
	}

	proto.MarshalTextString(task)
	return nil
}

func list() error {
	return nil
}
