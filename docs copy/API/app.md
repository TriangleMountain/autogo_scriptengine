# app - 应用
---
app模块提供一系列函数，用于使用其他应用、与其他应用交互。例如启动应用、打开文件、发送意图等。

同时提供了方便的进阶函数startActivity和sendBroadcast，用他们可完成app模块没有内置的和其他应用的交互。

以下是 `app` 包中定义的 `IntentOptions` 结构体及其字段说明：

| **字段名**      | **类型**            | **说明**                           |
|------------------|---------------------|------------------------------------|
| `Action`         | `string`           | Intent 的动作，例如 `android.intent.action.VIEW`。 |
| `Type`           | `string`           | Intent 的数据类型，例如 `text/plain` 或 `image/*`。 |
| `Data`           | `string`           | Intent 的数据，例如文件路径或 URI。 |
| `Category`       | `[]string`         | Intent 的类别。                   |
| `PackageName`    | `string`           | 应用包名，用于指定目标应用。       |
| `Extras`         | `map[string]string`| Intent 的额外参数（键值对）。      |
| `Flags`          | `[]string`         | Intent 的标志，例如 `FLAG_ACTIVITY_NEW_TASK`。 |


## CurrentPackage
<hr style="margin: 0;">

获取当前页面的应用包名。

```go
packageName := app.CurrentPackage()
```

## CurrentActivity
<hr style="margin: 0;">

获取当前页面的应用类名。

```go
activityName := app.CurrentActivity()
```

## Launch
<hr style="margin: 0;">

通过应用包名启动应用。

- `packageName` {string} 应用包名，也支持"包名/类名"格式
- `displayId` {int} 屏幕ID

```go
success := app.Launch("com.tencent.mm", 0)
```

## GetList
<hr style="margin: 0;">

获取手机中所有应用列表。

- `includeSystemApps` {bool} 是否需要包含系统应用

```go
list := app.GetList(true)
```

## GetName
<hr style="margin: 0;">

获取指定包名应用的应用名称。

- `packageName` {string} 应用包名

```go
name := app.GetName("bin.mt.plus")
```

## GetIcon
<hr style="margin: 0;">

获取应用图标。

- `packageName` {string} 应用包名

```go
data := app.GetIcon("com.tencent.mm")
```

## GetVersion
<hr style="margin: 0;">

获取应用版本号。

- `packageName` {string} 应用包名

```go
version := app.GetVersion("com.tencent.mm")
```

## OpenSetting
<hr style="margin: 0;">

打开应用的详情页（设置页）。

- `packageName` {string} 应用包名

```go
success := app.OpenSetting("com.tencent.mm")
```

## ViewFile
<hr style="margin: 0;">

用其他应用查看文件。文件不存在的情况由查看文件的应用处理。

- `path` {string} 文件路径

```go
app.ViewFile("/sdcard/example.txt")
```

## EditFile
<hr style="margin: 0;">

用其他应用编辑文件。文件不存在的情况由编辑文件的应用处理。

- `path` {string} 文件路径

```go
app.EditFile("/sdcard/example.txt")
```

## Uninstall
<hr style="margin: 0;">

卸载应用。

- `packageName` {string} 应用包名

```go
app.Uninstall("com.tencent.mm")
```

## Install
<hr style="margin: 0;">

安装应用（也支持xapk）。

- `path` {string} APK 文件路径

```go
app.Install("/sdcard/app.apk")
```

## IsInstalled
<hr style="margin: 0;">

判断是否已经安装某个应用。

- `packageName` {string} 应用包名

```go
installed := app.IsInstalled("com.tencent.mm")
```

## Clear
<hr style="margin: 0;">

清除应用数据。

- `packageName` {string} 应用包名

```go
app.Clear("com.tencent.mm")
```

## ForceStop
<hr style="margin: 0;">

强制停止应用。

- `packageName` {string} 应用包名

```go
app.ForceStop("com.tencent.mm")
```

## Disable
<hr style="margin: 0;">

禁用应用。

- `packageName` {string} 应用包名

```go
app.Disable("com.tencent.mm")
```

## Enable
<hr style="margin: 0;">

启用应用。

- `packageName` {string} 应用包名

```go
app.Enable("com.tencent.mm")
```

## EnableAccessibility
<hr style="margin: 0;">

启用无障碍服务。

- `packageName` {string} 应用包名

```go
app.EnableAccessibility("org.autojs.autoxjs.v6")
```

## DisableAccessibility
<hr style="margin: 0;">

关闭无障碍服务。

- `packageName` {string} 应用包名

```go
app.DisableAccessibility("org.autojs.autoxjs.v6")
```

## IgnoreBattOpt
<hr style="margin: 0;">

忽略应用电池优化。

- `packageName` {string} 应用包名

```go
app.IgnoreBattOpt("com.tencent.mm")
```

## GetBrowserPackage
<hr style="margin: 0;">

获取系统默认浏览器包名。

```go
packageName := app.GetBrowserPackage()
```

## OpenUrl
<hr style="margin: 0;">

用浏览器打开指定的网址。

- `url` {string} 网站地址

```go
app.OpenUrl("https://example.com")
```

## StartActivity
<hr style="margin: 0;">

根据选项构造一个 Intent，并启动该 Activity。

- `options` {IntentOptions} Intent 选项

```go
app.StartActivity(app.IntentOptions{
	Action: "SEND",
	Type:   "text/plain",
	Data:   "file:///sdcard/1.txt",
})
```

## SendBroadcast
<hr style="margin: 0;">

根据选项构造一个 Intent，并发送广播。

- `options` {IntentOptions} Intent 选项

```go
app.SendBroadcast(options)
```

## StartService
<hr style="margin: 0;">

根据选项构造一个 Intent，并启动服务。

- `options` {IntentOptions} Intent 选项

```go
app.StartService(options)
``` 