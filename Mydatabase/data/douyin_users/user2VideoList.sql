#序号为2的用户进行的个人视频列表的测试
#存放本用户视频列表
#这个仅仅作为douyin_users数据库的样例
USE douyin_users;
DROP TABLE IF EXISTS `2`;
CREATE TABLE `2`(
    video_id BIGINT(20) NOT NULL,
    user_id BIGINT(20) NOT NULL,
    PRIMARY KEY(video_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `2`(
    video_id,
    user_id
)VALUES(
    3,
    2
)