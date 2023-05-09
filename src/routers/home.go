package routers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// verifica se o servidor está funcionando ou não.
func SetupRouterHome(r *gin.Engine) {

	r.GET("/", home)
}

func home(c *gin.Context) {
	// Carrega o arquivo HTML
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	// Renderiza o arquivo HTML e escreve a resposta
	err = tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}
}
