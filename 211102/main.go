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

//描述
//22娘和33娘接到了小电视君的扭蛋任务：
//一共有两台扭蛋机，编号分别为扭蛋机2号和扭蛋机3号，22娘使用扭蛋机2号，33娘使用扭蛋机3号。
//扭蛋机都不需要投币，但有一项特殊能力：
//扭蛋机2号：如果塞x（x范围为>=0正整数）个扭蛋进去，然后就可以扭到2x+1个
//扭蛋机3号：如果塞x（x范围为>=0正整数）个扭蛋进去，然后就可以扭到2x+2个
//22娘和33娘手中没有扭蛋，需要你帮她们设计一个方案，两人“轮流扭”（谁先开始不限，扭到的蛋可以交给对方使用），用“最少”的次数，使她们能够最后恰好扭到N个交给小电视君。
//输入描述：
//输入一个正整数，表示小电视君需要的N个扭蛋。
//输出描述：
//输出一个字符串，每个字符表示扭蛋机，字符只能包含"2"和"3"。

func GetNDan(n int) string{
	if n ==0{
		return ""
	}
	var step string
	if n%2 ==1{
		m := (n-1)/2
		step = GetNDan(m) + "2"
	}else{
		m := (n-2)/2
		step = GetNDan(m) + "3"
	}
	return step
}

//描述
//一个数组有 N 个元素，求连续子数组的最大和。 例如：[-1,2,1]，和最大的连续子数组为[2,1]，其和为 3
//输入描述：
//输入为两行。 第一行一个整数n(1 <= n <= 100000)，表示一共有n个元素 第二行为n个数，即每个元素,每个整数都在32位int范围内。以空格分隔。
//输出描述：
//所有连续子数组中和最大的值。

//func MaxSubArr(arr []int) int{
//	max := arr[0]
//	subMax := MaxSubArr(arr[1:])
//	if subMax>0{
//		if max <0{
//			if subMax> max{
//				return subMax
//			}else{
//				return max
//			}
//		}else{
//			if subMax>0{
//				return max + subMax
//			}else{
//				return max
//			}
//		}
//	}
//}

// 返回开始结束索引
func MaxSubArr(arr []int,inLeft bool) (int,int, int){
	if len(arr)==1{
		return 0,1,arr[0]
	}
	midIndex := len(arr)/2
	leftStart,leftEnd,leftMax := MaxSubArr(arr[:midIndex],true)
	rightStart,rightEnd,rightMax := MaxSubArr(arr[midIndex:],false)

	fmt.Println("输入的数组",arr)
	fmt.Println("左侧数据",leftStart,leftEnd,leftMax)
	fmt.Println("右侧数据",rightStart+midIndex,rightEnd+midIndex,rightMax)

	tmp := leftMax+rightMax
	for i:=leftEnd+1;i<midIndex;i++{
		tmp += arr[i]
	}
	for j:= 0;j<rightStart;j++{
		tmp+=arr[midIndex+j]
	}
	fmt.Println("总值", tmp)
	if (tmp>leftMax) &&(tmp>rightMax){
		fmt.Println("返回值",leftStart,rightEnd+midIndex,tmp)
		return leftStart,rightEnd+midIndex,tmp
	}else{
		if leftMax>rightMax{
			fmt.Println("返回值",leftStart,leftEnd,leftMax)
			return leftStart,leftEnd,leftMax
		}else if leftMax<rightMax{
			fmt.Println("返回值",rightStart+midIndex,rightEnd+midIndex,rightMax)
			return rightStart+midIndex,rightEnd+midIndex,rightMax
		}else{
			if inLeft{
				fmt.Println("返回值",leftStart,leftEnd,leftMax)
				return leftStart,leftEnd,leftMax
			}else{
				fmt.Println("返回值",rightStart+midIndex,rightEnd+midIndex,rightMax)
				return rightStart+midIndex,rightEnd+midIndex,rightMax
			}
		}
	}
}

func MaxSubArr2(arr []int)int{
	n := len(arr)
	max:=arr[0]
	sum:=arr[0]
	for i:= 1; i < n; i++ {
		num := arr[i]
		if sum<0{
			sum = num
		}else{
			sum += num
		}
		if sum>max{
			max=sum
		}
	}
	return max
}


//有一个长为 n 的数组 A ，求满足 0 ≤ a ≤ b < n 的 A[b] - A[a] 的最大值。
//
//给定数组 A 及它的大小 n ，请返回最大差值。
//
//数据范围：  ，数组中的值满足

func Maxdiff(arr []int)int{
	max :=arr[0]
	min := arr[0]
	diff := 0
	for i:=0;i<len(arr);i++{
		elm := arr[i]
		if elm>max{
			max = elm
			if max>min{
				if (max-min)>diff{
					diff=max-min
				}
			}
		}
		if elm<min{
			min = elm
			max = elm
		}
	}
	return diff
}


//在一组数的编码中，若任意两个相邻的代码只有一位二进制数不同， 则称这种编码为格雷码(Gray Code)，请编写一个函数，使用递归的方法生成N位的格雷码。
//
//给定一个整数n，请返回n位的格雷码，顺序为从0开始。

func GrayCode(n int) []string{
	if n==1{
		return []string{"0","1"}
	}
	codeList := GrayCode(n-1)
	ret := make([]string,2*len(codeList))
	for i:=0;i<len(codeList);i++{
		ret[i] = "0" + codeList[i]
		ret[i+len(codeList)]= "1"+codeList[i]
	}
	return ret
}

//描述
//小明在越南旅游，参加了当地的娱乐活动。小明运气很好，拿到了大奖， 到了最后的拿奖金环节。小明发现桌子上放着一列红包，每个红包上写着奖金数额。
//现在主持人给要求小明在这一列红包之间“切”2刀，将这一列红包“切”成3组，并且第一组的奖金之和等于最后一组奖金和（允许任意一组的红包集合是空）。最终第一组红包的奖金之和就是小明能拿到的总奖金。小明想知道最多能拿到的奖金是多少，你能帮他算算吗。
//
//举例解释：桌子上放了红包  1, 2, 3, 4, 7, 10。小明在“4,7”之间、“7,10” 之间各切一刀，将红包分成3组 [1, 2, 3, 4]   [7]   [10]，其中第一组奖金之和=第三组奖金之和=10，所以小明可以拿到10越南盾。
//输入描述：
//第一行包含一个正整数n，(1<=n<= 200 000)，表示有多少个红包。
//
//第二行包含n个正整数d[i]，表示每个红包包含的奖金数额。其中1<= d[i] <= 1000 000 000
//输出描述：
//小明可以拿到的总奖金
//示例1

func GetWonPrice(arr []int) int{
	money := 0

	rightIndex :=len(arr)
	if rightIndex ==0{
		return 0
	}
	leftSum :=arr[0]
	rightSum:=0
	leftIndex:=0
	for leftIndex<rightIndex{
		if leftSum>rightSum {
			rightIndex--
			rightSum += arr[rightIndex]
		}else if leftSum<rightSum{
			leftIndex++
			leftSum += arr[leftIndex]
		}else{
			break
		}
	}
	if leftSum!=rightSum{
		return 0
	}else{
		money += leftSum
		submon :=GetWonPrice(arr[leftIndex+1:rightIndex])
		return leftSum + submon
	}
}