package userOpt

import (
	"douyin/pkg/gateway/api/internal/logic/userOpt"
	"douyin/pkg/gateway/api/internal/svc"
	"douyin/pkg/gateway/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userOpt.NewGetCommentListLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
