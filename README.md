## Request 参数获取

```golang
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
