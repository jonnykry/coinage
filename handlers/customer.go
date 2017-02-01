package handlers

import (
	"github.com/pborman/uuid"
	"gopkg.in/alexcesaro/statsd.v2"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ghmeier/bloodlines/handlers"
	"github.com/jonnykry/coinage/helpers"
	"github.com/jonnykry/coinage/models"
)

type CustomerI interface {
	New(ctx *gin.Context)
	View(ctx *gin.Context)
	ViewAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
	UpdatePayment(ctx *gin.Context)
	Subscribe(ctx *gin.Context)
}

type Customer struct {
	*handlers.BaseHandler
	Helper *helpers.Customer
}

func NewCustomer(ctx *handlers.GatewayContext) CustomerI {
	stats := ctx.Stats.Clone(statsd.Prefix("api.customer"))
	return &Customer{
		Helper:      helpers.NewCustomer(ctx.Sql, ctx.Stripe),
		BaseHandler: &handlers.BaseHandler{Stats: stats},
	}
}

func (c *Customer) New(ctx *gin.Context) {
	var json models.Customer
	err := ctx.BindJSON(json)
	if err != nil {
		c.UserError(ctx, "ERROR: unable to parse body", err)
		return
	}

	c.Success(ctx, nil)
}

func (c *Customer) ViewAll(ctx *gin.Context) {
	c.Success(ctx, nil)
}

func (c *Customer) View(ctx *gin.Context) {
	id := ctx.Param("id")

	customer, err := c.Helper.View(uuid.Parse(id))
	if err != nil {
		c.ServerError(ctx, err, nil)
		return
	}

	c.Success(ctx, customer)
}

func (c *Customer) Subscribe(ctx *gin.Context) {
	c.Success(ctx, nil)
}

func (c *Customer) UpdatePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	token, ok := ctx.GetQuery("token")
	if !ok {
		c.UserError(ctx, "ERROR: token is a required parameter", nil)
	}

	err := c.Helper.AddSource(uuid.Parse(id), token)
	if err != nil {
		c.ServerError(ctx, err, &gin.H{id: id, token: token})
		return
	}

	c.Success(ctx, nil)
}

func (c *Customer) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.Helper.Delete(uuid.Parse(id))
	if err != nil {
		c.ServerError(ctx, err, nil)
		return
	}

	c.Success(ctx, nil)
}