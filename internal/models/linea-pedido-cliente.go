package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	uuid "github.com/satori/go.uuid"
)

type LineaPedidoCliente struct {
	LineasPosicion      uuid.UUID  `json:"LineasPosicion" gorm:"type:uniqueidentifier;column:LineasPosicion" noovo:"primaryKey"`
	RazonSocial         string     `json:"RazonSocial" gorm:"column:RazonSocial"`
	NumeroAlbaran       uint       `json:"NumeroAlbaran" gorm:"column:NumeroAlbaran"`
	CodigoArticulo      string     `json:"CodigoArticulo" gorm:"column:CodigoArticulo"`
	CodigoEmpresa       uint       `json:"CodigoEmpresa" gorm:"column:CodigoEmpresa"`
	DescripcionArticulo string     `json:"DescripcionArticulo" gorm:"column:DescripcionArticulo"`
	Partida             string     `json:"Partida" gorm:"column:Partida"`
	UnidadesServidas    float32    `json:"UnidadesServidas" gorm:"column:UnidadesServidas"`
	UnidadesPedidas     float32    `json:"UnidadesPedidas" gorm:"column:UnidadesPedidas"`
	TipoBolsa           string     `json:"TipoBolsa" gorm:"column:Tipo_bolsa"`
	Precio              float32    `json:"Precio" gorm:"column:Precio"`
	CantidadDeCubos     uint       `json:"CantidadDeCubos" gorm:"column:Cant_Cubos"`
	UnidadesPorCubo     uint       `json:"UnidadesPorCubo" gorm:"column:Unid_Cubos"`
	CantidadDeBolsas    uint       `json:"CantidadDeBolsas" gorm:"column:Cant_Bolsas"`
	UnidadesPorBolsa    uint       `json:"UnidadesPorBolsa" gorm:"column:Unid_Bolsas"`
	NumeroLote          string     `json:"lote" gorm:"column:NUMERO_lote"`
	EjercicioPedido     uint       `json:"EjercicioPedido" gorm:"column:EjercicioPedido"`
	SeriePedido         string     `json:"SeriePedido" gorm:"column:SeriePedido"`
	NumeroPedido        uint       `json:"NumeroPedido" gorm:"column:NumeroPedido"`
	FechaPedido         *time.Time `json:"FechaPedido" gorm:"column:FechaPedido" noovo:"defaultorder;datelimit"`
	FechaNecesaria      *time.Time `json:"FechaNecesaria" gorm:"column:FechaNecesaria"`
	FechaEntrega        *time.Time `json:"FechaEntrega" gorm:"column:FechaEntrega"`
	FechaTope           *time.Time `json:"FechaTope" gorm:"column:FechaTope"`
	Estado              float32    `json:"Estado" gorm:"column:Estado"`
	ISPresentacion      string     `json:"ISPresentacion" gorm:"column:IS_Presentacion"`
	ISCantPresentacion  float32    `json:"ISCantPresentacion" gorm:"column:IS_CantPresentacion"`
	ISTipoEnvase        string     `json:"ISTipoEnvase" gorm:"column:IS_TipoEnvase"`
	ISCantEnvase        float32    `json:"ISCantEnvase" gorm:"column:IS_CantEnvase"`
	Unidxenv            float32    `json:"Unidxenv" gorm:"column:unidxenv"`
	ISEtiqueta          float32    `json:"ISEtiqueta" gorm:"column:IS_Etiqueta"`
	ISMarca             float32    `json:"ISMarca" gorm:"column:IS_Marca"`
	ISProceso           string     `json:"ISProceso" gorm:"column:IS_Proceso"`
	ISComentario        string     `json:"ISComentario" gorm:"column:IS_Comentario"`

	UnidadesPendientesFabricar float32 `json:"UnidadesPendientesFabricar" gorm:"column:UnidadesPendientesFabricar"`
	UnidadesPendientes         float32 `json:"UnidadesPendientes" gorm:"column:UnidadesPendientes"`
	UnidadesPendAnterior       float32 `json:"UnidadesPendAnterior" gorm:"column:UnidadesPendAnterior"`
	UnidadesaServir            float32 `json:"UnidadesaServir" gorm:"column:UnidadesaServir"`
	Unidades2_                 float32 `json:"Unidades2_" gorm:"column:Unidades2_"`
	Unidades2aServir_          float32 `json:"Unidades2aServir_" gorm:"column:Unidades2aServir_"`
}

type LineaPedidoClientePatch struct {
	UnidadesPendientesFabricar float32 `json:"UnidadesPendientesFabricar" gorm:"column:UnidadesPendientesFabricar"`
	UnidadesPedidas            float32 `json:"UnidadesPedidas" gorm:"column:UnidadesPedidas"`
	UnidadesPendientes         float32 `json:"UnidadesPendientes" gorm:"column:UnidadesPendientes"`
	UnidadesPendAnterior       float32 `json:"UnidadesPendAnterior" gorm:"column:UnidadesPendAnterior"`
	UnidadesaServir            float32 `json:"UnidadesaServir" gorm:"column:UnidadesaServir"`
	Unidades2_                 float32 `json:"Unidades2_" gorm:"column:Unidades2_"`
	Unidades2aServir_          float32 `json:"Unidades2aServir_" gorm:"column:Unidades2aServir_"`
}

func (LineaPedidoCliente) TableName() string {
	return "dbo.LineasPedidoCliente"
}

func (a LineaPedidoCliente) Join(query *gorm.DB) *gorm.DB {
	join := `LEFT JOIN
                     ( SELECT NumeroPedido, EjercicioPedido, RazonSocial, CodigoEmpresa, SeriePedido
                       FROM CabeceraPedidoCliente ) cpc
                  ON (     LineasPedidoCliente.NumeroPedido = cpc.NumeroPedido
                       AND LineasPedidoCliente.EjercicioPedido = cpc.EjercicioPedido
                       AND LineasPedidoCliente.CodigoEmpresa = cpc.CodigoEmpresa
                       AND LineasPedidoCliente.SeriePedido = cpc.SeriePedido )`

	return query.Select("LineasPedidoCliente.*, cpc.RazonSocial").Joins(join)
}
