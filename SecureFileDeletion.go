package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// Get the file name to be securely deleted from the user
	fmt.Print("Enter the name of the file you want to securely delete: ")
	var fileName string
	_, err := fmt.Scanln(&fileName)
	if err != nil {
		fmt.Println("Error getting the file name:", err)
		return
	}

	// Perform the secure file deletion
	err = secureDeleteFile(fileName)
	if err != nil {
		fmt.Println("File could not be securely deleted:", err)
	} else {
		fmt.Println("File securely deleted.")
	}
}

func secureDeleteFile(fileName string) error {

	// Open the file and securely delete its content
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// Generate random data of the same size as the file and overwrite the file's content
	randomData := make([]byte, fileSize)
	rand.Read(randomData)
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.Write(randomData)
	if err != nil {
		return err
	}
	file.Sync()

	// Delete the file
	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}
