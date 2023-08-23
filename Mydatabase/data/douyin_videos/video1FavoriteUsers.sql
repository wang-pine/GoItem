#序号为1的视频被点赞的统计数据
#存放本视频的点赞用户
USE douyin_videos;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    favorite_user_id BIGINT(20) NOT NULL,
    video_id BIGINT(20) NOT NULL ,
    is_delete int(1) NOT NULL default 0
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(
    favorite_user_id,
    video_id
)VALUES(
    2,
    1
)