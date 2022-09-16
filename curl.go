package curl

import "C"
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Client *http.Client

func NewClient() {
	t := http.Transport{
		IdleConnTimeout:     60 * time.Second, // Time out for reuse connection
		TLSHandshakeTimeout: 5 * time.Second,
	}

	Client = &http.Client{
		Transport: &t,
	}
}

func GetAPI(c context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func PostAPIJson(c context.Context, url string, data interface{}) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, 200*time.Millisecond)
	defer cancel()

	dataByte, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		log.Println(fmt.Sprintf("API Post team Data url = %s data = %v err = %s", url, data, err), "", nil)
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

func PostAPI(c context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}

func DeleteAPI(c context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}

func PostAPIJsonTonken(c context.Context, url string, data interface{}, token string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel()
	dataByte, _ := json.Marshal(data)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}

func PostAPIJsonForSearch(c context.Context, url string, data interface{}) ([]byte, error) {
	ctx, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel()

	dataByte, _ := json.Marshal(data)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}
