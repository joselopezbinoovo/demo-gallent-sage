package models

import (
	"testing"
	"vg-sage/internal/models"
)

func TestFindArticulos(t *testing.T) {
	articles := models.FindArticulos()
	if articles == nil {
		t.Error("No articles found")
	}
	for _, item := range articles {
		if item.CodigoArticulo == "" || item.DescripcionArticulo == "" {
			t.Error("No valid articles")
		}
	}
}

func TestFindArticulosByCode(t *testing.T) {
	articles := models.FindArticuloByCodigo("10R1014W1206")
	if articles == nil {
		t.Error("No articles found")
	}
	for _, item := range articles {
		if item.CodigoArticulo == "" || item.DescripcionArticulo == "" {
			t.Error("No valid articles")
		}
	}
}
