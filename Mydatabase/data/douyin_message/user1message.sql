#用户1的收发信息数据库
use douyin_message;
DROP TABLE IF EXISTS `1`;
CREATE TABLE `1`(
    message_id BIGINT(20) NOT NULL AUTO_INCREMENT,
    to_user_id BIGINT(20) NOT NULL,
    from_user_id BIGINT(20) NOT NULL,
    content VARCHAR(120) NOT NULL,
    create_time VARCHAR(60)NOT NULL,
    PRIMARY KEY(message_id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
INSERT INTO `1`(message_id,
                to_user_id,
                from_user_id,
                content,
                create_time) VALUES(
                1,
                1,
                2,
                '忽忽嘿嘿欸嘿',
                '2023-08-01 13:33::33'
                );
INSERT INTO `1`(message_id,
                to_user_id,
                from_user_id,
                content,
                create_time) VALUES(
                2,
                2,
                1,
                'xixixixixi',
                '2023-08-01 13:35::33'
                );