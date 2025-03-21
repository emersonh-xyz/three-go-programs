// This program demonstrates how to recover from a panic in a goroutine.
// It uses a deferred function to handle the panic and prevent the program from crashing.
// The main function continues to run even if the goroutine panics.

package main

import (
    "fmt"
    "time"
)

func main() {
    // Start a goroutine that might panic
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered from panic:", r)
            }
        }()

        fmt.Println("Starting a task...")
        time.Sleep(1 * time.Second)
        panic("Something went wrong in the goroutine!") // Simulate a panic
    }()

    // Main function continues running
    fmt.Println("Main function is doing other work...")
    time.Sleep(2 * time.Second) // Allow the routine to complete
    fmt.Println("Main function finished.")
}