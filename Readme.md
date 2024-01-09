# VG-Sage  ![](https://github.com/Binoovo/vg-sage/workflows/build.yml/badge.svg) 
[Docker File](https://github.com/Binoovo/vg-sage/packages)

Project that create api endpoint for the vg database

## Get started
- Create a new file in root called .env
- Modify the .env with your environment variables ( [.env.example](.env.example) is a template )
- With make
	- Run make start
- Without make 
	- Run ```go build cmd/vg-sage/main.go```
	- Execute the output

## How works
Make request to the proper endpoint '/api/{endpoint}'.

Supported endpoints:
- pedidos_proveedor
- articulos
- linea_albaran_cliente

Filter structure
```
type Filter struct {
	Where  map[string]interface{} `json:"where,omitempty"`
	Limit  int                    `json:"limit,omitempty"`
	Offset int                    `json:"offset,omitempty"`
	Order  string                 `json:"order,omitempty"`
}
```

Request support filter json query ex:
``` 
url: http://127.0.0.1:8000/api/wtf?filter=%22%7B'Where'%3A%7B'CodigoArticulo'%3A'10R1014W1206'%7D%7D%22
query : "{'Where':{'CodigoArticulo':'10R1014W1206'}}"
```

## Project Structure
- cmd: Main and exes folder
- internal
    - controllers: definition of main controller
    - models: definition of the models that relate to the database
    - server: definicion of the routes


## Important files

### controller.go
Generic controller that create an endpoint with support of filter ex:
```
articles := make([]models.Articulo, 0)
router.HandleFunc("/api/articulos", controllers.FindAll(&articles)).Methods("GET")
```

### model.go
Generic function definition that support filter to search in the database

### E.g. model definition
```
type Articulo struct {
	CodigoArticulo      string    `json:"CodigoArticulo" gorm:"column:CodigoArticulo"`
	CodigoEmpresa       uint      `json:"CodigoEmpresa"  gorm:"column:CodigoEmpresa"`
	DescripcionArticulo string    `json:"DescripcionArticulo"  gorm:"column:DescripcionArticulo"`
	FechaAlta           time.Time `gorm:"column:FechaAlta" noovo:"defaultorder"`
	IdArticulo          string    `gorm:"column:IdArticulo"`
}

func (a Articulo) TableName() string {
	return "dbo.Articulos"
}

```
