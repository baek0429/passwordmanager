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
	c := parseCommands([]string{"", "create", "daum", "baek0429", "baek12345"})
	c.run()
}

func _TestCheckingCName(t *testing.T) {
	t.Log(checkIfCompanyNameExists("naver"))
	t.Log(checkIfCompanyNameExists("naver1"))
}

func _TestDelete(t *testing.T) {
	deleteLineFromCompanyName("naver")
}

func _TestKeywordSearch(t *testing.T) {
	strs := []string{"hello baek", "hello world", "baek word", "kkkkkk world hello"}
	result := searchWithKeyword("hello", strs)
	t.Log(result)
	t.Log(len(result))
}

func _TestSearchWithCompanyName(t *testing.T) {
	strs := []string{"hello baek", "hello world", "baek word", "kkkkkk world hello"}
	result := searchWithCompanyName("hello", strs)
	t.Log(result)
	t.Log(len(result))
}
