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
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentMethod, nil
		}

		data := pm
		paymentMethod = append(paymentMethod, PaymentMethod{
			PaymentMethod: data.PaymentMethod,
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
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentMethodText, err
		}

		data := pm
		paymentMethodText = append(paymentMethodText, PaymentMethodText{
			PaymentMethod:     data.PaymentMethod,
			Language:          data.Language,
			PaymentMethodName: data.PaymentMethodName,
		})
	}

	return &paymentMethodText, nil
}

func ConvertToPaymentMethodTexts(rows *sql.Rows) (*[]PaymentMethodText, error) {
	defer rows.Close()
	paymentMethodText := make([]PaymentMethodText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentMethodTexts{}

		err := rows.Scan(
			&pm.PaymentMethod,
			&pm.Language,
			&pm.PaymentMethodName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentMethodText, err
		}

		data := pm
		paymentMethodText = append(paymentMethodText, PaymentMethodText{
			PaymentMethod:     data.PaymentMethod,
			Language:          data.Language,
			PaymentMethodName: data.PaymentMethodName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &paymentMethodText, nil
	}

	return &paymentMethodText, nil
}
