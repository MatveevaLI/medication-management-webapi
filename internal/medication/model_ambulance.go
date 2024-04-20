/*
 * Medication Management API
 *
 * Medication management for the Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: ladaivanna.matveeva@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package medication

type Ambulance struct {

	// Unique identifier of the ambulance
	Id string `json:"id"`

	// Human readable display name of the ambulance
	Name string `json:"name"`

	RoomNumber string `json:"roomNumber"`

	MedicationList []MedicationListEntry `json:"medicationList,omitempty"`
}
