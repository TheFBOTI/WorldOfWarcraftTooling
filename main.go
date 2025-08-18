package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func introduction() {
	fmt.Println("Hello and welcome to the world of Warcraft local tooling. \nPlease use the directional arrow keys to navigate the following menus. \nPlease also note that currently the directional arrow keys up and down may trigger pc notification noises and therefore best to use left or right. \nThanks! ")
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

	selectMenuWithKeyboard(initialMenu)
}

func selectMenuWithKeyboard(menuName []string) {
	// Create a prompt
	prompt := promptui.Select{
		Label: "Select Keyboard",
		Items: menuName,
	}
	// Run the prompt and get the selected index
	selected, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You selected: %s\n", menuName[selected])

}
