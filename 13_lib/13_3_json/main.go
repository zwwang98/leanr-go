package main

import (
	"encoding/json"
	"fmt"
)

// Intellij IDEA查看json数据的小技巧
// 1.新建scratch file，选择json
// 2.将json数据粘贴到新建的scratch file中
// 3.在Intellij IDEA toolbar选择code->reformat code

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Order
//  1. 在struct内部，首字母大写的字段才是public的，因此，在利用json包对struct处理时，只有首字母大写的字段会被处理，其余的则会被忽略。
//  2. 但是在利用json进行数据传输时，不是所有语言都像go一样的命名习惯，我们想所有字段都变成小写的。
//     这可以通过json tag做到，json.Marshal()就会选择json tag里的命名。
//  3. 加入了omitempty之后，在json.Marshal()时，如果该字段为空，则不会在最终的json里出现。
type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"item"`
	TotalPrice float64     `json:"total_price"`
}

func unmarshall() {
	s := `{"id":"1234","item":[{"id":"item_1","name":"learn go","price":15},{"id":"item_1","name":"interview","price":10}],"total_price":20}`
	var o Order
	// 由于json.Unmarshal()函数传参类型要求，要先将字符串s转化为[]byte，并且第二个参数要用一个目标类型的变量的地址，这是因为
	// unmarshal的结果会存储到这个变量去
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(interface{}(err))
	}
	fmt.Printf("%+v\n", o)
}

func main() {
	o := Order{
		ID:         "1234",
		TotalPrice: 20,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_1",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	buf, err := json.Marshal(o)
	if err != nil {
		panic(interface{}(err))
	}

	fmt.Printf("%s\n", buf)

	unmarshall()
}
