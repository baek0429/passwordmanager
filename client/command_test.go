package client

import (
	"testing"
)

func TestNode(t *testing.T) {
	c := &CommandNode{}
	d := &CommandNode{}
	c.AddCommandNode(d, 0)
	c.AddCommandNode(d, 1)
	c.AddCommandNode(d, 2)
}
