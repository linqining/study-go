package _10904

import (
	"sync"
	"testing"
)

// NOTE
func TestQueryClient_DoQuery(t *testing.T) {
	client := QueryClient{mutex: &sync.Mutex{}, cache: map[string]*CacheEntry{}}
	resp, err := client.DoQuery("test")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(string(resp))

	// 第二次从缓存获取
	resp, err = client.DoQuery("test")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(string(resp))

}

func TestQueryClientOnce_DoQueryOnce(t *testing.T) {
	client := QueryClientOnce{mutex: &sync.Mutex{}, cache: map[string]*CacheEntryOnce{}}
	resp, err := client.DoQueryOnce("test")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(string(resp))

	// 第二次从缓存获取
	resp, err = client.DoQueryOnce("test")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(string(resp))
}