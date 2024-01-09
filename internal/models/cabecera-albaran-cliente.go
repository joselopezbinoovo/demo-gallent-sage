package models

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type CabeceraAlbaranCliente struct {
	EjercicioAlbaran uint   `json:"EjercicioAlbaran" gorm:"column:EjercicioAlbaran"`
	SerieAlbaran     string `json:"SerieAlbaran" gorm:"column:SerieAlbaran"`
	NumeroAlbaran    uint   `json:"NumeroAlbaran" gorm:"column:NumeroAlbaran"`
	RazonSocial      string `json:"RazonSocial" gorm:"column:RazonSocial"`
	CodigoEmpresa    uint   `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"`
}

func (CabeceraAlbaranCliente) TableName() string {
	return "dbo.CabeceraAlbaranCliente"
}
