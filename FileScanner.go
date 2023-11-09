package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// Database for all Clients
var clientDB = map[string]string{
	"VapeV4": "17da27502307d04049ae11abac77bbcc14cf9a35dd54d879930b528ba4b9384f",
}

func calculateSHA256(fileData []byte) (string, error) {
	hash := sha256.New()
	_, err := hash.Write(fileData)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func detectClient(filePath string) (string, bool) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return "", false
	}

	fileSHA256, err := calculateSHA256(fileData)
	if err != nil {
		fmt.Printf("Error calculating SHA256 hash: %s\n", err)
		return "", false
	}

	for clientName, clientHash := range clientDB {
		if fileSHA256 == clientHash {
			return clientName, true
		}
	}

	return "", false
}

func scanDirectory(dirPath string) {
	startTime := time.Now()
	filesScanned := 0

	err := filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing file: %s\n", err)
			return nil
		}
		if !info.IsDir() {
			clientName, isInfected := detectClient(filePath)
			if isInfected {
				fmt.Printf("Clients Detected: %s (%s)\n", filePath, clientName)
			}
			filesScanned++
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error scanning directory: %s\n", err)
	}

	elapsed := time.Since(startTime)
	filesPerSecond := float64(filesScanned) / elapsed.Seconds()

	fmt.Printf("Scan took: %s\n", elapsed)
	fmt.Printf("Files scanned per second: %.2f\n", filesPerSecond)
}
