package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/console"
	"github.com/dop251/goja"
)

func injectConsoleMethods(engine *JSEngine) {
	vm := engine.GetVM()

	consoleObj := vm.NewObject()
	vm.Set("console_window", consoleObj)

	consoleObj.Set("new", func(call goja.FunctionCall) goja.Value {
		c := console.New()
		return vm.ToValue(c)
	})

	consoleObj.Set("setWindowSize", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		width := int(call.Argument(1).ToInteger())
		height := int(call.Argument(2).ToInteger())
		c.SetWindowSize(width, height)
		return vm.ToValue(c)
	})

	consoleObj.Set("setWindowPosition", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		x := int(call.Argument(1).ToInteger())
		y := int(call.Argument(2).ToInteger())
		c.SetWindowPosition(x, y)
		return vm.ToValue(c)
	})

	consoleObj.Set("setWindowColor", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		color := call.Argument(1).String()
		c.SetWindowColor(color)
		return vm.ToValue(c)
	})

	consoleObj.Set("setTextColor", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		color := call.Argument(1).String()
		c.SetTextColor(color)
		return vm.ToValue(c)
	})

	consoleObj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		size := int(call.Argument(1).ToInteger())
		c.SetTextSize(size)
		return vm.ToValue(c)
	})

	consoleObj.Set("println", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		var args []any
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		c.Println(args...)
		return goja.Undefined()
	})

	consoleObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Clear()
		return goja.Undefined()
	})

	consoleObj.Set("show", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Show()
		return goja.Undefined()
	})

	consoleObj.Set("hide", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Hide()
		return goja.Undefined()
	})

	consoleObj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		result := c.IsVisible()
		return vm.ToValue(result)
	})

	consoleObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Destroy()
		return goja.Undefined()
	})

	// 注册方法到文档
	engine.RegisterMethod("console_window.new", "创建一个新的Console对象", console.New, true)
	engine.RegisterMethod("console_window.setWindowSize", "设置控制台窗口的大小", func(c *console.Console, width, height int) *console.Console {
		return c.SetWindowSize(width, height)
	}, true)
	engine.RegisterMethod("console_window.setWindowPosition", "设置控制台窗口的位置", func(c *console.Console, x, y int) *console.Console {
		return c.SetWindowPosition(x, y)
	}, true)
	engine.RegisterMethod("console_window.setWindowColor", "设置控制台窗口的背景颜色", func(c *console.Console, color string) *console.Console {
		return c.SetWindowColor(color)
	}, true)
	engine.RegisterMethod("console_window.setTextColor", "设置控制台文本的颜色", func(c *console.Console, color string) *console.Console {
		return c.SetTextColor(color)
	}, true)
	engine.RegisterMethod("console_window.setTextSize", "设置控制台文本的字体大小", func(c *console.Console, size int) *console.Console {
		return c.SetTextSize(size)
	}, true)
	engine.RegisterMethod("console_window.println", "在控制台中打印一行内容", func(c *console.Console, a ...any) {
		c.Println(a...)
	}, true)
	engine.RegisterMethod("console_window.clear", "清空控制台的所有内容", func(c *console.Console) {
		c.Clear()
	}, true)
	engine.RegisterMethod("console_window.show", "显示控制台窗口", func(c *console.Console) {
		c.Show()
	}, true)
	engine.RegisterMethod("console_window.hide", "隐藏控制台窗口", func(c *console.Console) {
		c.Hide()
	}, true)
	engine.RegisterMethod("console_window.isVisible", "返回控制台窗口是否可见", func(c *console.Console) bool {
		return c.IsVisible()
	}, true)
	engine.RegisterMethod("console_window.destroy", "销毁控制台对象，释放资源", func(c *console.Console) {
		c.Destroy()
	}, true)
}
