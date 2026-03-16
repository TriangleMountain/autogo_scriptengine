# media - 多媒体
---
提供多媒体编程的支持，目前仅支持媒体文件扫描，后续会加入更多功能。

## ScanFile
<hr style="margin: 0;">

扫描指定文件，将其加入媒体库中。

- `path` {string} 要扫描的文件路径。

```go
media.ScanFile("/sdcard/1.png")
```

## PlayMP3
<hr style="margin: 0;">

播放指定路径的 MP3 音频文件。

- `path` {string} 要播放的 MP3 文件路径。

```go
media.PlayMP3("/sdcard/music.mp3")
```

## SendSMS
<hr style="margin: 0;">

向指定手机号发送短信。

- `number` {string} 接收短信的目标手机号。
- `message` {string} 要发送的短信内容。

```go
media.SendSMS("10086", "Hello AutoGo")
```