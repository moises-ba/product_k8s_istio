package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type priceResponse struct {
	Message string
	Price   float64
}

func main() {

	products := [5]string{"martelo", "serrote"}

	r := gin.Default()
	r.GET("/product/:productId", func(c *gin.Context) {

		paramProduct := c.Param("productId")
		var product string

		for _, p := range products {
			if p == paramProduct {
				product = p
			}
		}

		if product == "" {
			c.JSON(400, gin.H{
				"message": "Produto " + paramProduct + " não encontrado.",
			})
		} else {

			endPointPrice := os.Getenv("MS_PRICE_EP") //"http://ms-product-price/price/"

			urlPreco := endPointPrice + product

			log.Println("URL servico de preço: " + urlPreco)

			res, err := http.Get(urlPreco)
			if err != nil {
				log.Println("Falha ao buscar preço do produto: " + err.Error())

				c.JSON(200, gin.H{
					"product": product,
					"message": "Falha ao buscar preço do produto" + err.Error(),
				})
			} else {

				defer res.Body.Close()

				priceResponse := priceResponse{}
				err := json.NewDecoder(res.Body).Decode(&priceResponse)
				if err != nil {
					log.Println("Falha ao ler corpo da resposta do preço: " + err.Error())

					body, err2 := ioutil.ReadAll(res.Body)
					if err2 != nil {
						log.Println(err2)
					} else {
						log.Println("erro retornado pelo serviço de preços: " + string(body))
					}

					c.JSON(200, gin.H{
						"product": product,
						"message": "Falha ao ler corpo da resposta do preço" + err.Error(),
					})
				} else {
					c.JSON(200, gin.H{
						"product": product,
						"price":   priceResponse.Price,
					})
				}

			}

		}

	})
	r.Run(":8080")

}
