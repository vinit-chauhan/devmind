package utils

import "os"

func WriteToFile(filePath string, data []byte) error {
	// Open the file for writing, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the data to the file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func WriteToStdout(data []byte) error {
	// Write the data to stdout
	_, err := os.Stdout.Write(data)
	if err != nil {
		return err
	}

	return nil
}
