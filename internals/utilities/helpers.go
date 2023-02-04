package Utility

import (
	"github.com/google/uuid"
	"shortener/internals/config"
)

func ValidMainLink(link string) string {
	if link[:7] != "http://" && link[:8] != "https://" {
		return "https://" + link
	}
	return link
}

func GenerateShortLink() string {
	return config.AppConfig.Domain + "/s/" + uuid.New().String()[:6]
}
