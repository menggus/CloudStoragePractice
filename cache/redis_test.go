package cache

import (
	"testing"
)

func TestRedis(t *testing.T) {
	res := NewRedis()
	t.Log("测试redis客户端的连接")
	{
		if res == nil {
			t.Fail()
		}
	}

}
