// This program demonstrates the use of data structures and control structures in Go.
// It includes an array, a slice, and control structures (for loop and if).
// It also defines a struct to represent a player with a name and score.

package main

import "fmt"

// Define a struct for Player
type Player struct {
    Name  string
    Score int
}

func main() {
    // Data structure: Array
    numbers := [5]int{10, 20, 30, 40, 50}

    // Data structure: Slice
    players := []Player{
        {"Alice", 100},
        {"Bob", 200},
        {"Charlie", 150},
    }

    // Control structure: For loop
    fmt.Println("Array elements:")
    for i, num := range numbers {
        fmt.Printf("Index %d: %d\n", i, num)
    }

    // Control structure: If-else
    fmt.Println("\nPlayer scores:")
    for _, player := range players {
        if player.Score > 150 {
            fmt.Printf("%s has a high score of %d\n", player.Name, player.Score)
        } else {
            fmt.Printf("%s has a score of %d\n", player.Name, player.Score)
        }
    }
}