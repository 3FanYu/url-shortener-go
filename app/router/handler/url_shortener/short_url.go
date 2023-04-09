package url_shortener

import (
	"context"

	"github.com/3FanYu/url-shortener-go/domain/url_shortener"
	pb "github.com/3FanYu/url-shortener-go/proto/short_url"
	uc "github.com/3FanYu/url-shortener-go/usecase/url_shortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type shortUrlHandler struct {
	*pb.UnimplementedUrlShortenerServer

	usecase uc.ShortUrlUsecase
}

func NewHandler(usecase uc.ShortUrlUsecase) pb.UrlShortenerServer {
	return &shortUrlHandler{usecase: usecase}
}

func (h *shortUrlHandler) CreateShortUrl(ctx context.Context, req *pb.CreateShortUrlReq) (*pb.CreateShortUrlResp, error) {
	url := req.GetUrl()
	shortUrl := &url_shortener.ShortUrl{Url: url}

	_, err := h.usecase.Create(ctx, shortUrl)
	if err != nil {
		return nil, err
	}

	return &pb.CreateShortUrlResp{ShortUrl: convertToPb(shortUrl)}, nil
}

func (h *shortUrlHandler) RedirectToShortUrl(ctx context.Context, req *pb.RedirectToShortUrlReq) (*pb.RedirectToShortUrlResp, error) {
	key := req.GetShortUrl()
	shortUrl, err := h.usecase.RedirectToShortUrl(ctx, key)
	if err != nil {
		return nil, err
	}
	header := metadata.Pairs("Location", shortUrl.Url)
	grpc.SendHeader(ctx, header)

	return &pb.RedirectToShortUrlResp{ShortUrl: convertToPb(shortUrl)}, nil
}

// convertToPb converts url_shortener.ShortUrl to pb.ShortUrl
func convertToPb(r *url_shortener.ShortUrl) *pb.ShortUrl {
	return &pb.ShortUrl{
		ShortUrl:    r.Key,
		OriginalUrl: r.Url,
	}
}
