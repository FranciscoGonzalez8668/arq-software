package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"ioutil"
	"net/http"
)

func main(){
	var categorias Category
	for true{
	cats,err:=GetCategories("MLA")
	if err != nil{
		fmt.Fprintln("Error 404")
		break
	}
	categorias[i]=cats
	i++
	}
	fmt.Fprintln(categorias)
	resp:=http.Get("https://api.mercadolibre.com")
}
func GetCategories(siteID string)(Categories, error){
	response, err:=http.Get("https://api.mercadolibre.com/sites/MLA/Categories") //completar
	bytes:=ioutil.ReadAll(response.bytes)
	
	var cats Category
	err:=json.Unmarshal(bytes, &cats)

	return cats, nil
}
type Category struct{
	Id string 'json: "id"'
	Name string 'json: "name"'
}
type Categories []category

//invocacion correspondiente mediante http
//get    resp:= http.Get(url)
//Post   resp:=http.Post(url,body)



//leer los datos
//

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

