package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to the amazing plane walker!")
	board := Board{}.initialize()
	player := Player{}.initialize("Darren")
	scanner := bufio.NewScanner(os.Stdin)
	foo := Foo{}

	for scanner.Scan() {
		fmt.Print("> ")
		text := scanner.Text()
		if text == "exit" {
			fmt.Println("You teleport out!")
			break
		} else {
			switch text {
			case "e":
				if board.canMove(&player, text) {
					player.moveEast().displayCoordinates()
				} else {
					hitWall(text)
					player.displayCoordinates()
				}
			case "s":
				if board.canMove(&player, text) {
					player.moveSouth().displayCoordinates()
				} else {
					hitWall(text)
					player.displayCoordinates()
				}
			case "n":
				if board.canMove(&player, text) {
					player.moveNorth().displayCoordinates()
				} else {
					hitWall(text)
					player.displayCoordinates()
				}
			case "w":
				if board.canMove(&player, text) {
					player.moveWest().displayCoordinates()
				} else {
					hitWall(text)
					player.displayCoordinates()
				}
			default:
				fmt.Println("Unrecognized command.  Please use e, n, w or s. \"exit\" to quit")
			}
		}
	}

	if scanner.Err() != nil {
		// handle error.
	}
}

// Board Defines a board
type Board struct {
	xMax int
	xMin int
	yMax int
	yMin int
}

func (board Board) initialize() Board {
	board.xMax = 10
	board.yMax = 10
	board.xMin = 0
	board.yMin = 0
	return board
}

func (board *Board) canMove(player *Player, direction string) bool {
	switch direction {
	case "e":
		return player.currentPosition.x < board.xMax
	case "w":
		return player.currentPosition.x > board.xMin
	case "n":
		return player.currentPosition.y < board.yMax
	case "s":
		return player.currentPosition.y > board.yMin
	}
	return false
}

// Player Defines the player
type Player struct {
	name            string
	currentPosition Position
}

func (p Player) initialize(name string) Player {
	p.name = name
	p.currentPosition.x = 0
	p.currentPosition.y = 0
	fmt.Printf("Spawned at position %d, %d\n", 0, 0)
	return p
}

func (p *Player) moveEast() *Player {
	fmt.Println("You move East")
	p.currentPosition.x++
	return p
}

func (p *Player) moveWest() *Player {
	fmt.Println("You move West")
	p.currentPosition.x--
	return p
}

func (p *Player) moveNorth() *Player {
	fmt.Println("You move North")
	p.currentPosition.y++
	return p
}

func (p *Player) moveSouth() *Player {
	fmt.Println("You move South")
	p.currentPosition.y--
	return p
}

func (p *Player) displayCoordinates() {
	val := fmt.Sprintf("You are in room %v", p.currentPosition)
	fmt.Println(val)
	return
}

// Position Defines an x and y position
type Position struct {
	x int
	y int
}

func hitWall(direction string) {
	val := fmt.Sprintf("You hit a wall moving %s", direction)
	fmt.Println(val)
	return
}
