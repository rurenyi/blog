CREATE TABLE IF NOT EXISTS USERS(  
    ID INT AUTO_INCREMENT COMMENT "用户ID",
    USERNAME VARCHAR(20) NOT NULL COMMENT "用户名",
    PASSWORD VARCHAR(32) NOT NULL COMMENT "密码",  
    PRIMARY KEY ( ID ),
    UNIQUE KEY Index_name ( USERNAME )
) DEFAULT CHARSET = utf8mb4 COMMENT "用户表";


CREATE Table if NOT exists blogs(
    BLOG_ID INT AUTO_INCREMENT COMMENT "文章id",
    WRITER_ID INT NOT NULL COMMENT "作者id",
    TITLE VARCHAR(50) NOT NULL COMMENT "文章标题",
    ARTICLE TEXT NOT NULL COMMENT "文章",
    CREATE_TIME DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
    MODIFIED_NAME DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "修改时间",
    PRIMARY KEY ( BLOG_ID ),
    KEY idx_blog ( WRITER_ID )
) DEFAULT CHARSET = utf8mb4 COMMENT "博客";