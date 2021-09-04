package _10904

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)


type CacheEntry struct {
	data []byte
	wait chan struct{}
}

type QueryClient struct {
	cache map[string]*CacheEntry
	mutex *sync.Mutex
}

func (c *QueryClient) DoQuery(name string) ([]byte, error) {
	// 检查操作是否已启动
	c.mutex.Lock()
	if cached, found := c.cache[name]; found {
		c.mutex.Unlock()
		// 等待完成
		val, ok := <-cached.wait
		log.Println("从缓存中读取，此时管道关闭可以读取数据")
		log.Println(val)
		log.Println(ok)
		return cached.data, nil
	}

	entry := &CacheEntry{
		//data: result,
		wait: make(chan struct{}),
	}
	c.cache[name] = entry
	defer c.mutex.Unlock()

	// 如果未缓存，则发出请求
	resp, err := http.Get("https://www.baidu.com/" + url.QueryEscape(name))
	if err!=nil{
		return nil, err
	}
	defer resp.Body.Close()
	//err = resp.Body.Close()
	//if err!=nil{
	//	return nil, err
	//}

	// 为简洁起见，省略了错误处理和 resp.Body.Close
	entry.data, err = ioutil.ReadAll(resp.Body)



	// 关闭 channel，传递操作完成信号
	// 立即返回
	close(entry.wait)

	val, ok := <-entry.wait
	log.Println("获取缓存关闭管道后")
	log.Println(val)
	log.Println(ok)

	return entry.data,nil
}



type CacheEntryOnce struct {
	data []byte
	once *sync.Once
}

type QueryClientOnce struct {
	cache map[string]*CacheEntryOnce
	mutex *sync.Mutex
}

func (c *QueryClientOnce) DoQueryOnce(name string) ([]byte ,error){
	c.mutex.Lock()
	entry, found := c.cache[name]
	if !found {
		// 如果在缓存中未找到，创建新的 entry
		entry = &CacheEntryOnce{
			once: new(sync.Once),
		}
		c.cache[name] = entry
	}
	c.mutex.Unlock()

	// 现在，当我们调用 .Do 时，如果有一个正在同步进行的操作
	// 它将一直阻塞，直到完成（并填充 entry.data）
	// 或者如果操作之前已经完成过一次
	// 本次调用不会进行操作，也不会阻塞
	entry.once.Do(func() {
		resp, err := http.Get("https://www.baidu.com/" + url.QueryEscape(name))
		if err!=nil{
			log.Println(err)
            return
		}
		defer resp.Body.Close()
		// 为简洁起见，省略了错误处理和 resp.Body.Close
		entry.data, err = ioutil.ReadAll(resp.Body)
	})

	return entry.data,nil
}