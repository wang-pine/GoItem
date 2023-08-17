#存储用户的所有信息
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
INSERT INTO userInfo(
    id,
    name,
    follow_count,
    follower_count,
    avator,
    background_image,
    signature,
    total_favorited,
    work_count,
    favorite_count) VALUES(
        1,
        '王五',
        0,
        1,
        'http://192.168.3.10:8888/static/A_1.jpg',
        'http://192.168.3.10:8888/static/2.jpg',
        '好好学习天天向上',
        1000,
        2,
        0);
INSERT INTO userInfo(
    id,
    name,
    follow_count,
    follower_count,
    avator,
    background_image,
    signature,
    total_favorited,
    work_count,
    favorite_count) VALUES(
        2,
        '张三',
        0,
        1,
        'http://192.168.3.10:8888/static/A_2.jpg',
        'http://192.168.3.10:8888/static/1.jpg',
        '抖音Daisiki~~~',
        100000,
        1,
        0);