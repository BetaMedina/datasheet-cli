package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func resyncPayload(jobId string) (*bytes.Buffer, error) {
	postBody, err := json.Marshal(map[string]string{
		"job_id": jobId,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes.NewBuffer(postBody), nil
}

func RunReprocess(jobId string, kind string, actionType string) {
	payload, err := resyncPayload(jobId)
	if err != nil {
		fmt.Println("Failed to generate payload")
		return
	}
	fmt.Println("\nResyncing item...")
	response, err := http.Post(fmt.Sprintf("%s/v2/%s/%s", os.Getenv("BASE_URL"), kind, actionType), "application/json", payload)
	if err != nil {
		fmt.Println("Unsuccessful resync", err.Error())
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	if response.StatusCode == http.StatusOK {
		fmt.Println("\nItem resynced successfully")
		return
	}
	fmt.Println("\nResync failed with status:", response.Status, "and body:", string(body))
}
