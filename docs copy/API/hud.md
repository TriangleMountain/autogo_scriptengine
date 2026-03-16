# hud - 悬浮显示
---
提供悬浮显示功能，支持多实例、彩色文本显示等功能。

以下是 `hud` 包中定义的 `TextItem` 结构体及其字段说明：

| **字段名**   | **类型**      | **说明**                                      |
|--------------|---------------|-----------------------------------------------|
| `TextColor`  | `color.Color` | 文字颜色。格式如 `"#FFFFFF"`。 |
| `Text`       | `string`      | 显示的文本内容。                              |

## New
<hr style="margin: 0;">

创建一个新的 HUD 实例。

**返回** {*HUD} HUD 实例指针

```go
h := hud.New()
```

## SetPosition
<hr style="margin: 0;">

设置 HUD 的位置和大小。

- `x1` {int} 左上角横坐标
- `y1` {int} 左上角纵坐标
- `x2` {int} 右下角横坐标
- `y2` {int} 右下角纵坐标

**返回** {*HUD} HUD 实例指针

```go
h.SetPosition(100, 100, 400, 150)
```

## SetBackgroundColor
<hr style="margin: 0;">

设置 HUD 的背景颜色。

- `color` {string} 背景颜色的十六进制字符串，格式如 "#2D2D30" 或 "#2D2D3080"（带透明度）

**返回** {*HUD} HUD 实例指针

```go
h.SetBackgroundColor("#2D2D30")
h.SetBackgroundColor("#00000080")  // 半透明黑色
```

## SetTextSize
<hr style="margin: 0;">

设置 HUD 的字体大小。

- `size` {int} 字体大小（推荐范围：30-60）

**返回** {*HUD} HUD 实例指针

```go
h.SetTextSize(45)
```

## SetText
<hr style="margin: 0;">

设置 HUD 显示的文本内容（支持多色文本）。

- `items` {[]TextItem} 文本项数组，每个元素包含颜色和文本

**返回** {*HUD} HUD 实例指针

```go
h.SetText([]hud.TextItem{
    {TextColor: "#00FF00", Text: "HP: "},
    {TextColor: "#FFFFFF", Text: "100/100"},
})
```

## Show
<hr style="margin: 0;">

显示 HUD。

**返回** {*HUD} HUD 实例指针

```go
h.Show()
```

## Hide
<hr style="margin: 0;">

隐藏 HUD。

**返回** {*HUD} HUD 实例指针

```go
h.Hide()
```

## IsVisible
<hr style="margin: 0;">

检查 HUD 是否可见。

**返回** {bool} true 表示可见，false 表示隐藏

```go
if h.IsVisible() {
    // HUD 当前可见
}
```

## Destroy
<hr style="margin: 0;">

销毁 HUD 实例，释放资源。

```go
h.Destroy()
```

## 示例
<hr style="margin: 0;">

```go
// 创建并配置 HUD
h := hud.New()
h.SetPosition(50, 700, 450, 750)
h.SetBackgroundColor("#3D000080")
h.SetTextSize(45)

for {
    // 设置多色文本
    h.SetText([]hud.TextItem{
        {TextColor: "#00FF00", Text: "当前时间: "},
        {TextColor: "#FFFFFF", Text: time.Now().Format("2006-01-02 15:04:05")},
    })
    utils.Sleep(1000)
}
```