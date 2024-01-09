package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	uuid "github.com/satori/go.uuid"
)

type CabeceraPedidoCliente struct {
	CodigoEmpresa        uint       `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"`
	EjercicioPedido      uint       `json:"EjercicioPedido" gorm:"column:EjercicioPedido"`
	SeriePedido          string     `json:"SeriePedido" gorm:"column:SeriePedido"`
	NumeroPedido         uint       `json:"NumeroPedido" gorm:"column:NumeroPedido"`
	FechaPedido          *time.Time `json:"FechaPedido" gorm:"column:FechaPedido" noovo:"defaultorder;datelimit"`
	NumeroLineas         uint       `json:"NumeroLineas" gorm:"column:NumeroLineas"`
	CodigoCliente        uint       `json:"CodigoCliente" gorm:"column:CodigoCliente"`
	RazonSocial          string     `json:"RazonSocial" gorm:"column:RazonSocial"`
	ObservacionesCliente string     `json:"ObservacionesCliente" gorm:"column:ObservacionesCliente"`
	ObservacionesPedido  string     `json:"ObservacionesPedido" gorm:"column:ObservacionesPedido"`
	FechaNecesaria       *time.Time `json:"FechaNecesaria" gorm:"column:IS_FechaPlanificada"`
	FechaEntrega         *time.Time `json:"FechaEntrega" gorm:"column:FechaEntrega"`
	FechaTope            *time.Time `json:"FechaTope" gorm:"column:FechaTope"`
	Estado               uint       `json:"Estado" gorm:"column:Estado"`

	IdPedidoCli uuid.UUID `json:"IdPedidoCli" gorm:"column:IdPedidoCli" noovo:"primaryKey"`
}

type CabeceraPedidoClientePatch struct {
	FechaTope *time.Time `json:"FechaTope" gorm:"column:FechaTope"`
}

func (CabeceraPedidoCliente) TableName() string {
	return "dbo.CabeceraPedidoCliente"
}
