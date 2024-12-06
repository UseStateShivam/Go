package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cmdFlags := CmdFlags{}
	flag.StringVar(&cmdFlags.Add, "add", "", "Add command, specify a title")
	flag.StringVar(&cmdFlags.Edit, "edit", "", "Edit command, edit any title eg- id:new_title")
	flag.IntVar(&cmdFlags.Del, "del", -1, "Del command, remove any todo")
	flag.IntVar(&cmdFlags.Toggle, "toggle", -1, "Toggle command, toggle the completed value")
	flag.BoolVar(&cmdFlags.List, "list", false, "List command, list all todos")
	flag.Parse()
	return &cmdFlags
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.List:
		todos.Print()
	case cf.Del!= -1:
		todos.Delete(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: invalid edit format, expected id:new_title")
			os.Exit(1)	
		}
		index, err := strconv.Atoi(parts[0])
		if err!= nil || index < 0 || index >= len(*todos) {
            fmt.Printf("Error: invalid id %d, must be between 0 and %d\n", index, len(*todos)-1)
            os.Exit(1)    
        }
		todos.Edit(index, parts[1])
	case cf.Toggle!= -1:
		todos.Toggle(cf.Toggle)
	default:
		fmt.Println("Error: no command provided")
        os.Exit(1)
	}
}

// #Working heheh :)
// my first go application
// imma start with DSA from tomorrow :)
// DSA IN GOOOOOOOO
// 😭😭
// see yall tomorrow :))))