package publish

import (
	"douyin/pkg/gateway/api/internal/logic/publish"
	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPublishVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPubVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publish.NewGetPublishVideoListLogic(r.Context(), svcCtx)
		resp, err := l.GetPublishVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
