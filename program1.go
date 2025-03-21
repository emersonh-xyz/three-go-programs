// This program is a number guessing game where the user has to guess a random number between 1 and 100.
// It demonstrates the use of variables, control structures, and basic input/output in Go.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

 // Four data types
 var attempts int = 0               
 var sumOfGuesses int = 0            
 var averageGuess float64           
 var playerName string              
 var guessedCorrectly bool = false  

    // Prompt for player name
    fmt.Print("Enter your name: ")
    fmt.Scanln(&playerName)

    // Generate a random number
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1

    fmt.Printf("Hello %s! I have selected a random number between 1 and 100. Can you guess it?\n", playerName)

    var guess int
    for !guessedCorrectly{
        fmt.Print("Enter your guess: ")
        fmt.Scanln(&guess)

        attempts++
        sumOfGuesses += guess // Add the guessed number to the sum

        if guess < target {
            fmt.Println("Too low! Try again.")
        } else if guess > target {
            fmt.Println("Too high! Try again.")
        } else {
            guessedCorrectly = true
        }
    }

    // Calculate average attempts
    averageGuess = float64(sumOfGuesses) / float64(attempts)

    // Provide feedback about the guesses
    fmt.Printf("You guessed the number in %d attempts.\n", attempts)
    fmt.Printf("The sum of all your guesses is: %d\n", sumOfGuesses)
    fmt.Printf("The average of your guesses is: %.2f\n", averageGuess)
}