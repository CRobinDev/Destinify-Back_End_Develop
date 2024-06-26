package rest

import (
	"github.com/CRobin69/Destinify-Back_End_Develop/model"
	"github.com/CRobin69/Destinify-Back_End_Develop/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) GetTicketByID(ctx *gin.Context) {
	id := ctx.Param("id")
	param := model.TicketParam{ID: uuid.MustParse(id)}
	ticket, err := r.service.TicketService.GetTicketByID(param)
	if err != nil {
		helper.Error(ctx, http.StatusInternalServerError, "failed to get ticket", err)
		return
	}

	helper.Success(ctx, http.StatusOK, "success get ticket", ticket)
}

func (r *Rest) BuyTicket(ctx *gin.Context) {
	var ticketBuy model.TicketBuy
	if err := ctx.ShouldBindJSON(&ticketBuy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User authentication failed"})
		return
	}
	
	ticketBuy.UserID = user.ID

	order, tickets, err := r.service.TicketService.BuyTickets(ticketBuy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to buy tickets", "details": err.Error()})
		return
	}

	var ticketIDs []uuid.UUID
	for _, ticket := range tickets {
		ticketIDs = append(ticketIDs, ticket.ID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tickets created successfully", 
		"OrderID": order.ID,
		"TotalPrice": order.TotalPrice,
		"TicketIDs": ticketIDs})
}
