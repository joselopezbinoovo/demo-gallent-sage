package models

import (
	"testing"
	"vg-sage/internal/models"
)

func TestFindLineaAlbaranCliente(t *testing.T) {
	lineas := models.FindLineaAlbaranCliente()
	if lineas == nil {
		t.Error("No lineas found")
	}
	for _, item := range lineas {
		if item.DescripcionArticulo == "" || item.CodigoArticulo == "" {
			t.Error("No valid lineas")
		}
	}
}

func TestFindLineaAlbaranClienteByClienteEjercicio(t *testing.T) {
	lineas := models.FindLineaAlbaranClienteByClienteEjercicio("12073", "0")
	if lineas == nil {
		t.Error("No lineas found")
	}
	for _, item := range lineas {
		if item.DescripcionArticulo == "" || item.CodigoArticulo == "" {
			t.Error("No valid lineas")
		}
	}
}
