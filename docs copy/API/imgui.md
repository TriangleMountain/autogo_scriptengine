# imgui - 即时模式图形用户界面
---
提供基于 Dear ImGui 的图形用户界面功能。由于 ImGui 方法数量众多，完整的方法列表请参照 [Dear ImGui 官方文档](https://github.com/ocornut/imgui)

## 基础示例

```go
package main

import (
    "fmt"
    "github.com/Dasongzi1366/AutoGo/imgui"
)

func main() {
    // 初始化
    imgui.Init()
    
    // 状态变量
    counter := 0
    showWindow := true
    
    // 主循环
    imgui.Run(func() {
        // 设置窗口
        imgui.SetNextWindowSizeV(imgui.Vec2{X: 500, Y: 400}, imgui.CondOnce)
        imgui.SetNextWindowPosV(imgui.Vec2{X: 100, Y: 100}, imgui.CondOnce, imgui.Vec2{X: 0, Y: 0})
        
        // 创建窗口
        imgui.BeginV("示例窗口", &showWindow, 0)
        
        // 标题
        imgui.Text("ImGui 示例程序")
        imgui.Separator()
        imgui.Spacing()
        
        // 计数器
        imgui.Text(fmt.Sprintf("计数器: %d", counter))
        
        // 按钮
        if imgui.Button("增加") {
            counter++
        }
        imgui.SameLine()
        if imgui.Button("减少") {
            counter--
        }
        imgui.SameLine()
        if imgui.Button("重置") {
            counter = 0
        }
        
        imgui.Spacing()
        imgui.Separator()
        imgui.Spacing()
        
        // 样式化按钮
        imgui.Text("样式化按钮：")
        imgui.PushStyleColorVec4(imgui.ColButton, imgui.Vec4{X: 0.2, Y: 0.8, Z: 0.2, W: 1.0})
        if imgui.Button("绿色按钮") {
            // 绿色按钮的操作
        }
        imgui.PopStyleColor()
        
        imgui.SameLine()
        imgui.PushStyleColorVec4(imgui.ColButton, imgui.Vec4{X: 0.8, Y: 0.2, Z: 0.2, W: 1.0})
        if imgui.Button("红色按钮") {
            // 红色按钮的操作
        }
        imgui.PopStyleColor()
        
        // 结束窗口
        imgui.End()
    })

    // 阻塞主进程防止程序退出
	select {}
}
```