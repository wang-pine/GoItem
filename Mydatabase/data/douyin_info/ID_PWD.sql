#存储基本的用户的账号+密码
USE douyin_info;
DROP TABLE IF EXISTS ID_PWD;
CREATE TABLE ID_PWD(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    PWD VARCHAR(32) DEFAULT '',
    PRIMARY KEY(id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
#用户1密码wwss123
INSERT INTO ID_PWD(PWD) VALUES('3a40ce100855d88d4c6b194770b8c859');
#用户2密码123456
INSERT INTO ID_PWD(PWD) VALUES('f447b20a7fcbf53a5d5be013ea0b15af');