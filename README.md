# AutoGo ScriptEngine

[AutoGo](https://github.com/Dasongzi1366/AutoGo) 的脚本引擎扩展方案，为 AutoGo 提供 JavaScript 和 Lua 脚本语言支持，让开发者可以用熟悉的脚本语言编写自动化任务。

## 为什么选择 ScriptEngine

1. **降低准入门槛** - 使用脚本语言开发，无需深入理解 Go 语言和 Android 开发，降低学习成本
2. **代码保护** - 脚本代码易于混淆加密，有效保护业务逻辑
3. **热更新支持** - 脚本可动态加载，无需重新编译即可更新功能

## 功能特性

- **双引擎支持**：同时支持 JavaScript 和 Lua 脚本语言
- **丰富的 API**：提供应用管理、设备控制、图像识别、OCR 等多种功能
- **方法注册系统**：支持动态注册、重写和恢复方法
- **协程支持**：Lua 引擎支持协程操作
- **文档生成**：可自动生成 API 文档

### 支持的功能模块

| 模块 | 说明 |
|------|------|
| `app` | 应用管理（启动、安装、卸载、强制停止等） |
| `device` | 设备信息（分辨率、SDK 版本、屏幕方向等） |
| `motion` | 触摸操作（点击、滑动、手势等） |
| `files` | 文件操作（读写、复制、删除等） |
| `images` | 图像处理（截图、找色、找图等） |
| `storages` | 数据存储（键值对存储） |
| `system` | 系统功能（剪贴板、通知等） |
| `http` | 网络请求（GET、POST 等） |
| `media` | 媒体控制（音量、播放等） |
| `opencv` | 计算机视觉（图像处理、特征检测等） |
| `ppocr` | OCR 文字识别 |

## 环境要求

- Go 1.25.4 或更高版本
- AutoGo 框架（已在项目中集成）
- Android 设备（用于实际运行自动化脚本）

## 安装

```bash
go get github.com/ZingYao/autogo_scriptengine
```

## 快速开始

### JavaScript 引擎

```go
package main

import (
    "fmt"
    js_engine "github.com/ZingYao/autogo_scriptengine/js_engine"
)

func main() {
    // 获取 JavaScript 引擎实例
    engine := js_engine.GetEngine()
    defer js_engine.Close()

    // 执行 JavaScript 脚本
    script := `
        // 获取当前应用包名
        var packageName = app.currentPackage();
        console.log("当前应用: " + packageName);

        // 获取设备信息
        console.log("设备分辨率: " + device.width + "x" + device.height);

        // 点击屏幕
        click(500, 1000, 1);

        // 延时
        sleep(1000);
    `

    err := engine.ExecuteString(script)
    if err != nil {
        fmt.Printf("执行错误: %v\n", err)
    }
}
```

### Lua 引擎

```go
package main

import (
    "fmt"
    lua_engine "github.com/ZingYao/autogo_scriptengine/lua_engine"
)

func main() {
    // 获取 Lua 引擎实例
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()

    // 执行 Lua 脚本
    script := `
        -- 获取当前应用包名
        local packageName = app_currentPackage()
        print("当前应用: " .. packageName)

        -- 获取设备信息
        print("设备分辨率: " .. device.width .. "x" .. device.height)

        -- 点击屏幕
        click(500, 1000, 1)

        -- 延时
        sleep(1000)
    `

    err := engine.ExecuteString(script)
    if err != nil {
        fmt.Printf("执行错误: %v\n", err)
    }
}
```

## 执行脚本文件

### JavaScript

```go
engine := js_engine.GetEngine()
defer js_engine.Close()

// 执行脚本文件
err := engine.ExecuteFile("/path/to/script.js")
if err != nil {
    fmt.Printf("执行错误: %v\n", err)
}
```

### Lua

```go
engine := lua_engine.GetEngine()
defer lua_engine.Close()

// 执行脚本文件
err := engine.ExecuteFile("/path/to/script.lua")
if err != nil {
    fmt.Printf("执行错误: %v\n", err)
}
```

## 高级用法

### 方法注册

```go
engine := js_engine.GetEngine()

// 注册自定义方法
engine.RegisterMethod("myMethod", "我的自定义方法", func(param string) string {
    return "处理结果: " + param
}, true)
```

### 方法重写（Lua 示例）

```lua
-- 重写 click 方法，添加日志
local originalClick = click

function click(x, y, fingerID)
    print("点击坐标: (" .. x .. ", " .. y .. ")")
    originalClick(x, y, fingerID)
end
```

### Lua 协程

```lua
-- 创建协程
local coId = createCoroutine([[
    print("协程开始")
    yieldCoroutine(1)
    print("协程继续")
]])

-- 恢复协程执行
local results, status = resumeCoroutine(coId)
print("协程状态: " .. status)
```

## API 示例

### 应用管理

```javascript
// JavaScript
app.launch("com.example.app", 0);        // 启动应用
app.forceStop("com.example.app");        // 强制停止
app.isInstalled("com.example.app");      // 检查是否安装
app.uninstall("com.example.app");        // 卸载应用
```

```lua
-- Lua
app_launch("com.example.app", 0)         -- 启动应用
app_forceStop("com.example.app")         -- 强制停止
app_isInstalled("com.example.app")       -- 检查是否安装
app_uninstall("com.example.app")         -- 卸载应用
```

### 图像识别

```javascript
// JavaScript - 找色
var result = images.findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0);
if (result.x !== -1) {
    click(result.x, result.y, 1);
}

// OCR 文字识别
var text = ppocr.ocr(0, 0, 1080, 1920, "");
console.log(text);
```

```lua
-- Lua - 找色
local x, y = images_findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0)
if x ~= -1 then
    click(x, y, 1)
end

-- OCR 文字识别
local results = ppocr_ocr(0, 0, 1080, 1920, "")
for i, result in ipairs(results) do
    print(result["标签"])
end
```

### 文件操作

```lua
-- Lua
files_write("/sdcard/test.txt", "Hello World")  -- 写入文件
local content = files_read("/sdcard/test.txt")   -- 读取文件
files_exists("/sdcard/test.txt")                 -- 检查是否存在
```

### 数据存储

```lua
-- Lua
storages_put("myData", "key1", "value1")         -- 存储数据
local value = storages_get("myData", "key1")     -- 获取数据
storages_contains("myData", "key1")              -- 检查键是否存在
```

## 生成 API 文档

项目提供了自动生成 API 文档的功能：

```go
// JavaScript 引擎文档生成
docGen := js_engine.NewDocumentationGenerator()
docGen.SaveJSDocumentation("js_api.js")
docGen.SaveMarkdownDocumentation("js_api.md")

// Lua 引擎文档生成
docGen := lua_engine.NewDocumentationGenerator()
docGen.SaveLuaDocumentation("lua_api.lua")
docGen.SaveMarkdownDocumentation("lua_api.md")
```

## 项目结构

```
autogo_scriptengine/
├── go.mod
├── go.sum
├── js_engine/                 # JavaScript 引擎
│   ├── js_engine.go          # 引擎核心
│   ├── app_inject.go         # 应用管理 API
│   ├── device_inject.go      # 设备信息 API
│   ├── motion_inject.go      # 触摸操作 API
│   ├── files_inject.go       # 文件操作 API
│   ├── images_inject.go      # 图像处理 API
│   ├── storages_inject.go    # 数据存储 API
│   ├── system_inject.go      # 系统功能 API
│   ├── https_inject.go       # 网络请求 API
│   ├── media_inject.go       # 媒体控制 API
│   ├── opencv_inject.go      # OpenCV API
│   ├── ppocr_inject.go       # OCR API
│   └── ...
└── lua_engine/               # Lua 引擎
    ├── lua_engine.go         # 引擎核心
    ├── coroutine.go          # 协程支持
    ├── app_inject.go         # 应用管理 API
    ├── device_inject.go      # 设备信息 API
    ├── motion_inject.go      # 触摸操作 API
    └── ...
```

## 依赖

- [AutoGo](https://github.com/Dasongzi1366/AutoGo) - Android 自动化框架（核心依赖）
- [goja](https://github.com/dop251/goja) - JavaScript 解释器
- [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 解释器

## 与 AutoGo 的关系

本项目是 AutoGo 的扩展方案，通过封装 AutoGo 提供的原生 API，为开发者提供更灵活的脚本编写方式：

- **AutoGo** - 提供 Android 自动化的核心能力（无障碍服务、图像识别、触摸模拟等）
- **ScriptEngine** - 为 AutoGo 添加脚本语言支持，让开发者可以用 JavaScript 或 Lua 编写自动化脚本

## 许可证

MIT License
