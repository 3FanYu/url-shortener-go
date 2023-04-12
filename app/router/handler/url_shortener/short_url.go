package url_shortener

import (
	"context"
	"fmt"
	"time"

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
	curTime := time.Now()
	shortUrl := &url_shortener.ShortUrl{
		Url:       url,
		CreatedAt: curTime,
		UpdatedAt: curTime,
	}

	r, err := h.usecase.FindOrCreate(ctx, shortUrl)
	if err != nil {
		return nil, err
	}
	fmt.Println(r)

	return &pb.CreateShortUrlResp{ShortUrl: convertToPb(addUrlPrefix(r))}, nil
}

func (h *shortUrlHandler) RedirectToShortUrl(ctx context.Context, req *pb.RedirectToShortUrlReq) (*pb.RedirectToShortUrlResp, error) {
	key := req.GetShortUrl()
	shortUrl, err := h.usecase.RedirectToShortUrl(ctx, key)
	if err != nil {
		return nil, err
	}
	header := metadata.Pairs("Location", shortUrl.Url)
	grpc.SendHeader(ctx, header)

	return &pb.RedirectToShortUrlResp{ShortUrl: convertToPb(addUrlPrefix(shortUrl))}, nil
}

// convertToPb converts url_shortener.ShortUrl to pb.ShortUrl
func convertToPb(r *url_shortener.ShortUrl) *pb.ShortUrl {
	return &pb.ShortUrl{
		ShortUrl:    r.Key,
		OriginalUrl: r.Url,
	}
}

// addUrlPrefix adds url prefix to shortUrl
func addUrlPrefix(shortUrl *url_shortener.ShortUrl) *url_shortener.ShortUrl {
	shortUrl.Key = "http://localhost:3000/" + shortUrl.Key
	return shortUrl
}
