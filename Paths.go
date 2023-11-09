package main

import (
	"os"
	"path/filepath"
)

func getMainDriveSymbol() string {
	return filepath.VolumeName(os.Getenv("SystemDrive")) + string(filepath.Separator)
}

func getDesktopPath() string {
	return filepath.Join(os.Getenv("USERPROFILE"), "Desktop")
}

func getDownloadPath() string {
	return filepath.Join(os.Getenv("USERPROFILE"), "Downloads")
}

func getMainUserPath() string {
	return os.Getenv("USERPROFILE")
}
