package cmd

import (
	"errors"
	"strings"
)

func (c *CLI) dispatch(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("no arguments")
	}
	for i := 0; i < len(args); i++ {
		args[i] = strings.TrimSpace(args[i])
	}
	switch args[0] {
	case "pwrc":
		if len(args) != 3 {
			return "Command expects 3 arguments. Ex: pwrc <user> <password>", nil
		}
		msg := c.db.PasswordResetCLI(args[1], args[2])

		return msg, nil
	case "gpwr":
		if len(args) != 2 {
			return "Command expects 2 arguments. Ex: gpwr <user>", nil
		}
		msg := c.db.GeneratePasswordResetRequest(args[1])
		return msg, nil
	case "ls":
		if len(args) != 2 {
			return "Command expects 2 arguments. Ex: ls users", nil
		}
		msg := c.db.ListObjects(args[1])
		return msg, nil
	case "help":
		return "help run", nil
	default:
		return "Invalid command", nil
	}
}
