package publish

import (
	"douyin/pkg/gateway/internal/logic/publish"
	"douyin/pkg/gateway/internal/svc"
	"douyin/pkg/gateway/internal/types"
	"errors"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PubVideoReq

		if !r.Form.Has("title") || !r.Form.Has("token") || r.MultipartForm.File == nil || r.MultipartForm.File["data"] == nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("missing parameter title,token or file"))
			return
		}
		req.Token = r.Form.Get("token")
		req.Title = r.Form.Get("title")

		file, err := r.MultipartForm.File["data"][0].Open()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Data, err = io.ReadAll(file)
		if err != nil {
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
