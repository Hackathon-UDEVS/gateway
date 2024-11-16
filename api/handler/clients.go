package handler

import (
	"context"
	pb "gateway/internal/genproto/clients"
	"gateway/internal/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
