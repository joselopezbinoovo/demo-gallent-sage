package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	uuid "github.com/satori/go.uuid"
)

type LineaPedidoProveedor struct {
	CodigoEmpresa       string     `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"`
	RazonSocial         string     `json:"RazonSocial" gorm:"column:RazonSocial"`
	EjercicioPedido     string     `json:"EjercicioPedido" gorm:"column:EjercicioPedido"`
	SeriePedido         string     `json:"SeriePedido" gorm:"column:SeriePedido"`
	NumeroPedido        uint       `json:"NumeroPedido" gorm:"column:NumeroPedido"`
	Orden               string     `json:"Orden" gorm:"column:Orden"`
	FechaRegistro       *time.Time `json:"FechaRegistro" gorm:"column:FechaRegistro" noovo:"datelimit"`
	FechaPedido         *time.Time `json:"FechaPedido" gorm:"column:FechaPedido"`
	CodigoArticulo      string     `json:"CodigoArticulo" gorm:"column:CodigoArticulo"`
	DescripcionArticulo string     `json:"DescripcionArticulo" gorm:"column:DescripcionArticulo"`
	UnidadesPedidas     float32    `json:"UnidadesPedidas" gorm:"column:UnidadesPedidas"`
	CodigoAlmacen       string     `json:"CodigoAlmacen" gorm:"column:CodigoAlmacen"`
	Estado              uint       `json:"Estado" gorm:"column:Estado"`
	CodigoProveedor     string     `json:"CodigoProveedor" gorm:"column:CodigoProveedor"`
	Partida             string     `json:"Partida" gorm:"column:Partida"`

	LineasPosicion uuid.UUID `json:"LineasPosicion" gorm:"column:LineasPosicion" noovo:"primaryKey"`
	LineaPedidoCli uuid.UUID `json:"LineaPedidoCli" gorm:"column:LineaPedidoCli"`
}

// una estructura diferente para poder limitar que campos se escriben en cliente
type LineaPedidoProveedorPatch struct {
	Partida string `json:"Partida,omitempty" gorm:"column:Partida"`
}

func (a LineaPedidoProveedor) TableName() string {
	return "dbo.LineasPedidoProveedor"
}

func (a LineaPedidoProveedor) Join(query *gorm.DB) *gorm.DB {
	join := `LEFT JOIN
                 ( SELECT NumeroPedido, EjercicioPedido, CodigoEmpresa, SeriePedido, CodigoProveedor, RazonSocial
                   FROM CabeceraPedidoProveedor) cpp
             ON (    LineasPedidoProveedor.NumeroPedido = cpp.NumeroPedido
                 AND LineasPedidoProveedor.EjercicioPedido = cpp.EjercicioPedido
                 AND LineasPedidoProveedor.CodigoEmpresa = cpp.CodigoEmpresa
                 AND LineasPedidoProveedor.SeriePedido = cpp.SeriePedido )`

    return db.Select("LineasPedidoProveedor.*, cpp.CodigoProveedor, cpp.RazonSocial").Joins(join)
}
