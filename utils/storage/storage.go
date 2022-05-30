package storage

import (
	"api-store/utils"
	"bytes"
	"fmt"

	storage_go "github.com/supabase-community/storage-go"
)

func SetupStorage() {
	url := utils.GetEnv("STORAGE_URL", "http://localhost:9000")
	key := utils.GetEnv("STORAGE_TOKEN", "")

	client := storage_go.NewClient(url, key, nil)

	var options storage_go.BucketOptions
	options.Public = false

	fmt.Println("Storage client created")
	client.CreateBucket("api-store", options)
}

func UploadFiles(file []byte, fileName string) (string, error) {
	url := utils.GetEnv("STORAGE_URL", "http://localhost:9000")
	key := utils.GetEnv("STORAGE_TOKEN", "")

	client := storage_go.NewClient(url, key, nil)
	client.CreateBucket("api-store", storage_go.BucketOptions{})

	files := bytes.NewReader(file)

	filePath := "api-store/" + fileName
	err := client.UploadFile("users", filePath, files)

	fmt.Println("File uploaded", err)

	return filePath, nil
}

func UploadBase64(file []byte, name string) (storage_go.FileUploadResponse, error) {
	url := utils.GetEnv("STORAGE_URL", "http://localhost:9000")
	key := utils.GetEnv("STORAGE_TOKEN", "")

	client := storage_go.NewClient(url, key, nil)
	client.CreateBucket("api-store", storage_go.BucketOptions{})

	filePath := "api-store/" + name
	err := client.UploadFile("users", filePath, bytes.NewReader(file))

	//Get Respond

	fmt.Println("File uploaded", err)

	return err, nil
}
