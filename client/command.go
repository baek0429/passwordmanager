package client

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	FILENAME = "ENCRYPTED"

	C_CREATE = 1
	C_SHOW   = 2
	C_DELETE = 3
	C_ABOUT  = 9

	C_CREATE_IST = ""
	C_SHOW_IST   = ""
	C_DELETE_IST = ""
	C_ABOUT_IST  = ""

	ERR_MSG   = "Error, check your command by printing '[command] [-h]'"
	ABOUT_MSG = ""
	CREDIT    = "Copyright 2015 Chungseok Baek csbaek0429@gmail.com"
)

type Command struct {
	Type        int
	Instruction string
	Arguments   []string
	Flags       []string
}

func intialize(strs []string) {
	c := parseCommands(strs)
	c.run()
}

func (c *Command) run() {
	f, err := os.OpenFile(FILENAME, os.O_APPEND, 0600) // open file
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // defer close
	action := c.Type
	switch action {
	// 0			1		2		3	  4
	// [path.exe] [command] [arg1] [arg2] [arg3]
	case 1: //create
		if len(c.Arguments) < 3 {
			log.Fatal(errors.New("More arguments needed!"))
		}
	case 3: //delete
		if len(c.Arguments) < 1 {
			log.Fatal(errors.New("More arguments needed!"))
		}
	case 2: //show
		if len(c.Arguments) < 1 {
			log.Fatal(errors.New("More arguments needed!"))
		}
		// read data from file
		// search company
		return
	case 9:
		fmt.Println(ABOUT_MSG, CREDIT)
		return
	default:
		fmt.Println(ABOUT_MSG, ERR_MSG)
	}
	return
}

func parseCommands(strs []string) *Command {
	// prepare empty command
	c := &Command{}

	// parse first command and assign it to Type
	i := 1
	expectedCommand := ""
	if len(strs) > 1 {
		expectedCommand = strs[i]
	} else if len(strs) == 1 {
		fmt.Println(ABOUT_MSG)
		return c
	}
	switch expectedCommand {
	case "create":
		c.Type = C_CREATE
		c.Instruction = C_CREATE_IST
	case "show":
		c.Type = C_SHOW
		c.Instruction = C_SHOW_IST
	case "delete":
		c.Type = C_DELETE
		c.Instruction = C_DELETE_IST
	case "about":
		c.Type = C_ABOUT
		c.Instruction = C_ABOUT_IST
	default:
		fmt.Println(ERR_MSG)
		return c
	}

	// parse flags and arguments
	for i := 2; i < len(strs); i++ {
		if []rune(strs[i])[0] == '-' {
			c.addFlag(strs[i])
		} else {
			c.addArgument(strs[i])
		}
	}
	return c
}

// add Arguments
func (c *Command) addArgument(strs ...string) {
	c.Arguments = append(c.Arguments, strs...)
}

// add flags
func (c *Command) addFlag(strs ...string) {
	c.Flags = append(c.Flags, strs...)
}
