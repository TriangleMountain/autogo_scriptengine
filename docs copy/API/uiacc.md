# uiacc - 节点操作
---

提供基于辅助功能服务的控件定位、交互操作等功能。无需开启APP的无障碍服务。

以下是 `uiacc` 包中定义的 `Rect` 结构体及其字段说明：

| **字段名**  | **类型**    | **说明**                           |
|-------------|-------------|------------------------------------|
| `Left`      | `int`       | 矩形的左边界。                    |
| `Right`     | `int`       | 矩形的右边界。                    |
| `Top`       | `int`       | 矩形的上边界。                    |
| `Bottom`    | `int`       | 矩形的下边界。                    |
| `CenterX`   | `int`       | 矩形的中心 X 坐标。               |
| `CenterY`   | `int`       | 矩形的中心 Y 坐标。               |
| `Width`     | `int`       | 矩形的宽度。                      |
| `Height`    | `int`       | 矩形的高度。                      |

## New
<hr style="margin: 0;">

创建一个 Accessibility 对象。返回实例对象`*Uiacc`

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕，操作虚拟屏幕节点需要安卓版本大于等于11

```go
acc := uiacc.New(0)
```

## *Uiacc.Text
<hr style="margin: 0;">

设置选择器的 `text` 属性。

- `value` {string} 文本值。

```go
acc.Text("example text")
```

## *Uiacc.TextContains
<hr style="margin: 0;">

设置选择器的 `textContains` 属性。

- `value` {string} 包含的文本值。

```go
acc.TextContains("example")
```

## *Uiacc.TextStartsWith
<hr style="margin: 0;">

设置选择器的 `textStartsWith` 属性。

- `value` {string} 以此文本开头。

```go
acc.TextStartsWith("example")
```

## *Uiacc.TextEndsWith
<hr style="margin: 0;">

设置选择器的 `textEndsWith` 属性。

- `value` {string} 以此文本结尾。

```go
acc.TextEndsWith("example")
```

## *Uiacc.TextMatches
<hr style="margin: 0;">

设置选择器的 `textMatches` 属性，用于匹配符合指定正则表达式的控件。

- `value` {string} 正则表达式。

```go
acc.TextMatches("^example.*")
```

## *Uiacc.Desc
<hr style="margin: 0;">

设置选择器的 `desc` 属性，用于匹配描述等于指定文本的控件。

- `value` {string} 描述的文本值。

```go
acc.Desc("example description")
```

## *Uiacc.DescContains
<hr style="margin: 0;">

设置选择器的 `descContains` 属性，用于匹配描述包含指定文本的控件。

- `value` {string} 包含的描述文本值。

```go
acc.DescContains("example")
```

## *Uiacc.DescStartsWith
<hr style="margin: 0;">

设置选择器的 `descStartsWith` 属性，用于匹配描述以指定文本开头的控件。

- `value` {string} 描述文本的开头。

```go
acc.DescStartsWith("example")
```

## *Uiacc.DescEndsWith
<hr style="margin: 0;">

设置选择器的 `descEndsWith` 属性，用于匹配描述以指定文本结尾的控件。

- `value` {string} 描述文本的结尾。

```go
acc.DescEndsWith("example")
```

## *Uiacc.DescMatches
<hr style="margin: 0;">

设置选择器的 `descMatches` 属性，用于匹配描述符合指定正则表达式的控件。

- `value` {string} 正则表达式。

```go
acc.DescMatches("^example.*")
```

## *Uiacc.Id
<hr style="margin: 0;">

设置选择器的 `id` 属性，用于匹配 ID 等于指定值的控件。

- `value` {string} ID 值。

```go
acc.Id("example_id")
```

## *Uiacc.IdContains
<hr style="margin: 0;">

设置选择器的 `idContains` 属性，用于匹配 ID 包含指定值的控件。

- `value` {string} 包含的 ID 值。

```go
acc.IdContains("example")
```

## *Uiacc.IdStartsWith
<hr style="margin: 0;">

设置选择器的 `idStartsWith` 属性，用于匹配 ID 以指定值开头的控件。

- `value` {string} ID 的开头值。

```go
acc.IdStartsWith("example")
```

## *Uiacc.IdEndsWith
<hr style="margin: 0;">

设置选择器的 `idEndsWith` 属性，用于匹配 ID 以指定值结尾的控件。

- `value` {string} ID 的结尾值。

```go
acc.IdEndsWith("example")
```

## *Uiacc.IdMatches
<hr style="margin: 0;">

设置选择器的 `idMatches` 属性，用于匹配 ID 符合指定正则表达式的控件。

- `value` {string} 正则表达式。

```go
acc.IdMatches("^example.*")
```

## *Uiacc.ClassName
<hr style="margin: 0;">

设置选择器的 `className` 属性，用于匹配类名等于指定值的控件。

- `value` {string} 类名的值。

```go
acc.ClassName("example_class")
```

## *Uiacc.ClassNameContains
<hr style="margin: 0;">

设置选择器的 `classNameContains` 属性，用于匹配类名包含指定值的控件。

- `value` {string} 包含的类名值。

```go
acc.ClassNameContains("example")
```

## *Uiacc.ClassNameStartsWith
<hr style="margin: 0;">

设置选择器的 `classNameStartsWith` 属性，用于匹配类名以指定值开头的控件。

- `value` {string} 类名的开头值。

```go
acc.ClassNameStartsWith("example")
```

## *Uiacc.ClassNameEndsWith
<hr style="margin: 0;">

设置选择器的 `classNameEndsWith` 属性，用于匹配类名以指定值结尾的控件。

- `value` {string} 类名的结尾值。

```go
acc.ClassNameEndsWith("example")
```

## *Uiacc.ClassNameMatches
<hr style="margin: 0;">

设置选择器的 `classNameMatches` 属性，用于匹配类名符合指定正则表达式的控件。

- `value` {string} 正则表达式。

```go
acc.ClassNameMatches("^example.*")
```

## *Uiacc.PackageName
<hr style="margin: 0;">

设置选择器的 `packageName` 属性，用于匹配包名等于指定值的控件。

- `value` {string} 包名的值。

```go
acc.PackageName("com.example")
```

## *Uiacc.PackageNameContains
<hr style="margin: 0;">

设置选择器的 `packageNameContains` 属性，用于匹配包名包含指定值的控件。

- `value` {string} 包含的包名值。

```go
acc.PackageNameContains("example")
```

## *Uiacc.PackageNameStartsWith
<hr style="margin: 0;">

设置选择器的 `packageNameStartsWith` 属性，用于匹配包名以指定值开头的控件。

- `value` {string} 包名的开头值。

```go
acc.PackageNameStartsWith("com.example")
```

## *Uiacc.PackageNameEndsWith
<hr style="margin: 0;">

设置选择器的 `packageNameEndsWith` 属性，用于匹配包名以指定值结尾的控件。

- `value` {string} 包名的结尾值。

```go
acc.PackageNameEndsWith("example")
```

## *Uiacc.PackageNameMatches
<hr style="margin: 0;">

设置选择器的 `packageNameMatches` 属性，用于匹配包名符合指定正则表达式的控件。

- `value` {string} 正则表达式。

```go
acc.PackageNameMatches("^com\.example.*")
```

## *Uiacc.Bounds
<hr style="margin: 0;">

设置选择器的 `bounds` 属性，用于匹配控件在屏幕上的范围。

- `left, top, right, bottom` {int} 控件的屏幕边界。

```go
acc.Bounds(0, 0, 100, 100)
```

## *Uiacc.BoundsInside
<hr style="margin: 0;">

设置选择器的 `boundsInside` 属性，用于匹配控件在屏幕内的范围。

- `left, top, right, bottom` {int} 屏幕内的范围。

```go
acc.BoundsInside(0, 0, 500, 500)
```

## *Uiacc.BoundsContains
<hr style="margin: 0;">

设置选择器的 `boundsContains` 属性，用于匹配控件包含在指定范围内。

- `left, top, right, bottom` {int} 包含的范围。

```go
acc.BoundsContains(50, 50, 300, 300)
```

## *Uiacc.DrawingOrder
<hr style="margin: 0;">

设置选择器的 `drawingOrder` 属性，用于匹配控件在父控件中的绘制顺序。

- `value` {int} 绘制顺序。

```go
acc.DrawingOrder(2)
```

## *Uiacc.Clickable
<hr style="margin: 0;">

设置选择器的 `clickable` 属性，用于匹配控件是否可点击。

- `value` {bool} 是否可点击。

```go
acc.Clickable(true)
```

## *Uiacc.LongClickable
<hr style="margin: 0;">

设置选择器的 `longClickable` 属性，用于匹配控件是否可长按。

- `value` {bool} 是否可长按。

```go
acc.LongClickable(true)
```

## *Uiacc.Checkable
<hr style="margin: 0;">

设置选择器的 `checkable` 属性，用于匹配控件是否可选中。

- `value` {bool} 是否可选中。

```go
acc.Checkable(false)
```

## *Uiacc.Selected
<hr style="margin: 0;">

设置选择器的 `selected` 属性，用于匹配控件是否被选中。

- `value` {bool} 是否被选中。

```go
acc.Selected(true)
```

## *Uiacc.Enabled
<hr style="margin: 0;">

设置选择器的 `enabled` 属性，用于匹配控件是否启用。

- `value` {bool} 是否启用。

```go
acc.Enabled(true)
```

## *Uiacc.Scrollable
<hr style="margin: 0;">

设置选择器的 `scrollable` 属性，用于匹配控件是否可滚动。

- `value` {bool} 是否可滚动。

```go
acc.Scrollable(false)
```

## *Uiacc.Editable
<hr style="margin: 0;">

设置选择器的 `editable` 属性，用于匹配控件是否可编辑。

- `value` {bool} 是否可编辑。

```go
acc.Editable(true)
```

## *Uiacc.MultiLine
<hr style="margin: 0;">

设置选择器的 `multiLine` 属性，用于匹配控件是否多行。

- `value` {bool} 是否多行。

```go
acc.MultiLine(false)
```

## *Uiacc.Checked
<hr style="margin: 0;">

设置选择器的 `checked` 属性，用于匹配控件是否被勾选。

- `value` {bool} 是否勾选。

```go
acc.Checked(true)
```

## *Uiacc.Focusable
<hr style="margin: 0;">

设置选择器的 `focusable` 属性，用于匹配控件是否可聚焦。

- `value` {bool} 是否可聚焦。

```go
acc.Focusable(true)
```

## *Uiacc.Dismissable
<hr style="margin: 0;">

设置选择器的 `dismissable` 属性，用于匹配控件是否可解散。

- `value` {bool} 是否可解散。

```go
acc.Dismissable(false)
```

## *Uiacc.Focused
<hr style="margin: 0;">

设置选择器的 `focused` 属性，用于匹配控件是否是辅助功能焦点。

- `value` {bool} 是否为辅助功能焦点。

```go
acc.Focused(true)
```

## *Uiacc.ContextClickable
<hr style="margin: 0;">

设置选择器的 `contextClickable` 属性，用于匹配控件是否是上下文点击。

- `value` {bool} 是否为上下文点击。

```go
acc.ContextClickable(false)
```

## *Uiacc.Index
<hr style="margin: 0;">

设置选择器的 `index` 属性，用于匹配控件在父控件中的索引。

- `value` {int} 索引值。

```go
acc.Index(1)
```

## *Uiacc.Visible
<hr style="margin: 0;">

设置选择器的 `visible` 属性，用于匹配控件是否可见。

- `value` {bool} 是否可见。

```go
acc.Visible(true)
```

## *Uiacc.Password
<hr style="margin: 0;">

设置选择器的 `password` 属性，用于匹配控件是否为密码字段。

- `value` {bool} 是否为密码字段。

```go
acc.Password(false)
```

## *Uiacc.Click
<hr style="margin: 0;">

点击屏幕上的文本。

- `text` {string} 目标文本。

```go
acc.Click("目标文本")
```

## *Uiacc.WaitFor
<hr style="margin: 0;">

等待控件出现，返回控件对象 `*UiObject` 。

- `timeout` {int} 超时时间（毫秒）。`0` 表示无限等待。

```go
obj := acc.Text("hello").WaitFor(3000)
```

## *Uiacc.FindOnce
<hr style="margin: 0;">

查找单个控件，成功返回控件对象 `*UiObject` 。

```go
obj := acc.Text("hello").FindOnce()
```

## *Uiacc.Find
<hr style="margin: 0;">

查找所有符合条件的控件。返回控件对象数组 `[]*UiObject` 。

```go
objects := acc.Text("hello").Find()
```

## *Uiacc.Release
<hr style="margin: 0;">

释放无障碍服务资源 。

```go
uiacc.Release()
```

## *UiObject.Click
<hr style="margin: 0;">

点击该控件，并返回是否点击成功。

```go
success := uiObject.Click()
fmt.Println("点击成功:", success)
```

## *UiObject.ClickCenter
<hr style="margin: 0;">

使用控件坐标点击该控件的中点。

```go
success := uiObject.ClickCenter()
fmt.Println("点击中心成功:", success)
```

## *UiObject.ClickLongClick
<hr style="margin: 0;">

长按该控件，并返回是否点击成功。

```go
success := uiObject.ClickLongClick()
fmt.Println("长按成功:", success)
```

## *UiObject.Copy
<hr style="margin: 0;">

对输入框文本的选中内容进行复制，并返回是否操作成功。

```go
success := uiObject.Copy()
fmt.Println("复制成功:", success)
```

## *UiObject.Cut
<hr style="margin: 0;">

对输入框文本的选中内容进行剪切，并返回是否操作成功。

```go
success := uiObject.Cut()
fmt.Println("剪切成功:", success)
```

## *UiObject.Paste
<hr style="margin: 0;">

对输入框控件进行粘贴操作，把剪贴板内容粘贴到输入框中，并返回是否操作成功。

```go
success := uiObject.Paste()
fmt.Println("粘贴成功:", success)
```

## *UiObject.ScrollForward
<hr style="margin: 0;">

对控件执行向前滑动的操作，并返回是否操作成功。

```go
success := uiObject.ScrollForward()
fmt.Println("向前滑动成功:", success)
```

## *UiObject.ScrollBackward
<hr style="margin: 0;">

对控件执行向后滑动的操作，并返回是否操作成功。

```go
success := uiObject.ScrollBackward()
fmt.Println("向后滑动成功:", success)
```

## *UiObject.Collapse
<hr style="margin: 0;">

对控件执行折叠操作，并返回是否操作成功。

```go
success := uiObject.Collapse()
fmt.Println("折叠成功:", success)
```

## *UiObject.Expand
<hr style="margin: 0;">

对控件执行展开操作，并返回是否操作成功。

```go
success := uiObject.Expand()
fmt.Println("展开成功:", success)
```

## *UiObject.Show
<hr style="margin: 0;">

执行显示操作，并返回是否操作成功。

```go
success := uiObject.Show()
fmt.Println("显示成功:", success)
```

## *UiObject.Select
<hr style="margin: 0;">

对控件执行"选中"操作，并返回是否操作成功。

```go
selected := uiObject.Select()
fmt.Println("控件是否选中成功:", selected)
```

## *UiObject.ClearSelect
<hr style="margin: 0;">

清除控件的选中状态，并返回是否操作成功。

```go
cleared := uiObject.ClearSelect()
fmt.Println("控件是否清除选中成功:", cleared)
```

## *UiObject.SetSelection
<hr style="margin: 0;">

对输入框控件设置选中的文字内容，并返回是否操作成功。

- `start` {int} 选中内容的起始位置。
- `end` {int} 选中内容的结束位置。

```go
success := uiObject.SetSelection(0, 5)
fmt.Println("设置选中内容是否成功:", success)
```

## *UiObject.SetVisibleToUser
<hr style="margin: 0;">

设置控件是否可见。

- `isVisible` {bool} 是否可见。

```go
success := uiObject.SetVisibleToUser(false)
fmt.Println("设置控件不可见是否成功:", success)
```

## *UiObject.SetText
<hr style="margin: 0;">

设置输入框控件的文本内容，并返回是否设置成功。

- `str` {string} 文本内容。

```go
success := uiObject.SetText("example text")
fmt.Println("设置文本是否成功:", success)
```

## *UiObject.GetClickable
<hr style="margin: 0;">

获取控件的 `clickable` 属性。

```go
clickable := uiObject.GetClickable()
fmt.Println("控件是否可点击:", clickable)
```

## *UiObject.GetLongClickable
<hr style="margin: 0;">

获取控件的 `longClickable` 属性。

```go
longClickable := uiObject.GetLongClickable()
fmt.Println("控件是否支持长按:", longClickable)
```

## *UiObject.GetCheckable
<hr style="margin: 0;">

获取控件的 `checkable` 属性。

```go
checkable := uiObject.GetCheckable()
fmt.Println("控件是否可选中:", checkable)
```

## *UiObject.GetSelected
<hr style="margin: 0;">

获取控件的 `selected` 属性。

```go
selected := uiObject.GetSelected()
fmt.Println("控件是否被选中:", selected)
```

## *UiObject.GetEnabled
<hr style="margin: 0;">

获取控件的 `enabled` 属性。

```go
enabled := uiObject.GetEnabled()
fmt.Println("控件是否启用:", enabled)
```

## *UiObject.GetScrollable
<hr style="margin: 0;">

获取控件的 `scrollable` 属性。

```go
scrollable := uiObject.GetScrollable()
fmt.Println("控件是否可滚动:", scrollable)
```

## *UiObject.GetEditable
<hr style="margin: 0;">

获取控件的 `editable` 属性。

```go
editable := uiObject.GetEditable()
fmt.Println("控件是否可编辑:", editable)
```

## *UiObject.GetMultiLine
<hr style="margin: 0;">

获取控件的 `multiLine` 属性。

```go
multiLine := uiObject.GetMultiLine()
fmt.Println("控件是否多行:", multiLine)
```

## *UiObject.GetChecked
<hr style="margin: 0;">

获取控件的 `checked` 属性。

```go
checked := uiObject.GetChecked()
fmt.Println("控件是否被勾选:", checked)
```

## *UiObject.GetFocused
<hr style="margin: 0;">

获取控件的 `focused` 属性。

```go
focused := uiObject.GetFocused()
fmt.Println("控件是否获得了输入焦点:", focusable)
```

## *UiObject.GetFocusable
<hr style="margin: 0;">

获取控件的 `focusable` 属性。

```go
focusable := uiObject.GetFocusable()
fmt.Println("控件是否可聚焦:", focusable)
```

## *UiObject.GetDismissable
<hr style="margin: 0;">

获取控件的 `dismissable` 属性。

```go
dismissable := uiObject.GetDismissable()
fmt.Println("控件是否可解散:", dismissable)
```

## *UiObject.GetContextClickable
<hr style="margin: 0;">

获取控件的 `contextClickable` 属性。

```go
contextClickable := uiObject.GetContextClickable()
fmt.Println("控件是否支持上下文点击:", contextClickable)
```

## *UiObject.GetVisible
<hr style="margin: 0;">

获取控件的 `visible` 属性。

```go
visible := uiObject.GetVisible()
fmt.Println("控件是否可见:", visible)
```

## *UiObject.GetPassword
<hr style="margin: 0;">

获取控件的 `password` 属性。

```go
password := uiObject.GetPassword()
fmt.Println("控件是否为密码字段:", password)
```

## *UiObject.GetAccessibilityFocused
<hr style="margin: 0;">

获取控件的 `AccessibilityFocused` 属性。

```go
focused := uiObject.GetAccessibilityFocused()
fmt.Println("控件是否为辅助功能焦点:", focused)
```

## *UiObject.GetChildCount
<hr style="margin: 0;">

获取控件的子控件数目。

```go
childCount := uiObject.GetChildCount()
fmt.Println("子控件数量:", childCount)
```

## *UiObject.GetDrawingOrder
<hr style="margin: 0;">

获取控件在父控件中的绘制次序。

```go
drawingOrder := uiObject.GetDrawingOrder()
fmt.Println("控件绘制次序:", drawingOrder)
```

## *UiObject.GetIndex
<hr style="margin: 0;">

获取控件在父控件中的索引。

```go
index := uiObject.GetIndex()
fmt.Println("控件在父控件中的索引:", index)
```

## *UiObject.GetBounds
<hr style="margin: 0;">

获取控件在屏幕上的范围。

```go
bounds := uiObject.GetBounds()
fmt.Printf("控件范围: %v\n", bounds)
```

## *UiObject.GetBoundsInParent
<hr style="margin: 0;">

获取控件在父控件中的范围。

```go
bounds := uiObject.GetBoundsInParent()
fmt.Println("控件在父控件中的范围:", bounds)
```

## *UiObject.GetId
<hr style="margin: 0;">

获取控件的 ID。

```go
id := uiObject.GetId()
fmt.Println("控件 ID:", id)
```

## *UiObject.GetText
<hr style="margin: 0;">

获取控件的文本内容。

```go
text := uiObject.GetText()
fmt.Println("控件文本内容:", text)
```

## *UiObject.GetDesc
<hr style="margin: 0;">

获取控件的描述内容。

```go
desc := uiObject.GetDesc()
fmt.Println("控件描述内容:", desc)
```

## *UiObject.GetPackageName
<hr style="margin: 0;">

获取控件的包名。

```go
packageName := uiObject.GetPackageName()
fmt.Println("控件包名:", packageName)
```

## *UiObject.GetClassName
<hr style="margin: 0;">

获取控件的类名。

```go
className := uiObject.GetClassName()
fmt.Println("控件类名:", className)
```

## *UiObject.GetParent
<hr style="margin: 0;">

获取控件的父控件。

```go
parent := uiObject.GetParent()
fmt.Println("控件的父控件:", parent)
```

## *UiObject.GetChild
<hr style="margin: 0;">

获取控件的指定索引的子控件。

- `index` {int} 子控件的索引。

```go
child := uiObject.GetChild(0)
fmt.Println("第一个子控件:", child)
```

## *UiObject.GetChildren
<hr style="margin: 0;">

获取控件的所有子控件。返回控件对象数组 `[]*UiObject` 。

```go
children := uiObject.GetChildren()
for index, child := range children {
    fmt.Printf("子控件 %d: %v\n", index+1, child)
}
```

## *UiObject.ToString
<hr style="margin: 0;">

将节点对象转文本。

```go
str := uiObject.ToString()
fmt.Println("节点文本:", str)
```