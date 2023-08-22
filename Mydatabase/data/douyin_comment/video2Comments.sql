#记录video2的用户评论
#序号为1的用户喜欢的视频id
#存放本用户喜欢的视频
USE douyin_comment;
DROP TABLE IF EXISTS `2`;
CREATE TABLE `2`(
    comment_id BIGINT(20) NOT NULL AUTO_INCREMENT,
    status BOOLEAN NOT NULL,
    user_id BIGINT(20) NOT NULL ,
    content VARCHAR(120) NOT NULL,
    date VARCHAR(30) NOT NULL,
    PRIMARY KEY(comment_id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
#使用status来记录本条评论是否被删除
INSERT INTO `2`(
    comment_id,
    status,
    user_id,
    content,
    date
)VALUES(
    1,
    1,
    2,
    '你们好哇',
    '08-22'
)