package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/invoices"
	"github.com/bootcamp-go/desafio-cierre-db.git/pkg"
	"github.com/gin-gonic/gin"
)

type Invoices struct {
	s invoices.Service
}

func NewHandlerInvoices(s invoices.Service) *Invoices {
	return &Invoices{s}
}

func (i *Invoices) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := i.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (i *Invoices) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices := domain.Invoices{}
		err := ctx.ShouldBindJSON(&invoices)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = i.s.Create(&invoices)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": invoices})
	}
}

func (i *Invoices) GetRepareDB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorize")
			return
		}

		dat, err := i.s.RepareInvoices()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": dat})
	}
}

func (i *Invoices) GetRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorize")
			return
		}
		ret, err := pkg.RecoveryInvoices()

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		for _, v := range ret {
			err = i.s.Create(&v)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}

		}

		ctx.JSON(200, gin.H{"data": "Recovery"})
	}

}
