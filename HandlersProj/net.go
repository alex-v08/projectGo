package main


import (
    "fmt"
    "net/http"
)

func main() {
    // Definir el handler para la ruta "/"
    http.HandleFunc("/", handler)

    // Iniciar el servidor en el puerto 8080
    fmt.Println("Servidor escuchando en http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

// Handler para la ruta "/"
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola desde el servidor!")
}
