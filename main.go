package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Servir arquivos estáticos (CSS, imagens)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("static/styles"))))
	http.Handle("/imagens/", http.StripPrefix("/imagens/", http.FileServer(http.Dir("static/imagens"))))

	// Exibir a parte HTML do site
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	// Iniciar o servidor
	fmt.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
