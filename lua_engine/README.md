# AutoGo Lua Engine

AutoGo Lua Engine 是一个高性能的 Lua 脚本引擎，为 AutoGo 框架提供了完整的 Lua 脚本支持。它允许用户使用 Lua 语言编写自动化脚本，并提供了丰富的 API 接口。

## 特性

- **完整的 API 注入**: 将 AutoGo 的所有功能模块注入到 Lua 引擎中
- **方法管理**: 支持动态注册、移除、列出方法
- **方法重写**: 允许用 Lua 函数重写已注册的方法
- **线程安全**: 所有操作都是线程安全的
- **丰富的功能**: 支持应用管理、设备控制、触摸操作、文件操作、图像处理、文字识别等

## 快速开始

### 初始化引擎

```go
import "github.com/Dasongzi1366/AutoGo/lua_engine"

func main() {
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    // 使用引擎...
}
```

### 执行 Lua 代码

```go
// 执行 Lua 字符串
err := lua_engine.ExecuteString(`
    local packageName = app_currentPackage()
    print("当前应用包名: " .. packageName)
    
    click(500, 1000, 1)
`)

// 执行 Lua 文件
err = lua_engine.ExecuteFile("/path/to/script.lua")
```

## API 模块

### 基础函数

```lua
-- 延迟执行
sleep(1000) -- 延迟 1000 毫秒
```

### 应用管理 (app)

```lua
-- 获取当前应用包名
local packageName = app_currentPackage()

-- 获取当前应用类名
local activity = app_currentActivity()

-- 启动应用
app_launch("com.example.app", 0)

-- 打开应用设置
app_openAppSetting("com.example.app")

-- 查看文件
app_viewFile("/sdcard/test.txt")

-- 编辑文件
app_editFile("/sdcard/test.txt")

-- 卸载应用
app_uninstall("com.example.app")

-- 安装应用
app_install("/sdcard/app.apk")

-- 检查应用是否已安装
if app_isInstalled("com.example.app") then
    print("应用已安装")
end

-- 清除应用数据
app_clear("com.example.app")

-- 强制停止应用
app_forceStop("com.example.app")

-- 禁用应用
app_disable("com.example.app")

-- 忽略电池优化
app_ignoreBattOpt("com.example.app")

-- 启用应用
app_enable("com.example.app")

-- 获取默认浏览器包名
local browser = app_getBrowserPackage()

-- 打开 URL
app_openUrl("https://example.com")
```

### 设备管理 (device)

```lua
-- 获取设备信息
print("分辨率: " .. device.width(0) .. "x" .. device.height(0))
print("SDK 版本: " .. device.sdkInt())
print("CPU 架构: " .. device.cpuAbi())

-- 获取设备标识
local imei = device_getImei()
local androidId = device_getAndroidId()
local wifiMac = device_getWifiMac()
local wlanMac = device_getWlanMac()
local ip = device_getIp()

-- 音量控制
local musicVolume = device_getMusicVolume()
device_setMusicVolume(50)

-- 电池信息
local battery = device_getBattery()
local batteryStatus = device_getBatteryStatus()

-- 屏幕控制
if device_isScreenOn() then
    print("屏幕已点亮")
end

if device_isScreenUnlock() then
    print("屏幕已解锁")
end

device_wakeUp()
device_keepScreenOn()

-- 震动
device_vibrate(1000)
device_cancelVibration()
```

### 触摸操作 (touch)

```lua
-- 基本触摸操作
touchDown(500, 1000, 1, 0)
touchMove(600, 1100, 1, 0)
touchUp(600, 1100, 1, 0)

-- 点击
click(500, 1000, 1, 0)

-- 长按
longClick(500, 1000, 2000, 1, 0)

-- 滑动
swipe(500, 1000, 600, 1100, 500, 1, 0)

-- 系统按键
home(0)
back(0)
recents(0)
powerDialog()
notifications()
quickSettings()
volumeUp(0)
volumeDown(0)
camera()

-- 自定义按键
keyAction(3, 0) -- KEYCODE_HOME
```

### 文件操作 (files)

```lua
-- 检查文件/文件夹
if files_isFile("/sdcard/test.txt") then
    print("是文件")
end

if files_isDir("/sdcard/Download") then
    print("是文件夹")
end

-- 创建文件/文件夹
files_create("/sdcard/test.txt")

-- 检查文件是否存在
if files_exists("/sdcard/test.txt") then
    print("文件存在")
end

-- 读写文件
local content = files_read("/sdcard/test.txt")
files_write("/sdcard/test.txt", "Hello, World!")
files_append("/sdcard/test.txt", "\nAppend text")

-- 复制/移动/重命名/删除
files_copy("/sdcard/source.txt", "/sdcard/dest.txt")
files_move("/sdcard/source.txt", "/sdcard/newlocation.txt")
files_rename("/sdcard/old.txt", "/sdcard/new.txt")
files_remove("/sdcard/test.txt")

-- 获取文件信息
local name = files_getName("/sdcard/test.txt")
local nameWithoutExt = files_getNameWithoutExtension("/sdcard/test.txt")
local ext = files_getExtension("/sdcard/test.txt")

-- 获取绝对路径
local absPath = files_path("./test.txt")

-- 列出目录
local files = files_listDir("/sdcard")
for i, file in ipairs(files) do
    print(file)
end
```

### 图像处理 (images)

```lua
-- 获取像素颜色
local color = images_pixel(500, 1000, 0)
print("颜色: " .. color)

-- 截取屏幕
local img = images_captureScreen(0, 0, 1080, 1920, 0)

-- 比较颜色
if images_cmpColor(500, 1000, "#FF0000", 0.9, 0) then
    print("颜色匹配")
end

-- 查找颜色
local x, y = images_findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0)
if x ~= -1 and y ~= -1 then
    print("找到颜色在: " .. x .. ", " .. y)
    click(x, y, 1, 0)
end

-- 获取颜色数量
local count = images_getColorCountInRegion(0, 0, 1080, 1920, "#FF0000", 0.9, 0)
print("颜色数量: " .. count)

-- 多点颜色检测
if images_detectsMultiColors("0,0,#FF0000,10,10,#00FF00", 0.9, 0) then
    print("多点颜色匹配")
end

-- 查找多点颜色
local x, y = images_findMultiColors(0, 0, 1080, 1920, "0,0,#FF0000,10,10,#00FF00", 0.9, 0, 0)

-- 读取图片
local img = images_readFromPath("/sdcard/image.png")
local img = images_readFromUrl("https://example.com/image.png")
local img = images_readFromBase64("iVBORw0KG...")
local img = images_readFromBytes(data)

-- 保存图片
images_save(img, "/sdcard/output.png", 90)

-- 编码图片
local base64 = images_encodeToBase64(img, "png", 90)
local bytes = images_encodeToBytes(img, "png", 90)

-- 图像处理
local clipped = images_clip(img, 100, 100, 200, 200)
local resized = images_resize(img, 500, 500)
local rotated = images_rotate(img, 90)
local gray = images_grayscale(img)
local threshold = images_applyThreshold(img, 128, 255, "BINARY")
local adaptive = images_applyAdaptiveThreshold(img, 255, "GAUSSIAN_C", "BINARY", 11, 2)
local binary = images_applyBinarization(img, 128)
```

### 存储管理 (storages)

```lua
-- 存储键值对
storages_put("myTable", "key1", "value1")
storages_put("myTable", "key2", "value2")

-- 读取键值
local value = storages_get("myTable", "key1")
print("key1 = " .. value)

-- 检查键是否存在
if storages_contains("myTable", "key2") then
    print("key2 存在")
end

-- 获取所有键值对
local allData = storages_getAll("myTable")
for k, v in pairs(allData) do
    print(k .. " = " .. v)
end

-- 删除键
storages_remove("myTable", "key1")

-- 清空表
storages_clear("myTable")
```

### 系统管理 (system)

```lua
-- 获取进程 ID
local pid = system_getPid("myapp")

-- 获取内存使用
local memory = system_getMemoryUsage(pid)

-- 获取 CPU 使用率
local cpu = system_getCpuUsage(pid)

-- 重启自身
system_restartSelf()

-- 设置开机自启
system_setBootStart(true)
```

### 网络请求 (http)

```lua
-- GET 请求
local code, data = http_get("https://example.com", 5000)
print("状态码: " .. code)
print("响应: " .. data)

-- POST Multipart 请求
local fileData = files_readBytes("/sdcard/image.png")
local code, data = http_postMultipart("https://example.com/upload", "image.png", fileData, 5000)
```

### 媒体管理 (media)

```lua
-- 扫描媒体文件
media_scanFile("/sdcard/image.png")
```

### 图像识别 (opencv)

```lua
-- 查找图片
local template = files_readBytes("/sdcard/template.png")
local x, y = opencv_findImage(0, 0, 1080, 1920, template, false, 1.0, 0.8, 0)
if x ~= -1 and y ~= -1 then
    print("找到图片在: " .. x .. ", " .. y)
    click(x, y, 1, 0)
end
```

### 文字识别 (ppocr)

```lua
-- 识别屏幕文字
local results = ppocr_ocr(0, 0, 1080, 1920, "")
for i, result in ipairs(results) do
    print("文本: " .. result["标签"])
    print("位置: (" .. result["X"] .. ", " .. result["Y"] .. ")")
    print("大小: " .. result["宽"] .. "x" .. result["高"])
    print("精度: " .. result["精度"])
    print("中心: (" .. result["CenterX"] .. ", " .. result["CenterY"] .. ")")
end

-- 识别 Base64 图片
local results = ppocr_ocrFromBase64(base64Str, "")

-- 识别文件图片
local results = ppocr_ocrFromPath("/sdcard/image.png", "")
```

## 方法管理

### 注册新方法

```lua
registerMethod("myMethod", "我的自定义方法", nil, true)

function myMethod(param)
    print("自定义方法被调用: " .. param)
    return "返回值"
end
```

### 移除方法

```lua
local success = unregisterMethod("myMethod")
```

### 列出所有方法

```lua
local methods = listMethods()
for i, method in ipairs(methods) do
    print(method["name"] .. " - " .. method["description"])
    print("可重写: " .. tostring(method["overridable"]))
    print("已重写: " .. tostring(method["overridden"]))
end
```

### 重写方法

```lua
-- 方法 1: 直接重写
local originalClick = click

function click(x, y, fingerID)
    print("点击: (" .. x .. ", " .. y .. ")")
    originalClick(x, y, fingerID)
end

-- 方法 2: 使用 overrideMethod
overrideMethod("click", function(x, y, fingerID)
    print("点击: (" .. x .. ", " .. y .. ")")
    -- 调用原始实现
end)
```

### 恢复方法

```lua
local success = restoreMethod("click")
```

## 生成文档

### 生成 Lua 文档

```go
docGen := lua_engine.NewDocumentationGenerator()
err := docGen.SaveLuaDocumentation("lua_api.lua")
```

### 生成 Markdown 文档

```go
docGen := lua_engine.NewDocumentationGenerator()
err := docGen.SaveMarkdownDocumentation("lua_api.md")
```

## 完整示例

### 将脚本嵌入到程序并运行

以下是一个完整的 Demo，展示如何将 Lua 脚本嵌入到 Go 程序中，打包到产物中，运行时释放并执行：

#### 1. 使用 Go embed 嵌入脚本文件（Go 1.16+）

```go
package main

import (
    "embed"
    "fmt"
    "os"
    "path/filepath"
    
    "github.com/Dasongzi1366/AutoGo/files"
    "github.com/Dasongzi1366/AutoGo/lua_engine"
)

//go:embed scripts/*
var scriptFS embed.FS

// EmbeddedScriptManager 嵌入式脚本管理器
type EmbeddedScriptManager struct {
    scriptDir  string
    extracted bool
}

// NewEmbeddedScriptManager 创建嵌入式脚本管理器
func NewEmbeddedScriptManager(scriptDir string) *EmbeddedScriptManager {
    return &EmbeddedScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    }
}

// ExtractScripts 提取嵌入的脚本到文件系统
func (esm *EmbeddedScriptManager) ExtractScripts() error {
    if esm.extracted {
        return nil
    }
    
    // 确保目标目录存在
    if !files.Exists(esm.scriptDir) {
        if err := files.Create(esm.scriptDir); err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        }
    }
    
    // 读取嵌入的脚本目录
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return fmt.Errorf("读取嵌入目录失败: %v", err)
    }
    
    // 提取所有脚本文件
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        
        srcPath := filepath.Join("scripts", entry.Name())
        dstPath := filepath.Join(esm.scriptDir, entry.Name())
        
        // 读取嵌入的文件内容
        content, err := scriptFS.ReadFile(srcPath)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", srcPath, err)
        }
        
        // 写入到文件系统
        if err := files.Write(dstPath, string(content)); err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        }
        
        fmt.Printf("已提取脚本: %s\n", dstPath)
    }
    
    esm.extracted = true
    return nil
}

// ListScripts 列出所有嵌入的脚本
func (esm *EmbeddedScriptManager) ListScripts() ([]string, error) {
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return nil, fmt.Errorf("读取嵌入目录失败: %v", err)
    }
    
    var scripts []string
    for _, entry := range entries {
        if !entry.IsDir() {
            scripts = append(scripts, entry.Name())
        }
    }
    
    return scripts, nil
}

// GetScriptContent 获取脚本内容（不释放到文件系统）
func (esm *EmbeddedScriptManager) GetScriptContent(name string) (string, error) {
    srcPath := filepath.Join("scripts", name)
    content, err := scriptFS.ReadFile(srcPath)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", srcPath, err)
    }
    
    return string(content), nil
}

// RunScript 运行脚本（从嵌入的文件系统直接执行）
func (esm *EmbeddedScriptManager) RunScript(name string) error {
    content, err := esm.GetScriptContent(name)
    if err != nil {
        return err
    }
    
    if err := lua_engine.ExecuteString(content); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

// RunScriptFromFile 运行脚本（从释放的文件执行）
func (esm *EmbeddedScriptManager) RunScriptFromFile(name string) error {
    if !esm.extracted {
        if err := esm.ExtractScripts(); err != nil {
            return err
        }
    }
    
    scriptPath := filepath.Join(esm.scriptDir, name)
    if err := lua_engine.ExecuteFile(scriptPath); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

func main() {
    // 初始化 Lua 引擎
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    // 创建嵌入式脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewEmbeddedScriptManager(scriptDir)
    
    // 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    }
    
    for i, script := range scripts {
        fmt.Printf("  %d. %s\n", i+1, script)
    }
    
    // 方法1: 直接从嵌入的文件系统执行脚本（不释放到文件系统）
    fmt.Println("\n方法1: 直接从嵌入的文件系统执行脚本")
    if err := scriptManager.RunScript("demo.lua"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
    
    // 方法2: 提取脚本到文件系统后执行
    fmt.Println("\n方法2: 提取脚本到文件系统后执行")
    if err := scriptManager.ExtractScripts(); err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    }
    
    if err := scriptManager.RunScriptFromFile("demo.lua"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
}
```

#### 2. 脚本文件结构

```
project/
├── main.go
└── scripts/
    ├── demo.lua
    ├── auto_login.lua
    ├── find_color.lua
    └── ocr_text.lua
```

#### 3. 示例脚本文件

**scripts/demo.lua**
```lua
-- Demo 脚本
print("=== Demo 脚本开始执行 ===")

-- 获取设备信息
local width = device.width(0)
local height = device.height(0)
print("屏幕分辨率: " .. width .. "x" .. height)

-- 获取当前应用
local packageName = app_currentPackage()
print("当前应用: " .. packageName)

-- 点击屏幕中心
click(width/2, height/2, 1, 0)

print("=== Demo 脚本执行完成 ===")
return true
```

**scripts/auto_login.lua**
```lua
-- 自动登录脚本
function autoLogin(username, password)
    print("开始自动登录...")
    
    -- 检查当前应用
    local currentApp = app_currentPackage()
    if currentApp ~= "com.example.app" then
        print("启动应用...")
        app_launch("com.example.app", 0)
        sleep(3000)
    end
    
    -- 查找用户名输入框
    local x, y = images_findColor(0, 0, device.width(0), device.height(0), "#FF0000", 0.9, 0, 0)
    if x ~= -1 then
        click(x, y, 1, 0)
        sleep(500)
    end
    
    -- 查找密码输入框
    local x, y = images_findColor(0, 0, device.width(0), device.height(0), "#00FF00", 0.9, 0, 0)
    if x ~= -1 then
        click(x, y, 1, 0)
        sleep(500)
    end
    
    -- 查找登录按钮
    local x, y = images_findColor(0, 0, device.width(0), device.height(0), "#0000FF", 0.9, 0, 0)
    if x ~= -1 then
        click(x, y, 1, 0)
        sleep(2000)
    end
    
    print("登录完成")
    return true
end

return autoLogin("user123", "pass456")
```

**scripts/find_color.lua**
```lua
-- 查找颜色并点击
function findAndClick(color, sim)
    local x, y = images_findColor(0, 0, device.width(0), device.height(0), color, sim, 0, 0)
    if x ~= -1 and y ~= -1 then
        click(x, y, 1, 0)
        print("找到颜色 " .. color .. " 在: (" .. x .. ", " .. y .. ")")
        return true
    end
    print("未找到颜色: " .. color)
    return false
end

return findAndClick("#FF0000", 0.9)
```

**scripts/ocr_text.lua**
```lua
-- OCR 文字识别并点击
function findTextAndClick(text)
    local results = ppocr_ocr(0, 0, device.width(0), device.height(0), "", 0)
    for i, result in ipairs(results) do
        if string.find(result["标签"], text) then
            click(result["CenterX"], result["CenterY"], 1, 0)
            print("找到文字 '" .. text .. "' 在: (" .. result["CenterX"] .. ", " .. result["CenterY"] .. ")")
            return true
        end
    end
    print("未找到文字: " .. text)
    return false
end

return findTextAndClick("确定")
```

#### 4. 使用 go-bindata 嵌入脚本文件（兼容旧版本）

```bash
# 安装 go-bindata
go get -u github.com/go-bindata/go-bindata/...

# 生成嵌入的文件
go-bindata -o scripts.go -pkg main scripts/
```

```go
package main

import (
    "fmt"
    "path/filepath"
    
    "github.com/Dasongzi1366/AutoGo/files"
    "github.com/Dasongzi1366/AutoGo/lua_engine"
)

// BindataScriptManager 使用 go-bindata 的脚本管理器
type BindataScriptManager struct {
    scriptDir  string
    extracted bool
}

// NewBindataScriptManager 创建脚本管理器
func NewBindataScriptManager(scriptDir string) *BindataScriptManager {
    return &BindataScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    }
}

// ExtractScripts 提取嵌入的脚本到文件系统
func (bsm *BindataScriptManager) ExtractScripts() error {
    if bsm.extracted {
        return nil
    }
    
    // 确保目标目录存在
    if !files.Exists(bsm.scriptDir) {
        if err := files.Create(bsm.scriptDir); err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        }
    }
    
    // 遍历所有嵌入的文件
    for _, name := range AssetNames() {
        if filepath.Ext(name) != ".lua" {
            continue
        }
        
        // 读取嵌入的文件内容
        content, err := Asset(name)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", name, err)
        }
        
        // 写入到文件系统
        dstPath := filepath.Join(bsm.scriptDir, filepath.Base(name))
        if err := files.Write(dstPath, string(content)); err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        }
        
        fmt.Printf("已提取脚本: %s\n", dstPath)
    }
    
    bsm.extracted = true
    return nil
}

// ListScripts 列出所有嵌入的脚本
func (bsm *BindataScriptManager) ListScripts() ([]string, error) {
    var scripts []string
    for _, name := range AssetNames() {
        if filepath.Ext(name) == ".lua" {
            scripts = append(scripts, filepath.Base(name))
        }
    }
    return scripts, nil
}

// GetScriptContent 获取脚本内容
func (bsm *BindataScriptManager) GetScriptContent(name string) (string, error) {
    content, err := Asset(name)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", name, err)
    }
    return string(content), nil
}

// RunScript 运行脚本
func (bsm *BindataScriptManager) RunScript(name string) error {
    content, err := bsm.GetScriptContent(name)
    if err != nil {
        return err
    }
    
    if err := lua_engine.ExecuteString(content); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

func main() {
    // 初始化 Lua 引擎
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    // 创建脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewBindataScriptManager(scriptDir)
    
    // 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    }
    
    for i, script := range scripts {
        fmt.Printf("  %d. %s\n", i+1, script)
    }
    
    // 提取脚本到文件系统
    if err := scriptManager.ExtractScripts(); err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    }
    
    // 运行脚本
    if err := scriptManager.RunScript("scripts/demo.lua"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
}
```

#### 5. 使用 vfsgen 生成虚拟文件系统

```bash
# 安装 vfsgen
go get -u github.com/shurcooL/vfsgen
```

```go
//go:build ignore
// +build ignore

package main

import (
    "io/fs"
    "log"
    "net/http"
    "os"
    
    "github.com/shurcooL/vfsgen"
)

func main() {
    var httpFS = http.Dir("scripts")
    
    err := vfsgen.Generate(httpFS, vfsgen.Options{
        PackageName:  "main",
        BuildTags:    "!dev",
        VariableName: "scriptFS",
        Filename:     "scripts_vfs.go",
    })
    
    if err != nil {
        log.Fatalln(err)
    }
}
```

```go
package main

import (
    "embed"
    "fmt"
    "io/fs"
    "path/filepath"
    
    "github.com/Dasongzi1366/AutoGo/files"
    "github.com/Dasongzi1366/AutoGo/lua_engine"
)

//go:generate go run vfsgen.go

// scriptFS 是生成的虚拟文件系统
//go:build !dev
// +build !dev
var scriptFS fs.FS

// VfsScriptManager 使用虚拟文件系统的脚本管理器
type VfsScriptManager struct {
    scriptDir  string
    extracted bool
}

// NewVfsScriptManager 创建脚本管理器
func NewVfsScriptManager(scriptDir string) *VfsScriptManager {
    return &VfsScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    }
}

// ExtractScripts 提取嵌入的脚本到文件系统
func (vsm *VfsScriptManager) ExtractScripts() error {
    if vsm.extracted {
        return nil
    }
    
    // 确保目标目录存在
    if !files.Exists(vsm.scriptDir) {
        if err := files.Create(vsm.scriptDir); err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        }
    }
    
    // 读取虚拟文件系统
    entries, err := fs.ReadDir(scriptFS, ".")
    if err != nil {
        return fmt.Errorf("读取虚拟文件系统失败: %v", err)
    }
    
    // 提取所有脚本文件
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        
        // 读取文件内容
        content, err := fs.ReadFile(scriptFS, entry.Name())
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", entry.Name(), err)
        }
        
        // 写入到文件系统
        dstPath := filepath.Join(vsm.scriptDir, entry.Name())
        if err := files.Write(dstPath, string(content)); err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        }
        
        fmt.Printf("已提取脚本: %s\n", dstPath)
    }
    
    vsm.extracted = true
    return nil
}

// ListScripts 列出所有嵌入的脚本
func (vsm *VfsScriptManager) ListScripts() ([]string, error) {
    entries, err := fs.ReadDir(scriptFS, ".")
    if err != nil {
        return nil, fmt.Errorf("读取虚拟文件系统失败: %v", err)
    }
    
    var scripts []string
    for _, entry := range entries {
        if !entry.IsDir() && filepath.Ext(entry.Name()) == ".lua" {
            scripts = append(scripts, entry.Name())
        }
    }
    
    return scripts, nil
}

// GetScriptContent 获取脚本内容
func (vsm *VfsScriptManager) GetScriptContent(name string) (string, error) {
    content, err := fs.ReadFile(scriptFS, name)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", name, err)
    }
    return string(content), nil
}

// RunScript 运行脚本
func (vsm *VfsScriptManager) RunScript(name string) error {
    content, err := vsm.GetScriptContent(name)
    if err != nil {
        return err
    }
    
    if err := lua_engine.ExecuteString(content); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

func main() {
    // 初始化 Lua 引擎
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    // 创建脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewVfsScriptManager(scriptDir)
    
    // 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    }
    
    for i, script := range scripts {
        fmt.Printf("  %d. %s\n", i+1, script)
    }
    
    // 提取脚本到文件系统
    if err := scriptManager.ExtractScripts(); err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    }
    
    // 运行脚本
    if err := scriptManager.RunScript("demo.lua"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
}
```

#### 6. 完整的构建和部署流程

```bash
# 1. 创建脚本目录
mkdir -p scripts

# 2. 将脚本文件放入 scripts 目录
# demo.lua, auto_login.lua, find_color.lua, ocr_text.lua

# 3. 生成嵌入的文件（选择其中一种方法）

# 方法1: 使用 embed（推荐，Go 1.16+）
# 不需要额外步骤，直接使用 //go:embed 指令

# 方法2: 使用 go-bindata
go-bindata -o scripts.go -pkg main scripts/

# 方法3: 使用 vfsgen
go run vfsgen.go

# 4. 编译程序
go build -o autogo

# 5. 运行程序（脚本已嵌入到可执行文件中）
./autogo
```

#### 7. 高级用法：动态脚本加载和热更新

```go
package main

import (
    "embed"
    "fmt"
    "path/filepath"
    
    "github.com/Dasongzi1366/AutoGo/files"
    "github.com/Dasongzi1366/AutoGo/lua_engine"
)

//go:embed scripts/*
var scriptFS embed.FS

// AdvancedScriptManager 高级脚本管理器
type AdvancedScriptManager struct {
    scriptDir    string
    useEmbedded  bool
    scriptCache  map[string]string
}

// NewAdvancedScriptManager 创建高级脚本管理器
func NewAdvancedScriptManager(scriptDir string, useEmbedded bool) *AdvancedScriptManager {
    return &AdvancedScriptManager{
        scriptDir:   scriptDir,
        useEmbedded: useEmbedded,
        scriptCache: make(map[string]string),
    }
}

// LoadScript 加载脚本（优先从文件系统加载，失败则从嵌入的文件系统加载）
func (asm *AdvancedScriptManager) LoadScript(name string) (string, error) {
    // 检查缓存
    if content, ok := asm.scriptCache[name]; ok {
        return content, nil
    }
    
    var content string
    var err error
    
    // 优先从文件系统加载
    if !asm.useEmbedded {
        scriptPath := filepath.Join(asm.scriptDir, name)
        if files.Exists(scriptPath) {
            content, err = files.Read(scriptPath)
            if err == nil {
                asm.scriptCache[name] = content
                fmt.Printf("从文件系统加载脚本: %s\n", name)
                return content, nil
            }
        }
    }
    
    // 从嵌入的文件系统加载
    srcPath := filepath.Join("scripts", name)
    data, err := scriptFS.ReadFile(srcPath)
    if err != nil {
        return "", fmt.Errorf("加载脚本失败 %s: %v", name, err)
    }
    
    content = string(data)
    asm.scriptCache[name] = content
    fmt.Printf("从嵌入文件系统加载脚本: %s\n", name)
    return content, nil
}

// RunScript 运行脚本
func (asm *AdvancedScriptManager) RunScript(name string) error {
    content, err := asm.LoadScript(name)
    if err != nil {
        return err
    }
    
    if err := lua_engine.ExecuteString(content); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

// ClearCache 清除脚本缓存
func (asm *AdvancedScriptManager) ClearCache() {
    asm.scriptCache = make(map[string]string)
}

// ExtractEmbeddedScripts 提取所有嵌入的脚本到文件系统
func (asm *AdvancedScriptManager) ExtractEmbeddedScripts() error {
    if !files.Exists(asm.scriptDir) {
        if err := files.Create(asm.scriptDir); err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        }
    }
    
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return fmt.Errorf("读取嵌入目录失败: %v", err)
    }
    
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        
        srcPath := filepath.Join("scripts", entry.Name())
        dstPath := filepath.Join(asm.scriptDir, entry.Name())
        
        content, err := scriptFS.ReadFile(srcPath)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", srcPath, err)
        }
        
        if err := files.Write(dstPath, string(content)); err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        }
        
        fmt.Printf("已提取脚本: %s\n", dstPath)
    }
    
    return nil
}

func main() {
    // 初始化 Lua 引擎
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    // 创建高级脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    
    // 开发模式：优先从文件系统加载（支持热更新）
    scriptManager := NewAdvancedScriptManager(scriptDir, false)
    
    // 生产模式：只从嵌入的文件系统加载
    // scriptManager := NewAdvancedScriptManager(scriptDir, true)
    
    // 提取嵌入的脚本到文件系统（可选）
    if err := scriptManager.ExtractEmbeddedScripts(); err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
    }
    
    // 运行脚本
    fmt.Println("运行脚本...")
    if err := scriptManager.RunScript("demo.lua"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
    
    if err := scriptManager.RunScript("auto_login.lua"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
}
```

这个 Demo 展示了如何：
1. 使用 Go embed 将脚本文件嵌入到程序中
2. 使用 go-bindata 兼容旧版本 Go
3. 使用 vfsgen 生成虚拟文件系统
4. 运行时释放脚本到文件系统
5. 直接从嵌入的文件系统执行脚本（无需释放）
6. 支持动态脚本加载和热更新
7. 完整的构建和部署流程

这样可以将所有脚本打包到可执行文件中，无需额外分发脚本文件，同时支持在运行时释放和执行。

### 基础示例

```lua
-- AutoGo Lua 脚本示例

-- 1. 获取当前应用信息
local packageName = app_currentPackage()
local activity = app_currentActivity()
print("当前应用: " .. packageName)
print("当前活动: " .. activity)

-- 2. 检查屏幕状态
if not device_isScreenOn() then
    device_wakeUp()
end

-- 3. 等待屏幕解锁
while not device_isScreenUnlock() do
    sleep(1000)
end

-- 4. 查找并点击按钮
local x, y = images_findColor(0, 0, device.width(0), device.height(0), "#FF0000", 0.9, 0, 0)
if x ~= -1 and y ~= -1 then
    click(x, y, 1, 0)
else
    print("未找到目标颜色")
end

-- 5. 识别文字
local results = ppocr_ocr(0, 0, device.width(0), device.height(0), "", 0)
for i, result in ipairs(results) do
    if string.find(result["标签"], "确定") then
        click(result["CenterX"], result["CenterY"], 1, 0)
        break
    end
end

-- 6. 保存数据
storages_put("myData", "lastRun", os.time())

-- 7. 发送网络请求
local code, data = http_get("https://api.example.com/status", 5000)
if code == 200 then
    print("请求成功")
end
```

## 注意事项

1. **线程安全**: 所有操作都是线程安全的，可以在多个 goroutine 中使用
2. **资源管理**: 使用完毕后记得调用 `lua_engine.Close()` 释放资源
3. **错误处理**: 执行 Lua 代码时注意检查错误返回值
4. **方法重写**: 重写方法时要小心，避免影响其他功能
5. **性能考虑**: 频繁的图像操作可能会影响性能，建议适当添加延迟

## 许可证

本项目遵循 AutoGo 项目的许可证。
