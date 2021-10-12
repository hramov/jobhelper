package device_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/files"
	"github.com/hramov/jobhelper/src/modules/logger"
)

type ApiResponse struct {
	OK     bool `json:"ok"`
	Result struct {
		FileID       string `json:"file_id"`
		FileUniqueId string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
	}
}

func UploadTagImageUrl(device_id uint, file_id string) error {

	if device_id == 0 {
		return fmt.Errorf("Не выбрано оборудование")
	}

	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", os.Getenv("TOKEN"), (file_id)))
	if err != nil {
		logger.Log("TGBot:HandleQuery", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	apiResp := ApiResponse{}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Println(err)
		return err
	}

	realImagePath := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", os.Getenv("TOKEN"), apiResp.Result.FilePath)
	resp, err = http.Get(realImagePath)
	if err != nil {
		logger.Log("TGBot:HandleQuery", err.Error())
		return err
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	imagePath := fmt.Sprintf("uploads/%d.jpg", device_id)

	err = files.UploadFile(imagePath, body)
	if err != nil {
		logger.Log("TGBot:HandleQuery", err.Error())
		return err
	}

	err = deviceEntity.UploadImage(device_id, imagePath)
	if err != nil {
		logger.Log("TGBot:HandleQuery", err.Error())
		return err
	}
	return nil
}
