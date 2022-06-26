package goovi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const baseUrl = "https://api.goovi.id/v1"
const keyConstant = "X-API-KEY"

//Transcode Video and Audio
func CreateTranscode(key string, request CreateTranscodeRequest, client *http.Client) error {
	c := client
	if c == nil {
		c = &http.Client{}
	}

	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(b)

	r, err := http.NewRequest(http.MethodPost, baseUrl+"/transcode", bodyReader)
	if err != nil {
		return err
	}

	r.Header.Set(keyConstant, key)

	response, err := client.Do(r)

	if response.StatusCode > 400 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		return errors.New(bodyString)
	}

	return nil
}

//Compress Image to webp and original image
func CreateCompress(key string, request CreateCompressRequest, client *http.Client) error {
	c := client
	if c == nil {
		c = &http.Client{}
	}

	b, err := json.Marshal(request)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(b)

	r, err := http.NewRequest(http.MethodPost, baseUrl+"/compress", bodyReader)
	if err != nil {
		return err
	}

	r.Header.Set(keyConstant, key)

	response, err := client.Do(r)

	if response.StatusCode > 400 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		return errors.New(bodyString)
	}
	
	return nil
}
