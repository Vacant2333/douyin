package userOpt

import (
	"douyin/common/xerr"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"douyin/pkg/comment/api/internal/logic/userOpt"
	"douyin/pkg/comment/api/internal/svc"
	"douyin/pkg/comment/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentOptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentOptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 参数校验
		validate := validator.New()
		if err := validate.StructCtx(r.Context(), req); err != nil {
			errMsg := fmt.Sprintf("Parameter verification failed :%v", err)
			var resp = types.CommentOptRes{
				Status: types.Status{
					Code: xerr.ERR,
					Msg:  errMsg,
				},
			}
			httpx.ErrorCtx(r.Context(), w, resp)
			return
		}

		l := userOpt.NewCommentOptLogic(r.Context(), svcCtx)
		resp, err := l.CommentOpt(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
