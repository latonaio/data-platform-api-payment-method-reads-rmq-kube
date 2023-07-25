package dpfm_api_output_formatter

import (
	"data-platform-api-payment-method-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToPaymentMethod(rows *sql.Rows) (*[]PaymentMethod, error) {
	defer rows.Close()
	paymentMethod := make([]PaymentMethod, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentMethod{}

		err := rows.Scan(
			&pm.PaymentMethod,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentMethod, nil
		}

		data := pm
		paymentMethod = append(paymentMethod, PaymentMethod{
			PaymentMethod:			data.PaymentMethod,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &paymentMethod, nil
}

func ConvertToPaymentMethodText(rows *sql.Rows) (*[]PaymentMethodText, error) {
	defer rows.Close()
	paymentMethodText := make([]PaymentMethodText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentMethodText{}

		err := rows.Scan(
			&pm.PaymentMethod,
			&pm.Language,
			&pm.PaymentMethodName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentMethodText, err
		}

		data := pm
		paymentMethodText = append(paymentMethodText, PaymentMethodText{
			PaymentMethod:     		data.PaymentMethod,
			Language:          		data.Language,
			PaymentMethodName:		data.PaymentMethodName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &paymentMethodText, nil
}
