
# storages - 本地存储
---
提供了保存简单数据、用户配置等的支持。保存的数据除非应用被卸载或者被主动删除，否则会一直保留。

## Get
<hr style="margin: 0;">

从本地存储中取出键值为 `key` 的数据。

- `table` {string} 要操作的表。
- `key` {string} 要查询的键。

```go
value := storages.Get("data", "username")
fmt.Println("获取到的值:", value)
```

## Put
<hr style="margin: 0;">

把值 `value` 保存到本地存储中。

- `table` {string} 要操作的表。
- `key` {string} 要保存的键。
- `value` {string} 要保存的值。

```go
storages.Put("data", "username", "JohnDoe")
```

## Remove
<hr style="margin: 0;">

移除键值为 `key` 的数据。

- `table` {string} 要操作的表。
- `key` {string} 要移除的键。

```go
storages.Remove("data", "username")
```

## Contains
<hr style="margin: 0;">

返回该本地存储是否包含键值为 `key` 的数据。

- `table` {string} 要操作的表。
- `key` {string} 要检查的键。

```go
exists := storages.Contains("data", "username")
if exists {
    fmt.Println("键存在!")
} else {
    fmt.Println("键不存在!")
}
```

## GetAll
<hr style="margin: 0;">

获取该本地存储所有键值对。

- `table` {string} 要操作的表。

```go
data := storages.GetAll("data")
for k, v := range data {
	fmt.Printf("%s = %s\n", k, v)
}
```

## Clear
<hr style="margin: 0;">

移除该本地存储的所有数据。

- `table` {string} 要操作的表。

```go
storages.Clear("data")
fmt.Println("所有存储数据已被清除!")
```
