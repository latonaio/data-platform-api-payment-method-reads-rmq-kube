package requests

type PaymentMethodText struct {
	PaymentMethod     	string  `json:"PaymentMethod"`
	Language          	string  `json:"Language"`
	PaymentMethodName	string  `json:"PaymentMethodName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
