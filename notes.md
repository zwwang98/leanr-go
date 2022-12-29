### 如何扩展系统类型或者别人的类型
1. 定义别名
   1. 最简单
   2. 但是后续修改为别的方法时（比如使用组合）需要做很大修改
2. 使用组合
   1. 最常用
3. 使用内嵌
   1. 省下代码
   2. 语法糖
   3. 不方便新手理解
   4. 只有的确能省下很多代码时才这么用

### 常用命令行
1. 查看全局环境变量
```command
go env
```
2. 局部修改环境变量，只在项目处修改某个环境变量的值，不影响全局的环境变量
```command
export {VARIABLE_NAME}={VALUE}
```
3. 清理`go.mod`
```command
go mode tidy
```

## 5 依赖管理
三个阶段
1. gopath
2. govendor
3. go mo，最好

### 5-1 gopath
1. 将所有依赖都放在gopath(`~/go`)
2. 项目寻找依赖时，会从`goroot`和`gopath`两个地方找
3. 问题
   1. 所有不同项目的依赖都放在一个目录下，很大
   2. 版本管理困难，不同项目对同一个依赖的版本兼容性不一样，使用gopath很难管理

### 5-2 govendor
1. 在每个项目目录下新建一个`vendor`目录，在这里存放这个项目需要的所以依赖
2. 每次项目寻找依赖时，先到`vendor`下寻找，找不到才到`goroot`和`gopath`两个地方找
3. 有大量针对`govendor`的第三方依赖管理工具：glide、dep、go dep
4. 问题是

### 5-3 gomodule
1. `go.mod`会将所有的依赖存储在统一的地方，但是会按照`mod`分类，并且会有不同版本
2. 通过`go get`下载的依赖会记录在`go.mod`中，比如
```terminal
require (
	go.uber.org/atomic v1.10.0 // indirect，indirect表示没用到
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)
```
还会生成一个`go.sum`
3. 添加依赖有两种方式
   1. `go get`
   2. 直接在代码里`import`，编译时会自动下载
4.Go mod下如何引用本项目的包呢？
大家只要记住这个公式即可：

import的路径 = go.mod下的module name + 包相对于go.mod的相对目录

举例：
我们的课程示例代码中：https://git.imooc.com/coding-180/coding-180/src/master/go.mod

module imooc.com/ccmouse/learngo
tree包在：https://git.imooc.com/coding-180/coding-180/src/master/lang/tree

相对于go.mod的相对路径为：lang/tree

所以引用tree包时：

import "imooc.com/ccmouse/learngo/lang/tree"
在项目内的任一地方，都是用这个路径进行import。
5. 更新依赖
   1. `go get` + version
   2. `go mod tidy`

### 5-4 迁移到`gomodule`
1. 如果没有`go.mod`文件，先init
```
go mod init {module name}
```
2. 将当前目录下的所有文件都build，去拉取所有依赖
```
go build ./...
```

## 6 - 接口
### 6-1 如何实现接口
1. 首先定义接口
```go
package retriever

type retriever interface {
   GET(url string) string
}
```
2. 所有实现了这个函数的类型都视为实现了这个接口，不需要像Java一样显式地声明`implements`。
   比如下面的`infra.Retriever`和`test.Retriever`都实现了这个函数，都是实现了这个接口
```go
package infra

import (
   "fmt"
   "io/ioutil"
   "net/http"
)

type Retriever struct{}

func (Retriever) GET(url string) string {
   resp, err := http.Get(url)
   if err != nil {
      panic(err)
   }

   defer resp.Body.Close()

   bytes, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      panic(err)
   }
   return string(bytes)
}
```

```go
package test

type Retriever struct{}

func (Retriever) GET(url string) string {
	return "fake content"
}
```

### 6-2 Duck Typing
没有Java的继承、多态，这些在Go里都是通过接口来实现的。
1. "像鸭子，那么就是鸭子"
2. 描述事物的外部行为而不是内部结构
3. 严格来说go属于结构化类型系统，不属于duck typing，因为duck typing需要动态绑定，而go是在编译时就确定了的
4. go的duck typing好处在于
   1. 具有类似Python、C++的duck typing的灵活性
   2. 同时又具有Java的类型检查

#### Python Duck Typing

#### C++ Duck Typing
编译的时候才能知道传入的`retriever`有没有实现`get`，打代码的时候不知道
```
template <class R>
string download(const R& retriever) {
   return retriever.get("www.google.com")
}
```

#### Java 没有 Duck Typing，但是有类似的
- 传入的参数必须实现Retriever接口
- 不是duck typing
- 无法实现多个接口，`extends`只能用一个，有一种解决办法是`apache polygen`，但是很难用，Java就不是这么设计的
```java
<R extends Retriever>
String download(R r) {
  return r.get("www.google.com");
}
```

### 接口变量
1. 接口变量自带指针
2. 接口变量同样采用值传递，几乎不需要使用接口类型的指针
3. 指针接受者实现只能用指针来调用方法，值接受者则都可以

### 
1. `interface{}`表示任何类型

### 接口的组合


### 6-6 go常用接口
#### 6-6-1 `Stringer` Interface && `String() string` Method
相当于Java的`toString()`，
```go
type Stringer interface {
   String() string
}
```

#### 6-6-2 `Reader` Interface && `Writer` Interface
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
```

比如 `File` 就实现了 `Reader` interface 和 `Writer` interface。
很多底层的读写相关的工作都要写成 `Reader` 或者 `Writer`，这样就可以和系统的一系列函数合作使用。