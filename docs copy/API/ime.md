# ime - 输入法
---
ime模块提供了一系列函数，用于控制输入法行为，实现文本输入等功能。

## InputText
<hr style="margin: 0;">

使用输入法输入文本。

- `text` {string} 需要输入的文本
- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
ime.InputText("Hello, World!", 0)
```

## GetClipText
<hr style="margin: 0;">

获取剪贴板文本内容。

```go
text := ime.GetClipText()
fmt.Println("剪贴板内容:", text)
```

## SetClipText
<hr style="margin: 0;">

设置剪贴板文本内容。

- `text` {string} 要设置的文本

```go
ime.SetClipText("这是要复制到剪贴板的内容")
``` 

## GetIMEList
<hr style="margin: 0;">

获取输入法列表。

```go
fmt.Println("输入法列表:",ime.GetIMEList())
``` 

## SetCurrentIME
<hr style="margin: 0;">

设置系统当前输入法。

- `packageName` {string} 要设置为当前输入法的应用包名

```go
ime.SetCurrentIME("com.android.inputmethod.latin")
``` 
