package js_engine

import (
	"fmt"
	"os"
)

func main() {
	// 初始化 JavaScript 引擎
	engine := GetEngine()
	defer Close()

	// 生成 JavaScript API 文档
	docGen := NewDocumentationGenerator()

	// 生成 JavaScript 文档
	fmt.Println("生成 JavaScript API 文档...")
	err := docGen.SaveJSDocumentation("js_api.js")
	if err != nil {
		fmt.Printf("生成 JavaScript 文档失败: %v\n", err)
		return
	}
	fmt.Println("✓ JavaScript 文档已生成: js_api.js")

	// 生成 Markdown 文档
	fmt.Println("生成 Markdown API 文档...")
	err = docGen.SaveMarkdownDocumentation("js_api.md")
	if err != nil {
		fmt.Printf("生成 Markdown 文档失败: %v\n", err)
		return
	}
	fmt.Println("✓ Markdown 文档已生成: js_api.md")

	// 导出方法列表为 JSON
	fmt.Println("导出方法列表...")
	registry := engine.GetRegistry()
	jsonData, err := registry.ExportMethodsJSON()
	if err != nil {
		fmt.Printf("导出方法列表失败: %v\n", err)
		return
	}

	err = os.WriteFile("methods.json", []byte(jsonData), 0644)
	if err != nil {
		fmt.Printf("写入方法列表失败: %v\n", err)
		return
	}
	fmt.Println("✓ 方法列表已导出: methods.json")

	// 导出方法列表为 JavaScript 对象
	fmt.Println("导出方法列表为 JavaScript 对象...")
	jsObject := registry.ExportMethodsJSObject()
	err = os.WriteFile("methods.js", []byte(jsObject), 0644)
	if err != nil {
		fmt.Printf("写入 JavaScript 对象失败: %v\n", err)
		return
	}
	fmt.Println("✓ JavaScript 对象已导出: methods.js")

	fmt.Println("\n所有文档生成完成!")
}
