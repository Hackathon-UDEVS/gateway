package handler

import (
	"context"
	"gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"

	pb "gateway/internal/genproto/contractors"
)

// SubmitBid godoc
// @Summary Submit a bid
// @Description Submits a bid for a contractor
// @Accept json
// @Produce json
// @Param bid body pb.SubmitBidRequest true "Bid submission request"
// @Success 200 {object} pb.SubmitBidResponse "Bid submitted successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /submit-bid [post]
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

// GetListOfBids godoc
// @Summary Get list of bids
// @Description Retrieves a list of bids for a tender
// @Accept json
// @Produce json
// @Param tenderId query string true "Tender ID"
// @Success 200 {object} pb.GetBidsResponse "List of bids retrieved successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /bids [get]
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
