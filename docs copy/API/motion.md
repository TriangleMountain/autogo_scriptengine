# motion - 操作
---
motion模块提供了一系列模拟用户操作的函数，如点击、滑动、按键等。

## TouchDown
<hr style="margin: 0;">

模拟触摸屏按下操作。

- `x` {int} 触摸点的X坐标
- `y` {int} 触摸点的Y坐标
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.TouchDown(500, 600, 0, 0)
```

## TouchMove
<hr style="margin: 0;">

模拟触摸屏移动操作。

- `x` {int} 移动到的X坐标
- `y` {int} 移动到的Y坐标
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.TouchMove(550, 650, 0, 0)
```

## TouchUp
<hr style="margin: 0;">

模拟触摸屏抬起操作。

- `x` {int} 抬起点的X坐标
- `y` {int} 抬起点的Y坐标
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.TouchUp(550, 650, 0, 0)
```

## Click
<hr style="margin: 0;">

模拟单击操作。

- `x` {int} 单击点的X坐标
- `y` {int} 单击点的Y坐标
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.Click(500, 600, 0, 0)
```

## LongClick
<hr style="margin: 0;">

模拟长按操作。

- `x` {int} 长按点的X坐标
- `y` {int} 长按点的Y坐标
- `duration` {int} 长按持续时间（毫秒）
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.LongClick(500, 600, 1000, 0, 0)  // 长按1秒
```

## Swipe
<hr style="margin: 0;">

模拟滑动操作。

- `x1` {int} 起始点的X坐标
- `y1` {int} 起始点的Y坐标
- `x2` {int} 结束点的X坐标
- `y2` {int} 结束点的Y坐标
- `duration` {int} 滑动持续时间（毫秒）
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.Swipe(300, 800, 300, 200, 500, 0, 0)  // 从下往上滑动
```

## Swipe2
<hr style="margin: 0;">

使用贝塞尔曲线方式进行滑动（轨迹更加自然）。

- `x1` {int} 起始点的X坐标
- `y1` {int} 起始点的Y坐标
- `x2` {int} 结束点的X坐标
- `y2` {int} 结束点的Y坐标
- `duration` {int} 滑动持续时间（毫秒）
- `fingerID` {int} 触摸点的指针ID（0-9）
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.Swipe2(300, 800, 300, 200, 500, 0, 0)  // 从下往上滑动，轨迹更自然
```

## Home
<hr style="margin: 0;">

模拟按下Home键。

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.Home(0)
```

## Back
<hr style="margin: 0;">

模拟按下返回键。

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.Back(0)
```

## Recents
<hr style="margin: 0;">

显示最近任务。

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.Recents(0)
```

## PowerDialog
<hr style="margin: 0;">

弹出电源键菜单。

```go
motion.PowerDialog()
```

## Notifications
<hr style="margin: 0;">

拉出通知栏。

```go
motion.Notifications()
```

## QuickSettings
<hr style="margin: 0;">

显示快速设置（下拉通知栏到底）。

```go
motion.QuickSettings()
```

## VolumeUp
<hr style="margin: 0;">

按下音量上键。

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.VolumeUp(0)
```

## VolumeDown
<hr style="margin: 0;">

按下音量下键。

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.VolumeDown(0)
```

## Camera
<hr style="margin: 0;">

模拟按下照相键。

```go
motion.Camera()
```

## KeyAction
<hr style="margin: 0;">

模拟按下指定按键。

- `code` {int} 按键代码，参考KEYCODE_常量
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
motion.KeyAction(motion.KEYCODE_ENTER, 0)  // 按下回车键
``` 