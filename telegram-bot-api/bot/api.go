/*
Package bot telegram机器人接口
Created by chenguolin 2018-12-13
*/
package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang/telegram-bot-api/common"
	"golang/telegram-bot-api/types"
)

// BotAPI allows you to interact with the Telegram Bot API.
type BotAPI struct {
	Token           string       `json:"token"`
	Debug           bool         `json:"debug"`
	Buffer          int          `json:"buffer"`
	Self            types.User   `json:"-"` //转成JSON的时候这个字段不导出
	Client          *http.Client `json:"-"` //转成JSON的时候这个字段不导出
	shutdownChannel chan interface{}
}

// NewBotAPI creates a new BotAPI instance.
// It requires a token, provided by @BotFather on Telegram.
func NewBotAPI(token string) (*BotAPI, error) {
	bot := &BotAPI{
		Token:           token,
		Client:          &http.Client{},
		Buffer:          100,
		shutdownChannel: make(chan interface{}),
	}

	self, err := bot.GetMe()
	if err != nil {
		return nil, err
	}
	bot.Self = self

	return bot, nil
}

// MakeRequest makes a request to a specific endpoint with our token.
func (bot *BotAPI) MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	method := fmt.Sprintf(common.APIEndpoint, bot.Token, endpoint)

	resp, err := bot.Client.PostForm(method, params)
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	bytes, err := bot.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return apiResp, err
	}

	if bot.Debug {
		log.Printf("%s resp: %s", endpoint, bytes)
	}

	if !apiResp.Ok {
		parameters := ResponseParameters{}
		if apiResp.Parameters != nil {
			parameters = *apiResp.Parameters
		}
		return apiResp, Error{apiResp.Description, parameters}
	}

	return apiResp, nil
}

// decodeAPIResponse decode response and return slice of bytes if debug enabled.
// If debug disabled, just decode http.Response.Body stream to APIResponse struct
// for efficient memory usage
func (bot *BotAPI) decodeAPIResponse(responseBody io.Reader, resp *APIResponse) (_ []byte, err error) {
	if !bot.Debug {
		dec := json.NewDecoder(responseBody)
		err = dec.Decode(resp)
		return
	}

	// if debug, read reponse body
	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	return data, nil
}

// makeMessageRequest makes a request to a method that returns a Message.
func (bot *BotAPI) makeMessageRequest(endpoint string, params url.Values) (types.Message, error) {
	resp, err := bot.MakeRequest(endpoint, params)
	if err != nil {
		return types.Message{}, err
	}

	var message types.Message
	json.Unmarshal(resp.Result, &message)

	bot.debugLog(endpoint, params, message)

	return message, nil
}

// GetMe fetches the currently authenticated bot.
// This method is called upon creation to validate the token,
// and so you may get this data from BotAPI.Self without the need for
// another request.
func (bot *BotAPI) GetMe() (types.User, error) {
	resp, err := bot.MakeRequest("getMe", nil)
	if err != nil {
		return types.User{}, err
	}

	var user types.User
	json.Unmarshal(resp.Result, &user)
	bot.debugLog("getMe", nil, user)

	return user, nil
}

// IsMessageToMe returns true if message directed to this bot.
// It requires the Message.
func (bot *BotAPI) IsMessageToMe(message types.Message) bool {
	return strings.Contains(message.Text, "@"+bot.Self.UserName)
}

// debugLog checks if the bot is currently running in debug mode, and if
// so will display information about the request and response in the
// debug log.
func (bot *BotAPI) debugLog(context string, v url.Values, message interface{}) {
	if bot.Debug {
		log.Printf("%s req : %+v\n", context, v)
		log.Printf("%s resp: %+v\n", context, message)
	}
}
