package publish

import (
	"net/http"

	"douyin/pkg/favorite/api/internal/logic/publish"
	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PubVideoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publish.NewPublishVideoLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
