package cards

import "L9/pkg/bank/types"

func Total(cards []types.Card) (payments []types.Payment) {
	payment := types.Payment{}
	for _, operation := range cards {
		if operation.Active && operation.Balance >= 0 {
			payment.Balance = operation.Balance
			payment.Type = operation.Name
			payment.Number = operation.PAN

			payments = append(payments, payment)
		}
	}

	return payments
}
