#存储视频的所有信息
USE douyin_info;
DROP TABLE IF EXISTS videoInfo;
CREATE TABLE videoInfo(
    video_id BIGINT(20) NOT NULL AUTO_INCREMENT,
    author_id BIGINT(20) NOT NULL,
    author_name VARCHAR(20) DEFAULT '',
    author_follow_count BIGINT DEFAULT '0',
    author_follower_count BIGINT DEFAULT '0',
    author_avator VARCHAR(30) DEFAULT '',
    author_background_image VARCHAR(30) DEFAULT '',
    author_signature VARCHAR(120) DEFAULT '',
    author_total_favorited BIGINT(20) DEFAULT '0',
    author_work_count BIGINT(20) DEFAULT '0',
    author_favorite_count BIGINT(20) DEFAULT '0',
    video_play_url VARCHAR(30) DEFAULT '',
    video_cover_url VARCHAR(30) DEFAULT '',
    video_favorite_count BIGINT(20) DEFAULT '0',
    video_comment_count BIGINT(20) DEFAULT '0',
    video_title VARCHAR(30) DEFAULT '',
    video_time VARCHAR(30) DEFAULT '',
    PRIMARY KEY(video_id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
INSERT INTO videoInfo(
    author_id,
    video_play_url,
    video_cover_url,
    video_title,
    video_time
)VALUES(
    1,
    'here',
    'there',
    '哈哈哈哈',
    '202308122243'
)
#备注:我觉的这么设置时间戳是不安全的。这个需要想想办法，这样不好
#应该做一个约束，不能这么随意！！！！