package api

import (
	"deepl_api/global"
	"deepl_api/model"
	"deepl_api/model/deepl"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func Translate(ctx *gin.Context) {
	var req model.Request
	var deeplResp deepl.Response
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Error translating, %v\n", err)
	}
	global.GLO_REQ_CH <- []string{req.Text, req.SourceLang, req.TargetLang}

	select {
	case _resp := <-global.GLO_RESP_CH:
		defer _resp.Body.Close()

		body, _ := io.ReadAll(_resp.Body)
		_ = json.Unmarshal(body, &deeplResp)
		if deeplResp.Error.Code == -32600 {
			log.Println(deeplResp.Error.Message)
			ctx.JSON(http.StatusNotAcceptable, model.Response{
				Code: http.StatusNotAcceptable,
				Data: "Invalid targetLang",
			})
			return
		}

		if _resp.StatusCode == http.StatusTooManyRequests {
			ctx.JSON(http.StatusTooManyRequests, model.Response{
				Code: http.StatusTooManyRequests,
				Data: "Too Many Requests",
			})
		} else {
			ctx.JSON(http.StatusOK, model.Response{
				Code: http.StatusOK,
				Data: deeplResp.Result.Texts[0].Text,
			})
		}
	}

}

func TranslateGet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "POST {\"text\": \"input your content\", \"source_lang\": \"auto\", \"target_lang\": \"ZH\"} to /translate\n\n\nhttps://github.com/zggsong/stranslate\n")
}
