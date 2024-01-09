package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Articulo struct {
	CodigoArticulo      string     `json:"CodigoArticulo" gorm:"column:CodigoArticulo"`
	CodigoEmpresa       uint       `json:"CodigoEmpresa"  gorm:"column:CodigoEmpresa"`
	DescripcionArticulo string     `json:"DescripcionArticulo"  gorm:"column:DescripcionArticulo"`
	FechaAlta           *time.Time `gorm:"column:FechaAlta" noovo:"defaultorder"`
	IdArticulo          string     `gorm:"column:IdArticulo"`
	CodigoSubFamilia2   string     `json:"CodigoSubFamilia2"  gorm:"column:CodigoSubFamilia2"`
	Cod_Metraje         string     `json:"Cod_Metraje"  gorm:"column:Cod_Metraje"`
	Cod_Calibre         string     `json:"Cod_Calibre"  gorm:"column:Cod_Calibre"`
	Cod_EspProd         uint       `json:"Cod_EspProd"  gorm:"column:Cod_EspProd"`
	CC_CodPaisProd      string     `json:"CC_CodPaisProd"  gorm:"column:CC_CodPaisProd"`
	CC_CodOrdenProd     string     `json:"CC_CodOrdenProd"  gorm:"column:CC_CodOrdenProd"`
}

func (a Articulo) TableName() string {
	return "dbo.Articulos"
}
