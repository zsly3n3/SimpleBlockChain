package blockChain

import (
	"fmt"
	"sync"
	trans "SimpleBlockChain/transaction"
)

type blockChain struct {	
	Chain []block //块数组
	diff int //挖矿难度系数
	pendingTransactions []trans.Transaction //该链中等待处理的交易列表
	miningReward float64 //挖矿成功的奖励数目
	miningRewardData *rewardData //保存挖矿成功后的一些数据
}

type rewardData struct {
	 m *sync.RWMutex //读写互斥量
	 miningRewardAddress string //挖矿者的钱包地址
	 roundOver bool //是否一轮结束
}

/*创建链*/
func Create() *blockChain{
	bc:= new(blockChain)
	bc.Chain = make([]block,1)
	bc.pendingTransactions = make([]trans.Transaction,0)//在新链上,无任何要处理的交易列表
	block:=bc.createGenesisBlock()//生成创世块
	bc.Chain[0]= *block
	bc.diff = 4 
	bc.miningReward = 100.0
	bc.miningRewardData = &rewardData{
		m: new(sync.RWMutex),
		roundOver:false,miningRewardAddress:""}
	return bc
}

func (bc *blockChain) createGenesisBlock() *block{
	transactions := make([]trans.Transaction,0)//在创世块前,没有任何交易列表
    return createBlock(transactions,"")
}

/*获取链上最后一个块*/
func (bc *blockChain) getLatestBlock() *block{
    return &bc.Chain[len(bc.Chain)-1]
}

/*挖矿*/
func (bc *blockChain)MinePendingTransactions(miningRewardAddress string,diff int){
	b:=createBlock(bc.pendingTransactions,bc.getLatestBlock().hash)//传入待处理的交易列表和上一块的hash值
	b.mineBlock(diff,bc,miningRewardAddress)//传入难度系数,链地址,挖矿者地址后进行POW
	bc.miningRewardData.m.RLock()//加读锁
	address:= bc.miningRewardData.miningRewardAddress
	bc.miningRewardData.m.RUnlock()
	if miningRewardAddress == address{//该钱包地址等于成功地址
		bc.Chain=append(bc.Chain,*b)//在链上添加该块
		bc.pendingTransactions =  make([]trans.Transaction,1)//重置待处理交易列表,初始一条交易信息
		t:=trans.Create("",miningRewardAddress,bc.miningReward)//发送奖励
		bc.pendingTransactions[0]=*t
		nonce_str:=fmt.Sprintf("%d", b.nonce)
		fmt.Println(miningRewardAddress+" Mining successed ! nonce: "+nonce_str)
	}
}

/*创建交易数据*/
func (bc *blockChain)CreateTransaction(fromAddress string,toAddress string,amount float64){
	t:=trans.Create(fromAddress,toAddress,amount)
	bc.pendingTransactions=append(bc.pendingTransactions,*t)//添加到待处理列表
}

/*获取该地址的余额*/
func (bc *blockChain)GetBalanceOfAddress(miningRewardAddress string) float64{
	balance:=0.0//余额初始为0
	for  _, b := range bc.Chain{//遍历整条链的块
		for _, t := range b.transactions{//遍历块中的交易列表
			if t.FromAddress == miningRewardAddress{//如果地址是发起方,减少余额
				balance -= t.Amount
			}
			if t.ToAddress == miningRewardAddress{//如果地址是接收方,增加余额
				balance += t.Amount
			}
		}
	}
	return balance
}

/*验证该链是否有效,来确保没有人篡改过区块链*/
func (bc *blockChain) IsChainValid() bool{
	 length:= len(bc.Chain)
	 for i := 1; i < length ; i++ {//遍历所有块
		currentBlock:=bc.Chain[i] 
		previousBlock:=bc.Chain[i-1]
		if currentBlock.hash != currentBlock.calculateHash(){//每个区块的hash是否正确
			return false
		}
		if currentBlock.PreviousHash != previousBlock.hash{//每个区块PreviousHash是否等于上一个块的hash
			return false
		}
     }
	 return true
}

/*重置一些数据*/
func (bc *blockChain) ReSetData(){
	 bc.miningRewardData.roundOver = false
}





