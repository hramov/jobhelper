package files

import (
	"fmt"
	"os"

	"github.com/hramov/jobhelper/src/modules/logger"
)

func CheckFile(fileName string) {
	_, err := os.Open(fileName)
	if err != nil {
		logger.Log("File Manager", fmt.Sprintf("Cannot find file: %s", fileName))
		_, err := os.Create(fileName)
		if err != nil {
			logger.Log("File Manager", err.Error())
			return
		}
		logger.Log("File Manager", fmt.Sprintf("Successfully created file: %s", fileName))
	}
	logger.Log("File Manager", fmt.Sprintf("File %s exists", fileName))
}

func UploadFile(fileName string, file []byte) error {
	image, err := os.Create(fileName)
	if err != nil {
		logger.Log("File Manager", err.Error())
	}
	defer image.Close()

	_, err = image.Write(file)

	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(fileName string) error {
	return os.Remove(fileName)
}
