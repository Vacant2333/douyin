package feed

import (
	"douyin/pkg/gateway/internal/logic/feed"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := feed.NewFeedVideoListLogic(r.Context(), svcCtx)
		resp, err := l.FeedVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
