package model

import (
	"time"
)

//API ...
type API struct {
	APIName            string `json: "apiTeste"`
	APIDescription     string `json: "apiDescription"`
	APIServer          string `json: "apiServer"`
	APINumeroTentativa int    `json: "apiNumeroTentativa"`
	APIUser            string `json: "apiUser"`
	APIPassword        string `json: "apiPassword"`
	APIToken           string `json: "apiToken"`
}

//DocJSON ...
type DocJSON map[string]interface{}

//ModelDocumento ...
type ModelDocumento struct {
	DataCriacao time.Time   `json:"dataCriacao" mapstructure:"dataCriacao" form:"dataCriacao"`
	DataUpdate  time.Time   `json:"dataUpdate" mapstructure:"dataUpdate" form:"dataUpdate"`
	Documento   interface{} `json:"documento" mapstructure:"documento" form:"documento"`
	Deletado    bool        `json:"deletado" mapstructure:"deletado" form:"deletado"`
}

//ModelSearchDatabase ...
type ModelSearchDatabase struct {
	Collection string                     `json:"collection" mapstructure:"collection" form:"collection"`
	Where      []ModelSearchDatabaseWhere `json:"where" mapstructure:"where" form:"where"`
}

//ModelSearchDatabaseWhere ...
type ModelSearchDatabaseWhere struct {
	Var1     string      `json:"var1" mapstructure:"var1" form:"var1"`
	Operador string      `json:"operador" mapstructure:"operador" form:"operador"`
	Var2     interface{} `json:"var2" mapstructure:"var2" form:"var2"`
}

//ReturnUniqueData ...
type UniqueData struct {
	ID         string `json:"id" mapstructure:"id" form:"id"`
	Collection string `json:"collection" mapstructure:"collection" form:"collection"`
}
