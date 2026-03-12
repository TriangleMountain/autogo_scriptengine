package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/https"
	"github.com/dop251/goja"
)

func injectHttpsMethods(engine *JSEngine) {
	vm := engine.GetVM()

	httpObj := vm.NewObject()
	vm.Set("http", httpObj)

	httpObj.Set("get", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		timeout := int(call.Argument(1).ToInteger())
		code, data := https.Get(url, timeout)
		arr := vm.NewArray()
		arr.Set("0", code)
		if data != nil {
			arr.Set("1", string(data))
		} else {
			arr.Set("1", goja.Null())
		}
		return arr
	})

	httpObj.Set("postMultipart", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		fileName := call.Argument(1).String()
		fileData := call.Argument(2).Export().([]byte)
		timeout := int(call.Argument(3).ToInteger())
		code, responseData := https.PostMultipart(url, fileName, fileData, timeout)
		arr := vm.NewArray()
		arr.Set("0", code)
		if responseData != nil {
			arr.Set("1", string(responseData))
		} else {
			arr.Set("1", goja.Null())
		}
		return arr
	})

	engine.RegisterMethod("http.get", "发送GET请求", func(url string, timeout int) (int, []byte) { return https.Get(url, timeout) }, true)
	engine.RegisterMethod("http.postMultipart", "发送Multipart POST请求", func(url, fileName string, fileData []byte, timeout int) (int, []byte) {
		return https.PostMultipart(url, fileName, fileData, timeout)
	}, true)
}
