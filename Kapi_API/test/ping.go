package test 

import (
	"encoding/json"
	"github.com/brainfucker/zero"
	"fmt"

)

// PingInput ds
type PingInput struct {
	Ping string `json:"ping"`
}

// PingParse ds 
func PingParse(srv *zero.Server) *PingInput {
	input := PingInput{}
	err := json.Unmarshal(srv.GetBody(), &input)
	if err == nil {
		return &input
	}
	return nil
}

// returns pong
func Pong(srv *zero.Server) {
	ping := PingParse(srv)

	srv.Resp(ping)
	fmt.Println(ping, " pong")
}