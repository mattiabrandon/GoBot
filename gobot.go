// Package gobot provides an easy to use interface to interact with
// the BotAPI.
package gobot

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

// GoBot is the main struct of the gobot package.
type GoBot struct {
	client   *fasthttp.Client
	Timeout  int
	baseURL  string
	handlers []func(bot *GoBot, update *Update)
}

// Init is the function used to create the gobot main instance.
func Init(token string) (*GoBot, error) {
	log.Println("Copyright (c) 2020 Mattia Brandon <https://github.com/mattiabrandon>")
	bot := GoBot{
		client:   &fasthttp.Client{},
		Timeout:  12,
		baseURL:  "https://api.telegram.org/bot" + token + "/",
		handlers: []func(bot *GoBot, update *Update){},
	}
	getMe, err := bot.GetMe()

	if err != nil {
		return nil, err
	}
	log.Printf("Logged in as %s (%d)\n", getMe.Username, getMe.ID)
	return &bot, nil
}

// Request is the function used to call all the methods from the BotAPI
// and return their values.
func (bot *GoBot) Request(method string, params interface{}) (json.RawMessage, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(bot.baseURL + method)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")

	if params != nil {
		encodedParams, err := json.Marshal(params)

		if err != nil {
			return nil, err
		}
		req.SetBody(encodedParams)
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := bot.client.Do(req, resp); err != nil {
		return nil, err
	}
	decodedBody := Body{}

	if err := json.Unmarshal(resp.Body(), &decodedBody); err != nil {
		return nil, err
	} else if !decodedBody.Ok {
		return nil, &Error{
			Description:        decodedBody.Description,
			ErrorCode:          decodedBody.ErrorCode,
			ResponseParameters: decodedBody.ResponseParameters,
		}
	}
	return decodedBody.Result, nil
}

// AddHandler is the function used to add an updates handler
// to the list.
func (bot *GoBot) AddHandler(handler func(bot *GoBot, update *Update)) {
	bot.handlers = append(bot.handlers, handler)
}

func (bot *GoBot) handleUpdate(update *Update) {
	for _, handler := range bot.handlers {
		handler(bot, update)
	}
}

// Loop is the function used to start the update loop and their
// handling process.
func (bot *GoBot) Loop() error {
	log.Println("Starting the loop...")
	offset := 0

	for {
		updates, err := bot.GetUpdates(GetUpdatesParams{
			Offset:  offset,
			Timeout: bot.Timeout,
		})

		if err != nil {
			return err
		}

		for _, update := range updates {
			offset = update.UpdateID + 1
			go bot.handleUpdate(update)
		}
	}
}
