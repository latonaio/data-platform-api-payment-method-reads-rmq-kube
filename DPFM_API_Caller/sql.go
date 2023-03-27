package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-payment-method-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-payment-method-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var paymentMethod *[]dpfm_api_output_formatter.PaymentMethod
	var paymentMethodText *[]dpfm_api_output_formatter.PaymentMethodText
	for _, fn := range accepter {
		switch fn {
		case "PaymentMethod":
			func() {
				paymentMethod = c.PaymentMethod(mtx, input, output, errs, log)
			}()
		case "PaymentMethodText":
			func() {
				paymentMethodText = c.PaymentMethodText(mtx, input, output, errs, log)
			}()
		case "PaymentMethodTexts":
			func() {
				paymentMethodText = c.PaymentMethodTexts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		PaymentMethod:     paymentMethod,
		PaymentMethodText: paymentMethodText,
	}

	return data
}

func (c *DPFMAPICaller) PaymentMethod(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentMethod {
	paymentMethod := input.PaymentMethod.PaymentMethod

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_method_payment_method_data
		WHERE PaymentMethod = ?;`, paymentMethod,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentMethod(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PaymentMethodText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentMethodText {
	var args []interface{}
	paymentMethod := input.PaymentMethod.PaymentMethod
	paymentMethodText := input.PaymentMethod.PaymentMethodText

	cnt := 0
	for _, v := range paymentMethodText {
		args = append(args, paymentMethod, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_method_payment_method_text_data
		WHERE (PaymentMethod, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentMethodText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PaymentMethodTexts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentMethodText {
	var args []interface{}
	paymentMethodText := input.PaymentMethod.PaymentMethodText

	cnt := 0
	for _, v := range paymentMethodText {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_method_payment_method_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToPaymentMethodTexts(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
