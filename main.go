package main

import (
	"deepl_api/server"
	"fmt"
	"log"
	"net/http"
)

var (
	PORT = 8000
	URL  = "https://www2.deepl.com/jsonrpc"
)

func main() {
	apiServer := server.NewAPIServer(URL, 10)

	mux := http.NewServeMux()
	mux.Handle("/translate", apiServer)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST {\"text\": \"input your content\", \"source_lang\": \"auto\", \"target_lang\": \"ZH\"} to /translate\n\n\ngithub.com/zggsong/stranslate\n"))
	})
	// 启动HTTP服务器
	fmt.Printf("deepl server starte at %v...\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), mux))
}
