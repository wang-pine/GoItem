#序号为1的用户进行的个人视频列表的测试
#存放本用户视频列表
#这个仅仅作为douyin_users数据库的样例
USE douyin_users;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    user_id BIGINT(20) NOT NULL AUTO_INCREMENT,
    video_id BIGINT(20) NOT NULL,
    PRIMARY KEY(user_id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(
    user_id,
    video_id
)VALUES(
    1,
    1
)