package model

// propiedades de un cliente
// Note además las anotaciones encerradas e m ` esto ayudará a realizar automaticamente la conversion a json
// o viceversa
type Customer struct {
	ID   string `json:"ID"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
