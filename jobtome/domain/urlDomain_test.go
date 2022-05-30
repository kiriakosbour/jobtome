package domain

import (
	"testing"
)

func TestShortLinkGenerator(t *testing.T) {
	var domain Urls
	initialLink_1 := "https://www.marvel.com/unlimited?cid=SEM_Google_20200302_unlimited_Brand&gclid=CjwKCAjw7cGUBhA9EiwArBAvoqcoCF_lveSuGwn0e5Gz-H_kv4escI-5SYse_2Taw4fyqxD8SqTgrBoCwTwQAvD_BwE"
	shortLink_1 := domain.GenerateShortLink(initialLink_1)

	initialLink_2 := "https://www.dccomics.com/videos/justice-league-batman-comic-book-origins"
	shortLink_2 := domain.GenerateShortLink(initialLink_2)

	initialLink_3 := "https://www.starwars.com/news/category/forces-of-destiny"
	shortLink_3 := domain.GenerateShortLink(initialLink_3)
	if shortLink_1 != "ikuNMe7z" {
		t.Errorf("got %s,wanted %s", shortLink_1, "ikuNMe7z")
	}
	if shortLink_2 != "TuYva5qP" {
		t.Errorf("got %s,wanted %s", shortLink_2, "TuYva5qP")
	}
	if shortLink_3 != "cHDEioWh" {
		t.Errorf("got %s,wanted %s", shortLink_2, "cHDEioWh")
	}
}
