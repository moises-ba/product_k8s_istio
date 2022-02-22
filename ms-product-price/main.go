package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	products := map[string]float64{
		"martelo": 1.9,
		"serrote": 5.0,
	}

	r := gin.Default()

	r.GET("/price/:productId", func(c *gin.Context) {

		paramProduct := c.Param("productId")
		price, ok := products[paramProduct]

		if ok {
			c.JSON(200, gin.H{
				"product": paramProduct,
				"price":   price,
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Preço para o produto " + paramProduct + "não encontrado.",
			})
		}

	})
	r.Run(":8081")

}
