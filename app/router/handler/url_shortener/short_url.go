package url_shortener

import (
	"context"

	"github.com/3FanYu/url-shortener-go/domain/url_shortener"
	pb "github.com/3FanYu/url-shortener-go/proto/short_url"
	uc "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
)

type shortUrlHanlder struct {
	*pb.UnimplementedShortUrlServer

	usecase uc.ShortUrlUsecase
}

func NewUrlHandler(usecase uc.ShortUrlUsecase) pb.ShortUrlServer {
	return &shortUrlHanlder{usecase: usecase}
}

func (h *shortUrlHanlder) CreateShortUrl(ctx context.Context, req *pb.CreateShortUrlReq) (*pb.CreateShortUrlResp, error) {
	url := req.GetUrl()

	err := h.usecase.Create(ctx, url)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{}, nil
}
