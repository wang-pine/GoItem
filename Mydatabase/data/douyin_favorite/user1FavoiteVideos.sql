#序号为1的用户喜欢的视频id
#存放本用户喜欢的视频
USE douyin_favorite;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    favorite_video_id BIGINT(20) NOT NULL,
    user_id BIGINT(20) NOT NULL ,
    PRIMARY KEY(favorite_video_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(
    favorite_video_id,
    user_id
)VALUES(
    2,
    1
)