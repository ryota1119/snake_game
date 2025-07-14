package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/ryota1119/snake_game/internal/snake_game"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err = s.Init(); err != nil {
		log.Fatal(err)
	}
	defer s.Fini()

	g := snake.NewGame(s)
	go g.RunInputLoop()
	g.Loop()
}
