package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	cats, err := GetCategories("MLA")
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	viewCategories(cats)
}
func GetCategories(siteID string) (categories, error) {
	response, err := http.Get("https://api.mercadolibre.com/sites/MLA/categories") //completar
	vacio := categories{}
	if err != nil {
		err := errors.New("404 Not found")

		return vacio, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err := errors.New("404 Not found")

		return vacio, err
	}
	var cats categories
	err = json.Unmarshal(bytes, &cats)
	if err != nil {
		panic(err)
	}
	return cats, nil
}

type categories []struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func viewCategories(cats categories) {
	i := 0
	for i < len(cats) {
		fmt.Println("\n Id:", cats[i].Id)
		fmt.Println("Name:", cats[i].Name)

		i++
	}
}

//invocacion correspondiente mediante http
//get    resp:= http.Get(url)
//Post   resp:=http.Post(url,body)

//leer los datos
//ioutil.ReadAll(response.Body)

// una vez obtenido los bytes
// var myVar myStruct

/* ejemplo json
{
	"nombre": "Emiliano",
	"edad": 27
}
en go seria
type Persona struct{
	Nombre string
	Edad int
}

debemos tener en cuenta las anotaciones especificadas a la derecha de cada uno de los atributos de nuestra estructura e invocar
el metodo del paquete de encoding/json

y usar la funcion err := json.Unmarshal(bytes, &myVar

utilizar el appi   https://api.mercadolibre.com/sites/MLA/search?q=Motorola)
documentacion PKG.Go.dev/net/http

{"site_id":"MLA","country_default_time_zone":"GMT-03:00","query":"Motorola)",

configuraicon apr
primero en git se debe mapear
*/
