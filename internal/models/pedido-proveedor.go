package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	uuid "github.com/satori/go.uuid"
)

type PedidoProveedor struct {
	EjercicioPedido string     `json:"EjercicioPedido" gorm:"column:EjercicioPedido"`
	SeriePedido     string     `json:"SeriePedido" gorm:"column:SeriePedido"`
	NumeroPedido    uint       `json:"NumeroPedido" gorm:"column:NumeroPedido"`
	Proveedor       string     `json:"Proveedor" gorm:"column:Nombre"`
	FechaNecesaria  *time.Time `json:"FechaNecesaria" gorm:"column:FechaNecesaria" noovo:"random;datelimit"`
	FechaPedido     *time.Time `json:"FechaPedido" gorm:"column:FechaPedido"`
	FechaTope       *time.Time `json:"FechaTope" gorm:"column:FechaTope" noovo:"defaultorder"`
	NumeroLineas    uint       `json:"NumeroLineas" gorm:"column:NumeroLineas"`
	CodigoProveedor string     `json:"CodigoProveedor" gorm:"column:CodigoProveedor"`
	IdPedidoPro     uuid.UUID  `json:"IdPedidoPro" gorm:"column:IdPedidoPro"`
 	Estado			uint	   `json:"Estado" gorm:"column:Estado"`  
	CodigoEmpresa   uint       `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"` 
}

func (a PedidoProveedor) TableName() string {
	return "dbo.CabeceraPedidoProveedor"
}
