package url_shortener

import (
	"context"

	pb "github.com/3FanYu/url-shortener-go/proto/short_url"
	uc "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
)

type shortUrlHanlder struct {
	*pb.UnimplementedUrlShortenerServer

	usecase uc.ShortUrlUsecase
}

func NewHandler(usecase uc.ShortUrlUsecase) pb.UrlShortenerServer {
	return &shortUrlHanlder{usecase: usecase}
}

func (h *shortUrlHanlder) CreateShortUrl(ctx context.Context, req *pb.CreateShortUrlReq) (*pb.CreateShortUrlResp, error) {
	url := req.GetUrl()

	err := h.usecase.Create(ctx, url)
	if err != nil {
		return nil, err
	}

	return &pb.CreateShortUrlResp{}, nil
}
