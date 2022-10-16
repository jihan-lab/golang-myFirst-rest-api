package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// car
// {
// 	"id" : "1",
// 	"brand" : "Honda",
// 	"car_type" : "City"
// }

type car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"car_type"`
}

var cars = []car{
	{ID: "1", Brand: "Honday", Type: "City"},
	{ID: "2", Brand: "Yamaha", Type: "King"},
}

func main() {
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	// GET /cars - list cars
	r.GET("/cars", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cars)
	})

	// POST /cars - create car
	r.POST("/cars", func(ctx *gin.Context) {
		var car car
		if err := ctx.ShouldBindJSON(&car); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		cars = append(cars, car)
		ctx.JSON(http.StatusCreated, car)
	})

	// DELETE /cars/:id - delete car
	r.DELETE("/cars/:car_id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				cars = append(cars[:i], cars[i+1:]...)
				break
			}
		}
	})

	// PUT /cars/:id - update car
	r.PUT("/cars/:car_id", func(ctx *gin.Context) {
		var c car

		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				if err := ctx.ShouldBindJSON(&c); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}
				cars[i] = c
			}
		}
	})

	// GET /cars/:id - get one car
	r.GET("/cars/:car_id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				ctx.JSON(http.StatusOK, cars[i])
			}
		}
	})

	r.Run()
}
