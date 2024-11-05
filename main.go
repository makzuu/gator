package main

import (
	"errors"
	"fmt"
	"github.com/makzuu/gator/internal/config"
	"log"
	"os"
)

type state struct {
	conf *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	availableCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.availableCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.availableCommands[cmd.name]
	if !ok {
		return errors.New("error command does not exist")
	}
	return f(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("error username required")
	}
	name := cmd.args[0]
	err := s.conf.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("current user set to: %v\n", name)
	return nil
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	s := &state{}
	s.conf = &conf
	cmds := &commands{map[string]func(*state, command) error{}}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatalln("error not enough arguments provided")
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{cmdName, cmdArgs}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
