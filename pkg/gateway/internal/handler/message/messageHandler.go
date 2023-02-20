package message

import (
	"net/http"

	"douyin/pkg/gateway/internal/logic/message"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewMessageLogic(r.Context(), svcCtx)
		resp, err := l.Message(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
