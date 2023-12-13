package healthchecks

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Notify notify to healthchecks.io a state (use constants provided)
func (C *Config) Notify(event string) error {

	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("%s%s", C.Url, event)

	_, err := client.Head(url)
	return err
}

// NotifyWithLog notify to healthchecks.io a state (use constants provided) and adds a message.
// NB: healthchecks.io will accept only the first 100kB of the string.
func (C *Config) NotifyWithLog(event string, text string) error {

	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("%s%s", C.Url, event)
	body := strings.NewReader(text)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	return nil
}

// Panic is an util function that accepts an error and notify a fail status with the error message before panicking.
func (C *Config) Panic(err any) {

	message := fmt.Sprintf("%v", err)

	if e := C.NotifyWithLog(FAIL, message); e != nil {
		log.Println(e)
	}

	panic(err)
}
