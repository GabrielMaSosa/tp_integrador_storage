package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/bootcamp-go/desafio-cierre-db.git/pkg"
	"github.com/gin-gonic/gin"
)

type Products struct {
	s products.Service
}

func NewHandlerProducts(s products.Service) *Products {
	return &Products{s}
}

func (p *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, products)
	}
}

func (p *Products) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := domain.Product{}
		err := ctx.ShouldBindJSON(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = p.s.Create(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": products})
	}
}

func (p *Products) GetRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorize")
			return
		}
		ret, err := pkg.RecoveryProducts()

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		for _, v := range ret {
			err = p.s.Create(&v)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}

		}

		ctx.JSON(200, gin.H{"data": "Recovery"})
	}

}

func (p *Products) GetTotalTop() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		ret, err := p.s.GetTopProducts()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": ret})
	}

}
