package handler

import (
	"net/http"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/sales"
	"github.com/bootcamp-go/desafio-cierre-db.git/pkg"
	"github.com/gin-gonic/gin"
)

type Sales struct {
	s sales.Service
}

func NewHandlerSales(s sales.Service) *Sales {
	return &Sales{s}
}

func (s *Sales) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := s.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (s *Sales) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sale := domain.Sales{}
		err := ctx.ShouldBindJSON(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = s.s.Create(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": sale})
	}
}

func (s *Sales) GetRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		pass := ctx.GetHeader("password")
		if pass != "123456" {
			ctx.String(http.StatusUnauthorized, "Unauthorize")
			return
		}
		ret, err := pkg.RecoverySales()

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		for _, v := range ret {
			err = s.s.Create(&v)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}

		}

		ctx.JSON(200, gin.H{"data": "Recovery"})
	}

}
