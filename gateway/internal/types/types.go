// Code generated by goctl. DO NOT EDIT.
package types

type IndexResp struct {
	Resp
	Github    string   `json:"Github"`
	Author    []string `json:"Author"`
	Timestamp string   `json:"Timestamp"`
}

type Resp struct {
	Code int32  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
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

type FriendUser struct {
	User
	Message string `json:"message"`
	MsgType int64  `json:"msgType"`
}

type Message struct {
	Id         int64  `json:"id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	ToUserId   int64  `json:"to_user_id"`
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
	ActionType int32  `form:"action_type,range=[1:2]"` // 1: favorite, 2: unfavorite
}

type FavoriteActionResp struct {
	Resp
}

type FavoriteListReq struct {
	Token  string `form:"token,optional"`
	UserId int64  `form:"user_id"`
}

type FovoriteListResp struct {
	Resp
	VideoList []Video `json:"video_list"`
}

type CommentActionReq struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type,range=[1:2]"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id,optional"`
}

type CommentActionResp struct {
	Resp
	Comment Comment `json:"comment"`
}

type CommentListReq struct {
	Token   string `form:"token,optional"`
	VideoId int64  `form:"video_id"`
}

type CommentListResp struct {
	Resp
	CommentList []Comment `json:"comment_list"`
}

type PublishListReq struct {
	Token  string `form:"token,optional"`
	UserId int64  `form:"user_id"`
}

type PublishListResp struct {
	Resp
	VideoList []Video `json:"video_list"`
}

type FeedReq struct {
	Token      string `form:"token,optional"`
	LatestTime int64  `form:"latest_time,optional"`
}

type FeedResp struct {
	Resp
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time,optional"`
}

type UserReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token,optional"`
}

type UserResp struct {
	Resp
	User User `json:"user"`
}

type RelationActionReq struct {
	Token     string `form:"token"`
	ToUserId  int64  `form:"to_user_id"`
	ActionTyp int32  `form:"action_type,range=[1:2]"`
}

type RelationActionResp struct {
	Resp
}

type FollowerListReq struct {
	Token  string `form:"token,optional"`
	UserId int64  `form:"user_id"`
}

type FollowerListResp struct {
	Resp
	UserList []User `json:"user_list"`
}

type FansListReq struct {
	Token  string `form:"token,optional"`
	UserId int64  `form:"user_id"`
}

type FansListResp struct {
	Resp
	UserList []User `json:"user_list"`
}

type FriendListReq struct {
	Token  string `form:"token,optional"`
	UserId int64  `form:"user_id""` // ! can only be self
}

type FriendListResp struct {
	Resp
	UserList []FriendUser `json:"user_list"`
}

type ChatActionReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
	Content    string `form:"content,optional"`
}

type ChatActionResp struct {
	Resp
}

type ChatMessageReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	PreMsgTime int64  `form:"pre_msg_time,optional"`
}

type ChatMessageResp struct {
	Resp
	MessageList []Message `json:"message_list"`
}
