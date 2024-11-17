package handler

import (
	"context"
	"github.com/Hackaton-UDEVS/gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Hackaton-UDEVS/gateway/internal/genproto/tender-service"
)

// SubmitBid godoc
// @Summary Submit a bid
// @Description Submits a bid for a contractor
// @Accept json
// @Produce json
// @tags bid
// @Param bid body tender_service.SubmitBidRequest true "Bid submission request"
// @Success 200 {object} tender_service.BidResponse "Bid submitted successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /submit-bid [post]
func (h *Handler) SubmitBid(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req tender_service.SubmitBidRequest

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
// @tags bid
// @Param tenderId query string true "Tender ID"
// @Success 200 {object} tender_service.BidsListResponse "List of bids retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /bids [get]
func (h *Handler) GetListOfBids(r *gin.Context) {
	log, _ := logs.NewLogger()

	req := tender_service.GetBidsRequest{}
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

// GetMyBids godoc
// @Summary Get list of bids by the current user
// @Description Retrieves all bids made by the authenticated user on various tenders
// @Accept json
// @Produce json
// @tags bid
// @Param limit query int false "Limit for pagination"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} tender_service.BidsListResponse "List of bids retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /bids/my [get]
func (h *Handler) GetMyBids(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req tender_service.GetMyBidsRequest

	if err := r.ShouldBindQuery(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.Contractor.GetMyBids(context.Background(), &req)
	if err != nil {
		log.Error("Error while get bids for tender")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Successfully get bids for tender")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}
