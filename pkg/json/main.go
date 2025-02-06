package json

import "os"

func GetJsonData(fileDir string) (*[]byte, error) {
	file, err := os.ReadFile(fileDir)
	if err != nil {
		return nil, err
	}

	return &file, nil
}

func WriteToJson(jsonData *[]byte, fileDir string) error {
	file, err := os.Create(fileDir)
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
