# https - 网络请求
---
https模块提供了发送HTTP/HTTPS请求的功能，可用于与网络服务进行交互，获取网页内容，上传文件等。

## Get
<hr style="margin: 0;">

发送GET请求并返回响应状态码和数据。

- `url` {string} 请求的URL
- `timeout` {int} 请求的超时时间（毫秒），如果为0则不设置超时

```go
code, data := https.Get("https://example.com", 5000)
if code == 200 {
    fmt.Println("请求成功：", string(data))
} else {
    fmt.Println("请求失败，状态码：", code)
}
```

## Post
<hr style="margin: 0;">

发送POST请求并返回响应状态码和数据。支持自定义请求头和请求体，适用于发送JSON、XML等格式的数据。

- `url` {string} 请求的URL
- `data` {[]byte} 请求体数据（如JSON序列化后的字节数组）
- `headers` {map[string]string} 自定义请求头，如果为nil或未设置Content-Type，默认使用application/json
- `timeout` {int} 请求的超时时间（毫秒），如果为0则不设置超时

```go
// 构造JSON请求体
reqData := map[string]interface{}{
    "name": "张三",
    "age":  25,
}
jsonData, _ := json.Marshal(reqData)

// 设置请求头
headers := map[string]string{
    "Content-Type":  "application/json",
    "Authorization": "Bearer your-token",
}

// 发送请求
code, data := https.Post("https://example.com/api/user", jsonData, headers, 10000)
if code == 200 {
    fmt.Println("请求成功：", string(data))
} else {
    fmt.Println("请求失败，状态码：", code)
}
``` 

## PostMultipart
<hr style="margin: 0;">

发送带有文件的POST请求（multipart/form-data格式）并返回响应状态码和数据。

- `url` {string} 请求的URL
- `fileName` {string} 文件名
- `fileData` {[]byte} 文件数据（字节数组）
- `timeout` {int} 请求的超时时间（毫秒），如果为0则不设置超时

```go
// 读取文件数据
fileData := files.ReadBytes("/sdcard/image.jpg")
// 发送请求
code, data := https.PostMultipart("https://example.com/upload", "image.jpg", fileData, 10000)
if code == 200 {
    fmt.Println("上传成功：", string(data))
} else {
    fmt.Println("上传失败，状态码：", code)
}
```