package model

import (
	"time"
)

//Especialidade ...
type Especialidade struct {
	Especialidade string
}

//Medico ...
type Medico struct {
	CRM            int             `json:"CRM" mapstructure:"CRM" form:"CRM"`
	Usuario        Usuario         `json:"Usuario" mapstructure:"Usuario" form:"Usuario"`
	Especialidades []Especialidade `json:"Especialidade" mapstructure:"Especialidade" form:"Especialidade"`
}

//Usuario ...
type Usuario struct {
	Nome           string    `json:"Nome" mapstructure:"Nome" form:"Nome"`
	CPF            string    `json:"CPF" mapstructure:"CPF" form:"CPF"`
	DataNascimento time.Time `json:"DataNascimento" mapstructure:"DataNascimento" form:"DataNascimento"`
	Email          string    `json:"Email" mapstructure:"Email" form:"Email"`
	Telefone       string    `json:"Telefone" mapstructure:"Telefone" form:"Telefone"`
	Endereco       string    `json:"Endereco" mapstructure:"Endereco" form:"Endereco"`
	CEP            int       `json:"CEP" mapstructure:"CEP" form:"CEP"`
}

//UBS ...
type UBS struct {
	UBS      string
	Endereco string
	CEP      int
	Telefone string
}

//Funcionario ...
type Funcionario struct {
	Cargo   string
	Usuario Usuario
}

//Agenda ...
type Agenda struct {
	HoraInicio time.Time
	HoraFim    time.Time
}

//Notificar ...
type Notificar struct {
	Recebida    string
	Visualizada string
	Resposta    string
}

//Agendamento ...
type Agendamento struct {
	Status          string
	Atendimento     time.Time
	Agendamento     time.Time
	QuantidadeTempo int
}

//Paciente ...
type Paciente struct {
	SIS          int64
	NomePaciente string
}
