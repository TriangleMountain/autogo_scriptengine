//go:build ignore

package lua_engine

import (
	"fmt"
	"os"
)

func ExampleUsage() {
	engine := GetEngine()

	fmt.Println("=== AutoGo Lua Engine 示例 ===\n")

	fmt.Println("1. 执行 Lua 代码:")
	script := `
-- 获取当前应用包名
local packageName = app_currentPackage()
print("当前应用包名: " .. packageName)

-- 获取设备信息
print("设备分辨率: " .. device.width .. "x" .. device.height)
print("设备SDK版本: " .. device.sdkInt)

-- 点击屏幕
click(500, 1000, 1)
`
	err := engine.ExecuteString(script)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n2. 方法管理:")
	registry := engine.GetRegistry()
	methods := registry.ListMethods()
	fmt.Printf("已注册方法数量: %d\n", len(methods))

	fmt.Println("\n3. 方法重写示例:")
	overrideScript := `
-- 重写 click 方法，添加日志
local originalClick = click

function click(x, y, fingerID)
    print("点击: (" .. x .. ", " .. y .. ")")
    originalClick(x, y, fingerID)
end

-- 测试重写的方法
click(100, 200, 1)
`
	err = engine.ExecuteString(overrideScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n4. 图像识别示例:")
	imageScript := `
-- 查找颜色
local x, y = images_findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0)
if x ~= -1 and y ~= -1 then
    print("找到颜色在: " .. x .. ", " .. y)
    click(x, y, 1)
else
    print("未找到颜色")
end

-- 识别文字
local results = ppocr_ocr(0, 0, 1080, 1920, "")
print("识别到 " .. #results .. " 个文本")
for i, result in ipairs(results) do
    print(result["标签"] .. " (" .. result["精度"] .. ")")
end
`
	err = engine.ExecuteString(imageScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n5. 生成文档:")
	docGen := NewDocumentationGenerator()

	luaDoc := docGen.GenerateLuaDocumentation()
	err = os.WriteFile("lua_api.lua", []byte(luaDoc), 0644)
	if err != nil {
		fmt.Printf("生成 Lua 文档错误: %v\n", err)
	} else {
		fmt.Println("Lua API 文档已生成: lua_api.lua")
	}

	markdownDoc := docGen.GenerateMarkdownDocumentation()
	err = os.WriteFile("lua_api.md", []byte(markdownDoc), 0644)
	if err != nil {
		fmt.Printf("生成 Markdown 文档错误: %v\n", err)
	} else {
		fmt.Println("Markdown API 文档已生成: lua_api.md")
	}

	fmt.Println("\n=== 示例完成 ===")
}

func ExampleAdvancedUsage() {
	engine := GetEngine()

	fmt.Println("=== AutoGo Lua Engine 高级示例 ===\n")

	fmt.Println("1. 注册自定义方法:")
	registerScript := `
-- 注册自定义方法
registerMethod("myCustomMethod", "我的自定义方法", nil, true)

function myCustomMethod(param)
    print("自定义方法被调用: " .. param)
    return "返回值: " .. param
end

-- 使用自定义方法
local result = myCustomMethod("测试参数")
print(result)
`
	err := engine.ExecuteString(registerScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n2. 文件操作示例:")
	fileScript := `
-- 检查文件是否存在
if files_exists("/sdcard/test.txt") then
    print("文件存在")
    local content = files_read("/sdcard/test.txt")
    print("文件内容: " .. content)
else
    print("文件不存在")
    files_write("/sdcard/test.txt", "Hello from Lua!")
    print("文件已创建")
end

-- 列出目录
local files = files_listDir("/sdcard")
print("目录文件数量: " .. #files)
`
	err = engine.ExecuteString(fileScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n3. 存储操作示例:")
	storageScript := `
-- 存储数据
storages_put("myTable", "key1", "value1")
storages_put("myTable", "key2", "value2")

-- 读取数据
local value1 = storages_get("myTable", "key1")
print("key1 = " .. value1)

-- 检查键是否存在
if storages_contains("myTable", "key2") then
    print("key2 存在")
end

-- 获取所有数据
local allData = storages_getAll("myTable")
for k, v in pairs(allData) do
    print(k .. " = " .. v)
end
`
	err = engine.ExecuteString(storageScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n4. 网络请求示例:")
	networkScript := `
-- 发送 GET 请求
local code, data = http_get("https://example.com", 5000)
print("状态码: " .. code)
print("响应长度: " .. string.len(data))
`
	err = engine.ExecuteString(networkScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n5. 列出所有方法:")
	listScript := `
local methods = listMethods()
print("已注册方法:")
for i, method in ipairs(methods) do
    print(i .. ". " .. method["name"] .. " - " .. method["description"])
    print("   可重写: " .. tostring(method["overridable"]))
    print("   已重写: " .. tostring(method["overridden"]))
end
`
	err = engine.ExecuteString(listScript)
	if err != nil {
		fmt.Printf("执行错误: %v\n", err)
	}

	fmt.Println("\n=== 高级示例完成 ===")
}

func RunExample() {
	ExampleUsage()
	fmt.Println("\n")
	ExampleAdvancedUsage()
}
