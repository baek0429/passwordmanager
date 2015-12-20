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
	action := c.Type
	switch action {
	// 0			1		2		3	  4
	// [path.exe] [command] [arg1] [arg2] [arg3]
	case 1: //create
		if len(c.Arguments) < 3 {
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

	case 3: //delete
		if len(c.Arguments) < 1 {
			log.Fatal(errors.New("More arguments needed!"))
			err := deleteLineFromCompanyName(c.Arguments[1])
			if err != nil {
				log.Fatal(err)
			}
		}
	case 2: //show
		if len(c.Arguments) < 1 {
			log.Fatal(errors.New("More arguments needed!"))
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
			fmt.Println(v)
		}
		return
	case 9:
		fmt.Println(ABOUT_MSG, CREDIT)
		return
	default:
		fmt.Println(ABOUT_MSG, ERR_MSG)
	}
	return
}

func deleteLineFromCompanyName(cname string) error {
	input, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return err
	}
	re := regexp.MustCompile("(?m)^.*" + cname + ".*$[\r\n]+")
	res := re.ReplaceAllString(string(input), "")
	ioutil.WriteFile(FILENAME, []byte(res), 0666)
	return nil
}

func searchWithCompanyName(cname string, strs []string) []string {
	var result []string
	for _, str := range strs {
		eachColumn := strings.Split(str, " ")
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
	str = str + company + blank + encrypted.Key + blank + encrypted.Value + time.Now().String() + newline
	_, err = f.Write([]byte(str))
	return err
}

func checkIfCompanyNameExists(str string) bool {
	data, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return false
	}
	re := regexp.MustCompile("(?m)^.*" + str + ".*$[\r\n]+")
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
