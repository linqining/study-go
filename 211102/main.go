package _11102

import (
	"bytes"
	"fmt"
	"sync"
)

func InitSlice(){
	var a []int
	a = append(a,1)
	fmt.Println(a)
}

func ChanPanic(){
	ch := make(chan int,1000)
	ch<-1
	for{
		<-ch

		ch<-1
	}
}

func PrintNums(){
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++{
		go func(){
			for {
				num,ok := <-ch
				if ok{
					fmt.Println(num)
					if num==100{
						close(ch)
					}else{
						ch<-(num+1)
					}
				}else{
					wg.Done()
					return
				}
			}
		}()
	}
	ch<-1
	wg.Wait()
}

//描述
//从非负整数序列 0, 1, 2, ..., n中给出包含其中n个数的子序列，请找出未出现在该子序列中的那个数。
//输入描述：
//输入为n+1个非负整数，用空格分开。
//其中：首个数字为非负整数序列的最大值n，后面n个数字为子序列中包含的数字。
//输出描述：
//输出为1个数字，即未出现在子序列中的那个数。
type Bitmap struct {
	words  []uint64
	length int
}

func (bitmap *Bitmap) Has(num int) bool {
	word, bit := num/64, uint(num%64)
	return word < len(bitmap.words) && (bitmap.words[word]&(1<<bit)) != 0
}



func (bitmap *Bitmap) Add(num int) {
	word, bit := num/64, uint(num%64)
	for word >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}
	// 判断num是否已经存在bitmap中
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
		bitmap.length++
	}
}

func (bitmap *Bitmap) Len() int {
	return bitmap.length
}

func (bitmap *Bitmap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range bitmap.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if v&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*uint(i)+j)
			}
		}
	}
	buf.WriteByte('}')
	fmt.Fprintf(&buf,"\nLength: %d", bitmap.length)
	return buf.String()
}


func FindMiss(numArr []int)int{
	bm := &Bitmap{}
	max := numArr[0]

	for _,item:=range numArr{
		bm.Add(item)
	}
	counter := 0
	for i:=0;i<=max;i++{
		if !bm.Has(i){
			counter ++
		}
	}
	return counter
}