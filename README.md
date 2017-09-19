# Speakeasy

![Speakeasy](img/speakeasy.png)

A very simple package for sending SMS messages and making calls with a Twilio account.

## Methods

### SMS(to, body)
Generate an outbound SMS message

### Call(to, twimlUrl)
Generate an outbound Twilio call to the supplied number. Twilio will attempt to parse TwiML served up at the supplied twimlUrl.

## Usage

```go

package main

import (
  "fmt"

  "github.com/zackproser/speakeasy"
)

func main() {
  //Instantiate Speakeasy with your account SID, AuthToken and Twilio number
  s := speakeasy.New("AC49a43r78fh717463fce21dsfue6ae", "8211129a9d43c587eftxbdh39c859666", "+555-555-5555")

  //Send an SMS
  res, err := s.SMS("+14158675309", "Hello from Speakeasy")

  if err != nil {
    fmt.Printf("Error sending SMS: %v", err)
  }

  fmt.Printf("Response: %v", res)

  //Make a phone call
  callRes, callErr := s.Call("+15103267023", "http://twimlets.com/echo?Twiml=%3CResponse%3E%3CSay%3EWelcome+to+speak+easy.%3C%2FSay%3E%3C%2FResponse%3E")

  if callErr != nil {
    fmt.Printf("Error making Call: %v", callErr)
  }

  fmt.Printf("Response: %v", callRes)
}

```

### Documentation

[Godocs](https://godoc.org/github.com/zackproser/speakeasy)