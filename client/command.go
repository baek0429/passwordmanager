package client

import "log"

const (
	COMMAND  = 0
	ARGUMENT = 1
	FLAG     = 2
)

type CommandNode struct {
	NextCommand *CommandNode
	Arguments   []*CommandNode
	Flags       []*CommandNode
}

func (c *CommandNode) AddCommandNode(node *CommandNode, cst int) *CommandNode {
	switch {
	default:
		log.Println(node, ", with mode,", cst)
	case cst == COMMAND:
		c.NextCommand = node
		log.Println("added node and returns the next node")
		return c.NextCommand
	case cst == ARGUMENT:
		c.Arguments = append(c.Arguments, node)
		log.Println(c.Arguments)
		return c
	case cst == FLAG:
		c.Arguments = append(c.Flags, node)
		log.Println(c.Arguments)
		return c
	}
	return c
}
