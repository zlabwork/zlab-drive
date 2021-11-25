#### 2.1 文件列表

| 类型 | URL  |
| ------ | ------ |
| POST | /files/{:id:} |

响应 demo
```json
{
    "code":200,
    "message":"success",
    "data":[
        {
            "uuid":"b152f302-f436-436f-9843-de9da01a5d74",
            "mime":"folder",
            "size":0,
            "hash":"",
            "name":"风景",
            "key":"L-mjjuaZrw",
            "attr":"",
            "file_ctime":1633756532,
            "file_mtime":1633756532,
            "ctime":1633756532,
            "mtime":1633756532
        }
    ]
}
```


#### 2.2 操作接口

| 类型 | URL  |
| ------ | ------ |
| POST | /do/{key} |


| 参数名 | 类型 | 描述 | 是否必须 |
| ------ | ------ | ------ | ------ |
| action | string | 操作类型 delete, move | 是 |
| to | string | 要移动目标位置，经过 RawURLEncoding，如: L0dhbyDph5HkuL0 | 否 |


响应 demo
```json
{
    "code": 200,
    "message": "success"
}
```
