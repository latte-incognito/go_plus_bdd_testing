package main

import (
	"fmt"
	"github.com/cucumber/messages-go/v10"
	"io/ioutil"
	"net/http"
)

type responseSt struct {
	resp *http.Response
}

func (b *responseSt) iSendRequestToExternal(method, url string) error {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	resp, err := client.Do(req)

	b.resp = resp

	return err
}

func (b *responseSt) theExternalResponseCodeShouldBe(code int) error {
	if code != b.resp.StatusCode {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, b.resp.StatusCode)
	}
	return nil
}

func (b *responseSt) theExternalResponseShouldMatchJson(body *messages.PickleStepArgument_PickleDocString) (err error) {
	var expected, actual interface{}

	// re-encode actual response too
	bodyBytes, err := ioutil.ReadAll(b.resp.Body)

	encodeAndCompare([]byte(body.Content), &expected, bodyBytes, &actual)

	return nil
}
