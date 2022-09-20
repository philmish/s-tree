package kvdb

import (
	"fmt"
	"strings"

	"github.com/philmish/s-tree/pkg"
)

type command struct {
	cmd  string
	args []string
}

func (c *command) execute(t *pkg.RadixTree) string {
	switch c.cmd {
	case "ping":
		return "PONG"
	case "GET":
		if len(c.args) == 0 {
			return "ERROR Not Enough args"
		}
		return getKey(c.args[0], t)
	case "SET":
		if len(c.args) < 2 {
			return "ERROR Not enough args"
		}
		return addKeyVal(c.args[0], c.args[1], t)
	default:
		return "ERROR Unknown command"
	}
}

func parseCommand(data []byte) (*command, error) {
	msg := string(data)
	parts := strings.Split(msg, " ")
	if len(parts) < 1 {
		return nil, fmt.Errorf("Error parsing: %s", msg)
	}
	cmd := command{cmd: parts[0], args: make([]string, 0)}
	if len(parts) > 1 {
		cmd.args = parts[1:]
	}
	return &cmd, nil
}

func getKey(key string, t *pkg.RadixTree) string {
	node, err := t.ThreadSafeSearchNode([]byte(key))
	if err != nil {
		return fmt.Sprintf("ERROR %s", err.Error())
	}
	if len(node.Children) < 1 {
		return "NOT FOUND"
	}
	value := node.Children[0].Value
	return fmt.Sprintf("RESULT %s", string(value))
}

func addKeyVal(key, val string, t *pkg.RadixTree) string {
	vals := [][]byte{[]byte(key), []byte(val)}
	node, err := t.ThreadSafeSearchNode([]byte(key))
	if err != nil {
		err = t.ThreadSafeAddBranch(vals)
		if err != nil {
			return fmt.Sprintf("ERROR %s", err.Error())
		}
		return "RESULT SUCCESS"
	}
	if len(node.Children) > 0 {
		node.Children = []*pkg.Node{}
	}
	err = t.ThreadSafeRadixAdd(vals)
	if err != nil {
		return fmt.Sprintf("ERROR %s", err.Error())
	}
	return "RESULT SUCCESS"
}
