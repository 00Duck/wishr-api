package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ResponseModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// EncodeJSON encodes JSON into a write response. Fails with reply to client containing error code
func (env *Env) encodeResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		errResponse(w, env.Log, 500, err)
	}
}

// DecodeJSON decodes JSON into an interface. Fails with reply to client containing error code
func (env *Env) decodeRequest(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	if r.Body == nil {
		errResponse(w, env.Log, 400, errors.New("Request body is empty"))
		return false
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		env.Log.Println("Decode error: " + err.Error())
		errResponse(w, env.Log, 400, errors.New("There was a problem decoding your request."))
		return false
	}
	return true
}

func errResponse(w http.ResponseWriter, logger *log.Logger, code int, errMsg error) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(&ResponseModel{Message: "Error: " + errMsg.Error()})
	if err != nil {
		logger.Println("EncodeResponse could not send error response to client: " + err.Error())
	}
}
