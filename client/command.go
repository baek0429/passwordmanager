package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	FILENAME       = "ENCRYPTED"
	LAST_PUSH_DATE = "LAST_PUSH"

	C_CREATE  = 1
	C_SHOW    = 2
	C_DELETE  = 3
	C_PUSH    = 4
	C_REPLACE = 5
	C_ABOUT   = 9

	C_CREATE_IST  = "[create] [company name] [id] [password]"
	C_SHOW_IST    = "[show] [company name] or [show] for all lists"
	C_DELETE_IST  = "[delete] [company name]"
	C_ABOUT_IST   = "[about]"
	C_REPLACE_IST = "[replace] [company name] [id] [password]"
	C_PUSH_IST    = "[push] [url] [id] [password]"

	F_HELP = "-h"

	ERR_MSG      = "Error, check your command by printing '[command] [-h]'"
	COMMAND_LIST = "create, delete, show"
	ABOUT_MSG    = "This program safely stores your passwords."
	CREDIT       = "Copyright 2015 Chungseok Baek csbaek0429@gmail.com"
)

type Command struct {
	Type        int
	Instruction string
	Arguments   []string
	Flags       []string
}

// entry point of command.go
// note that it has *Command pointer receiver
func (c *Command) Run() {
	action := c.Type
	if !c.flagProcess() { // flag process returns true if it needs to proceed the following action
		return
	}
	switch action {
	// 0			1		2		3	  4
	// [path.exe] [command] [arg1] [arg2] [arg3]
	case 1: //create
		create(c)
	case 3: //delete
		delete(c)
	case 2: //show
		show(c)
	case 4: // push
		push(c)
	case 5: // replace
		replace(c)
	case 9:
		fmt.Println(ABOUT_MSG, CREDIT)
		return
	default:
		fmt.Println(ABOUT_MSG, ERR_MSG, COMMAND_LIST)
	}
	return
}

func ParseCommands(strs []string) *Command {
	// prepare empty command
	c := &Command{}

	// parse first command and assign it to Type
	i := 1
	expectedCommand := ""
	if len(strs) > 1 {
		expectedCommand = strs[i]
	} else if len(strs) == 1 {
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
	case "replace":
		c.Type = C_REPLACE
		c.Instruction = C_REPLACE_IST
	case "push":
		c.Type = C_PUSH
		c.Instruction = C_PUSH_IST
	default:
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

func replace(c *Command) { // the worst case is 'delete' succeeds and 'create' fails
	delete(c)
	create(c)
	return
}

func show(c *Command) {
	if len(c.Arguments) < 1 { // argument check
		b, err := ioutil.ReadFile(FILENAME)
		if err != nil {
			fmt.Println(err)
		}
		split := strings.Split(string(b), "\n")
		for _, v := range split {
			cname := strings.Split(v, " ")
			fmt.Println(cname[0])
		}
		return
	}
	rows, err := readEncryptedDataFromFile() // read data from file
	if err != nil {
		panic(err)
	}
	result := searchWithCompanyName(c.Arguments[0], rows) // search company
	if len(result) == 0 {
		fmt.Println("none was found with that name")
		return
	}
	for _, v := range result {
		eachColumn := strings.Split(v, " ")
		encrypted := &EncryptedPassword{
			Key:   eachColumn[1],
			Value: eachColumn[2],
		}
		decrypted := encrypted.SimpleDecrypt()
		fmt.Println(eachColumn[0], decrypted.String())
	}
	return
}

func delete(c *Command) {
	if len(c.Arguments) < 1 { // check the number of arguments
		log.Fatal(errors.New("More arguments needed!"))
	}
	if !checkIfCompanyNameExists(c.Arguments[0]) { // search by company name
		fmt.Println("No company was found by that name")
		return
	}
	err := deleteLineFromCompanyName(c.Arguments[0]) // delete line by company name
	if err != nil {
		panic(err)
	}
	return
}

func create(c *Command) {
	if len(c.Arguments) < 3 { // check the number of arguments
		log.Fatal(errors.New("More arguments needed!"))
		return
	}
	if checkIfCompanyNameExists(c.Arguments[0]) { // check if company exists
		fmt.Println("Company Name Exist, use replace")
		return
	}

	var d = DecryptedPassword{ // prepare decrypted file
		Key:   c.Arguments[1],
		Value: c.Arguments[2],
	}
	var e = d.SimpleEncrypt() // encrypt
	err := writeEncryptedDataToFile(c.Arguments[0], e)
	if err != nil {
		panic(err)
	}
	return
}

// TODO: checkIfPushNeeded
func push(c *Command) {
	if checkIfPushNeeded() {
		uploadToServer()
	}
}

func (c *Command) flagProcess() bool {
	for _, v := range c.Flags {
		switch v {
		case F_HELP:
			fmt.Println(c.Instruction)
			return false
		}
	}
	return true
}

// TODO: upload ENCRYPTED
func uploadToServer() {
	fmt.Println("will be updated")
}

func checkIfPushNeeded() bool {
	lastPushDate, err := getLastPushDate()
	if err != nil {
		panic(err)
	}
	ti, err := os.Stat(FILENAME)
	if err != nil {
		panic(err)
	}
	modeTime := ti.ModTime()
	return !lastPushDate.After(modeTime)
}

func lastPushDateUpdateNow() error {
	f, err := os.OpenFile(LAST_PUSH_DATE, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := time.Now().MarshalBinary()
	if err != nil {
		return err
	}
	f.Write(b)
	return nil
}

func getLastPushDate() (*time.Time, error) {
	b, err := ioutil.ReadFile(LAST_PUSH_DATE)
	if err != nil {
		return new(time.Time), err
	}
	var ti time.Time
	err = ti.UnmarshalBinary(b)
	if err != nil {
		return new(time.Time), err
	}
	return &ti, nil
}

func deleteLineFromCompanyName(cname string) error {
	input, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return err
	}
	re := regexp.MustCompile("(?m)^.*" + cname + "\t\n\v\f\r.*$[\r\n]+")
	res := re.ReplaceAllString(string(input), "")
	err = ioutil.WriteFile(FILENAME, []byte(res), 0666)
	if err != nil {
		return err
	}
	return nil
}

func searchWithCompanyName(cname string, strs []string) []string {
	var result []string
	for _, str := range strs {
		eachColumn := strings.Split(str, " ") // searching is more 'generous' than other functions.
		if strings.Contains(eachColumn[0], cname) {
			result = append(result, str)
		}
	}
	return result
}

func searchWithKeyword(keyword string, strs []string) []string {
	var result []string
	for _, str := range strs {
		if strings.Contains(str, keyword) {
			result = append(result, str)
		}
	}
	return result
}

func writeEncryptedDataToFile(company string, encrypted *EncryptedPassword) error {
	f, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_APPEND, 0600) // open file
	if err != nil {
		return err
	}
	defer f.Close() // defer close
	str := ""
	blank := " "
	newline := "\n"
	// complete the format
	str = str + company + blank + encrypted.Key + blank + encrypted.Value + blank + time.Now().String() + newline
	_, err = f.Write([]byte(str))
	return err
}

func checkIfCompanyNameExists(str string) bool {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return false
	}
	re := regexp.MustCompile("(?m)^.*" + str + "\t\n\v\f\r.*$[\r\n]+")
	return re.Match(data)
}

func readEncryptedDataFromFile() ([]string, error) {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return nil, err
	}
	eachRow := strings.Split(string(data), "\n")
	return eachRow, nil
}

// add Arguments
func (c *Command) addArgument(strs ...string) {
	c.Arguments = append(c.Arguments, strs...)
}

// add flags
func (c *Command) addFlag(strs ...string) {
	c.Flags = append(c.Flags, strs...)
}
