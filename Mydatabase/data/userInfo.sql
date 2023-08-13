#存储用户的所有信息
USE douyin_info;
DROP TABLE IF EXISTS userInfo;
CREATE TABLE userInfo(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    name VARCHAR(20) DEFAULT '',
    follow_count BIGINT DEFAULT '0',
    follower_count BIGINT DEFAULT '0',
    avator VARCHAR(30) DEFAULT '',
    background_image VARCHAR(30) DEFAULT '',
    signature VARCHAR(120) DEFAULT '',
    total_favorited BIGINT(20) DEFAULT '0',
    work_count BIGINT(20) DEFAULT '0',
    favorite_count BIGINT(20) DEFAULT '0',
    PRIMARY KEY(id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
INSERT INTO userInfo(name) VALUES('王五')