### 接收 json - 使用 map && 自定义 struct
1. 如果是使用map，接收起来很方便。在go里用于接收json的map类型必须定义成下面这样
    ```go
   m := make(map[string]interface{})
   ```
   因为json的key一定是string，而value可能是数字、string、也可能是数组、object
2. 因为json接收map的类型限制，所以当我们从json map里提取数据时，需要不断地用type assertion来提取其中的数据，很麻烦
3. 所以推荐使用自定义struct来接收json