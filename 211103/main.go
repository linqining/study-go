package _11103

import (
	"fmt"
	"math/rand"
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

