package models

import (
	"testing"
	"vg-sage/internal/models"
)

func TestFindPedidosProveedor(t *testing.T) {
	pedidos := models.FindPedidosProveedor()
	if pedidos == nil {
		t.Error("No pedidos found")
	}
	for _, item := range pedidos {
		if item.NumeroPedido == 0 || item.Proveedor == "" {
			t.Error("No valid articles")
		}
	}
}

func TestFindPedidoProveedorById(t *testing.T) {
	pedido := models.FindPedidoProveedorById("1")
	if pedido == nil {
		t.Error("No pedido found")
	}
}
