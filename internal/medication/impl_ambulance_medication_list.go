package medication

import (
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_medication_list.go

// CreateMedicationListEntry - Saves new entry into medication list
func (this *implMedicationListAPI) CreateMedicationListEntry(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		var entry MedicationListEntry

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if entry.Id == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Medication ID is required",
			}, http.StatusBadRequest
		}

		if entry.Id == "" || entry.Id == "@new" {
			entry.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(ambulance.MedicationList, func(medication MedicationListEntry) bool {
			return entry.Id == medication.Id || entry.Id == medication.Id
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Entry already exists",
			}, http.StatusConflict
		}

		ambulance.MedicationList = append(ambulance.MedicationList, entry)

		entryIndx := slices.IndexFunc(ambulance.MedicationList, func(medication MedicationListEntry) bool {
			return entry.Id == medication.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save entry",
			}, http.StatusInternalServerError
		}

		return ambulance, ambulance.MedicationList[entryIndx], http.StatusOK
	})
}

// DeleteMedicationListEntry - Deletes a specific medication
func (this *implMedicationListAPI) DeleteMedicationListEntry(ctx *gin.Context) {
	// update ambulance document
	updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		entryId := ctx.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(ambulance.MedicationList, func(medication MedicationListEntry) bool {
			return entryId == medication.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		ambulance.MedicationList = append(ambulance.MedicationList[:entryIndx], ambulance.MedicationList[entryIndx+1:]...)
		return ambulance, nil, http.StatusNoContent
	})
}

// GetMedicationListEntries - Provides the ambulance medication list
func (this *implMedicationListAPI) GetMedicationListEntries(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		result := ambulance.MedicationList
		if result == nil {
			result = []MedicationListEntry{}
		}
		// return nil ambulance - no need to update it in db
		return nil, result, http.StatusOK
	})
}

// GetMedicationListEntry - Provides details about a specific medication
func (this *implMedicationListAPI) GetMedicationListEntry(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		entryId := ctx.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(ambulance.MedicationList, func(medication MedicationListEntry) bool {
			return entryId == medication.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		// return nil ambulance - no need to update it in db
		return nil, ambulance.MedicationList[entryIndx], http.StatusOK
	})
}

// UpdateMedicationListEntry - Updates specific medication details
func (this *implMedicationListAPI) UpdateMedicationListEntry(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(c *gin.Context, ambulance *Ambulance) (*Ambulance, interface{}, int) {
		var entry MedicationListEntry

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		entryId := ctx.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(ambulance.MedicationList, func(medication MedicationListEntry) bool {
			return entryId == medication.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		if entry.Id != "" {
			ambulance.MedicationList[entryIndx].Id = entry.Id
		}

		if entry.Id != "" {
			ambulance.MedicationList[entryIndx].Id = entry.Id
		}

		return ambulance, ambulance.MedicationList[entryIndx], http.StatusOK
	})
}
