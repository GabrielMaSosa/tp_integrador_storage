package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/customers"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/pkg"
	"github.com/gin-gonic/gin"
)

type Customers struct {
	s customers.Service
}

func NewHandlerCustomers(s customers.Service) *Customers {
	return &Customers{s}
}

func (c *Customers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := c.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, customers)
	}
}

func (c *Customers) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := domain.Customers{}
		err := ctx.ShouldBindJSON(&customer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = c.s.Create(&customer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": customer})
	}
}

func (c *Customers) GetRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorize")
			return
		}
		ret, err := pkg.RecoveryCustomers()

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		for _, v := range ret {
			err = c.s.Create(&v)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}

		}

		ctx.JSON(200, gin.H{"data": "Recovery"})
	}

}

func (c *Customers) GetTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		ret, err := c.s.TotalCustomers()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": ret})
	}

}

func (c *Customers) GetActiveTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		ret, err := c.s.TopCustomerActive()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": ret})
	}

}
