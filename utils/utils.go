package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func migrate() {
	featuresDir := "../features"

	entries, err := os.ReadDir(featuresDir)

	if err != nil {
		log.Fatalf("Failed to read features directory: %v\n", err)
	}

	log.Println("Running migrations...")

	for _, entry := range entries {
		if entry.IsDir() {
			migrationsPath := filepath.Join(featuresDir, entry.Name(), "migrations")

			if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
				continue
			}

			files, err := os.ReadDir(migrationsPath)
			if err != nil {
				log.Printf("Failed to read %s: %v\n", migrationsPath, err)
				continue
			}

			for _, file := range files {
				if filepath.Ext(file.Name()) == ".go" {
					filePath := filepath.Join(migrationsPath, file.Name())
					log.Printf("Running %s\n", filePath)

					cmd := exec.Command("go", "run", filePath)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						log.Printf("Failed to run %s: %v\n", filePath, err)
					}
				}
			}
		}
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("No argument passed")
	}

	for _, arg := range args {
		if arg == "--migrate" {
			migrate()
		}
	}
}
