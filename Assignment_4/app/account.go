package app

type Account struct {
	FirstName, LastName string
	AccountNumber       int
	Balance             int
}

var allAccts = make(map[int]Account)

func AddAccount(acc Account) map[int]Account {
	allAccts[acc.AccountNumber] = acc
	return allAccts
}

func RemoveAccount(accountNum int) map[int]Account {
	delete(allAccts, accountNum)
	return allAccts
}

func AccountBalance(accountNum int) int {
	return allAccts[accountNum].Balance
}

func DepositMoney(accountNum int, deposit int) int {
	acc := allAccts[accountNum]
	acc.Balance += deposit
	return acc.Balance
}

func WithdrawMoney(accountNum int, withdraw int) int {
	acc := allAccts[accountNum]
	acc.Balance -= withdraw
	return acc.Balance
}

func TransferMoney(srcAccountNum int, tgtAccountNum int, amount int) (int, int) {
	//srcAccount := allAccts[srcAccountNum]
	//tgtAccount := allAccts[tgtAccountNum]
	tgtBalance := DepositMoney(tgtAccountNum, amount)
	srcBalance := WithdrawMoney(srcAccountNum, amount)
	return tgtBalance, srcBalance
}

func ListAccounts() map[int]Account {
	return allAccts
}
