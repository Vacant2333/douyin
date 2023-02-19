package checkpwd

import (
	"fmt"
	"testing"
)

func TestCheckPassword(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := CheckPassword("@2019") // 程序输出的结果
	fmt.Printf("result:%v\n", got)
}
