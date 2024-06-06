package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wellmtx/session-finance/model/transaction"
	"github.com/wellmtx/session-finance/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	transactions := transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    1000,
			Type:      1,
			CreatedAt: util.StringToTime("2024-06-06 01:44:25"),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

type CreateTransactionDTO struct {
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Type   int     `json:"type"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var dto CreateTransactionDTO

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	transaction := transaction.Transaction{
		Title:     dto.Title,
		Amount:    dto.Amount,
		Type:      dto.Type,
		CreatedAt: util.StringToTime("2024-06-06T01:44:25"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}
