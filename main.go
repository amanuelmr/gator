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
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{cfg: &cfg}

	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}


	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
    	os.Exit(1)
	}

	cmd := command{
    name: os.Args[1],
    args: os.Args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
