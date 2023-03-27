package requests

type PaymentMethodTexts struct {
	PaymentMethod     string  `json:"PaymentMethod"`
	Language          string  `json:"Language"`
	PaymentMethodName *string `json:"PaymentMethodName"`
}
