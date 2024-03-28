package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type appConfig struct {
	NodeID string
	Date   string
}

var (
	port   string
	nodeID string
)

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	nodeID = os.Getenv("NODE_ID")
	if nodeID == "" {
		nodeID = "unknown"
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	logger = logger.With("app", "go_simple_website")
	logger = logger.With("node_id", nodeID)
	logger = logger.With("port", port)
	slog.SetDefault(logger)
}

func main() {
	assetsFs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetsFs))

	http.HandleFunc("/", serveTemplate)

	slog.Info("starting server...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	// load dynamic template
	dp := filepath.Join("templates", "dynamic.html")

	// get requested path
	up := filepath.Clean(r.URL.Path)
	if up == "/" {
		// default to index.html
		up = "index.html"
	}
	slog.Info("requested page", up)
	// load requested template
	fp := filepath.Join("templates", up)

	tmpl, err := template.ParseFiles(dp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "layout", appConfig{
		NodeID: nodeID,
		Date:   time.Now().Format(time.RFC3339),
	})
}
