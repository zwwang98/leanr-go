1. 命名一般是 camel-case
   1. 首字母大写表示 package-public
   2首字母小写表示 package-private
2. package - 包
   1. 每个目录一个包
   2. main package 包含程序入口 - main() 方法
   3. 为同种结构定义的方法需要放在同一个 package 下，但不一定要是同一个文件
3. 举例来说，这里我们
   1. 把 node 单独放在一个文件里
   2. 把 traverse 方法单独放在一个文件里
   3. 在 tree package 下新建一个 packag main，在这里放入 main() 作为整个 package 的 entry
