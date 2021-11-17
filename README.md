## reference
https://hub.fastgit.org/cloudreve/Cloudreve


## todo
* 支持存储端 本地、 阿里云 OSS、腾讯云 COS、亚马逊 S3、 OneDrive
* 支持客户端直传
* 多用户、用户分组
* 文件分享、过期时间
* 在线预览


## request

```go
// 获取 uri
vars := mux.Vars(r)
vars["name"]

// 获取参数 GET
vars := r.URL.Query()
vars.Get("name")

// 获取参数 POST PUT PATCH 
// application/x-www-form-urlencoded
name := r.PostFormValue("name")

// 获取参数 GET POST PUT PATCH
name := r.FormValue("name")

// 获取 body 任意请求类型
var reader io.Reader = r.Body
b, _ := ioutil.ReadAll(reader)
v, _ := url.ParseQuery(string(b))
```

## cookie
```go
// 读取 cookie
r.Cookies()

// 写入 cookie
cookie := &http.Cookie{Name: "userId", Value: "123456"}
http.SetCookie(w, cookie)
```
