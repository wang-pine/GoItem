#序号为2的用户喜欢的视频id
#存放本用户喜欢的视频
USE douyin_favorite;
DROP TABLE IF EXISTS `2`;
CREATE TABLE `2`(
    favorite_video_id BIGINT(20) NOT NULL,
    user_id BIGINT(20) NOT NULL ,
    PRIMARY KEY(favorite_video_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
INSERT INTO `2`(
    favorite_video_id,
    user_id
)VALUES(
    1,
    2
)