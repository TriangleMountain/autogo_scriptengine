# ppocr - 飞桨OCR
---
提供基于 PaddleOCR 的文字检测和识别功能。

以下是 `ppocr` 包中定义的 `Result` 结构体及其字段说明：

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

创建一个 Ppocr 实例对象。成功返回实例对象`*Ppocr`，如果加载失败则返回 nil。

- `version` {string} 模型版本，目前仅支持`v2`和`v5`

```go
ocr := ppocr.New("v5")
if ocr == nil {
    fmt.Println("初始化失败")
    return
}
fmt.Println("初始化成功")
```

## *Ppocr.Ocr
<hr style="margin: 0;">

在屏幕指定区域进行OCR文字识别。

- `x1` {int} 区域左上角的 x 坐标
- `y1` {int} 区域左上角的 y 坐标
- `x2` {int} 区域右下角的 x 坐标，当为 0 时表示使用屏幕最大宽度
- `y2` {int} 区域右下角的 y 坐标，当为 0 时表示使用屏幕最大高度
- `colorStr` {string} 文字颜色范围，格式如 "FFFFFF" 或 "FFFFFF-101010"，"-" 后表示偏色范围，如果不需要指定则直接传入空字符串`""`
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
results := ppocr.Ocr(0, 0, 0, 0, "000000", 0) // 识别主屏幕全屏的黑色文字
for _, result := range results {
    fmt.Println(result.Label) // 打印识别到的文字
}
```

## *Ppocr.OcrFromImage
<hr style="margin: 0;">

从图像对象进行OCR文字识别。

- `img` {*image.NRGBA} 要识别的图像对象
- `colorStr` {string} 文字颜色范围，格式如 "FFFFFF" 或 "FFFFFF-101010"

```go
img := images.ReadFromPath("/sdcard/screenshot.png")
results := ppocr.OcrFromImage(img, "000000")
```

## *Ppocr.OcrFromBase64
<hr style="margin: 0;">

从Base64编码的图像数据进行OCR文字识别。

- `b64` {string} Base64编码的图像数据
- `colorStr` {string} 文字颜色范围，格式如 "FFFFFF" 或 "FFFFFF-101010"

```go
results := ppocr.OcrFromBase64(base64String, "000000")
```

## *Ppocr.OcrFromPath
<hr style="margin: 0;">

从图像文件路径进行OCR文字识别。

- `path` {string} 图像文件路径
- `colorStr` {string} 文字颜色范围，格式如 "FFFFFF" 或 "FFFFFF-101010"

```go
results := ppocr.OcrFromPath("/sdcard/screenshot.png", "000000")
```

## *Ppocr.Close
<hr style="margin: 0;">

关闭 PPOCR 实例。

```go
ppocr.Close()
```