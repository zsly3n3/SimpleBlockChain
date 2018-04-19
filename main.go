package main

import (
	"fmt"
	"time"
	"SimpleBlockChain/blockChain"
)

func main() {
	   
		bc:=blockChain.Create()//创建一个链
		
		fmt.Println("Starting the miner...");
		bc.ReSetData()//挖矿前重置数据

		go bc.MinePendingTransactions("zsly-address1",3)//开启协程1 挖矿,奖励保存到等待交易列表中
		go bc.MinePendingTransactions("zsly-address2",3)//开启协程2 挖矿,奖励保存到等待交易列表中
		go bc.MinePendingTransactions("zsly-address3",3)//开启协程3 挖矿,奖励保存到等待交易列表中
		
		time.Sleep(time.Second*5)

		//获取余额
		b1:=bc.GetBalanceOfAddress("zsly-address1")
		b2:=bc.GetBalanceOfAddress("zsly-address2")
		b3:=bc.GetBalanceOfAddress("zsly-address3")
	
		/*下面余额都为0,并不是没人挖到矿,而是挖矿成功后,奖励保存到等待交易列表中,还没记录到链上*/
		fmt.Println(b1)//余额为0
		fmt.Println(b2)//余额为0
		fmt.Println(b3)//余额为0
		
		
		fmt.Println("Starting the miner Again...");
		bc.ReSetData()//挖矿前重置数据
		bc.MinePendingTransactions("zsly-address4",3)//继续挖矿,成功后,之前等待处理的交易数据已被记录到链上

		//获取余额
		b1=bc.GetBalanceOfAddress("zsly-address1")
		b2=bc.GetBalanceOfAddress("zsly-address2")
		b3=bc.GetBalanceOfAddress("zsly-address3")

		fmt.Println(b1)
		fmt.Println(b2)
		fmt.Println(b3)
}
