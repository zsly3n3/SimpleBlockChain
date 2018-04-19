package blockChain

import (
	"fmt"
	"time"
	"encoding/json"//json封装解析
	trans "SimpleBlockChain/transaction"
	"SimpleBlockChain/utils"
)

type block struct {
	timestamp int64
	transactions []trans.Transaction //等待处理的交易数据
	PreviousHash string //保存上一个块的hash值,用于校验
	hash string
	nonce int //挖矿时计算的次数
}

/*创建块*/
func createBlock(transactions []trans.Transaction,previousHash string) *block{
	b := new(block)
	b.timestamp = time.Now().Unix()
	b.transactions = transactions 
	b.PreviousHash = previousHash
	b.hash = b.calculateHash() 
	b.nonce = 0
	return b
}

/*计算当前块的hash值*/
func (b *block) calculateHash() string {
   trans_str:=""
   j, err := json.Marshal(b.transactions)//序列化
   if err== nil{
	   trans_str = string(j)
   }
   timestamp_str := fmt.Sprintf("%d", b.timestamp)
   nonce_str := fmt.Sprintf("%d", b.nonce)
   rs:= timestamp_str + trans_str + nonce_str
   return utils.GetMd5String(rs)
}

/*获取指定难度匹配的字符串*/
func getMatchString(diff int,tmp string) string{
	rs:=""
	for i:=0;i<diff;i++{
	   rs+=tmp
	}
	return rs
}

/*在区块链中进行POW*/
func (b *block) mineBlock(diff int,bc *blockChain,miningRewardAddress string){
	   if diff <= 0{ /*挖矿难度系数要大于0*/
		   diff = 1
		}
		length:=len(b.hash)
		if diff >= length{/*难度系数不能超过hash字符个数*/
		   diff=length
		}
		tmp:=getMatchString(diff,"0") /*获取全为0的字符串*/
		for{
			   bc.miningRewardData.m.RLock()//加读锁
			   roundOver:=bc.miningRewardData.roundOver
			   bc.miningRewardData.m.RUnlock()
			   if roundOver{//被通知这一轮结束,挖矿失败
				   nonce_str:=fmt.Sprintf("%d", b.nonce)
				   fmt.Println(miningRewardAddress+" round over ! nonce: "+nonce_str)
				   return
			   }
			   if b.hash[0:diff] == tmp{//hash匹配成功,即挖矿成功
				   bc.miningRewardData.m.Lock()//加写锁
				   bc.miningRewardData.roundOver = true//通知这一轮的其他人结束挖矿
				   bc.miningRewardData.miningRewardAddress = miningRewardAddress//保存该挖矿者的钱包地址
				   bc.miningRewardData.m.Unlock()
				   break
			   }
			   b.nonce++ //统计计算次数
			   b.hash = b.calculateHash()/*计算hash*/		
		}
}

