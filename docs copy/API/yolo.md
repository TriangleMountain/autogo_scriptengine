# yolo - 目标检测
---
提供基于 YOLO 的目标检测功能。

以下是 `yolo` 包中定义的 `Result` 结构体及其字段说明：

| **字段名**  | **类型**    | **说明**                              |
|-------------|-------------|---------------------------------------|
| `X`         | `int`       | 检测结果的左上角 X 坐标。             |
| `Y`         | `int`       | 检测结果的左上角 Y 坐标。             |
| `Width`     | `int`       | 检测结果的宽度。                      |
| `Height`    | `int`       | 检测结果的高度。                      |
| `Label`     | `string`    | 检测到的文字内容或标签。              |
| `Score`     | `float64`   | 检测结果的置信度，取值范围为 0-1。   |
| `CenterX`   | `int`       | 检测结果的中心 X 坐标。               |
| `CenterY`   | `int`       | 检测结果的中心 Y 坐标。               |

## New
<hr style="margin: 0;">

创建一个 Yolo 实例对象。成功返回实例对象`*Yolo`，如果加载失败则返回 nil。

- `version` {string} 模型版本，目前仅支持`v5`和`v8`
- `cpuThreadNum` {int} 用于模型推理的 CPU 线程数。
- `paramPath` {string} 模型参数文件路径。
- `binPath` {string} 模型二进制文件路径。
- `labels` {string} 标签文本，多个标签使用`,`进行隔开。

```go
yolo := yolo.New("v8", 4, "/data/local/tmp/param", "/data/local/tmp/bin", "person,bicycle,car")
if yolo == nil {
    fmt.Println("模型加载失败")
    return
}
fmt.Println("模型加载成功")
```

## *Yolo.Detect
<hr style="margin: 0;">

在屏幕指定区域进行目标检测。

- `x1` {int} 区域左上角的 x 坐标
- `y1` {int} 区域左上角的 y 坐标
- `x2` {int} 区域右下角的 x 坐标，当为 0 时表示使用屏幕最大宽度
- `y2` {int} 区域右下角的 y 坐标，当为 0 时表示使用屏幕最大高度
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
results := detector.Detect(0, 0, 0, 0, 0) // 在主屏幕全屏范围内检测目标
for _, result := range results {
    fmt.Printf("检测到 %s，置信度: %.2f\n", result.Label, result.Score)
}
```

## *Yolo.DetectFromImage
<hr style="margin: 0;">

从图像对象进行目标检测。

- `img` {*image.NRGBA} 要检测的图像对象

```go
img := images.ReadFromPath("/sdcard/photo.jpg")
results := detector.DetectFromImage(img)
```

## *Yolo.DetectFromBase64
<hr style="margin: 0;">

从Base64编码的图像数据进行目标检测。

- `b64` {string} Base64编码的图像数据
- `colorStr` {string} 保留参数，可以传入空字符串

```go
results := detector.DetectFromBase64(base64String, "")
```

## *Yolo.DetectFromPath
<hr style="margin: 0;">

从图像文件路径进行目标检测。

- `path` {string} 图像文件路径
- `colorStr` {string} 保留参数，可以传入空字符串

```go
results := detector.DetectFromPath("/sdcard/photo.png", "")
```

## *Yolo.Close
<hr style="margin: 0;">

关闭 YOLO 模型实例，释放相关资源。

```go
yolo.Close()
```