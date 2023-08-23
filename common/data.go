
package common

// 这个文件是数据库使用的结构体
type Videoinfo struct {
	VideoId               int64  `gorm:"type:int(20); not null" json:"video_id" binding:"required"`
	AuthorId              int64  `gorm:"type:int(20); not null" json:"author_id" binding:"required"`
	AuthorName            string `gorm:"type:varchar(20); not null" json:"author_name" binding:"required"`
	AuthorFollowCount     int64  `gorm:"type:int(20); not null" json:"author_follow_count" binding:"required"`
	AuthorFollowerCount   int64  `gorm:"type:int(20); not null" json:"author_follower_count" binding:"required"`
	AuthorAvator          string `gorm:"type:varchar(20); not null" json:"author_avator" binding:"required"`
	AuthorBackgroundImage string `gorm:"type:varchar(120); not null" json:"author_background_image" binding:"required"`
	AuthorSignature       string `gorm:"type:varchar(120); not null" json:"author_signature" binding:"required"`
	AuthorTotalFavorited  int64  `gorm:"type:int(20); not null" json:"author_total_favorited" binding:"required"`
	AuthorWorkCount       int64  `gorm:"type:int(20); not null" json:"author_work_count" binding:"required"`
	AuthorFavoriteCount   int64  `gorm:"type:int(20); not null" json:"author_favorite_count" binding:"required"`
	VideoPlayUrl          string `gorm:"type:varchar(120); not null" json:"video_play_url" binding:"required"`
	VideoCoverUrl         string `gorm:"type:varchar(120); not null" json:"video_cover_url" binding:"required"`
	VideoFavoriteCount    int64  `gorm:"type:int(20); not null" json:"video_favorite_count" binding:"required"`
	VideoCommentCount     int64  `gorm:"type:int(20); not null" json:"video_comment_count" binding:"required"`
	VideoTitle            string `gorm:"type:varchar(30); not null" json:"video_title" binding:"required"`
	VideoTime             string `gorm:"type:varchar(30); not null" json:"video_time" binding:"required"`
}
type Userinfo struct {
	Id              int64  `gorm:"type:varchar(20); not null" json:"id" binding:"required"`
	Name            string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
	FollowCount     int64  `gorm:"type:int(64); not null" json:"follow_count" binding:"required"`
	FollowerCount   int64  `gorm:"type:int(64); not null" json:"follower_count" binding:"required"`
	Avator          string `gorm:"type:varchar(256); not null" json:"avator" binding:"required"`
	BackgroundImage string `gorm:"type:varchar(256); not null" json:"background_image" binding:"required"`
	Signature       string `gorm:"type:varchar(64); not null" json:"signature" binding:"required"`
	TotalFavorited  int64  `gorm:"type:int(20); not null" json:"total_favorited" binding:"required"`
	WorkCount       int64  `gorm:"type:int(20); not null" json:"work_count" binding:"required"`
	FavoriteCount   int64  `gorm:"type:int(20); not null" json:"favorite_count" binding:"required"`
}

type Pare struct {
	VideoId  int64 `json:"video_id"`
	UserId   int64 `json:"user_id"`
	IsDelete int64 `json:"is_delete"`
}
