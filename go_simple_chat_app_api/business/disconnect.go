package business

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"

	"github.com/Faaizz/code_demos/go_simple_chat_app_api/db"
	"github.com/Faaizz/code_demos/go_simple_chat_app_api/types"
)

func DisconnectHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debugln("disconnecting user...")

	rBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode request into bytes"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}
	logger.Debugf("request body: \n%v", string(rBytes))

	var u types.User
	err = json.NewDecoder(bytes.NewReader(rBytes)).Decode(&u)
	if err != nil {
		logger.Errorln(err)
		msg := "could not decode input"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	err = db.Disconnect(u)
	if err != nil {
		logger.Errorln(err)
		msg := "could not disconnect"
		logger.Errorln(msg)
		w.WriteHeader(400)
		fmt.Fprint(w, msg)
		return
	}

	msg := fmt.Sprintf("disconnected user: %s", u.Username)
	logger.Debugln(msg)
	_, err = fmt.Fprintln(w, msg)
	if err != nil {
		logger.Errorln(err)
	}
}
