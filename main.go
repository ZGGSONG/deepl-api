package main

import (
	"deepl_api/server"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
)

var (
	HOST = "0.0.0.0"
	PORT = 8000
	URL  = "https://www2.deepl.com/jsonrpc"
)

func init() {
	color.Magenta("Configuration")
	color.Blue(">> address: %v", HOST)
	color.Blue(">> port: %v", PORT)
	color.Magenta("Routes:")
	color.Blue(">> (index) GET /")
	color.Blue(">> (text_translate) POST /translate")
	color.Magenta("Catchers:")
	color.Blue(">> (not_found) 404")
	color.Blue(">> (bad_request) 400")
	color.Blue(">> (internal_err) 500")
}

func main() {
	apiServer := server.NewAPIServer(URL, 10)

	mux := http.NewServeMux()
	mux.Handle("/translate", apiServer)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST {\"text\": \"input your content\", \"source_lang\": \"auto\", \"target_lang\": \"ZH\"} to /translate\n\n\ngithub.com/zggsong/stranslate\n"))
	})
	// 启动HTTP服务器
	color.Magenta("Rocket has launched from http://%v:%v\n", HOST, PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", HOST, PORT), mux))
}
