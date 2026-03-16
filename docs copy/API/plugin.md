# plugin - APK插件加载
---
plugin模块提供了动态加载外部APK插件的功能，通过JNI实现Go代码与Java代码的交互。可以加载APK中的类、创建实例并调用方法，支持参数自动转换和Bitmap内存管理。

## 类型说明

### Context
<hr style="margin: 0;">

`Context` 是一个占位符类型，用于表示Android Context对象。当作为参数传递时，plugin会自动获取Android应用上下文。

```go
type Context struct{}
```

**使用示例：**
```go
// 传递Context参数给Java方法
instance.Call("methodNeedContext", plugin.Context{}, "otherParam")
```

### AssetManager
<hr style="margin: 0;">

`AssetManager` 是一个占位符类型，用于表示Android AssetManager对象。当作为参数传递时，plugin会自动创建AssetManager实例并添加ClassLoader中保存的APK路径。

```go
type AssetManager struct{}
```

**使用示例：**
```go
// 创建AssetManager参数
assetMgr := plugin.NewAssetManager()

// 传递AssetManager参数给Java方法
instance.Call("loadResources", assetMgr, "resource_name")
```

### ClassLoader
<hr style="margin: 0;">

`ClassLoader` 表示一个已加载的APK类加载器，用于创建APK中的类实例。

**方法：**
- `NewInstance(className string, args ...interface{}) (*Instance, error)` - 创建类实例
- `Release()` - 释放类加载器资源

### Instance
<hr style="margin: 0;">

`Instance` 表示一个Java对象实例，用于调用对象的方法。

**方法：**
- `Call(methodName string, args ...interface{}) (interface{}, error)` - 调用实例方法
- `CallInt(methodName string, args ...interface{}) (int, error)` - 调用返回int的方法
- `CallLong(methodName string, args ...interface{}) (int64, error)` - 调用返回long的方法
- `CallFloat(methodName string, args ...interface{}) (float32, error)` - 调用返回float的方法
- `CallDouble(methodName string, args ...interface{}) (float64, error)` - 调用返回double的方法
- `CallBool(methodName string, args ...interface{}) (bool, error)` - 调用返回boolean的方法
- `CallString(methodName string, args ...interface{}) (string, error)` - 调用返回String的方法
- `CallVoid(methodName string, args ...interface{}) error` - 调用无返回值的方法
- `Release()` - 释放实例资源

## LoadApk
<hr style="margin: 0;">

加载外部APK插件文件，返回一个类加载器对象。加载时会自动提取APK中与当前Go二进制架构对应的SO库到Go二进制所在目录。

- `apkPath` {string} APK文件的绝对路径

**返回值：**
- `*ClassLoader` 类加载器对象，加载失败返回 `nil`

**架构映射：**
- `arm64` → `arm64-v8a`
- `amd64` → `x86_64`
- `386` → `x86`

```go
cl := plugin.LoadApk("/data/local/tmp/assets/my_plugin.apk")
if cl != nil {
    defer cl.Release()
    fmt.Println("APK加载成功")
}
```

## NewInstance
<hr style="margin: 0;">

创建APK中指定类的实例。支持普通类和静态内部类（使用`$`分隔）。

- `className` {string} 完整的类名，例如 "com.example.MyClass" 或 "com.example.MyClass$InnerClass"
- `args` {...interface{}} 构造函数参数（可选）

**返回值：**
- `*Instance` 实例对象
- `error` 错误信息，成功返回 `nil`

**支持的参数类型：**
- `int`, `int32` → Java `int`
- `int64` → Java `long`
- `float32` → Java `float`
- `float64` → Java `double`
- `bool` → Java `boolean`
- `string` → Java `String`
- `[]byte` → Java `byte[]`
- `*image.NRGBA` → `android.graphics.Bitmap`（自动转换）
- `plugin.Context` → `android.content.Context`（自动获取）
- `plugin.AssetManager` → `android.content.res.AssetManager`（自动创建并添加APK路径）

```go
// 创建无参构造的实例
instance, err := cl.NewInstance("com.example.MyClass")
if err != nil {
    fmt.Println("创建实例失败:", err)
    return
}
defer instance.Release()

// 创建带参数的实例
instance, err := cl.NewInstance("com.example.MyClass", "param1", int64(123), true)

// 创建静态内部类实例
instance, err := cl.NewInstance("com.example.MyClass$InnerClass")
```

## Call
<hr style="margin: 0;">

调用实例的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `interface{}` 方法返回值（可能为 `nil`）
- `error` 错误信息，成功返回 `nil`

**注意事项：**
- 如果Java方法需要 `long` 类型参数，Go端需要显式传递 `int64` 类型，不能直接传递 `int`
- 传递 `*image.NRGBA` 作为参数时，会自动转换为 `android.graphics.Bitmap`，并在方法调用后自动调用 `recycle()` 释放内存
- 返回值类型需要根据实际情况进行类型断言或使用类型安全的调用方法

```go
// 调用无参方法
result, err := instance.Call("getStatus")

// 调用带参数的方法
result, err := instance.Call("process", "input", int64(100))

// 传递图像参数（自动转为Bitmap）
img := images.CaptureScreen(0, 0, 0, 0, 0)
result, err := instance.Call("analyzeImage", img)

// 传递Context参数
result, err := instance.Call("methodNeedContext", plugin.Context{}, "param2")

// 传递AssetManager参数
assetMgr := plugin.NewAssetManager()
result, err := instance.Call("loadAsset", assetMgr, "asset_file.txt")
```

## CallInt
<hr style="margin: 0;">

调用返回 `int` 类型的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `int` 方法返回的整数值
- `error` 错误信息，成功返回 `nil`

```go
count, err := instance.CallInt("getCount")
if err != nil {
    fmt.Println("调用失败:", err)
} else {
    fmt.Printf("计数: %d\n", count)
}
```

## CallLong
<hr style="margin: 0;">

调用返回 `long` 类型的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `int64` 方法返回的长整数值
- `error` 错误信息，成功返回 `nil`

```go
timestamp, err := instance.CallLong("getTimestamp")
if err != nil {
    fmt.Println("调用失败:", err)
} else {
    fmt.Printf("时间戳: %d\n", timestamp)
}
```

## CallFloat
<hr style="margin: 0;">

调用返回 `float` 类型的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `float32` 方法返回的浮点数值
- `error` 错误信息，成功返回 `nil`

```go
score, err := instance.CallFloat("getScore")
if err != nil {
    fmt.Println("调用失败:", err)
} else {
    fmt.Printf("分数: %.2f\n", score)
}
```

## CallDouble
<hr style="margin: 0;">

调用返回 `double` 类型的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `float64` 方法返回的双精度浮点数值
- `error` 错误信息，成功返回 `nil`

```go
ratio, err := instance.CallDouble("getRatio")
if err != nil {
    fmt.Println("调用失败:", err)
} else {
    fmt.Printf("比率: %.4f\n", ratio)
}
```

## CallBool
<hr style="margin: 0;">

调用返回 `boolean` 类型的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `bool` 方法返回的布尔值
- `error` 错误信息，成功返回 `nil`

```go
isValid, err := instance.CallBool("validate", "input")
if err != nil {
    fmt.Println("调用失败:", err)
} else if isValid {
    fmt.Println("验证通过")
}
```

## CallString
<hr style="margin: 0;">

调用返回 `String` 类型的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `string` 方法返回的字符串
- `error` 错误信息，成功返回 `nil`

```go
name, err := instance.CallString("getName")
if err != nil {
    fmt.Println("调用失败:", err)
} else {
    fmt.Printf("名称: %s\n", name)
}
```

## CallVoid
<hr style="margin: 0;">

调用无返回值的方法。

- `methodName` {string} 方法名称
- `args` {...interface{}} 方法参数（可选）

**返回值：**
- `error` 错误信息，成功返回 `nil`

```go
err := instance.CallVoid("initialize", "config.xml")
if err != nil {
    fmt.Println("初始化失败:", err)
}
```

## Release (ClassLoader)
<hr style="margin: 0;">

释放类加载器占用的资源。调用后不能再使用该类加载器创建新实例。

```go
cl := plugin.LoadApk("/path/to/plugin.apk")
if cl != nil {
    defer cl.Release()
}
```

## Release (Instance)
<hr style="margin: 0;">

释放实例占用的资源。调用后不能再使用该实例调用方法。

```go
instance, err := cl.NewInstance("com.example.MyClass")
if err == nil {
    defer instance.Release()
}
```

## 完整示例
<hr style="margin: 0;">

```go
package main

import (
    "fmt"
    "github.com/Dasongzi1366/AutoGo/plugin"
    "github.com/Dasongzi1366/AutoGo/images"
)

func main() {
    // 加载APK插件
    cl := plugin.LoadApk("/data/local/tmp/assets/ocr_plugin.apk")
    if cl == nil {
        fmt.Println("加载APK失败")
        return
    }
    defer cl.Release()

    // 创建OCR引擎实例
    engine, err := cl.NewInstance("com.ddddocr.OCREngine")
    if err != nil {
        fmt.Printf("创建实例失败: %v\n", err)
        return
    }
    defer engine.Release()

    // 截取屏幕
    img := images.CaptureScreen(0, 0, 0, 0, 0)
    
    // 调用OCR识别（图像自动转为Bitmap）
    result, err := engine.CallString("recognition", img)
    if err != nil {
        fmt.Printf("识别失败: %v\n", err)
        return
    }
    
    fmt.Printf("识别结果: %s\n", result)
    
    // 调用需要Context的方法
    err = engine.CallVoid("saveToCache", plugin.Context{}, "data.txt")
    if err != nil {
        fmt.Printf("保存失败: %v\n", err)
    }
    
    // 使用AssetManager加载APK中的资源
    assetMgr := plugin.NewAssetManager()
    resourceData, err := engine.CallString("loadResource", assetMgr, "config.json")
    if err != nil {
        fmt.Printf("加载资源失败: %v\n", err)
    } else {
        fmt.Printf("资源内容: %s\n", resourceData)
    }
}
```

## 类型转换说明
<hr style="margin: 0;">

### Go类型 → Java类型映射

| Go类型 | Java类型 | 说明 |
|--------|----------|------|
| `int`, `int32` | `int` | 32位整数 |
| `int64` | `long` | 64位整数，**注意必须显式使用int64** |
| `float32` | `float` | 单精度浮点数 |
| `float64` | `double` | 双精度浮点数 |
| `bool` | `boolean` | 布尔值 |
| `string` | `String` | 字符串 |
| `[]byte` | `byte[]` | 字节数组 |
| `*image.NRGBA` | `android.graphics.Bitmap` | 图像（自动转换，自动释放） |
| `plugin.Context` | `android.content.Context` | 应用上下文（自动获取） |
| `plugin.AssetManager` | `android.content.res.AssetManager` | 资源管理器（自动创建并添加APK路径） |
