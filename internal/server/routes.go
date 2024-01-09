package server

import (
	"net/http"
	"vg-sage/internal/controllers"
	"vg-sage/internal/models"

	"github.com/gorilla/mux"
)

func routerConfig(router *mux.Router) {
	//TODO Find better way to pass types, no make, instead interface{}

	router.HandleFunc("/api/articulos", findArticles()).Methods("GET")

	router.HandleFunc("/api/articulo/{id}", findArticle()).Methods("GET")

	router.HandleFunc("/api/pedidos_proveedor", findPedidosProveedor()).Methods("GET")

	router.HandleFunc("/api/lineas_albaran_cliente", findLineasAlbaranCliente()).Methods("GET")

	router.HandleFunc("/api/lineas_pedidos_proveedor", findLineasPedidosProveedor()).Methods("GET")

	router.HandleFunc("/api/lineas_pedidos_proveedor/{id}", patchLineasPedidosProveedor()).Methods("PATCH")

	router.HandleFunc("/api/cabeceras_pedidos_cliente", findCabeceraPedidoCliente()).Methods("GET")

	router.HandleFunc("/api/cabeceras_pedidos_cliente/{id}", patchCabeceraPedidoCliente()).Methods("PATCH")

	router.HandleFunc("/api/lineas_pedidos_cliente", findLineasPedidosCliente()).Methods("GET")

	router.HandleFunc("/api/lineas_pedidos_cliente/{id}", patchLineasPedidosCliente()).Methods("PATCH")

	router.Use(authMiddleware)

}

func findArticles() func(w http.ResponseWriter, r *http.Request) {
	articles := make([]models.Articulo, 0)

	return controllers.FindAll(&articles)
}

func findArticle() func(w http.ResponseWriter, r *http.Request) {

	return controllers.FindByArticleCode(&models.Articulo{})
}

func findPedidosProveedor() func(w http.ResponseWriter, r *http.Request) {
	pedidosProveedor := make([]models.PedidoProveedor, 0)

	return controllers.FindAll(&pedidosProveedor)
}

func findLineasAlbaranCliente() func(w http.ResponseWriter, r *http.Request) {
	lineasAlbaranCliente := make([]models.LineaAlbaranCliente, 0)

	return controllers.FindAll(&lineasAlbaranCliente)
}

func findLineasPedidosProveedor() func(w http.ResponseWriter, r *http.Request) {
	lineaPedidoProveedor := make([]models.LineaPedidoProveedor, 0)

	return controllers.FindAll(&lineaPedidoProveedor)
}

func patchLineasPedidosProveedor() func(w http.ResponseWriter, r *http.Request) {
	lineaPedidoProveedor := make([]models.LineaPedidoProveedor, 0)

	return controllers.Patch(&lineaPedidoProveedor, models.LineaPedidoProveedorPatch{})
}

func findCabeceraPedidoCliente() func(w http.ResponseWriter, r *http.Request) {
	cabeceraPedidoCliente := make([]models.CabeceraPedidoCliente, 0)

	return controllers.FindAll(&cabeceraPedidoCliente)
}

func patchCabeceraPedidoCliente() func(w http.ResponseWriter, r *http.Request) {
	cabeceraPedidoCliente := make([]models.CabeceraPedidoCliente, 0)

	return controllers.Patch(&cabeceraPedidoCliente, models.CabeceraPedidoClientePatch{})
}

func findLineasPedidosCliente() func(w http.ResponseWriter, r *http.Request) {
	lineasPedidosCliente := make([]models.LineaPedidoCliente, 0)

	return controllers.FindAll(&lineasPedidosCliente)
}

func patchLineasPedidosCliente() func(w http.ResponseWriter, r *http.Request) {
	lineaPedidoCliente := make([]models.LineaPedidoCliente, 0)

	return controllers.Patch(&lineaPedidoCliente, models.LineaPedidoClientePatch{})
}
