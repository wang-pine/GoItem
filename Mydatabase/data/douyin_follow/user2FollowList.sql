#用户2关注的列表
USE douyin_follow;
DROP TABLE IF EXISTS `2`;
CREATE TABLE `2`(
    follow_id BIGINT(20) NOT NULL ,
    user_id BIGINT(20) NOT NULL,
    PRIMARY KEY(follow_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `2`(
    follow_id,
    user_id
)VALUES(
    1,
    2
)