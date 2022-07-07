package goovi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gooviid/gopkg/goovireq"
)

const baseUrl = "https://api.goovi.id/v1"
const keyConstant = "X-API-KEY"

//Transcode Video and Audio
func CreateTranscode(key string, request goovireq.CreateTranscodeRequest, client *http.Client) error {
	c := client
	if c == nil {
		c = http.DefaultClient
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
	r.Header.Set("Content-Type", "application/json")

	response, err := c.Do(r)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
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
func CreateCompress(key string, request goovireq.CreateCompressRequest, client *http.Client) error {
	c := client
	if c == nil {
		c = http.DefaultClient
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
	r.Header.Set("Content-Type", "application/json")

	response, err := c.Do(r)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
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
func GetTranscode(key string, id string, client *http.Client) (*TranscodeQueue, error) {
	c := client
	if c == nil {
		c = http.DefaultClient
	}

	r, err := http.NewRequest(http.MethodGet, baseUrl+"/transcode?id="+id, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set(keyConstant, key)

	response, err := c.Do(r)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		bodyString := string(bodyBytes)
		return nil, errors.New(bodyString)
	}
	res := new(TranscodeQueue)
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
