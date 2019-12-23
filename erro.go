package erro

import (
	"container/list"
	"encoding/json"
	"runtime"
	"strconv"
)

type Error struct {
	Info      interface{} `json:"info"`
	StackInfo string      `json:"stack_info"`
}

// new一个error
func NewError(info interface{}) *Error {
	return &Error{
		Info:      info,
		StackInfo: getStackInfo(),
	}
}

// string err
func (err *Error) String() string {
	content, _ := json.Marshal(err)
	return string(content)
}

// 获取错误堆栈信息， 从最外层到最里层
func getStackInfo() string {
	var i = 0
	var stackList list.List
	for {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stackList.PushFront("file: " + file + ";line: " + strconv.FormatInt(int64(line), 10) + "\n")
		i++
	}
	// 生成堆栈报错信息
	var stackInfo string
	e := stackList.Front()
	for i := 0; i < stackList.Len(); i++ {
		stackInfo += e.Value.(string)
		e = e.Next()
	}
	return stackInfo
}
