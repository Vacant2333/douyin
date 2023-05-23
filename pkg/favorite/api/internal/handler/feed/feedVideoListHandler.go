package feed

import (
	"net/http"

	"douyin/pkg/favorite/api/internal/logic/feed"
	"douyin/pkg/favorite/api/internal/svc"
	"douyin/pkg/favorite/api/internal/types"
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
