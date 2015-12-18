package main

import (
	client "./client"
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
		action: func() error { return errors.New("about execute") },
	},
	{
		key:    "show",
		value:  "[company key, ex:naver]",
		action: func() error { return errors.New(" error") },
	},
	{
		key:    "delete",
		value:  "[company key, ex:naver]",
		action: func() error { return errors.New("error") },
	},
	{
		key:   "create",
		value: "[company key, ex:naver] [id, ex: foo@naver.com] [password]",
		action: func() error {

			decrypted := client.DecryptedPassword{Key: "hello", Value: "world"}

			f, err := os.OpenFile("ENCRYPTED", os.O_CREATE, 0660)
			if err != nil {
				return err
			}

			f.Write([]byte(decrypted.Key))
			return errors.New("error")
		},
	},
	{
		key:    "replace",
		value:  "[comapany key, ex:naver] [id, ex: foo@naver.com] [password]",
		action: func() error { return errors.New("replace called") },
	},
}

var optionCommandTrees = []*OptionalCommand{
	{
		key:    "-h",
		value:  "Display how to use the command",
		action: func() error { return errors.New("error") },
	},
}

func main() {
	// retrieve optional commands
	optionalCommands, err := lookupOptionalCommand(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	// retrieve command
	command, err := lookupCommand(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// execute optional commands
	for _, v := range optionalCommands {
		err = v.Execute()
		if err != nil {
			fmt.Println(err)
			command.String()
		}
	}

	err = command.Execute()
	if err != nil {
		fmt.Println(err)
	}

	var a string
	fmt.Scanf("%s", &a)
}

func lookupCommand(str string) (*Command, error) {
	for _, v := range commandTree {
		if v.key == str {
			return v, nil
		}
	}
	return new(Command), errors.New("Command error")
}

func lookupOptionalCommand(strs []string) ([]OptionalCommand, error) {
	optionalCommands := []OptionalCommand{}
	for _, o := range strs {
		for _, v := range optionCommandTrees {
			if v.key == o {
				optionalCommands = append(optionalCommands, *v)
				break
			}
			// return optionalCommands, errors.New("OptionalCommandError")
		}
	}
	return optionalCommands, nil
}

/**
- interface CommandInterface
- struct Command
- struct OptionalCommand
**/

type CommandInterface interface {
	String()
	Execute() error
	Clean()
}

type OptionalCommand struct {
	key    string
	value  string
	action func() error
}

func (o *OptionalCommand) String() {
	fmt.Println(o.key, o.value)
}

func (o *OptionalCommand) Execute() error {
	return o.action()
}

func (o *OptionalCommand) Clean() {
	*o = OptionalCommand{}
}

type Command struct {
	key    string
	value  string
	action func() error
}

func (c *Command) Clean() {
	*c = Command{}
}

func (c *Command) String() {
	fmt.Println(c.key, c.value)
}

func (c *Command) Execute() error {
	return c.action()
}
