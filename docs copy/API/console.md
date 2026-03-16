# console - 控制台
---
提供用于控制台悬浮窗的控制接口，支持多实例、位置、大小、颜色设置以及内容打印等功能。

## New
<hr style="margin: 0;">

创建一个新的控制台实例。

**返回** {*Console} 控制台实例指针

```go
c := console.New()
```

## SetWindowSize
<hr style="margin: 0;">

设置控制台窗口的宽高。

- `width` {int} 控制台窗口的宽度
- `height` {int} 控制台窗口的高度

**返回** {*Console} 控制台实例指针

```go
c.SetWindowSize(800, 600)
```

## SetWindowPosition
<hr style="margin: 0;">

设置控制台窗口的位置。

- `x` {int} 控制台窗口左上角的横坐标
- `y` {int} 控制台窗口左上角的纵坐标

**返回** {*Console} 控制台实例指针

```go
c.SetWindowPosition(100, 200)
```

## SetWindowColor
<hr style="margin: 0;">

设置控制台窗口的背景颜色。

- `color` {string} 背景颜色的十六进制字符串，格式如 "#1E1F22"

**返回** {*Console} 控制台实例指针

```go
c.SetWindowColor("#1E1F22")
```

## SetTextColor
<hr style="margin: 0;">

设置控制台文字颜色。

- `color` {string} 文字颜色的十六进制字符串，格式如 "#FFFFFF"

**返回** {*Console} 控制台实例指针

```go
c.SetTextColor("#FFFFFF")
```

## SetTextSize
<hr style="margin: 0;">

设置控制台文字大小。

- `size` {int} 文字大小

```go
c.SetTextSize(50)
```

## Println
<hr style="margin: 0;">

打印文本到控制台。

- `a` {any} 要打印的参数，支持多个参数，行为类似 fmt.Println

```go
c.Println("Hello, world!")
c.Println("用户ID:", 123, "状态:", "在线")
```

## Clear
<hr style="margin: 0;">

清空控制台内容。

```go
c.Clear()
```

## Show
<hr style="margin: 0;">

显示控制台窗口。

```go
c.Show()
```

## Hide
<hr style="margin: 0;">

隐藏控制台窗口。

```go
c.Hide()
```

## IsVisible
<hr style="margin: 0;">

检查控制台是否可见。

**返回** {bool} true 表示可见，false 表示隐藏

```go
if c.IsVisible() {
    // 控制台当前可见
}
```

## Destroy
<hr style="margin: 0;">

销毁控制台实例，释放资源。

```go
c.Destroy()
```

## 示例
<hr style="margin: 0;">

```go
// 创建并配置控制台
c := console.New()
c.SetWindowPosition(50, 50)
c.SetWindowSize(700, 500)
c.SetWindowColor("#1E1F22")
c.SetTextColor("#00FF00")
c.SetTextSize(45)

// 打印日志
c.Println("控制台已就绪")

for {
    c.Println("当前时间:", time.Now().Format("2006-01-02 15:04:05"))
    utils.Sleep(1000)
}
```