package handler

import (
	"context"
	pb "gateway/internal/genproto/clients"
	"gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTender godoc
// @Summary Create a new tender
// @Description Creates a new tender and returns the created tender data
// @Accept json
// @Produce json
// @Param tender body pb.CreateTenderReq true "Tender creation request"
// @Success 200 {object} pb.CreateTenderResp "Tender created successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /create-tender [post]
func (h *Handler) CreateTender(r *gin.Context) {
	log, _ := logs.NewLogger()
	var req pb.CreateTenderReq
	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.Client.CreateTender(context.Background(), &req)
	if err != nil {
		log.Error("Error while creating Tender")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Created Tender")

	r.JSON(http.StatusOK, gin.H{"data": &resp})
}

// UpdateTender godoc
// @Summary Update an existing tender
// @Description Updates the status of an existing tender
// @Accept json
// @Produce json
// @Param tender body pb.UpdateTenderStatusReq true "Tender update request"
// @Success 200 {object} pb.UpdateTenderStatusResp "Tender updated successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /update-tender [put]
func (h *Handler) UpdateTender(r *gin.Context) {

	log, _ := logs.NewLogger()

	var req pb.UpdateTenderStatusReq
	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Clients.Client.UpdateTenderStatus(context.Background(), &req)
	if err != nil {
		log.Error("Error while updating Tender")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Updated Tender")
	r.JSON(http.StatusOK, gin.H{"data": &resp})

}

// DeleteTender godoc
// @Summary Delete a tender
// @Description Deletes a tender by ID
// @Accept json
// @Produce json
// @Param tender body pb.DeleteTenderReq true "Tender deletion request"
// @Success 200 {object} gin.H "Tender deleted successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /delete-tender[:id] [delete]
func (h *Handler) DeleteTender(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req pb.DeleteTenderReq
	if err := r.ShouldBindJSON(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Clients.Client.DeleteTender(context.Background(), &req)
	if err != nil {
		log.Error("Error while deleting Tender")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Deleted Tender")
	r.JSON(http.StatusOK, gin.H{"data": resp})

}

// GetTenders godoc
// @Summary Get tenders
// @Description Retrieves a list of tenders for the authenticated user
// @Accept json
// @Produce json
// @Param query body pb.GetMyTendersReq true "Query parameters"
// @Success 200 {object} pb.GetMyTendersResp "List of tenders retrieved successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /getAll-tenders [get]
func (h *Handler) GetTenders(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req pb.GetMyTendersReq

	if err := r.ShouldBindQuery(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Clients.Client.GetMyTenders(context.Background(), &req)
	if err != nil {
		log.Error("Error while getting Tenders")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Retrieved Tenders")
	r.JSON(http.StatusOK, gin.H{"data": &resp})

}

// SortTenders godoc
// @Summary Sort tenders
// @Description Retrieves a sorted list of all tenders
// @Accept json
// @Produce json
// @Param query body pb.GetAllTendersReq true "Query parameters"
// @Success 200 {object} pb.GetAllTendersResp "Sorted list of tenders retrieved successfully"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /tenders/sort [get]
func (h *Handler) SortTenders(r *gin.Context) {
	log, _ := logs.NewLogger()

	var req pb.GetAllTendersReq

	if err := r.ShouldBindQuery(&req); err != nil {
		log.Error("Error while parsing request")
		r.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.Clients.Client.GetAllTenders(context.Background(), &req)
	if err != nil {
		log.Error("Error while getting Tenders")
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Info("Retrieved Tenders")
	r.JSON(http.StatusOK, gin.H{"data": resp})
}
