package models

import (
	"testing"
	"vg-sage/internal/models"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestFind(t *testing.T) {
	t.Log("Given a model to search")

	articles := make([]models.Articulo, 0)
	filter := models.Filter{Where: map[string]interface{}{"CodigoArticulo": "10R1014W1206"}}
	err := models.Find(filter, &articles)

	if err != nil || len(articles) == 0 {
		t.Error("\tShould load articles passing model", err, ballotX)
	} else {
		if len(articles) > 0 {
			t.Log("\tShould load mode passing where", checkMark)
			if filter.Where["CodigoArticulo"] == articles[0].CodigoArticulo {
				t.Log("\t\tShould make filter where", checkMark)
			} else {
				t.Error("\t\tShould make filter where", ballotX)
			}
		}
	}
}

func TestFindWithLimit(t *testing.T) {
	t.Log("Given a model to search")

	articles := make([]models.Articulo, 0)
	filter := models.Filter{Limit: 2, Order: "FechaAlta"}
	err := models.Find(filter, &articles)
	if err != nil || len(articles) != 2 {
		t.Error("\tShould load articles passing limit", err, ballotX)
	} else {
		t.Log("\tShould load mode passing limit", checkMark)
	}
}

func TestFindWithLimitOffset(t *testing.T) {
	t.Log("Given a model to search with offset")

	articles := make([]models.Articulo, 0)
	filter := models.Filter{Limit: 2, Order: "FechaAlta", Offset: 2}
	err := models.Find(filter, &articles)
	if err != nil || len(articles) != 2 {
		t.Error("\tShould load articles passing limit", err, ballotX)
	} else {
		t.Log("\tShould load mode passing limit", checkMark)
	}
}

func TestFindOne(t *testing.T) {
	t.Log("Given a model to search one")

	article := &models.Articulo{}
	filter := models.Filter{}
	err := models.Find(filter, &article)
	if err != nil || article == nil {
		t.Error("\tShould load one model", err, ballotX)
	} else {
		t.Log("\tShould load one model", checkMark)
	}
}
