package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func renewPayload(jobId string) (*bytes.Buffer, error) {
	postBody, err := json.Marshal(map[string]string{
		"jobId": jobId,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes.NewBuffer(postBody), nil
}

func RunRenew(jobId string, kind string) {
	payload, err := renewPayload(jobId)
	if err != nil {
		fmt.Println("Failed to generate payload")
		return
	}
	_, err = http.Post(fmt.Sprintf("https://webhook.site/afe1b27b-70fe-4d75-a4f8-668252ed17d2/renew/%s", kind), "application/json", payload)
	if err != nil {
		fmt.Println("Unsuccessful renew", err.Error())
		return
	}
	fmt.Println("Job renew successfully")
}
