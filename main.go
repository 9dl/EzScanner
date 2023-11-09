package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("github/9dl -> EzScanner")
	fmt.Println("Choose a path to scan:")
	fmt.Println("[1] Desktop")
	fmt.Println("[2] Download")
	fmt.Println("[3] Main User")
	fmt.Println("[4] Main Drive")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	var dirPath string
	switch choice {
	case 1:
		dirPath = getDesktopPath()
	case 2:
		dirPath = getDownloadPath()
	case 3:
		dirPath = getMainUserPath()
	case 4:
		dirPath = getMainDriveSymbol()
	default:
		fmt.Println("Invalid choice. Exiting.")
		os.Exit(1)
	}

	fmt.Println("Scanning path:", dirPath)
	scanDirectory(dirPath)

	fmt.Print("Done.")
}
