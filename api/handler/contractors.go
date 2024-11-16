package handler

import (
	"context"
	"gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"

	pb "gateway/internal/genproto/contractors"
)

func (h *Handler) SubmitBid(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req pb.SubmitBidRequest

	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.Contractor.SubmitBid(context.Background(), &req)
	if err != nil {
		log.Error("Error while submit bid")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Successfully submit bid")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}

func (h *Handler) GetListOfBids(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req pb.GetBidsRequest

	if err := r.ShouldBindQuery(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.Contractor.GetBidsForTender(context.Background(), &req)
	if err != nil {
		log.Error("Error while get bids for tender")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Successfully get bids for tender")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}
