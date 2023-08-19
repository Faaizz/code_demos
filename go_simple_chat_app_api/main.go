package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/Faaizz/code_demos/go_simple_chat_app_api/business"
	"github.com/Faaizz/code_demos/go_simple_chat_app_api/db"
	"github.com/Faaizz/code_demos/go_simple_chat_app_api/misc"
	"github.com/Faaizz/code_demos/go_simple_chat_app_api/msg"
	"github.com/Faaizz/code_demos/go_simple_chat_app_api/types"
)

func main() {
	// setup logger
	logger := misc.Logger()

	// setup DB
	dbType := os.Getenv("DB_TYPE")
	logger.Infof("DB_TYPE: %s\n", dbType)
	tn := os.Getenv("DYNAMODB_TABLE_NAME")
	logger.Infof("DYNAMODB_TABLE_NAME: %s\n", tn)

	var dba types.DBAdapter

	switch dbType {
	case "", "DYNAMODB":
		dba = &types.DynamoDBAdapter{}

	default:
		dba = &types.DynamoDBAdapter{}
	}

	db.SetDatabaseAdapter(dba)
	err := db.CheckExists(tn)
	if err != nil {
		logger.Fatalf("table does not exist %v", err)
	}

	// setup message gateway adapter
	var mga types.MsgGwAdapter
	mga = &types.AWSApiGwAdapter{}

	msg.SetMsgGwAdapter(mga)

	// setup routing
	r := mux.NewRouter()

	r.HandleFunc("/connect", business.ConnectHandler).Methods("POST")
	r.HandleFunc("/username", business.UsernameHandler).Methods("POST")
	r.HandleFunc("/online", business.OnlineHandler).Methods("POST")
	r.HandleFunc("/disconnect", business.DisconnectHandler).Methods("POST")
	r.HandleFunc("/message", business.MessageHandler).Methods("POST")
	r.HandleFunc("/healthz", business.HealthHandler).Methods("GET")

	// setup middleware for AWS X-Ray
	xda := os.Getenv("AWS_XRAY_DAEMON_ADDRESS")
	if xda != "" {
		err = misc.ConfigXRay(xda)
		if err != nil {
			logger.Fatalf("could not configure X-Ray %v", err)
		}

		r.Use(misc.XRayMw)
		logger.Infoln("using X-Ray tracing")
	}

	// listen for connections
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}
	logger.Infof("HTTP_PORT: %s", port)

	listenIpPort := fmt.Sprintf(":%s", port)

	logger.Infoln("starting server")
	logger.Fatal(http.ListenAndServe(listenIpPort, r))
}
