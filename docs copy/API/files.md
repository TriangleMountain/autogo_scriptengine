# files - 文件操作
---
提供文件和文件夹的操作接口，例如读取、写入、移动等。

## IsFile
<hr style="margin: 0;">

判断路径是否是文件。

- `path` {string} 路径

```go
isFile := files.IsFile("/sdcard/example.txt")
```

## IsDir
<hr style="margin: 0;">

判断路径是否是文件夹。

- `path` {string} 路径

```go
isDir := files.IsDir("/sdcard/example_folder")
```

## IsEmptyDir
<hr style="margin: 0;">

判断文件夹是否为空。如果路径不是文件夹，返回 false。

- `path` {string} 文件夹路径

```go
isEmpty := files.IsEmptyDir("/sdcard/example_folder")
```

## Create
<hr style="margin: 0;">

创建文件或文件夹。如果文件已存在，返回 true。

- `path` {string} 路径

```go
success := files.Create("/sdcard/new_file.txt")
```

## Exists
<hr style="margin: 0;">

判断路径是否存在。

- `path` {string} 路径

```go
exists := files.Exists("/sdcard/example.txt")
```

## EnsureDir
<hr style="margin: 0;">

确保文件夹存在，如果不存在则创建。

- `path` {string} 路径

```go
success := files.EnsureDir("/sdcard/new_folder")
```

## Read
<hr style="margin: 0;">

读取文本文件的内容。

- `path` {string} 文件路径

```go
content := files.Read("/sdcard/example.txt")
```

## ReadBytes
<hr style="margin: 0;">

读取文件的字节数据。

- `path` {string} 文件路径

```go
data := files.ReadBytes("/sdcard/example.txt")
```

## Write
<hr style="margin: 0;">

将文本写入文件。如果文件不存在则创建，存在则覆盖。

- `path` {string} 文件路径
- `text` {string} 要写入的文本

```go
files.Write("/sdcard/example.txt", "Hello, World!")
```

## WriteBytes
<hr style="margin: 0;">

将字节数据写入文件。如果文件不存在则创建，存在则覆盖。

- `path` {string} 文件路径
- `bytes` {[]byte} 要写入的字节数据

```go
files.WriteBytes("/sdcard/example.txt", []byte("Hello, World!"))
```

## Append
<hr style="margin: 0;">

将文本追加到文件末尾。如果文件不存在则创建。

- `path` {string} 文件路径
- `text` {string} 要追加的文本

```go
files.Append("/sdcard/example.txt", "Appended text")
```

## AppendBytes
<hr style="margin: 0;">

将字节数据追加到文件末尾。如果文件不存在则创建。

- `path` {string} 文件路径
- `bytes` {[]byte} 要追加的字节数据

```go
files.AppendBytes("/sdcard/example.txt", []byte("Appended bytes"))
```

## Copy
<hr style="margin: 0;">

复制文件。

- `fromPath` {string} 源文件路径
- `toPath` {string} 目标文件路径

```go
success := files.Copy("/sdcard/source.txt", "/sdcard/destination.txt")
```

## Move
<hr style="margin: 0;">

移动文件。

- `fromPath` {string} 源文件路径
- `toPath` {string} 目标文件路径

```go
success := files.Move("/sdcard/source.txt", "/sdcard/destination.txt")
```

## Rename
<hr style="margin: 0;">

重命名文件。

- `path` {string} 文件路径
- `newName` {string} 新文件名

```go
success := files.Rename("/sdcard/example.txt", "new_name.txt")
```

## GetName
<hr style="margin: 0;">

获取文件名。

- `path` {string} 文件路径

```go
name := files.GetName("/sdcard/example.txt")
```

## GetNameWithoutExtension
<hr style="margin: 0;">

获取不含扩展名的文件名。

- `path` {string} 文件路径

```go
name := files.GetNameWithoutExtension("/sdcard/example.txt")
```

## GetExtension
<hr style="margin: 0;">

获取文件的扩展名。

- `path` {string} 文件路径

```go
extension := files.GetExtension("/sdcard/example.txt")
```

## GetMd5
<hr style="margin: 0;">

获取文件的MD5值。

- `path` {string} 文件路径

```go
md5 := files.GetMd5("/sdcard/example.txt")
```

## Remove
<hr style="margin: 0;">

删除文件或文件夹。如果是文件夹，则删除其所有内容。

- `path` {string} 文件路径或文件夹路径

```go
success := files.Remove("/sdcard/example.txt")
```

## Path
<hr style="margin: 0;">

将相对路径转换为绝对路径。

- `relativePath` {string} 相对路径

```go
absolutePath := files.Path("./example.txt")
```

## ListDir
<hr style="margin: 0;">

列出文件夹下的所有文件和文件夹。

- `path` {string} 文件夹路径

```go
entries := files.ListDir("/sdcard/example_folder")
```