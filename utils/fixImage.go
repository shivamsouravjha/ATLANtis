package utils

import (
	"context"
	"strings"
)

func UpdateAddimage(ctx context.Context, thumbnailUrl string) string {

	if strings.Contains(thumbnailUrl, "cdn.trell.co") {
		splitArray := strings.Split(thumbnailUrl, "/")
		thumbnail := splitArray[len(splitArray)-1]
		return thumbnail
	}
	return thumbnailUrl
}

func UpdateGetimage(ctx context.Context, thumbnail string) string {

	if !strings.Contains(thumbnail, "http") {
		thumbnailUrl := "https://cdn.trell.co/images/orig/" + thumbnail
		return thumbnailUrl
	}
	return thumbnail
}
func UpdateAvatar(ctx context.Context, thumbnailUrl string) string {

	if strings.Contains(thumbnailUrl, "cdn.trell.co") {
		splitArray := strings.Split(thumbnailUrl, "cdn.trell.co/")
		thumbnail := splitArray[len(splitArray)-1]
		return thumbnail
	}
	return thumbnailUrl
}

func GetAvatar(ctx context.Context, thumbnail string) string {

	if !strings.Contains(thumbnail, "http") && len(thumbnail) != 0 {
		thumbnailUrl := "https://cdn.trell.co/" + thumbnail
		return thumbnailUrl
	}
	return thumbnail
}
