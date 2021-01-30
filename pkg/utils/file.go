package utils

import "os"

func WriteFile(filename string, output []byte) error {
	file, err  := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(output)
	if err != nil {
		return err
	}
	return nil
}
