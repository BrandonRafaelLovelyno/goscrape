package json

import "os"

func WriteToJson(jsonData *[]byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(*jsonData)
	if err != nil {
		return err
	}

	return nil
}
