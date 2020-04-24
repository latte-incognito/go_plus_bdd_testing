package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) resetResponse(*messages.Pickle) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendrequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch endpoint {
	case "/version":
		getVersion(a.resp, req)
	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *messages.PickleStepArgument_PickleDocString) (err error) {
	var expected, actual interface{}

	encodeAndCompare([]byte(body.Content), &expected, a.resp.Body.Bytes(), &actual)

	return nil
}

func encodeAndCompare(exp2Unmarshal []byte, exp2Store interface{}, actual2Unmarshll []byte, actual2Store interface{}) (err error) {
	// re-encode expected response
	if err = json.Unmarshal(exp2Unmarshal, exp2Store); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(actual2Unmarshll, actual2Store); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(exp2Store, actual2Store) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", exp2Store, actual2Store)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	internal_api := &apiFeature{}

	s.BeforeScenario(internal_api.resetResponse)
	s.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, internal_api.iSendrequestTo)
	s.Step(`^the response code should be (\d+)$`, internal_api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, internal_api.theResponseShouldMatchJSON)

	external_api := &responseSt{}

	s.Step(`^I send "([^"]*)" request to external "([^"]*)"$`, external_api.iSendRequestToExternal)
	s.Step(`^the external response code should be (\d+)$`, external_api.theExternalResponseCodeShouldBe)
	s.Step(`^the external response should match json:$`, external_api.theExternalResponseShouldMatchJson)
}
