# 微信告警接口封装
<hr>


## Features
* 成员信息：查询、添加、删除
* 组管理：查询、添加、删除、更新
* 组成员管理：查询、添加、删除
* 发送信息：发送


### 描述

基于微信企业号注册登录,帮助用户快速实现告警方式；

### 安装说明

 - 编译安装
 
``` go build && ./xalarm ```

### 用户管理

 - 成员查询
```
curl -XGET 'http://127.0.0.1:8000/api/users/list'
```
 - 创建用户
 
```
curl -XPOST 'http://127.0.0.1:8000/api/users/create' -d '
{
    "userid": "1075741124@qq.com",
    "mobile": "1111111111"
}'

```
 - 删除用户
``` 
curl -XGET 'http://127.0.0.1:8000/api/users/delete?userid=1075741124@qq.com'
```

### 组管理

 - 组查询
``` 
curl -XGET 'http://127.0.0.1:8000/api/tags/list'
```

 - 添加组
``` 
curl -XPOST 'http://127.0.0.1:8000/api/tags/create' -d '
{
	"tagname": "ops"
}
'
```
 - 删除组
``` 
curl -XGET 'http://127.0.0.1:8000/api/tags/delete?tagid=1'
```

 - 更新组
``` 
curl -XPOST 'http://127.0.0.1:8000/api/tags/update?tagid=1' -d '
{
	"tagname": "newops"
}
'
```
### 组成员管理

 - 查询组成员
``` 
curl -XGET 'http://127.0.0.1:8000/api/tag/users/list?tagid=1'
```

 - 添加组成员
``` 
curl -XPOST 'http://127.0.0.1:8000/api/tag/users/create?tagid=1' -d '
{
    "userlist": ["1075741124@qq.com", "1075741125@qq.com"]
}
'
```
 - 删除组成员
``` 
curl -XPOST 'http://127.0.0.1:8000/api/tag/users/delete?tagid=1' -d '
{
    "userlist": ["1075741124@qq.com"]
}
'
```
### 发送信息
 - 按组发送
``` 
curl -XPOST 'http://127.0.0.1:8000/api/msg/send' -d '
{
    "touser":"",
    "totag":"1|2",
    "msg": "linux"
}
'
```
 - 按用户发送
``` 
curl -XPOST 'http://127.0.0.1:8000/api/alarm/wechat/send' -d '
{
    "touser":"1075741124@qq.com|1075741124@qq.com",
    "totag":"",
    "msg": "linux"
}
'
```
 - 按组&用户发送
``` 
curl -XPOST 'http://127.0.0.1:8000/api/msg/send' -d '
{
    "touser":"1075741124@qq.com",
    "totag":"1|2",
    "msg": "linux"
}
'
```