package msg

import (
	"context"

	"github.com/Faaizz/code_demos/go_simple_chat_app_api/types"
)

var mga types.MsgGwAdapter

func SetMsgGwAdapter(mgaInit types.MsgGwAdapter) {
	mga = mgaInit
}

func Message(cID, msg, fromUsername, url string) error {
	return mga.Message(context.TODO(), cID, msg, fromUsername, url)
}
