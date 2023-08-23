#用户1的关注列表
USE douyin_follow;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    follow_id BIGINT(20) NOT NULL ,
    user_id BIGINT(20) NOT NULL,
    PRIMARY KEY(follow_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(
    follow_id,
    user_id
)VALUES(
    2,
    1
)