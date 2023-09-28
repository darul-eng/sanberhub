package helpers

import (
	"sanberhub/models/api"
	"sanberhub/models/domain"
)

func ToStatementResponse(transaction domain.Transaction) api.StatementResponse {
	return api.StatementResponse{
		CreatedAt:       transaction.CreatedAt,
		TransactionCode: transaction.TransactionCode,
		Amount:          transaction.Amount,
	}
}

func ToStatementResponses(transactions []domain.Transaction) []api.StatementResponse {
	var transactionsResponses []api.StatementResponse
	for _, transaction := range transactions {
		transactionsResponses = append(transactionsResponses, ToStatementResponse(transaction))
	}

	return transactionsResponses
}
