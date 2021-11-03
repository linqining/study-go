package _11103

import (
	"fmt"
	"math/rand"
	"sort"
)

type point struct{
	x int
	y int
}

//小东所在公司要发年终奖，而小东恰好获得了最高福利，他要在公司年会上参与一个抽奖游戏，游戏在一个6*6的棋盘上进行，上面放着36个价值不等的礼物，每个小的棋盘上面放置着一个礼物，他需要从左上角开始游戏，每次只能向下或者向右移动一步，到达右下角停止，一路上的格子里的礼物小东都能拿到，请设计一个算法使小东拿到价值最高的礼物。
//
//给定一个6*6的矩阵board，其中每个元素为对应格子的礼物价值,左上角为[0,0],请返回能获得的最大价值，保证每个礼物价值大于100小于1000。

func GetMostBonus(){
	board := GenBoard()
	p := point{0,0}
	bonus,path :=GetBonus(board[0:6][0:6],p)
	fmt.Println("总奖金",bonus)
	fmt.Println("路径",path)
}

func GetBonus(board [][]int,p point)(int,[]point){
	path := []point{}
	path = append(path,p)
	length := len(board)
	currentBonus := board[p.x][p.y]
	//var minIdx int
	var rBonus,dBonus int
	var rPath,dPath []point
	// 向右移动
	if p.y >=length-1{
		// 不能向右移动了
		rBonus =0
	}else{
		pr := point{x:p.x,y:p.y+1}
		//传入的point是相对地址
		rBonus,rPath = GetBonus(board, pr)
	}

	if p.x >=length-1{
		// 不能向下移动了
		dBonus =0
	}else{
		//向下移动
		pd := point{x:p.x+1, y:p.y}
		//处理边界问题

		//传入的point是相对地址
		dBonus,dPath = GetBonus(board,pd)
	}
	if rBonus>dBonus{
		currentBonus += rBonus
		path = append(path,rPath...)
	}else if rBonus<dBonus{
		currentBonus += dBonus
		path = append(path,dPath...)
	}else{
		currentBonus+=rBonus
		path = append(path,rPath...)
	}
	return currentBonus,path
}

func GenBoard()[][]int{
	ret := make([][]int,6)
	for i:=0;i<6;i++{
		row := make([]int,6)
		for j:=0;j<6;j++{
			randAmount := 101+rand.Intn(899)
			row[j] = randAmount
			fmt.Print(randAmount,"\t")
		}
		ret[i]=row
		fmt.Println("")
	}
	return ret
}

func JumpToX(x int)int{
	if x ==0{
		return 0
	}
	if x ==1{
		return 1
	}
	step :=0
	if x%2 ==0{
		prevStep := JumpToX(x/2)
		step = 1 +prevStep
	}else{
		x = x-1
		remainStep := JumpToX(x)
		step = 1+ remainStep
	}
	return step
}



//小A参加了一个n人的活动，每个人都有一个唯一编号i(i>=0 & i<n)，其中m对相互认识，在活动中两个人可以通过互相都认识的一个人介绍认识。现在问活动结束后，小A最多会认识多少人？
//
//输入描述：
//第一行聚会的人数：n（n>=3 & n<10000）；
//第二行小A的编号: ai（ai >= 0 & ai < n)；
//第三互相认识的数目: m（m>=1 & m
//< n(n-1)/2）；
//第4到m+3行为互相认识的对，以','分割的编号。
//输出描述：
//输出小A最多会新认识的多少人？

type uset struct{
	data []int
}

func FindFriends(personCount int){
	unionSet := NewUset(10)
	fmt.Println(unionSet)
	unionSet.Union(1,0)
	unionSet.Union(3,1)
	unionSet.Union(4,1)
	unionSet.Union(5,3)
	unionSet.Union(6,1)
	unionSet.Union(6,5)
	for _,v :=range unionSet.data{
		v := v
		fmt.Println(v)
	}
	fmt.Println(unionSet.getFriendCount(5))
}

func NewUset(personCount int) *uset{
	unionSet := uset{data:make([]int, personCount)}
	for i,_ :=range unionSet.data{
		unionSet.data[i] = i
	}
	return &unionSet
}

func (u *uset)Union(a int, b int){
	if u.data[a]==u.data[b]{
		return
	}
	// 已经有主，全部投靠a
	if u.data[b]!=b{
		bParent := u.data[b]
		for i,v :=range u.data{
			if v==bParent{
				u.data[i] = u.data[a]
			}
		}
	}else{
		u.data[b] = u.data[a]
	}
}

func (u *uset)getFriendCount(num int)int{
	parent := u.data[num]
	var count int
	for _,v :=range u.data{
		if v==parent{
			count++
		}
	}
	if count>0 && parent==num{
		// 除去自己
		count --
	}
	// 还要减去自己已经认识的
	return count
}

func CountBitDiff( m int32, n int32) int{
	res := m^n
	count := 0
	var mask int32 =1
	for i:=0;i<31;i++{
		if res & (mask<<i)>0{
			count++
		}
	}
	return count
}

func RestaruantMaxIncome(tableCap []int, customerG [][]int) int {
	if len(tableCap)==0 || len(customerG)==0{
		return 0
	}
	tableIndex :=0
	// 先把桌子按照容量排序
	sort.Ints(tableCap)

	sort.Slice(customerG,func(i int,j int) bool {
		a := customerG[i]
		b := customerG[j]
		if a[0]<b[0]{
			return true
		}
		return false
	})
	// 如果桌子坐不下人，把桌子扔掉
	minCount := customerG[0][0]
	fmt.Println(customerG)

	for{
		if tableCap[tableIndex]<minCount{
			tableIndex++
			if tableCap[tableIndex]>minCount{
				break
			}
		}else{
			break
		}
	}

	// 一个坐得下的桌子都没有
	if tableIndex == len(tableCap){
		return 0
	}
	customerIndex := pickMostVC(customerG,tableCap[tableIndex])

	currentIncome := customerG[customerIndex][1]

	// 移除已安排座位的顾客
	newCustomerG := [][]int{}
	for i,v :=range customerG{
		v := v
		if i!=customerIndex{
			newCustomerG = append(newCustomerG,v)
		}
	}
	remainIncome := RestaruantMaxIncome(tableCap[tableIndex+1:],newCustomerG)

	return currentIncome+remainIncome


	// 逐个安排桌子如果一个桌子只能坐下一组人，那么就是唯一选项


	//如果桌子较大，并且有多组可以分配的，那么选择金额大的

	// 剩下的递归调用

}

func pickMostVC(customerGroup [][]int,tableCap int) int{
	maxIndex := 0
	maxVal:=0
	for i,v :=range customerGroup{
		if v[0]>tableCap{
			return maxIndex
		}
		if v[0]<=tableCap {
			if v[1]>maxVal{
				maxVal = v[1]
				maxIndex = i
			}
		}else{
			return maxIndex
		}
	}
	return maxIndex
}

type QScore struct{
	set int
	val int
}

//描述
//小M和大M要通过选择元素组成队列进行得分pk。目前存在队列S1,S2,S3...Sn,每个元素包括2个正整数属性set和value.从中选出任意K个元素S[i1],S[i2]...S[ik],保证顺序不变即i1 < i2 < i3< ... < ik,组成新的队列P1,P2,P3......Pk.我们通过一个机制评价队列的好坏:
//
//Base=P1.value+P2.value+...Pk.value,
//
//Bonus=10*t;t为新队列中Pi.set=P(i+1).set的i个数.
//
//最终得分Score=Base-Bonus;求Score的最大值和取最大值时新队列元素个数的最小值.
//
//输入描述：
//第一行包含一个数N(0 < N <= 500)
//
//接下来N行每一行两个数表示S1,S2,...,Sn的set和value值 (0 < set,value <= 20)
//输出描述：
//Score的最大值和新队列元素个数的最小值

func SerialMax(qList []QScore) (score int, lastq QScore){
	if len(qList)==1{
		lastq = qList[0]
		return lastq.val,lastq
	}
	lastq = qList[len(qList)-1]
	last2Score, last2p := SerialMax(qList[:len(qList)-1])
	var bonus,tmpScore int
	if last2p.set == lastq.set{
		bonus = 10
		tmpScore = last2Score + lastq.val - bonus
	}else{
		tmpScore = last2Score + lastq.val
	}

	if tmpScore> last2Score{
		return tmpScore,lastq
	}else{
		return last2Score, last2p
	}
}


type Node struct{
	val int
	left *Node
	right *Node
}

func buildTree(nodes []int)*Node{
	if len(nodes)==1{
		return &Node{val:nodes[0]}
	}
	midVal := nodes[0]
	mid := Node{val:midVal}
	leftNodes := nodes[1:(len(nodes)+1)/2]
	left := buildTree(leftNodes)
	rightNodes := nodes[(len(nodes)+1)/2:]
	right := buildTree(rightNodes)
	mid.left=left
	mid.right = right
	return &mid
}


//描述
//给出满二叉树，编写算法将其转化为求和树
//
//什么是求和树：二叉树的求和树， 是一颗同样结构的二叉树，其树中的每个节点将包含原始树中的左子树和右子树的和。
//
//二叉树：
//10
///      \
//-2        6
///   \      /  \
//8    -4    7    5
//
//求和树：
//20(4-2+12+6)
///      \
//4(8-4)      12(7+5)
///   \      /  \
//0      0    0    0
//
//二叉树给出前序和中序输入，求和树要求中序输出；
//所有处理数据不会大于int；
//
//输入描述：
//2行整数，第1行表示二叉树的前序遍历，第2行表示二叉树的中序遍历，以空格分割
//输出描述：
//1行整数，表示求和树的中序遍历，以空格分割

func SumTree(){
	node := buildTree([]int{10,-2,8,-4,6,7,5})
	sum,_:=node.Sum()
	fmt.Println(sum)
}

func (n *Node)Sum() (sum int, origin int){
	if n.left==nil && n.right==nil{
		return 0, n.val
	}
	origin = n.val
	leftSum, lOrigin := n.left.Sum()
	rightSum,rOrigin := n.right.Sum()
	return leftSum+rightSum+lOrigin+rOrigin,n.val
}