package domain

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func NewContact(name, phone string) Contact {
	return Contact{
		Name:  name,
		Phone: phone,
	}
}
