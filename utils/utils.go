package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// _, err := os.Getwd()
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }

	featuresDir := "../features"

	entries, err := os.ReadDir(featuresDir)

	log.Println(entries)
	if err != nil {
		fmt.Printf("Failed to read features directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Running migrations...")

	for _, entry := range entries {
		if entry.IsDir() {
			migrationsPath := filepath.Join(featuresDir, entry.Name(), "migrations")

			// Check if migrations directory exists
			if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
				continue // skip if no migrations folder
			}

			// Read migration .go files
			files, err := os.ReadDir(migrationsPath)
			if err != nil {
				fmt.Printf("Failed to read %s: %v\n", migrationsPath, err)
				continue
			}

			for _, file := range files {
				if filepath.Ext(file.Name()) == ".go" {
					filePath := filepath.Join(migrationsPath, file.Name())
					fmt.Printf("Running %s\n", filePath)

					// Run the migration file
					cmd := exec.Command("go", "run", filePath)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						fmt.Printf("Failed to run %s: %v\n", filePath, err)
					}
				}
			}
		}
	}
}
