package manipulador
import "html/template"

//Modelos aramzena os modelo de pagina ola
var ModelosOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal aramzena os modelo de pagina local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))