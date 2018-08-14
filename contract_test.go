package contract

import (
	"fmt"
	"os/exec"
)

type TxType int

const (
	SEND = 1 + iota
	DEPLOY
	INVOKE
	ISSUE
)

func ContractPrint() int {
	return 1
}

func Example_ContractPrint() {
	fmt.Println(ContractPrint())
	// Output:
}

type Transaction struct {
	txType       TxType
	senderAddr   string
	receiverAddr string
	amount       int64
	data         []interface{}
}

func NewTransaction(txType TxType, senderAddr string, receiverAddr string, amount int64, data []interface{}) *Transaction {
	tx := &Transaction{
		txType:       txType,
		senderAddr:   senderAddr,
		receiverAddr: receiverAddr,
		amount:       amount,
		data:         data,
	}
	return tx
}

type Digitalsignature struct {
}

type Account struct {
	balance     int64
	contract    Contract
	accountName string
	skx         string
	address     string
}

func NewAccount(name string) *Account {
	addr, _ := exec.Command("uuidgen").Output()
	a := &Account{
		balance:     0,
		accountName: name,
		address:     string(addr[:]),
	}
	return a
}

type Blockchain struct {
	globalLedger       map[string]*Account
	genesisAccountSkx  string
	genesisAccountAddr string
	blockArray         []Transaction
}

func NewBlockchain(genesisAmount int) *Blockchain {
	account := NewAccount("genesis")
	b := &Blockchain{
		globalLedger:       map[string]*Account{},
		genesisAccountAddr: account.address,
		genesisAccountSkx:  account.skx,
		blockArray:         []Transaction{},
	}

	return b
}

func (b *Blockchain) GetNewAccount(name string) (string, string) {
	var account = NewAccount(name)
	b.globalLedger[account.address] = account
	return account.skx, account.address
}

func (b *Blockchain) Send(tx Transaction) {
	if tx.txType == SEND {
		b.globalLedger[tx.senderAddr].balance -= tx.amount
		b.globalLedger[tx.receiverAddr].balance += tx.amount
		b.blockArray = append(b.blockArray, tx)
	}
}

func (b *Blockchain) invoke(tx Transaction) {
	if tx.txType == SEND {
	}
}

func (b *Blockchain) issue(contractAddr string) {
	b.globalLedger[contractAddr].balance += b.globalLedger[contractAddr].contract.issueCoin
}

func (b *Blockchain) newAccount(name string) (string, string) {
	account := NewAccount(name)
	b.globalLedger[account.address] = account
	return account.skx, account.address
}

func (b *Blockchain) GetBlockHeight() int {
	return len(b.blockArray)
}

func (b *Blockchain) GetBalance(accountAddr string) string {
	return ""
}

type Client struct {
	accountAddr string
	accountName string
	accountSkx  string
}

func NewClient(name string) *Client {

	c := &Client{
		accountAddr: "",
		accountName: name,
		accountSkx:  "",
	}
	return c
}

func (c *Client) GetBalance(b *Blockchain) string {
	return b.GetBalance(c.accountAddr)
}

func (c Client) sendTx(receiverAddr string, amount uint64, b *Blockchain) {
	if c.accountAddr != "" {
		var tx Transaction
		tx = NewTransaction(SEND, c.accountAddr, receiverAddr, amount, "")
		b.send(tx)
	}

}

func (c Client) invokeTx(b *Blockchain, contractAddr string, _ uint64, method func(), args []interface{}) *Client {
	if c.accountAddr != "" {
		d := [...]interface{}{method, args}
		tx := Transaction(INVOKE, c.accountAddr, contractAddr, d)
		b.invoke(tx)
	}
}

type Contract interface {
	transfer()
}

type P_Contract struct {
	own_addr     string
	sum          uint64
	account_addr string
}
