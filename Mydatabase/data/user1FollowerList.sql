#序号为1的用户的关注列表
#存放本用户的关注用户
USE douyin_followers;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    follower_id BIGINT(20) NOT NULL ,
    user_id BIGINT(20) NOT NULL,
    PRIMARY KEY(follower_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(
    follower_id,
    user_id
)VALUES(
    1,
    1
)