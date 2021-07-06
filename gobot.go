package gobot

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/valyala/fasthttp"
)

type GoBot struct {
	client   *fasthttp.Client
	Timeout  int
	baseURL  string
	handlers []Handler
}

type Handler struct {
	updateType reflect.Type
	callback   func(bot *GoBot, update *Update)
}

var updateTypeUpdate = reflect.TypeOf(&Update{})

func Init(token string) *GoBot {
	log.Println("Copyright (c) 2020 Mattia Brandon <https://github.com/mattiabrandon>")
	bot := GoBot{
		client:   &fasthttp.Client{},
		Timeout:  12,
		baseURL:  "https://api.telegram.org/bot" + token + "/",
		handlers: []Handler{},
	}
	return &bot
}

func InitTimeout(token string, timeout int) *GoBot {
	log.Println("Copyright (c) 2020 Mattia Brandon <https://github.com/mattiabrandon>")
	bot := GoBot{
		client:   &fasthttp.Client{},
		Timeout:  timeout,
		baseURL:  "https://api.telegram.org/bot" + token + "/",
		handlers: []Handler{},
	}
	return &bot
}

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
			Description: decodedBody.Description,
			ErrorCode:   decodedBody.ErrorCode,
			Parameters:  decodedBody.Parameters,
		}
	}
	return decodedBody.Result, nil
}

func (bot *GoBot) AddHandler(updateType interface{}, callback func(bot *GoBot, update *Update)) {
	if reflect.ValueOf(updateType).Type().Kind() != reflect.Ptr {
		panic("Update type must be a pointer")
	}
	bot.handlers = append(bot.handlers, Handler{reflect.TypeOf(updateType), callback})
}

func (bot *GoBot) handleUpdate(update *Update) {
	for _, handler := range bot.handlers {
		if handler.updateType == updateTypeUpdate {
			handler.callback(bot, update)
			continue
		}
		v := reflect.Indirect(reflect.ValueOf(update))

		for i := 1; i < v.NumField(); i++ {
			if v.Field(i).Type() == handler.updateType && !v.Field(i).IsNil() {
				handler.callback(bot, update)
				break
			}
		}
	}
}

func (bot *GoBot) Loop(returnError bool) error {
	log.Println("Starting the loop...")
	offset := 0

	for {
		updates, err := bot.GetUpdates(GetUpdatesParams{
			Offset:  offset,
			Timeout: bot.Timeout,
		})

		if err != nil {
			if returnError {
				return err
			}
			fmt.Println(err)
		}

		for _, update := range updates {
			offset = update.UpdateId + 1
			go bot.handleUpdate(update)
		}
	}
}
