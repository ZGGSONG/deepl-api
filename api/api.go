package api

import (
	"api4Deeplx/global"
	"api4Deeplx/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Translate(ctx *gin.Context) {
	var req model.Request
	var resp model.Response
	if err := ctx.BindJSON(&req); err != nil {
		log.Printf("Error translating, %v\n", err)
	}
	sourceText, targetLang := req.Text, req.TargetLang
	global.GLO_REQ_CH <- []string{sourceText, targetLang}

	select {
	case respText := <-global.GLO_RESP_CH:
		if respText[1] == "" {
			resp.Code = http.StatusOK
			resp.Data = respText[0]
			ctx.JSON(http.StatusOK, resp)
		} else {
			resp.Code = http.StatusGatewayTimeout
		}
	}

}

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "DeepL Translate Api\n\nPOST {\"text\": \"have a try\", \"source_lang\": \"auto\", \"target_lang\": \"ZH\"} to /translate\n\nhttps://github.com/zu1k")
}
