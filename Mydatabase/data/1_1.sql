#序号为1的视频被点赞的统计数据
#存放本视频的点赞用户
USE douyin_videos;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    video_id BIGINT(20) NOT NULL AUTO_INCREMENT,
    favorite_user_id BIGINT(20) NOT NULL,
    PRIMARY KEY(video_id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(
    video_id,
    favorite_user_id
)VALUES(
    1,
    1
)