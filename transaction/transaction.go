package transaction


type Transaction struct {
	FromAddress string  `json:"fromAddress"` //发起方钱包地址
	ToAddress string    `json:"toAddress"`   //接受方钱包地址
	Amount float64          `json:"amount"` //交易数目
}

/*创建交易数据*/
func Create(fromAddress string,toAddress string,amount float64) *Transaction{
	t := new(Transaction)
	t.FromAddress = fromAddress
	t.ToAddress = toAddress
	t.Amount = amount
	return t
}