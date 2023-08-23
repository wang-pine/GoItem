#序号为2的视频被点赞的统计数据
#存放本视频的点赞用户
USE douyin_videos;
DROP TABLE IF EXISTS `2`;
CREATE TABLE `2`(
    favorite_user_id BIGINT(20) NOT NULL,
    video_id BIGINT(20) NOT NULL ,
    is_delete int(1) NOT NULL default 0
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
INSERT INTO `2`(
    favorite_user_id,
    video_id
)VALUES(
    1,
    2
)