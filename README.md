# bmsg
bmsg
#### 站内信

##### 1: 未读消息 toUserId: 7003496620
```bash
/innerMsg/unreadDetail/7003496620?
```

##### 2: 全部消息 
```bash
/innerMsg/allDetail/7003714351?
```

##### 3: 已读消息
```bash
#/innerMsg/read/0
/innerMsg/readDetail/7003714351?categoryId=0
```
##### 4.分页设计
```bash
https://notifications.console.aliyun.com/message/getMessageList.json?pageNumber=2&pageSize=10&status=1
```

##### 4.1 分页测试
````sh
POST localhost:5008/v1/api/msg/read/to?pageSize=5&pageNumber=2
{
    "from_user_id": 100,
    "to_user_id": 1
}

{
    "data": [
        {
            "Id": 6,
            "FromUserId": 105,
            "ToUserId": 0,
            "CreatedAt": "2018-04-03T10:43:32.580822+08:00",
            "UpdateAt": "2018-04-03T10:43:32.580822+08:00",
            "Title": "Welcome to join APKPURE developer",
            "Message": "1.you xxx 2.xxxxx 3.xxxx",
            "IsDelete": false,
            "Status": "UNSEEN"
        },
        {
            "Id": 7,
            "FromUserId": 106,
            "ToUserId": 0,
            "CreatedAt": "2018-04-03T10:43:32.580827+08:00",
            "UpdateAt": "2018-04-03T10:43:32.580827+08:00",
            "Title": "Welcome to join APKPURE developer",
            "Message": "1.you xxx 2.xxxxx 3.xxxx",
            "IsDelete": false,
            "Status": "UNSEEN"
        },
        ...
        {
            "Id": 10,
            "FromUserId": 109,
            "ToUserId": 0,
            "CreatedAt": "2018-04-03T10:43:32.580835+08:00",
            "UpdateAt": "2018-04-03T10:43:32.580836+08:00",
            "Title": "Welcome to join APKPURE developer",
            "Message": "1.you xxx 2.xxxxx 3.xxxx",
            "IsDelete": false,
            "Status": "UNSEEN"
        }
    ],
    "err": "",
    "nums": 5
}
```