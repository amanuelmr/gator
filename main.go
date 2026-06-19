package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amanuelmandefro3/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments provided")
    	os.Exit(1)
	}

	 cmd := command{
    name: os.Args[1],
    args: os.Args[2:],
	}
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	cmds := commands{
		commands: map[string]func(*state, command) error{},
	}
	s := state{cfg: &cfg}
	
	cmds.register("login", handlerLogin)

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
    	os.Exit(1)
	}

	// err = s.cfg.SetUser("")
	// if err != nil {
	// 	log.Fatalf("couldn't set current user: %v", err)
	// }

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
