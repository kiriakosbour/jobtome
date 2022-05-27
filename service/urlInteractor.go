package service

import (
	"log"
	"time"
	"urlShortener/domain"
)

type UrlInteractor struct {
	repo domain.UrlRepoInt
}

func UrlInteractorInit(repository domain.UrlRepoInt) *UrlInteractor {
	return &UrlInteractor{repo: repository}
}

type UrlInteractorInt interface {
	CreateTheUrlShortingService(longUrl string, userId string) string
	RetrieveTheUrlShortingService(shortUrl string) string
	RetrieveAllTheUrlsSService() int
}

func (u *UrlInteractor) CreateTheUrlShortingService(longUrl string, userId string) string {
	var urlDomain domain.Urls
	shortLink := urlDomain.GenerateShortLink(longUrl, userId)
	urlDomain.OriginalUrl = longUrl
	urlDomain.ShortUrl = shortLink
	urlDomain.UserId = userId
	respFromRedis, err := u.repo.SetKey(urlDomain, 6*time.Hour)
	if err != nil {
		return ""
	}
	return respFromRedis.ShortUrl
}
func (u *UrlInteractor) RetrieveTheUrlShortingService(shortUrl string) string {
	var urlDomain domain.Urls
	urlDomain.ShortUrl = shortUrl
	respFromRedis, err := u.repo.GetKey(urlDomain)
	log.Print(shortUrl)
	if err != nil {
		return ""
	}
	return respFromRedis.OriginalUrl
}
func (u *UrlInteractor) RetrieveAllTheUrlsSService() int {

	mapOfUrls := u.repo.GetAllValues()

	return len(mapOfUrls)
}
