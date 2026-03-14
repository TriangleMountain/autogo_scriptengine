# 更新日志

本文档记录了 AutoGo ScriptEngine 的所有版本更新记录。

---

## v0.0.9 (2026-03-15)

### 完善 Windows 编译问题说明和模块依赖管理

- 更新主 README.md，详细说明 Windows 下超过 1 个 C 依赖库会导致编译命令过长
- 更新 js_engine/README.md，添加白名单使用示例和模块依赖列表
- 更新 lua_engine/README.md，添加白名单使用示例和模块依赖列表
- 说明这是受限于 AutoGo Extension 实现的问题，无法根本解决
- 提供使用白名单手动指定依赖的 Demo 代码
- 添加完整的模块列表，标注哪些模块包含 C 依赖

### 添加兼容性说明文档

- 更新主 README.md，添加 Android 版本兼容性和 Windows 开发环境说明
- 更新 js_engine/README.md，添加兼容性说明和问题解决方案
- 更新 lua_engine/README.md，添加兼容性说明和问题解决方案
- 说明某些包在某些 Android 版本下会有内存引用错误问题
- 说明 Windows 下引入 C 包会有命令行过长问题

### 更新 Go 版本

- 更新 Go 版本到 1.25.0

### 添加右侧悬浮目录功能，修复文档标题显示问题

- 更新 convert_to_html.py 脚本，为所有文档添加 headings 字段
- 修复右侧悬浮目录不显示的问题
- 添加 common/ 目录
- 添加 js_engine/model/require/ 目录
- 更新 README.md 文档
- 更新 js_engine 和 lua_engine 相关文件

### 支持 Lua 和 JavaScript 代码导入功能 (require)

- 实现 JavaScript 模块的 require 功能，支持动态加载模块
- 实现 Lua 模块的 require 功能，支持动态加载模块
- 添加模块搜索路径配置
- 支持从嵌入文件系统加载模块
- 添加模块缓存机制，提升性能
- 完善模块依赖管理文档

### 优化 HTML 文档交互体验

- 为左右两边的菜单添加收起/展开功能
- 实现鼠标贴边自动展开、鼠标移出自动收回
- 添加 100px 宽度的触发区域，提供充足的缓冲空间
- 收起时显示图标提示（左侧 ☰，右侧 ≡）
- 添加平滑的动画效果，图标在菜单展开时自动隐藏，收起时自动显示
- 优化触发区域样式，使用渐变背景和阴影效果提升视觉体验

---

## v0.0.5 (2026-03-13)

### 修复错误的包名信息

---

## v0.0.4 (2026-03-13)

### 重构项目结构，将 inject 文件移至 model 子目录

- 将所有 inject 文件从根目录移至 model 子目录
- 按模块类型组织文件结构（app, device, motion, images 等）
- 每个模块目录添加 README.md 说明文档
- 删除旧的 inject 文件
- 重构测试脚本结构
- 更新 types.go 和 documentation.go 以支持新的文件结构
- 添加 define 目录用于模型定义
- 删除 .DS_Store 文件

---

## v0.0.3 (2026-03-13)

### 更新 README 文档

- 添加 displayId 参数说明
- 更新所有 API 示例代码，添加 displayId 参数
- 更新 app.launch 示例，说明 displayId 参数
- 更新 images.findColor 示例，添加 displayId 参数
- 更新 click 示例，添加 displayId 参数
- 更新 uiacc.new 示例，说明 displayId 参数
- 更新 yolo.detect 示例，说明 displayId 参数

---

## v0.0.2 (2026-03-13)

### 修复 motion_inject

---

## v0.0.1 (2026-03-13)

### 版本 v0.0.1 - 修复 YOLO 检测方法参数问题

#### 重构测试文件结构并优化代码

- 将测试文件从根目录移动到 test_scripts/scripts/ 目录
- 优化 js_engine 和 lua_engine 中的注入代码
- 减少 js_engine/device_inject.go 代码行数，提升可读性
- 优化 js_engine/https_inject.go 和 js_engine/yolo_inject.go
- 增强 lua_engine/device_inject.go 功能
- 优化 lua_engine/yolo_inject.go 代码结构
