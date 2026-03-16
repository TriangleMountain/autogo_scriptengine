# vdisplay - 虚拟屏幕
---
vdisplay模块提供了虚拟屏幕的创建和管理功能，需要Android 10及以上版本。通过该模块可以创建虚拟屏幕，用于在虚拟屏幕中启动和操作应用程序，实现多应用并行操作和完美全分辨率。

## Create
<hr style="margin: 0;">

创建一个虚拟屏幕。

- `width` {int} 屏幕宽度
- `height` {int} 屏幕高度
- `dpi` {int} 屏幕的像素密度

**返回值：**
- `*Vdisplay` 虚拟屏幕对象，如果创建失败则返回 `nil`

**注意：**
- 可以通过Scrcpy执行命令查看虚拟屏幕的实时画面：`scrcpy --display-id=虚拟屏幕ID`

```go
v := vdisplay.Create(720, 1280, 320)
if v != nil {
    fmt.Printf("创建成功,屏幕ID:%d\n", v.GetDisplayId())
}
```

## GetDisplayId
<hr style="margin: 0;">

获取虚拟屏幕的ID。

```go
displayId := v.GetDisplayId()
fmt.Printf("虚拟屏幕ID: %d\n", displayId)
```

## LaunchApp
<hr style="margin: 0;">

在虚拟屏幕中启动指定的应用程序。

- `packageName` {string} 要启动的应用包名

```go
success := v.LaunchApp("com.example.app")
if success {
    fmt.Println("应用启动成功")
}
```

## SetTitle
<hr style="margin: 0;">

设置虚拟屏预览窗口的标题。

- `title` {string} 标题

```go
v.SetTitle("MT管理器")
```

## SetTouchCallback
<hr style="margin: 0;">

设置一个点击回调。

- `callback` {function} 当用户点击虚拟屏预览窗口时调用的函数，格式为 `func(x, y, action, displayId int)`，如果传入 `nil`，则会移除当前设置的回调

```go
v.SetTouchCallback(func(x, y, action, displayId int) {
    // 处理用户点击事件
})
```

## ShowPreviewWindow
<hr style="margin: 0;">

显示虚拟屏幕的预览窗口。预览窗口可以在ImgUI界面中显示虚拟屏幕的实时画面，并支持触摸操作。

- `rotated` {bool} 是否将预览窗口旋转90度显示(显示横屏游戏画面时可能需要)

```go
v.ShowPreviewWindow(false)
```

## HidePreviewWindow
<hr style="margin: 0;">

隐藏虚拟屏幕的预览窗口。

```go
v.HidePreviewWindow()
```

## SetPreviewWindowSize
<hr style="margin: 0;">

设置预览窗口的大小。

- `width` {int} 预览窗口宽度（像素）
- `height` {int} 预览窗口高度（像素）

```go
v.SetPreviewWindowSize(800, 600)
```

## SetPreviewWindowPos
<hr style="margin: 0;">

设置预览窗口的位置。

- `x` {int} 预览窗口X坐标
- `y` {int} 预览窗口Y坐标

```go
v.SetPreviewWindowPos(100, 100)
```

## Destroy
<hr style="margin: 0;">

销毁虚拟屏幕，释放相关资源。

```go
v.Destroy()
```