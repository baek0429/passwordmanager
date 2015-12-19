package client

import (
	"testing"
)

func _TestNodeBasic(t *testing.T) {
	t.Log(parseCommands([]string{"", "create"}))
	t.Log(parseCommands([]string{"", "create1"}))
	t.Log(parseCommands([]string{"", "create", "hello", "-h"}))
	t.Log(parseCommands([]string{"", "create", "hello", "world", "-h"}))
}

func TestRun(t *testing.T) {
	c := parseCommands([]string{"", "create", "naver", "baek0429", "baek12345"})
	c.run()
}
