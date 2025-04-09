package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		maze = append(maze, l)
	}

	return nil
}

func printMaze() {
	for _, l := range maze {
		fmt.Println(l)
	}
}

func main() {

	// initialize game

	// load resources
	if err := loadMaze("maze01.txt"); err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	// main game loop

	for {

		// update screen
		printMaze()

		// process input

		// process movement

		// process collisions

		// check game over

		break

	}

}
