# utils - 工具函数
---
提供一组常用的工具函数，包括日志记录、字符串与数据类型的转换、随机数生成等功能。

## Shell
<hr style="margin: 0;">

执行 shell 命令并返回输出。

- `cmd` {string} 要执行的命令。

```go
output := utils.Shell("ls -l")
fmt.Println("Command Output:", output)
```

## Toast
<hr style="margin: 0;">

显示 Toast 提示信息。

- `message` {string} 要显示的提示信息。

```go
utils.Toast("Hello AutoGo")
```

## Alert
<hr style="margin: 0;">

显示带标题、内容和按钮的弹窗，阻塞等待用户点击后返回按钮索引，安卓15及以上只有在APK模式下才能正常弹出。

- `title` {string} 弹窗标题。
- `content` {string} 弹窗内容。
- `btn1Text` {string} 第一个按钮的文字(通常为"取消")，输入空字符串默认不显示该按钮。
- `btn2Text` {string} 第二个按钮的文字(通常为"确定")。

```go
btnIndex := utils.Alert("确认操作", "重置脚本数据？", "取消", "确认")
if btnIndex == 1 {
	fmt.Println("用户点击了确认")
} else {
	fmt.Println("用户点击了取消")
}
```

## Sleep
<hr style="margin: 0;">

让当前线程暂停执行指定的时间。

- `i` {int} 暂停时间（毫秒）。

```go
utils.Sleep(500) // 暂停 500 毫秒
```

## Random
<hr style="margin: 0;">

返回指定范围内的真随机整数，包含最小值和最大值。

- `min` {int} 最小值。
- `max` {int} 最大值。

```go
randNum := utils.Random(1, 10)
fmt.Println("Random Number:", randNum)
```

## LogI
<hr style="margin: 0;">

记录一条 INFO 级别的日志。

- `label` {string} 日志标签，用于标识日志类别。
- `message` {...interface{}} 日志消息，描述具体的日志内容。

```go
utils.LogI("AppStart", "Application has started successfully.")
```

## LogE
<hr style="margin: 0;">

记录一条 ERROR 级别的日志。

- `label` {string} 日志标签，用于标识日志类别。
- `message` {...interface{}} 日志消息，描述具体的日志内容。

```go
utils.LogE("AppCrash", "Application encountered an unexpected error.")
```

## I2s
<hr style="margin: 0;">

将整数转换为字符串。

- `i` {int} 要转换的整数。

```go
str := utils.I2s(123)
fmt.Println("String Value:", str)
```

## S2i
<hr style="margin: 0;">

将字符串转换为整数。

- `s` {string} 要转换的字符串。

```go
num := utils.S2i("123")
fmt.Println("Integer Value:", num)
```

## F2s
<hr style="margin: 0;">

将浮点数转换为字符串。

- `f` {float64} 要转换的浮点数。

```go
str := utils.F2s(123.45)
fmt.Println("String Value:", str)
```

## S2f
<hr style="margin: 0;">

将字符串转换为浮点数。如果转换失败返回 0.0。

- `s` {string} 要转换的字符串。

```go
num := utils.S2f("123.45")
fmt.Println("Float Value:", num)
```

## B2s
<hr style="margin: 0;">

将布尔值转换为字符串 ("true" 或 "false")。

- `b` {bool} 要转换的布尔值。

```go
str := utils.B2s(true)
fmt.Println("Boolean as String:", str)
```

## S2b
<hr style="margin: 0;">

将字符串转换为布尔值。如果无法转换则返回 false。

- `s` {string} 要转换的字符串。

```go
boolVal := utils.S2b("true")
fmt.Println("Boolean Value:", boolVal)
```
