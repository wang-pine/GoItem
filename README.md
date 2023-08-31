# GoItem

just for byte dance Go-Project

# 仓库使用须知

1. main分支禁止提交
2. 想要对仓库进行更新，一定要先pull main分支
3. 修改完源代码之后，请在commit信息中加入个人信息，方便后续统计贡献度
4. 对源码进行修改操作之后，重新push回仓库之前，请先checkout到本地dev分支，然后进行一波add，commit操作，再push到github上的dev分支，最后请去github上提一个pull requests，等待合并入main分支（当然了，你也可以自行合并，没有设置限制，但是最好等待别人测试完之后对你的merge进行同意最好）

# 一、项目介绍

> 项目信息：极简版的抖音app的服务端，实现了要求的所有接口（第六届字节跳动青训营，2023暑假）
>
> 服务地址：https://1024code.com/~codecubes/mfmjglx
>
> 项目地址：https://github.com/wang-pine/GoItem
>
> 项目小组：go小分队（060120）
>
> 小组成员：王松（队长），乔旭东，戴学微，郑照鸿

# 二、项目分工

* 主要任务及分工

| 主要内容     | 实现者                                                                                |
| ------------ | ------------------------------------------------------------------------------------- |
| 项目基础架构 | 王松                                                                                  |
| mysql封装    | 王松，戴学微                                                                          |
| redis封装    | 戴学微                                                                                |
| 基本api      | 视频流（王松）注册登录（乔旭东+王松）用户信息（郑照鸿）投稿（戴学微）发布列表（王松） |
| 进阶api-I    | 赞+赞列表（戴学微）评论+评论列表（王松）                                              |
| 进阶api-II   | 关注关注列表+粉丝列表+好友列表（王松）发送消息+消息列表（郑照鸿）                     |
| 测试         | 王松，戴学微，郑照鸿                                                                  |
| 文档         | 王松                                                                                  |

* 主要贡献成员及github地址

| 王松   | **https://github.com/wang-pine**      |
| ------ | ------------------------------------------- |
| 戴学微 | **https://github.com/dxw-1997**       |
| 郑照鸿 | **https://github.com/Koreyer**        |
| 乔旭东 | **https://github.com/NothingEvering** |

# 三、项目实现

### 3.1 技术选型

> 抖音，是一个日活用户极多，数据量极其之大的app，为了应付海量的用户和数据，我们对其服务端进行了细致且缜密的分析。我们认为，用户量如此大的app，其数据库查询的开销是巨大的，为了应付数据库的开销，我们决定采用“存储过程”，“分库分表”，“小表驱动大表”和“建立索引”的做法来减轻查表产生的数据库消耗，针对需要大量查表的操作，如向数据库请求该用户的上传过视频的表，我们在对应的用户上传的视频库中存储了该用户累计上传的视频id，并以此作为主键，在快速查询完用户上传过后的视频id之后到视频总表中，以id依次查询用户视频的详细信息，由于查询的都是主键，速度会非常快，我们依照这种方式，通过分库分表，小表驱动大表，建立查询索引等等方式，有效的减少了用户在查询过程中的时间消耗，牺牲了一定的存储空间，但是考虑用户的体验才是第一优先级，节省了大量的服务器cpu开销。

1. 开发语言与基本框架： golang和gin。使用gin框架可以较为轻松的配置项目的路由信息
2. 数据库：redis和mysql，使用redis作为token服务，mysql用于存取基本信息
3. 配置热重载：viper。这个是为了让项目可以快速迁移，同时可以动态迁移，热迁移等等，且能够让项目的配置便于及时更新
4. golang与数据库通信：mysql-driver和gorm，充分考虑到mysql注入等等问题，并且在数据库权限的方面，douyin这个数据库用户只有最少量的权限，其中大量的权限如删除采用伪删除的策略，在删除上置一个tag，在读取的时候做tag的判断，有效防止即便被sql注入成功，也能保护表信息不受破坏，除非对方提权到root
5. 使用存储过程，数据库进一层封装，开发效率高，数据库相关操作不容易发生问题，且当数据库的相关函数测试完成后可以直接投入到服务端的使用当中，如果程序崩溃可以直接排除数据库操作的问题

### 3.2 架构设计

#### 3.2.1基本架构

![](https://bo9uwdfynq.feishu.cn/space/api/box/stream/download/asynccode/?code=NTAyNzMxOTMwYmVkNzQ5NTkzZmFiZDdhZGNjNDE3MjhfcFdJUHhad25SeUN5ZDZwaEtnYUJjSTBxUUhCQmxmbjZfVG9rZW46WWo3RWJjNlhJbzFJU2h4b2hYNGNwdWhWbmNmXzE2OTM0ODc0NTY6MTY5MzQ5MTA1Nl9WNA)

使用mysql存储用户基本信息，使用redis存储服务端生成的token

#### 3.2.2.架构分析

考虑到抖音的用户量非常大，任何小的查询操作都会造成巨大的数据库负载

我们采用了“分库分表”，“小表驱动大表”以及“建立索引”的方式减轻查表的负担，并且将大部分的查询操作进行了单用户分表，建立了大量的数据库和数据表，能够有效的减轻查表的负担

在此基础上，我们将数据库go-mysql进一步封装，使用存储过程

开发的过程中只需要调用封装好的数据库函数即可实现

#### 3.3.3.分表概览

![](https://bo9uwdfynq.feishu.cn/space/api/box/stream/download/asynccode/?code=YWIzYzM5ZDExNjM4MDgxMzMzNjhmZTliNDM5NmZmY2FfNFRlM29wU2t0a1lYQXYxZlRnNUljaFA0M2xCRmVxZ2dfVG9rZW46QllRTWJLaFdwb09Sd1h4YWQza2MyYzRybk1lXzE2OTM0ODc0NTY6MTY5MzQ5MTA1Nl9WNA)

为了提高效率，减少查表对服务器造成大量的开销，我们着力在数据库的分库分表上

1. 核心：douyin_info：存储了综合的总数据库，包括了三个表：id_pwd（用户的账号密码）,userinfo（用户的基本信息），videoinfo（视频的基本信息）。存储了用户和视频的基本信息
2. douyin_followers：关注者数据库，针对每个用户id进行分表，记录了每一个用户的关注者，主键是关注者的id
3. douyin_follow：用户的关注列表，针对每个用户id进行分表，主键是用户关注的人的id
4. douyin_user：每个用户的发布视频的分表，针对每个用户id进行分表，保存了每个用户发布的视频，主键是用户发布视频的id
5. douyin_videos：每个视频的点赞信息，对每一个视频id进行分表，记录了每一个视频中点赞用户的id，主键是点赞的用户id
6. douyin_comment：每个视频下的用户评论，对每一个视频id进行分表，以评论的id为主键，id的排序依照时间顺序，所以id可以代替时间戳使用
7. douyin_message：用户的信息数据，针对每个用户id进行分表，以信息的id为主键，每次进行信息数据库的维护都是同时对收发双方进行维护的
8. douyin_favorite：每个用户的点赞视频的id，对每个用户id进行分表，记录了用户点赞的视频的id，主键是点赞用户的id

### 3.3 项目代码样例

> 仅展示具有特征性质的样例代码，具体详细信息请阅读项目源代码

#### 3.3.1.mysql封装

简要的介绍了mysql封装的代码，介绍了存储用户投稿视频的代码，在mysql基础上使用对象存储进一步的封装

```Go
package Mydatabase
/*
********************
存储用户投稿的视频
********************
*/
import (
    "config"
    "database/sql"
    "fmt"
    "strconv"
    _ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbUsers *sql.DB
// 这里用来对单个用户的分表进行维护
// 单个用户的分表存放的是该用户上传的视频
func InitUsersDatabase() (err error) {
    fmt.Printf("正在初始化用户视频列表数据库...\n")
    dsn := "douyin:123456@tcp(" + config.GetDBAddr() + ")/douyin_users"
    dbUsers, err = sql.Open("mysql", dsn)
    //open函数是不会检查用户名和密码的
    if err != nil {
        return
    }
    err = dbUsers.Ping() //尝试对数据库进行链接
    if err != nil {
        return
    }
    fmt.Println("链接数据库成功")
    dbUsers.SetMaxIdleConns(100)
    //设置数据库连接池的最大连接数
    return
}

// 根据用户的id创建每个用户的分表
func MakeNewUserTable(id int64) (err error) {
    InitUsersDatabase()
    sqlStr := "CREATE TABLE `" + strconv.FormatInt(id, 10) + "`(" +
        "video_id BIGINT(20) NOT NULL," +
        "user_id BIGINT(20) NOT NULL," +
        "PRIMARY KEY(video_id)" +
        ")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
    _, err1 := dbUsers.Exec(sqlStr)
    if err1 != nil {
        fmt.Printf("make table error:%v\n", err)
        return err1
    }
    dbUsers.Close()
    return
}

// 创建完用户分表之后对用户分表插入视频id
// 这个表现为用户每次上传完一个视频之后，就把这个视频的id插入到与用户同名的数据表中
func InsertVideoIdToUserTable(videoId int64, userId int64) {
    InitUsersDatabase()
    sqlStr := "INSERT INTO `" + strconv.FormatInt(userId, 10) + "`(video_id,user_id)VALUES(" + strconv.FormatInt(videoId, 10) + "," + strconv.FormatInt(userId, 10) + ");"
    execDatabase(sqlStr)
    dbUsers.Close()
}

// 这是执行数据库语句的函数
// 用户不要调用
func execDatabase(sqlStr string) {
    ret, err := dbUsers.Exec(sqlStr)
    if err != nil {
        fmt.Printf("failed,err%v\n", err)
        return
    }
    id, err := ret.LastInsertId()
    if err != nil {
        fmt.Printf("get failed,err:%v\n", err)
        return
    }
    fmt.Println("运行成功的id是", id)
}

// 查询视频的id表
// 还需要向总库请求视频的具体列表
// id表是为了快速的知道用户的视频id，这样查起总表来可以更快
func GetUserVideosList(userId int64) (ret []int64, arrayLen int) {
    InitUsersDatabase()
    var UserVideoList []int64
    sqlStr := "SELECT video_id,user_id FROM `" + strconv.FormatInt(userId, 10) + "` WHERE video_id > 0"
    rows, err := dbUsers.Query(sqlStr)
    if err != nil {
        fmt.Printf("query failed, err:%v\n", err)
        return
    }
    defer func() {
        rows.Close() // 释放数据库连接
    }()
    var user_id int64
    var video_id int64
    for rows.Next() {
        err := rows.Scan(&video_id, &user_id)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n", err)
            return
        }
        fmt.Printf("scan success ,user id =%v", user_id)
        fmt.Printf("viideo id = %v\n", video_id)
        UserVideoList = append(UserVideoList, video_id)
    }
    dbUsers.Close()
    return UserVideoList, len(UserVideoList)
}
```

#### 3.3.2.路由配置

```Go
package tools
import (
    "config"
    "controller"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)
func InitRouter(r *gin.Engine) {
    // public directory is used to serve static resources
    r.Static("/static", "./public")
    r.LoadHTMLGlob("./templates/*")

    // home page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "Main website",
        })
    })
    apiRouter := r.Group("/douyin")
    // basic apis
    apiRouter.GET("/feed/", controller.Feed)
    apiRouter.GET("/user/", controller.UserInfo)
    apiRouter.POST("/user/register/", controller.Register)
    apiRouter.POST("/user/login/", controller.Login)
    apiRouter.POST("/publish/action/", controller.Publish)
    apiRouter.GET("/publish/list/", controller.PublishList)
    // extra apis - I
    apiRouter.POST("/favorite/action/", controller.FavoriteAction)
    apiRouter.GET("/favorite/list/", controller.FavoriteList)
    apiRouter.POST("/comment/action/", controller.CommentAction)
    apiRouter.GET("/comment/list/", controller.CommentList)
    // extra apis - II
    apiRouter.POST("/relation/action/", controller.RelationAction)
    apiRouter.GET("/relation/follow/list/", controller.FollowList)
    apiRouter.GET("/relation/follower/list/", controller.FollowerList)
    apiRouter.GET("/relation/friend/list/", controller.FriendList)
    apiRouter.GET("/message/chat/", controller.MessageChat)
    apiRouter.POST("/message/action/", controller.MessageAction)
    //获取配置表
    apiRouter.GET("/config", func(context *gin.Context) {
        fmt.Println("当前")
        context.JSON(
            200, gin.H{
                "local_addr": config.ConfigInfo.GetString("local.addr"),
                "mysql_ip":   config.ConfigInfo.GetString("mysql.IP"),
                "mysql_port": config.ConfigInfo.GetString("mysql.Port"),
            })
    })
}
```

#### 3.3.3.用户互动接口

```Go
package controller
import (
    "Mydatabase"
    "common"
    "fmt"
    "net/http"
    "service"
    "strconv"
    "github.com/gin-gonic/gin"
)

type UserListResponse struct {
    common.Response
    UserList []common.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
    token := c.Query("token")
    toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
    actionType := c.Query("action_type")
    ok, userId := service.SearchToken(token)
    if !ok {
        fmt.Println("Token not exist")
    }
    if actionType == "1" {
        //1表示对目标用户进行关注
        Mydatabase.InsertFollowIdToUserTable(toUserId, userId)
        Mydatabase.InsertFollowerIdToUserTable(userId, toUserId)
        //改库1.总库2.视频库中的用户信息
        //1
        userInfo1 := Mydatabase.QueryUserById(userId)
        userInfo1.FollowCount++
        userInfo2 := Mydatabase.QueryUserById(toUserId)
        userInfo2.FollowerCount++
        Mydatabase.UpdateUser(&userInfo1)
        Mydatabase.UpdateUser(&userInfo2)
        //2
        videoList := Mydatabase.QueryVideoByAuthorId(userId)
        var i int
        for i = 0; i < len(videoList); i++ {
            videoList[i].AuthorFollowCount++
            Mydatabase.UpdateVideoInfo(&videoList[i])
        }
        videoList2 := Mydatabase.QueryVideoByAuthorId(toUserId)
        var j int
        for j = 0; j < len(videoList2); j++ {
            videoList2[j].AuthorFollowerCount++
            Mydatabase.UpdateVideoInfo(&videoList2[j])
        }

        c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "关注成功"})
    } else if actionType == "2" {
        //2表示对目标用户进行取关
        Mydatabase.DeleteFollow(toUserId, userId)
        Mydatabase.DeleteFollower(userId, toUserId)
        //改库1.总库2.视频库中的用户信息
        //1
        userInfo1 := Mydatabase.QueryUserById(userId)
        userInfo1.FollowCount--
        userInfo2 := Mydatabase.QueryUserById(toUserId)
        userInfo2.FollowerCount--
        Mydatabase.UpdateUser(&userInfo1)
        Mydatabase.UpdateUser(&userInfo2)
        //2
        videoList := Mydatabase.QueryVideoByAuthorId(userId)
        var i int
        for i = 0; i < len(videoList); i++ {
            videoList[i].AuthorFollowCount--
            Mydatabase.UpdateVideoInfo(&videoList[i])
        }
        videoList2 := Mydatabase.QueryVideoByAuthorId(toUserId)
        var j int
        for j = 0; j < len(videoList2); j++ {
            videoList2[j].AuthorFollowerCount--
            Mydatabase.UpdateVideoInfo(&videoList2[j])
        }
        c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "取消关注成功"})
    } else {
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "关注失败"})
    }
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
    userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
    token := c.Query("token")
    _, id := service.SearchToken(token)
    if id != userId {
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "无token"})
        return
    }
    list, len := Mydatabase.GetUserFollowList(userId)
    var followList []common.User
    var i int
    for i = 0; i < len; i++ {
        userInfo := Mydatabase.QueryUserById(list[i])
        var user common.User
        service.ConvertUserInfoToUser(&userInfo, &user, list[i])
        followList = append(followList, user)
    }
    c.JSON(http.StatusOK, UserListResponse{
        Response: common.Response{
            StatusCode: 0,
        },
        UserList: followList,
    })
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
    userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
    token := c.Query("token")
    _, id := service.SearchToken(token)
    if id != userId {
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "无token"})
        return
    }
    list, len := Mydatabase.GetUserFollowersList(userId)
    var followerList []common.User
    var i int
    for i = 0; i < len; i++ {
        userInfo := Mydatabase.QueryUserById(list[i])
        var user common.User
        service.ConvertUserInfoToUser(&userInfo, &user, list[i])
        followerList = append(followerList, user)
    }
    c.JSON(http.StatusOK, UserListResponse{
        Response: common.Response{
            StatusCode: 0,
        },
        UserList: followerList,
    })
}

// FriendList all users have same friend list
// 朋友是互相关注的两个人为朋友
func FriendList(c *gin.Context) {
    userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
    token := c.Query("token")
    _, id := service.SearchToken(token)
    if id != userId {
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "无token"})
        return
    }
    list, len := Mydatabase.GetUserFollowList(userId)
    var i int
    //当关注列表中的用户有关注本user，即表示为相互关注，把其加入到friendList的id表中
    var friendList []int64
    var length int = 0
    for i = 0; i < len; i++ {
        if Mydatabase.IsFollow(userId, list[i]) {
            friendList = append(friendList, list[i])
            length++
        }
    }
    //加入到friendlist中之后进行user信息的返回
    var friendUserList []common.User
    var userInfo common.Userinfo
    var j int
    for j = 0; j < length; j++ {
        userInfo = Mydatabase.QueryUserById(friendList[j])
        var user common.User
        service.ConvertUserInfoToUser(&userInfo, &user, friendList[j])
        friendUserList = append(friendUserList, user)
    }
    c.JSON(http.StatusOK, UserListResponse{
        Response: common.Response{
            StatusCode: 0,
        },
        UserList: friendUserList,
    })
}
```

#### 3.3.4.token获取

```Go
package service

import (
    "Mydatabase/util"
    "crypto/md5"
    "encoding/hex"
    "strconv"
)

// 做一个哈希表，用于对Token进行维护
// 注意，这里可以使用redis进行替代
// 哈希表只是权宜之策
var tokenStore = make(map[string]int64)
var idStore = make(map[int64]string)

// "token":userID
// userId:"token"
// 维护两个哈希表实现键值对的快速查询
func PushToken(token string, userId int64) bool {
    userIdstr := strconv.FormatInt(userId, 10)
    tag, _ := util.Get("token")
    if tag == false {
        util.Set(token, userIdstr)
        util.Set(userIdstr, token)
    } else {
        return false
    }
    return true

}

func SearchToken(token string) (ok bool, userId int64) {
    tag, res := util.Get(token)
    nt64, err := strconv.ParseInt(res, 10, 64)
    if tag == false || err != nil {
        return false, 0
    } else {
        return true, nt64
    }
}
func SearchTokenById(Id int64) (ok bool, token string) {
    tag, res := util.Get(strconv.FormatInt(Id, 10))
    if tag == false {
        return false, ""
    } else {
        return true, res
    }
}

// MD5加密
func StringToMD5(PWD string) string {
    w := md5.New()
    w.Write([]byte(PWD))
    return hex.EncodeToString(w.Sum(nil))
}
func CreateUserToken(Id int64, password string) string {
    token := strconv.FormatInt(Id, 10) + StringToMD5(password)

    return token
}
```

#### 3.3.5.redis

```Go
package util
import (
    "Mydatabase"
    "context"
    "fmt"
    "time"
)
var client = Mydatabase.GetRedisClient()
var ctx = context.Background()

/*------------------------------------ 字符 操作 ------------------------------------*/
// Set 设置 key的值
func Set(key, value string) bool {
    result, err := client.Set(ctx, key, value, 0).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return result == "OK"
}
// SetEX 设置 key的值并指定过期时间
func SetEX(key, value string, ex time.Duration) bool {
    result, err := client.Set(ctx, key, value, ex).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return result == "OK"
}
// Get 获取 key的值
func Get(key string) (bool, string) {
    result, err := client.Get(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
        return false, ""
    }
    return true, result
}
// GetSet 设置新值获取旧值
func GetSet(key, value string) (bool, string) {
    oldValue, err := client.GetSet(ctx, key, value).Result()
    if err != nil {
        fmt.Println(err)
        return false, ""
    }
    return true, oldValue
}
// Incr key值每次加一 并返回新值
func Incr(key string) int64 {
    val, err := client.Incr(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}
// IncrBy key值每次加指定数值 并返回新值
func IncrBy(key string, incr int64) int64 {
    val, err := client.IncrBy(ctx, key, incr).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}
// IncrByFloat key值每次加指定浮点型数值 并返回新值
func IncrByFloat(key string, incrFloat float64) float64 {
    val, err := client.IncrByFloat(ctx, key, incrFloat).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}

// Decr key值每次递减 1 并返回新值
func Decr(key string) int64 {
    val, err := client.Decr(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}
// DecrBy key值每次递减指定数值 并返回新值
func DecrBy(key string, incr int64) int64 {
    val, err := client.DecrBy(ctx, key, incr).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}
// Del 删除 key
func Del(key string) bool {
    result, err := client.Del(ctx, key).Result()
    if err != nil {
        return false
    }
    return result == 1
}

// Expire 设置 key的过期时间
func Expire(key string, ex time.Duration) bool {
    result, err := client.Expire(ctx, key, ex).Result()
    if err != nil {
        return false
    }
    return result
}
/*------------------------------------ list 操作 ------------------------------------*/
// LPush 从列表左边插入数据，并返回列表长度
func LPush(key string, date ...interface{}) int64 {
    result, err := client.LPush(ctx, key, date).Result()
    if err != nil {
        fmt.Println(err)
    }
    return result
}
// RPush 从列表右边插入数据，并返回列表长度
func RPush(key string, date ...interface{}) int64 {
    result, err := client.RPush(ctx, key, date).Result()
    if err != nil {
        fmt.Println(err)
    }
    return result
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func LPop(key string) (bool, string) {
    val, err := client.LPop(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
        return false, ""
    }
    return true, val
}
// RPop 从列表右边删除第一个数据，并返回删除的数据
func RPop(key string) (bool, string) {
    val, err := client.RPop(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
        return false, ""
    }
    return true, val
}
// LIndex 根据索引坐标，查询列表中的数据
func LIndex(key string, index int64) (bool, string) {
    val, err := client.LIndex(ctx, key, index).Result()
    if err != nil {
        fmt.Println(err)
        return false, ""
    }
    return true, val
}
// LLen 返回列表长度
func LLen(key string) int64 {
    val, err := client.LLen(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}
// LRange 返回列表的一个范围内的数据，也可以返回全部数据
func LRange(key string, start, stop int64) []string {
    vales, err := client.LRange(ctx, key, start, stop).Result()
    if err != nil {
        fmt.Println(err)
    }
    return vales
}
// LRem 从列表左边开始，删除元素data， 如果出现重复元素，仅删除 count次
func LRem(key string, count int64, data interface{}) bool {
    _, err := client.LRem(ctx, key, count, data).Result()
    if err != nil {
        fmt.Println(err)
    }
    return true
}
// LInsert 在列表中 pivot 元素的后面插入 data
func LInsert(key string, pivot int64, data interface{}) bool {
    err := client.LInsert(ctx, key, "after", pivot, data).Err()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}
/*------------------------------------ set 操作 ------------------------------------*/
// SAdd 添加元素到集合中
func SAdd(key string, data ...interface{}) bool {
    err := client.SAdd(ctx, key, data).Err()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}
// SCard 获取集合元素个数
func SCard(key string) int64 {
    size, err := client.SCard(ctx, "key").Result()
    if err != nil {
        fmt.Println(err)
    }
    return size
}

// SIsMember 判断元素是否在集合中
func SIsMember(key string, data interface{}) bool {
    ok, err := client.SIsMember(ctx, key, data).Result()
    if err != nil {
        fmt.Println(err)
    }
    return ok
}
// SMembers 获取集合所有元素
func SMembers(key string) []string {
    es, err := client.SMembers(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return es
}
// SRem 删除 key集合中的 data元素
func SRem(key string, data ...interface{}) bool {
    _, err := client.SRem(ctx, key, data).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

// SPopN 随机返回集合中的 count个元素，并且删除这些元素
func SPopN(key string, count int64) []string {
    vales, err := client.SPopN(ctx, key, count).Result()
    if err != nil {
        fmt.Println(err)
    }
    return vales
}
/*------------------------------------ hash 操作 ------------------------------------*/
// HSet 根据 key和 field字段设置，field字段的值
func HSet(key, field, value string) bool {
    err := client.HSet(ctx, key, field, value).Err()
    if err != nil {
        return false
    }
    return true
}

// HGet 根据 key和 field字段，查询field字段的值
func HGet(key, field string) string {
    val, err := client.HGet(ctx, key, field).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}
// HMGet 根据key和多个字段名，批量查询多个 hash字段值
func HMGet(key string, fields ...string) []interface{} {
    vales, err := client.HMGet(ctx, key, fields...).Result()
    if err != nil {
        panic(err)
    }
    return vales
}
// HGetAll 根据 key查询所有字段和值
func HGetAll(key string) map[string]string {
    data, err := client.HGetAll(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return data
}
// HKeys 根据 key返回所有字段名
func HKeys(key string) []string {
    fields, err := client.HKeys(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return fields
}
// HLen 根据 key，查询hash的字段数量
func HLen(key string) int64 {
    size, err := client.HLen(ctx, key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return size
}
// HMSet 根据 key和多个字段名和字段值，批量设置 hash字段值
func HMSet(key string, data map[string]interface{}) bool {
    result, err := client.HMSet(ctx, key, data).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return result
}
// HSetNX 如果 field字段不存在，则设置 hash字段值
func HSetNX(key, field string, value interface{}) bool {
    result, err := client.HSetNX(ctx, key, field, value).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return result
}
// HDel 根据 key和字段名，删除 hash字段，支持批量删除
func HDel(key string, fields ...string) bool {
    _, err := client.HDel(ctx, key, fields...).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}
// HExists 检测 hash字段名是否存在
func HExists(key, field string) bool {
    result, err := client.HExists(ctx, key, field).Result()
    if err != nil {
        fmt.Println(err)
        return false
    }
    return result
}
```

#### 3.3.6.配置热重载

```Go
package config
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)
var ConfigInfo *viper.Viper
// config结构体，用于返回对应得config信息
type Config struct {
    LocalAddr string `json:"local_addr"`
    MysqlIP   string `json:"mysql_ip"`
    MysqlPort string `json:"mysql_port"`
}
// 热配置函数
func InitConfig() {
    ConfigInfo = initConfig()
    go dynamicConfig()
}
func initConfig() *viper.Viper {
    //v := viper.New()
    v := viper.New()
    v.SetConfigName("config")   // 设置文件名称（无后缀）
    v.SetConfigType("yaml")     // 设置后缀名 {"1.6以后的版本可以不设置该后缀"}
    v.AddConfigPath("./config") // 设置文件所在路径
    v.Set("verbose", true)      // 设置默认参数
    if err := v.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            panic(" Config file not found; ignore error if desired")
        } else {
            panic("Config file was found but another error was produced")
        }
    }
    return v
}
func dynamicConfig() {
    // 监控配置和重新获取配置
    ConfigInfo.WatchConfig()
    ConfigInfo.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed:", e.Name)
    })
}
func GetConfig() (config Config) {
    fmt.Println("正在获取配置文件...")
    url := "http://localhost:8080/douyin/config"
    spaceClient := http.Client{
        Timeout: time.Second * 2, // Maximum of 2 secs
    }
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("User-Agent", "spacecount-tutorial")
    res, getErr := spaceClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }
    if res.Body != nil {
        defer res.Body.Close()
    }
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }
    jsonErr := json.Unmarshal(body, &config)
    if jsonErr != nil {
        fmt.Println(jsonErr)
    }
    return config
}
func GetDBAddr() (dbAddr string) {
    config := GetConfig()
    dbAddr = config.MysqlIP + ":" + config.MysqlPort
    return dbAddr
}
func GetLocalAddr() string {
    config := GetConfig()
    return config.LocalAddr
}
```

#### 3.3.7.消息功能

```Go
package controller
import (
    "Mydatabase"
    "common"
    "fmt"
    "net/http"
    "service"
    "strconv"
    "github.com/gin-gonic/gin"
)
var tempChat = map[string][]common.Message{}
var messageIdSequence = int64(1)
type ChatResponse struct {
    common.Response
    MessageList []common.MessageRender `json:"message_list"`
}
// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
    token := c.Query("token")
    toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
    content := c.Query("content")
    ok, userId := service.SearchToken(token)
    if !ok {
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "token not exist"})
        return
    }
    _, err := Mydatabase.InsertMessage(userId, toUserId, content)
    if err != nil {
        fmt.Println("insert message error", err)
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "insert message error"})
        return
    }
    c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "消息发送成功"})
}
// 获取chatkey
func GetChatKey(userId int64, fromUserId int64, toUserId int64) (key string) {
    if fromUserId > toUserId {
        key = strconv.FormatInt(userId, 10) + "_" + strconv.FormatInt(fromUserId, 10) + "_" + strconv.FormatInt(toUserId, 10)
        return key
    } else {
        key = strconv.FormatInt(userId, 10) + "_" + strconv.FormatInt(toUserId, 10) + "_" + strconv.FormatInt(fromUserId, 10)
        return key
    }
}
// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
    token := c.Query("token")
    msg_time, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
    toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
    ok, userId := service.SearchToken(token)
    // userId = 2
    if !ok {
        fmt.Println("token error")
        c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "token not exist"})
        return
    }
    //注意，这样发送消息的致命缺陷就是消息会重复发送
    //所以需要用到消息队列
    messageList := Mydatabase.GetMessageList(userId, msg_time)
    //var res []common.Message
    var i int
    var messageRenderList []common.MessageRender
    var temp common.MessageRender
    for i = 0; i < len(messageList); i++ {
        if messageList[i].ToUserId == toUserId {
            temp.FromUserId = messageList[i].FromUserId
            temp.ToUserId = messageList[i].ToUserId
        } else {
            temp.FromUserId = messageList[i].ToUserId
            temp.ToUserId = messageList[i].FromUserId
        }
        temp.Id = messageList[i].Id
        temp.Content = messageList[i].Content
        temp.CreateTime = messageList[i].CreateTime
        messageRenderList = append(messageRenderList, temp)
    }
    c.JSON(http.StatusOK,
        ChatResponse{Response: common.Response{StatusCode: 0},
            MessageList: messageRenderList,
        })
    fmt.Println(messageRenderList)
}
```

# 四、测试结果

> 建议从功能测试和性能测试两部分分析，其中功能测试补充测试用例，性能测试补充性能分析报告、可优化点等内容。

***功能测试为必填**

## 4.1.功能测试

> 我们使用了存储过程的形式进行的项目架构，所以会在数据库上再进行一层封装，当我们进行单元测试的时候，不需要对gin框架内容进行具体的测试，只需要对数据库层上的封装进行测试即可，gin框架层依据结构体返回的是json文件

### 4.1.1. 基础功能

#### 1.视频流

视频流测试从返回的json进行分析（结合了publish操作，当视频库发生更新后，返回的json会立刻进行更新）

* 空数据库状态

![](https://bo9uwdfynq.feishu.cn/space/api/box/stream/download/asynccode/?code=ZjlmN2U3MmYyMDU2N2I1MjY4ZmM0MjE3NjBiZGYzNjNfQlpDYk44MnJxMlZmNkRaYlp4WDl5TGVVa1pISUUzaXpfVG9rZW46S01ZOGJLR0Vpb1RqN294WGdNTWNDbWhNbktnXzE2OTM0ODc0NTc6MTY5MzQ5MTA1N19WNA)

* 数据库添加数据（模拟投稿）

![](https://bo9uwdfynq.feishu.cn/space/api/box/stream/download/asynccode/?code=MTgwNzhlZDAwNTM5YzFlMDlmMzllNzdlMjMzNWQxYjlfTjYwYnVPbW4wQ0ZXMHJEZmtHWE5SWldlcjlYUmZPRG1fVG9rZW46UWFIM2JydkEzb0pvSUJ4a21ycmM3R2tpbndnXzE2OTM0ODc0NTc6MTY5MzQ5MTA1N19WNA)

可以见到feed中的json文件已经发生变化，说明视频流的测试已经成功了

#### 2.用户注册登录

注册的时候我们会生成用户的基础数据，由于用户的基本数据无法进行二次更改，所以在一开始就随机生成了用户的头像和背景图片（直接使用了随机数种子从后台的六张图片中随机生成）注册登录的时候会依据用户信息生成独一无二的用户token，存入到redis数据库中，此后用户端无论是发布视频还是进行操作只需要凭借token即可进行鉴权操作

```Go
package test

import (
    "Mydatabase"
    "fmt"
    "testing"
)

func TestInitPWDDatabase(t *testing.T) {
    err := Mydatabase.InitPWDDatabase()
    if err != nil {
        fmt.Println("init DB failed,err \n", err)
        t.Errorf("发生错误")
    }
}
func TestInsertNewUser(t *testing.T) {
    err,_ := Mydatabase.InsertNewUser("wodiaonimade")
    if err != nil {
        fmt.Println("init DB failed,err \n", err)
        t.Errorf("发生错误")
    }
}
func TestQueryUserPWD(t *testing.T) {
    pwd:=Mydatabase.QueryUserPWD(2)
    fmt.Println(pwd)
}
func TestJudgePWD(t *testing.T) {
    test :="wsasnan"
    ok := Mydatabase.JudgePWD(2,test)
    fmt.Println(ok)
}
```

#### 3.用户信息

从数据库中提取了用户的信息，以json的形式发送给前端

```Go
package test

import (
    "Mydatabase"
    "common"
    "fmt"
    "testing"
)

// 数据库连接测试
func TestGetDB(t *testing.T) {
    //e := newExpect(t)
    Mydatabase.GetDB()
}

// 根据id查询用户测试
func TestGetUserId(t *testing.T) {
    //e := newExpect(t)
    fmt.Println(Mydatabase.QueryUserById(1))
}

// 根据id查询用户测试
func TestGetUserName(t *testing.T) {
    //e := newExpect(t)
    fmt.Println(Mydatabase.QueryUserByName("王五"))
}

// 创建一个用户
func TestInsertUser(t *testing.T) {
    // e := newExpect(t)
    user := common.Userinfo{
        Id: 0, Name: "ddssjj", FollowCount: 1,
        FollowerCount: 20, Avator: "哈实习",
        BackgroundImage: "赫斯", Signature: "早上好",
        TotalFavorited: 87,
        WorkCount:      23,
        FavoriteCount:  67}
    res := Mydatabase.InsertUser(&user)

    fmt.Println(res)
}

func TestUpdateUser(t *testing.T) {
    // e := newExpect(t)
    user := common.Userinfo{
        Id: 3, Name: "大哥3355", FollowCount: 10,
        FollowerCount: 20, Avator: "ccc",
        BackgroundImage: "c", Signature: "早上好",
        TotalFavorited: 87,
        WorkCount:      23,
        FavoriteCount:  67}
    res := Mydatabase.UpdateUser(&user)

    fmt.Println(res)
}
```

#### 4.投稿与发布列表

投稿：向后端发送视频，后端将视频的地址存入视频数据库中以供播放，这里的视频的封面使用随机数种子从仓库中随机生成视频封面

发布列表：展示用户的视频列表，向前端传送的仍然是json的形式,主要测试的就是视频的信息功能

主要需要测试的就是视频信息的存储功能

```Go
package test

import (
    "Mydatabase"
    "common"
    "fmt"
    "math/rand"
    "testing"
)

func TestQueryVideoById(t *testing.T) {
    fmt.Println(Mydatabase.QueryVideoById(1))
}

func TestQueryVideoByAuthorId(t *testing.T) {
    fmt.Println(Mydatabase.QueryVideoByAuthorId(2))
}

func TestInsertVideoInfo(t *testing.T) {
    video := common.Videoinfo{
        VideoId: 2, AuthorId: 2, AuthorName: "测试数据",
        AuthorFollowCount: 21, AuthorFollowerCount: 23,
        AuthorAvator: "eferg", AuthorBackgroundImage: "深夜下哦美好看",
        AuthorSignature: "qg", AuthorTotalFavorited: 21, AuthorWorkCount: 765,
        AuthorFavoriteCount: 43, VideoPlayUrl: "更好的武功和",
        VideoCoverUrl: "ge", VideoFavoriteCount: 231,
        VideoCommentCount: 233, VideoTitle: "好看的时间噢批",
        VideoTime: "geq",
    }
    res := Mydatabase.InsertVideoInfo(&video)
    fmt.Println(res)
}

func TestQueryVideoByAuthorName(t *testing.T) {
    fmt.Println(Mydatabase.QueryVideoByAuthorName("测试数据"))
}

func TestQueryVideoByVideoTitle(t *testing.T) {
    fmt.Println(Mydatabase.QueryVideoByVideoTitle("哈哈哈哈"))
}

func TestQueryVideoIdByAuthorId(t *testing.T) {
    fmt.Println(Mydatabase.QueryVideoIdByAuthorId(1))
}

func TestQueryVideoIdByAuthorName(t *testing.T) {
    fmt.Println(Mydatabase.QueryVideoIdByAuthorName("测试数据"))
}

func TestDeleteByAuthorId(t *testing.T) {
    //注意，不要测试delete，没有授权
    fmt.Println(Mydatabase.DeleteByAuthorId(1))
}

func TestUpdateVideoInfo(t *testing.T) {
    video := common.Videoinfo{
        VideoId: 5, AuthorId: 2, AuthorName: "垃圾很多很多ddddd",

        AuthorFollowCount: 21, AuthorFollowerCount: 23,
        AuthorAvator: "的结局", AuthorBackgroundImage: "d334",
        AuthorSignature: "二分", AuthorTotalFavorited: 21, AuthorWorkCount: 765,
        AuthorFavoriteCount: 43, VideoPlayUrl: "ddfa",
        VideoCoverUrl: "而非我", VideoFavoriteCount: 231,
        VideoCommentCount: 233, VideoTitle: "dfwfw的我",
        VideoTime: "微服务",
    }
    fmt.Println(Mydatabase.UpdateVideoInfo(&video))
}
func TestDeleteByVideoId(t *testing.T) {
    fmt.Println(Mydatabase.DeleteByVideoId(2))
}

func TestRand(t *testing.T) {
    fmt.Println(rand.Int())
}
```

### 4.1.2.进阶功能1

#### 1.点赞与点赞列表

点赞：点赞操作分为两部分，点赞和取消赞

点赞列表：点赞列表的获取我们借助了是否点赞的tag，如果tag置1就是点赞，tag置0就是取消赞，最后获取点赞列表的时候只需要扫描tag即可，虽然会牺牲一部分的数据库空间，但是确保了数据库的安全

```Go
package test

import (
    "Mydatabase"
    "fmt"
    "testing"
)

func TestInitFavoriteDatabase(t *testing.T) {
    err := Mydatabase.InitFavoriteDatabase()
    if err != nil {
        fmt.Println("init DB failed,err \n", err)
        t.Errorf("发生错误")
    }
}

func TestFavoriteMake(t *testing.T) {
    err := Mydatabase.MakeNewFavoriteTable(5)
    if err != nil {
        fmt.Println("init DB failed,err \n", err)
        t.Errorf("发生错误")
    }
}

func TestFavoriteInsert(t *testing.T) {
    Mydatabase.InsertUserIdToFavoriteTable(2, 5)

}

func TestFavoriteDelete(t *testing.T) {
    Mydatabase.DeleteUserIdToFavoriteTable(2, 5)

}

func TestFavoriteQuery(t *testing.T) {
    fmt.Println(Mydatabase.GetFavoriteVideoList(5, 2))

}
func TestDelete(t *testing.T) {
    //Mydatabase.DeleteUserIdToVideoTable(10, 1)
    Mydatabase.InsertUserIdToVideoTable(10, 1)

}
```

#### 2.评论与评论列表

评论：针对了每个视频的id进行对视频分表，每条评论集中在视频下方

评论列表：扫描视频的评论数据库

```Go
package test

import (
    "Mydatabase"
    "testing"
)

func TestMakeCommentTable(t *testing.T) {
    Mydatabase.MakeCommentTable(3)
}
func TestInsertComment(t *testing.T) {
    Mydatabase.InsertComment(1, 2, "哈哈哈")
}
func TestDeleteComment(t *testing.T) {
    Mydatabase.DeleteComment(1,4)
}
func TestGetCommentList(t *testing.T) {
    Mydatabase.GetCommentList(1)
}
```

### 4.2.3.进阶功能2

#### 1.关注及关注、粉丝、好友列表

关注

关注列表

粉丝列表

好友列表

```Go
package test

import (
    "Mydatabase"
    "fmt"
    "testing"
)

func TestInitFollowDatabase(t *testing.T) {
    Mydatabase.InitFollowDatabase()
}
func TestMakeNewFollowTable(t *testing.T) {
    Mydatabase.MakeNewFollowTable(3)
}
func TestInsertFollowId(t *testing.T) {
    Mydatabase.InsertFollowIdToUserTable(2, 3)
}
func TestUserFollowList(t *testing.T) {
    list, len := Mydatabase.GetUserFollowList(3)
    var i int
    for i = 0; i < len; i++ {
        fmt.Println(list[i])
    }
}
func TestDeleteFollow(t *testing.T) {
    //Mydatabase.DeleteFollow(3, 1)
    //Mydatabase.DeleteFollow(4, 1)
    Mydatabase.GetUserFollowList(1)
}
```

```Go
package test

import (
    "Mydatabase"
    "fmt"
    "testing"
)
func TestInitFollowersFatabbase(t *testing.T) {
    Mydatabase.InitFollowersDatabase()
}
func TestMakeNewFollowerTable(t *testing.T) {
    Mydatabase.MakeNewFollowerTable(2)
}
func TestInsertFollowerIdToUserTable(t *testing.T) {
    Mydatabase.InsertFollowerIdToUserTable(17,2)
}
func TestGetUserFollowersList(t *testing.T) {
    list,len:=Mydatabase.GetUserFollowersList(2)
    var i int
    for i=0;i<len;i++{
        fmt.Println(list[i])
    }
}
func TestIsFollow(t *testing.T) {
    fmt.Println(Mydatabase.IsFollow(2,11))
    fmt.Println(Mydatabase.IsFollow(2,12))
}
func TestDeleteFollower(t *testing.T) {
    Mydatabase.DeleteFollower(3,1)
}
```

#### 2.消息及消息列表

发送消息

聊天记录

```Go
package test

import (
    "Mydatabase"
    "fmt"
    "testing"
)

func TestInitMessageDatabase(t *testing.T) {
    Mydatabase.InitMessageDatabase()
}
func TestMakeMessageTable(t *testing.T) {
    Mydatabase.MakeNewMessageTable(4)
}
func TestInsertMessage(t *testing.T) {
    Mydatabase.InsertMessage(1, 2, "大家好哇")
}
func TestMessageList(t *testing.T) {
    messageList := Mydatabase.GetMessageList(2, 0)
    length := len(messageList)
    var i int
    for i = 0; i < length; i++ {
        fmt.Println(messageList[i])
    }
}
```

## 4.2.可优化点

* 应当使用redis作为缓冲数据库，整体的架构应当改成
  ![](https://bo9uwdfynq.feishu.cn/space/api/box/stream/download/asynccode/?code=YjM0YTJkNTk2ZGMyYmQxYmQzMjkxMDY0MmY5OWQ5MGNfSWdjY0Y2emZ2U25OYlJRU1l1dTdKVnhzalpQb3FTZHNfVG9rZW46U0F3Y2JxekxGb0VwUnV4ampoc2N2TDJObmRnXzE2OTM0ODc0NTc6MTY5MzQ5MTA1N19WNA)

以redis充当持久化存储的缓冲层，即替代内存的作用，而mysql作为持久化存储。

这样用户的存储读取会变得更加迅速，用户体验会更好

而且当服务端程序发生崩溃的时候，未及时备份的数据不会丢失，数据更加的安全

* 介于linux服务器和我们使用的windows环境略有不同，本应当使用ffmpeg截取视频的第一帧作为封面，但是技术路线不是很相似，所以我们采用了随机生成视频封面的策略

# 五、Demo 演示视频

[极简版抖音客户端简单测试_哔哩哔哩_bilibili](https://www.bilibili.com/video/BV1Gh4y1m7fL/?vd_source=babe8305104e60531ced6a3686665d35)

# 六、项目总结与反思

> 1. 目前仍存在的问题
> 2. 已识别出的优化项
> 3. 架构演进的可能性
> 4. 项目过程中的反思与总结

1. 视频流功能。由于视频流的实现没有使用流化的操作，导致用户端在刷视频的时候必须要等待视频完整发送到用户端的时候才能进行播放，这样在当视频占用存储空间较大的时候用户会花费较长的时间进行等待，不利于用户的体验。采取的结局策略是可以将视频分片，流化发送给用户，可以让用户获得较好的观看体验
2. 架构的基本问题：

* 没有使用redis作为缓冲层而是使用了内存+基本数据结构作为用户端数据的缓冲，这样在面对大量信息的时候如果服务端发生崩溃，数据无法及时的备份到数据库中，容易造成数据的丢失。解决策略：用户发送的所有有可能改变用户基本信息状态的信息都应当及时的备份到以内存为核心的redis数据库中，redis不仅速度快，而且数据能够及时的迁移到硬盘中，即便服务程序崩溃，redis仍然可以稳定运行，不会造成数据丢失，同时redis可以稳定的将数据备份到以硬盘为核心的mysql数据库中，也能够减轻频繁的对mysql进行读写所造成的数据库压力

3. 对mysql数据库的读写频次过高：mysql只有每秒几十次的存取速度，当服务程序频繁存取mysql会造成mysql的压力过大。可以使用redis作为缓冲数据库，然后将redis数据备份到mysql中即可
4. 没有使用ffmpeg，而是使用了随机生成的封面，是一个遗憾，可以改进

# 七、其他补充资料

-  各功能对应的接口说明文档，地址：https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c

- 极简抖音 APP 包使用说明：https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7
- 服务端 Demo，仓库地址 https://github.com/RaymondCode/simple-demo
