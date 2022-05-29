package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreatorPost struct {
	Id        int64  `json:"id"`
	CreatorId int    `json:"creatorId"`
	PostTitle string `json:"postTitle"`
	PostUrl   string `json:"postUrl"`
	Thumbnail string `json:"thumbnail"`
	Position  uint64 `json:"position"`
	IsDeleted bool   `json:"isDeleted"`
}

type CreatorDetails struct {
	UserId             int                    `form:"userId" binding:"required"`
	UserName           string                 `form:"userName" binding:"required"`
	HandleName         string                 `form:"handleName" binding:"required"`
	UserAvatar         string                 `form:"userAvatar" binding:"required"`
	UserEmail          string                 `form:"userEmail" binding:"required"`
	Bio                string                 `form:"bio" binding:"required"`
	Status             int                    `form:"status" binding:"required"`
	SocialMediaHandles map[string]interface{} `form:"socialMediaHandles" binding:"required"`
	Categories         string                 `form:"categories" binding:"required"`
}

type EventDocument struct {
	EventType string             `json:"eventType"`
	UserId    int                `json:"userId"`
	PostId    int                `json:"postId"`
	TimeStamp int32              `json:"timeStamp"`
	ObjectId  primitive.ObjectID `bson:"_id" json:"objectId"`
}
type ClicksEventAnalytics struct {
	TotalClicks        int32       `json:"totalClicks"`
	PostAnalytics      map[int]int `json:"postAnalytics"`
	HighestClicksInDay int32       `json:"highestClicksInDay"`
}

type ViewsEventAnalytics struct {
	TotalViews         int32              `json:"totalViews"`
	GraphPoints        map[string]float64 `json:"GraphPoints"`
	HighestClicksInDay int32              `json:"highestViewsInDay"`
}

type EmailDetails struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
}
