package utils

import (
	"fmt"
	"goapp/models"
	"os"

	"github.com/bytedance/sonic"
)

func LoadConfig(filePath string) {
	// Implementation to load configuration from the specified file path
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	err = sonic.Unmarshal(fileBytes, &models.Config)
	if err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		return
	}
}
