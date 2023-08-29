########################
#本文件用于数据库的初始化#
########################
#使用root，创建用户douyin
create user douyin identified by '123456';
#创建douyin_info数据库
create database douyin_info;
#创建douyin_users数据库
create database douyin_users;
#创建douyin_videos数据库
create database douyin_videos;
#创建douyin_followers数据库
create database douyin_followers;
#创建douyin_favorite数据库
create database douyin_favorite;
#创建douyin_favorite数据库
create database douyin_comment;
#创建douyin_message数据库
create database douyin_message;
#创建douyin_folllow数据库

#创建用户数据库并授权
create database douyin_follow;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE ON douyin_info.* TO douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE ON douyin_users.* to douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE ON douyin_videos.* to douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE,DELETE ON douyin_followers.* TO douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE ON douyin_favorite.* TO douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE ON douyin_comment.* TO douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE ON douyin_message.* TO douyin;
GRANT CREATE,ALTER,INSERT,SELECT,UPDATE,DELETE ON douyin_follow.* TO douyin；

#创建账号密码数据库
USE douyin_info;
DROP TABLE IF EXISTS ID_PWD;
CREATE TABLE ID_PWD(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    PWD VARCHAR(32) DEFAULT '',
    PRIMARY KEY(id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

#创建用户信息数据库
USE douyin_info;
DROP TABLE IF EXISTS userInfo;
CREATE TABLE userInfo(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) DEFAULT '',
    follow_count BIGINT DEFAULT '0',
    follower_count BIGINT DEFAULT '0',
    avator VARCHAR(60) DEFAULT '',
    background_image VARCHAR(60) DEFAULT '',
    signature VARCHAR(120) DEFAULT '',
    total_favorited BIGINT(20) DEFAULT '0',
    work_count BIGINT(20) DEFAULT '0',
    favorite_count BIGINT(20) DEFAULT '0',
    PRIMARY KEY(id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

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
