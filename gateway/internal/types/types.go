// Code generated by goctl. DO NOT EDIT.
package types

type PingResp struct {
	Resp
	Msg string `json:"msg"`
}

type Resp struct {
	Code int32  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResp struct {
	Resp
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterResp struct {
	Resp
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type PublishActionReq struct {
	Token string `form:"token"`
	Title string `form:"title"`
}

type FavoriteActionReq struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"` // 1: favorite, 2: unfavorite
}

type FavoriteActionResp struct {
	Resp
}

type FavoriteListReq struct {
	Token  string `form:"token"`
	UserId int64  `form:"user_id"`
}

type FovoriteListResp struct {
	Resp
	VideoList []Video `json:"video_list"`
}

type CommentActionReq struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentActionResp struct {
	Resp
	Comment
}

type CommentListReq struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
}

type CommentListResp struct {
	Resp
	CommentList []Comment `json:"comment_list"`
}

type PublishListReq struct {
	Token  string `form:"token"`
	UserId int64  `form:"user_id"`
}

type PublishListResp struct {
	Resp
	VideoList []Video `json:"video_list"`
}

type FeedReq struct {
	Token      string `form:"token"`
	LatestTime int64  `form:"latest_time,optional"`
}

type FeedResp struct {
	Resp
	VideoList []Video `json:"video_list"`
}

type UserReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserResp struct {
	Resp
	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	IsFollow      bool  `json:"is_follow"`
}
