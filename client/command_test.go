package client

import (
	"testing"
)

func _TestNodeBasic(t *testing.T) {
	t.Log(ParseCommands([]string{"", "create"}))
	t.Log(ParseCommands([]string{"", "create1"}))
	t.Log(ParseCommands([]string{"", "create", "hello", "-h"}))
	t.Log(ParseCommands([]string{"", "create", "hello", "world", "-h"}))
}

func _TestRun(t *testing.T) {
	c := ParseCommands([]string{"", "create", "daum1", "baek0429", "baek12345"})
	c.Run()
}

func _TestLastPushDate(t *testing.T) {
	lastPushDateUpdateNow()
	ti, err := getLastPushDate()
	if err != nil {
		t.Log(err)
	}
	t.Log(ti.String())
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
