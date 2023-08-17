#存储视频的所有信息
USE douyin_info;
DROP TABLE IF EXISTS videoInfo;
CREATE TABLE videoInfo(
    video_id BIGINT(20) NOT NULL AUTO_INCREMENT,
    author_id BIGINT(20) NOT NULL,
    author_name VARCHAR(20) DEFAULT '',
    author_follow_count BIGINT DEFAULT '0',
    author_follower_count BIGINT DEFAULT '0',
    author_avator VARCHAR(60) DEFAULT '',
    author_background_image VARCHAR(60) DEFAULT '',
    author_signature VARCHAR(120) DEFAULT '',
    author_total_favorited BIGINT(20) DEFAULT '0',
    author_work_count BIGINT(20) DEFAULT '0',
    author_favorite_count BIGINT(20) DEFAULT '0',
    video_play_url VARCHAR(60) DEFAULT '',
    video_cover_url VARCHAR(60) DEFAULT '',
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
    'http://192.168.3.10:8888/static/sohu.mp4',
    'http://192.168.3.10:8888/static/1.jpg',
    '哈哈哈哈',
    '202308122243'
);
INSERT INTO videoInfo(
    author_id,
    video_play_url,
    video_cover_url,
    video_title,
    video_time
)VALUES(
    1,
    'http://192.168.3.10:8888/static/光年之外.mp4',
    'http://192.168.3.10:8888/static/2.jpg',
    '哈哈哈哈',
    '202308122243'
);
INSERT INTO videoInfo(
    author_id,
    video_play_url,
    video_cover_url,
    video_title,
    video_time
)VALUES(
    2,
    'http://192.168.3.10:8888/static/bear.mp4',
    'http://192.168.3.10:8888/static/3.jpg',
    '哈哈哈哈',
    '202308122243'
);
