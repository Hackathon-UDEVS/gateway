package handler

import (
	"context"
	"log"
	"net/http"

	tender_service "github.com/Hackaton-UDEVS/gateway/internal/genproto/tender-service"
	"github.com/gin-gonic/gin"
)

// CreateTender godoc
// @Summary Create a new tender
// @Description Creates a new tender and returns the created tender data
// @Accept json
// @Produce json
// @Tags tender
// @Param tender body tender_service.CreateTenderReq true "Tender creation request"
// @Success 200 {object} tender_service.ResponseMessage "Tender created successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /client/create-tender [post]
func (h *Handler) CreateTender(c *gin.Context) {
	// Parse request body
	var req tender_service.CreateTenderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error while parsing request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call service to create tender
	resp, err := h.Clients.Client.CreateTender(context.Background(), &req)
	if err != nil {
		log.Println("Error while creating tender:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tender"})
		return
	}

	// Respond with success
	log.Println("Tender created successfully")
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// UpdateTender godoc
// @Summary Update an existing tender
// @Description Updates the status of an existing tender
// @Accept json
// @Produce json
// @Tags tender
// @Param tender body tender_service.UpdateTenderStatusReq true "Tender update request"
// @Success 200 {object} tender_service.ResponseMessage "Tender updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /client/update-tender [put]
func (h *Handler) UpdateTender(c *gin.Context) {
	// Parse request body
	var req tender_service.UpdateTenderStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error while parsing request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call service to update tender
	resp, err := h.Clients.Client.UpdateTenderStatus(context.Background(), &req)
	if err != nil {
		log.Println("Error while updating tender:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tender"})
		return
	}

	// Respond with success
	log.Println("Tender updated successfully")
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// DeleteTender godoc
// @Summary Delete a tender
// @Description Deletes a tender by ID
// @Accept json
// @Produce json
// @Tags tender
// @Param id path string true "Tender ID"
// @Success 200 {object} string "Tender deleted successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /client/delete-tender/{id} [delete]
func (h *Handler) DeleteTender(c *gin.Context) {
	// Parse tender ID from path
	id := c.Param("id")
	if id == "" {
		log.Println("Tender ID is missing")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tender ID is required"})
		return
	}

	// Prepare request
	req := &tender_service.DeleteTenderReq{TenderId: id}

	// Call service to delete tender
	resp, err := h.Clients.Client.DeleteTender(context.Background(), req)
	if err != nil {
		log.Println("Error while deleting tender:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tender"})
		return
	}

	// Respond with success
	log.Println("Tender deleted successfully")
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// GetTenders godoc
// @Summary Get tenders
// @Description Retrieves a list of tenders for the authenticated user
// @Accept json
// @Produce json
// @Tags tender
// @Param query body tender_service.GetMyTendersReq true "Query parameters"
// @Success 200 {object} tender_service.ResponseMessage "List of tenders retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /client/getAll-tenders [get]
func (h *Handler) GetTenders(c *gin.Context) {
	// Parse query parameters
	var req tender_service.GetMyTendersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Println("Error while parsing query parameters:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// Call service to get tenders
	resp, err := h.Clients.Client.GetMyTenders(context.Background(), &req)
	if err != nil {
		log.Println("Error while retrieving tenders:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tenders"})
		return
	}

	// Respond with success
	log.Println("Tenders retrieved successfully")
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

// SortTenders godoc
// @Summary Sort tenders
// @Description Retrieves a sorted list of all tenders
// @Accept json
// @Produce json
// @Tags tender
// @Param query body tender_service.GetAllTendersReq true "Query parameters"
// @Success 200 {object} tender_service.TendersList "Sorted list of tenders retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /client/tenders/sort [get]
func (h *Handler) SortTenders(c *gin.Context) {
	// Parse query parameters
	var req tender_service.GetAllTendersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Println("Error while parsing query parameters:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// Call service to get sorted tenders
	resp, err := h.Clients.Client.GetAllTenders(context.Background(), &req)
	if err != nil {
		log.Println("Error while retrieving sorted tenders:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sorted tenders"})
		return
	}

	// Respond with success
	log.Println("Sorted tenders retrieved successfully")
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
