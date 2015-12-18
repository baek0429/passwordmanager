package main

import (
	// client "./client"
	"errors"
	"fmt"
	"os"
)

const (
	PROGRAM_DISCRIPTION = "---------------\nPasswordManager\nThis program encrypt/decrypte passwords.\nCredit: Chungseok Baek, csbaek0429@gmail.com\n--------------"
)

var commandTree = []*Command{
	{
		key:    "about",
		value:  PROGRAM_DISCRIPTION,
		action: func() error { return errors.New("error") },
	},
	{
		key:    "show",
		value:  "[company key, ex:naver]",
		action: func() error { return errors.New("error") },
	},
	{
		key:    "delete",
		value:  "[company key, ex:naver]",
		action: func() error { return errors.New("error") },
	},
	{
		key:    "create",
		value:  "[company key, ex:naver] [id, ex: foo@naver.com] [password]",
		action: func() error { return errors.New("error") },
	},
	{
		key:    "replace",
		value:  "[comapany key, ex:naver] [id, ex: foo@naver.com] [password]",
		action: func() error { return errors.New("error") },
	},
}

var optionCommandTree = []*OptionalCommand{
	{
		key:    "-h",
		value:  "Display how to use the command",
		action: func() error { return errors.New("error") },
	},
}

func main() {

}

type CommandInterface interface {
	String()
	Execute()
}

type OptionalCommand struct {
	key    string
	value  string
	action func() error
}

func (o *OptionalCommand) String() {
	fmt.Println(c.key, c.value)
}

func (o *OptionalCommand) Execute() {
	o.action
}

type Command struct {
	key    string
	value  string
	action func() error
}

func (c *Command) String() {
	fmt.Println(c.key, c.value)
}

func (c *Command) Execute() {
	c.action
}
