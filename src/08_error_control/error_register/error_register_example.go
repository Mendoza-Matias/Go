package main

import (
	"log"
	"os"
)

func main() {
	// Set a prefix for all log messages to make them identifiable
	log.SetPrefix("prefix(): ")

	// Open or create the log file named "info.log" with specific flags and permissions
	file, err := os.OpenFile(
		"info.log", // Name of the log file
		os.O_CREATE| // Create the file if it does not exist
			os.O_APPEND| // Append new log entries to the end of the file
			os.O_WRONLY, // Open the file in write-only mode
		0644, // File permissions: read/write for owner, read-only for others
	)

	// Ensure the file is closed properly when the function ends
	defer file.Close()

	// Handle any errors that occurred while opening the file
	if err != nil {
		// Log a fatal error message and exit the program
		log.Fatal(err)
	}

	// Set the log output destination to the opened file
	log.SetOutput(file)

	// Write a log message to the file
	log.Print("Hey, I am log")
}
