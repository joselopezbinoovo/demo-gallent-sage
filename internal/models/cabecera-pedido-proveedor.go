package models

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type CabeceraPedidoProveedor struct {
	EjercicioPedido uint   `json:"EjercicioPedido" gorm:"column:EjercicioPedido"`
	SeriePedido     string `json:"SeriePedido" gorm:"column:SeriePedido"`
	NumeroPedido    uint   `json:"NumeroPedido" gorm:"column:NumeroPedido"`
	RazonSocial     string `json:"RazonSocial" gorm:"column:RazonSocial"`
	CodigoEmpresa   uint   `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"`
	CodigoProveedor string `json:"CodigoProveedor" gorm:"column:CodigoProveedor"`
}

func (CabeceraPedidoProveedor) TableName() string {
	return "dbo.CabeceraPedidoProveedor"
}
