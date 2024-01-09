package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	uuid "github.com/satori/go.uuid"
)

type LineaAlbaranCliente struct {
	LineasPosicion   uuid.UUID `json:"LineasPosicion" gorm:"type:uniqueidentifier;column:LineasPosicion"`
	EjercicioAlbaran uint      `json:"EjercicioAlbaran" gorm:"column:EjercicioAlbaran"`
	SerieAlbaran     string    `json:"SerieAlbaran" gorm:"column:SerieAlbaran"`
	RazonSocial      string    `json:"RazonSocial" gorm:"column:RazonSocial"`

	NumeroAlbaran       uint    `json:"NumeroAlbaran" gorm:"column:NumeroAlbaran"`
	CodigoArticulo      string  `json:"CodigoArticulo" gorm:"column:CodigoArticulo"`
	CodigoEmpresa       uint    `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"`
	DescripcionArticulo string  `json:"DescripcionArticulo" gorm:"column:DescripcionArticulo"`
	Partida             string  `json:"Partida" gorm:"column:Partida"`
	UnidadesServidas    float32 `json:"UnidadesServidas" gorm:"column:UnidadesServidas"`
	TipoBolsa           string  `json:"TipoBolsa" gorm:"column:Tipo_bolsa"`
	Precio              float32 `json:"Precio" gorm:"column:Precio"`

	EjercicioPedido uint       `json:"EjercicioPedido" gorm:"column:EjercicioPedido"`
	SeriePedido     string     `json:"SeriePedido" gorm:"column:SeriePedido"`
	NumeroPedido    uint       `json:"NumeroPedido" gorm:"column:NumeroPedido"`
	FechaAlbaran    *time.Time `json:"FechaAlbaran" gorm:"column:FechaAlbaran" noovo:"defaultorder;datelimit"`
}

func (LineaAlbaranCliente) TableName() string {
	return "dbo.LineasAlbaranCliente"
}

func (a LineaAlbaranCliente) Join(query *gorm.DB) *gorm.DB {
	join := "JOIN CabeceraAlbaranCliente ON (LineasAlbaranCliente.NumeroAlbaran = CabeceraAlbaranCliente.NumeroAlbaran  AND LineasAlbaranCliente.EjercicioAlbaran =CabeceraAlbaranCliente.EjercicioAlbaran AND LineasAlbaranCliente.CodigoEmpresa = CabeceraAlbaranCliente.CodigoEmpresa  AND LineasAlbaranCliente.SerieAlbaran = CabeceraAlbaranCliente.SerieAlbaran)"
	return db.Select("LineasAlbaranCliente.*, CabeceraAlbaranCliente.RazonSocial").Joins(join)
}
