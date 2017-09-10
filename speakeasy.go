package speakeasy

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

//Twilio REST API Root
const TwilioBaseUrl = "https://api.twilio.com/2010-04-01"

type Speakeasy struct {
	SID          string
	AuthToken    string
	TwilioNumber string
	client       *http.Client
}

// Create and return a new Speakeasy client
// Panics if any of the required parameters is missing
func New(sid, token, fromNumber string) *Speakeasy {

	if sid == "" || token == "" || fromNumber == "" {
		panic(errors.New("Speakeasy must be initialized with sid, token and fromNumber"))
	}

	return &Speakeasy{
		SID:          sid,
		AuthToken:    token,
		TwilioNumber: fromNumber,
		client:       &http.Client{},
	}
}

// Returns a properly formatted resource URL for Twilio REST calls
func (s *Speakeasy) FormatTwilioUrl(ResourceType string) string {
	return fmt.Sprintf("%s/Accounts/%s/%s.json", TwilioBaseUrl, s.SID, ResourceType)
}

// Builds a properly formatted and authorized request to Twilio's API
func (s *Speakeasy) FormatRequest(data url.Values, ResourceType string) (*http.Request, error) {

	req, buildErr := http.NewRequest("POST", s.FormatTwilioUrl(ResourceType), bytes.NewBufferString(data.Encode()))
	req.SetBasicAuth(s.SID, s.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, buildErr
}

// Generates an outbound SMS message
// to is the recipient phone number
// body is the message to send
func (s *Speakeasy) SMS(to, body string) (*http.Response, error) {

	data := url.Values{}
	data.Set("From", s.TwilioNumber)
	data.Set("To", to)
	data.Set("Body", body)

	req, buildErr := s.FormatRequest(data, "Messages")

	if buildErr != nil {
		fmt.Printf("Error formatting SMS request: %v", buildErr)
	}

	return s.client.Do(req)
}

// Generates an outbound Twilio call to the supplied number
// to is the recipient phone number
// twimlUrl is the publicly accessible URL where Twilio can expect
// to find valid TwiML
//
// See Twilio docs on TwiML: https://www.twilio.com/docs/api/twiml
func (s *Speakeasy) Call(to, twimlUrl string) (*http.Response, error) {

	data := url.Values{}
	data.Set("From", s.TwilioNumber)
	data.Set("To", to)
	data.Set("Url", twimlUrl)

	req, buildErr := s.FormatRequest(data, "Calls")

	if buildErr != nil {
		fmt.Printf("Error formatting SMS request: %v", buildErr)
	}

	return s.client.Do(req)
}
