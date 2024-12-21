package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func resyncPayload(jobId string) (*bytes.Buffer, error) {
	postBody, err := json.Marshal(map[string]string{
		"jobId": jobId,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes.NewBuffer(postBody), nil
}

func RunResync(jobId string, kind string) {
	payload, err := resyncPayload(jobId)
	if err != nil {
		fmt.Println("Failed to generate payload")
		return
	}
	_, err = http.Post(fmt.Sprintf("https://webhook.site/afe1b27b-70fe-4d75-a4f8-668252ed17d2/%s", kind), "application/json", payload)
	if err != nil {
		fmt.Println("Unsuccessful resync", err.Error())
		return
	}
	fmt.Println("Item resynced successfully")
}
