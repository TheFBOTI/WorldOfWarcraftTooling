package main

import (
	"fmt"
)

func introduction() {
	fmt.Println("Hello and welcome to the world of Warcraft local tooling")
}

var initialMenu = []string{
	"Expansion",
	"Exit",
}

var expansionMenu = []string{
	"War Within",
}

var seasonMenu = []string{
	"Season Three/ Season 3",
}

var dungeonOrRaidMenu = []string{
	"Dungeons",
	"Raid(s)",
}

var dungeonsMenu = []string{
	"Halls of Attonement",
}

var raidMenu = []string{
	"Manaforge Omega",
}

// TODO - can we prevent loading all the bosses as individual menus and upload them as files and load them in.

func main() {
	introduction()
}
