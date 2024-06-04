package statemachine

import (
	"mpt/trie"
	"mpt/types"
)

type IMachine interface {
	Execute(state trie.ITrie, tx types.Transaction)
}

type StateMachine struct {
}

func (m StateMachine) Execute(state trie.ITrie, tx types.Transaction) {
	from := tx.From()
	to := tx.To
	value := tx.Value
	gasUsed := tx.Gas
	if gasUsed > 21000 {
		gasUsed = 21000
	}
	gasUsed = gasUsed * tx.GasPrice
	cost := value + gasUsed
	account, err := state.Load(from[:])
	if err != nil {
		return
	}
	if account.Amount < cost {
		return
	}
	account.Amount = account.Amount - cost
	state.Store(from[:], account)

	toAccount, err := state.Load(to[:])
	if err != nil {
		toAccount = types.Account{}
	}
	toAccount.Amount = toAccount.Amount + value
	state.Store(to[:], toAccount)
}
