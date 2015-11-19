package vindinium

import "math/rand"

type Direction string

var DIRS = []Direction{"Stay", "North", "South", "East", "West"}

func randDir() Direction {
	dir := DIRS[rand.Intn(len(DIRS))]

	return dir
}

type Bot interface {
	Move(state *State) Direction
}

type RandomBot struct{}

func (b *RandomBot) Move(state *State) Direction {
	return randDir()
}

type FighterBot struct{}

func (b *FighterBot) Move(state *State) Direction {
	// g := NewGame(state)
	// Do something awesome
	return randDir()
}

type NaiveBot struct{}

func (b *NaiveBot) Move(state *State) Direction {
	// g := NewGame(state)
	// Do something awesome
	return randDir()
}
