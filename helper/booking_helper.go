package helper

import (
	"Libeery/database"
	model "Libeery/model"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Function to make a booking
func MakeBooking(req model.BookingRequestDTO) error {
	db := database.GlobalDB

	// Check if the user already has a booking for the same session
	hasBooking := userHasBooking(req.UserID, req.SessionID)
	if hasBooking {
		return errors.New("user already has a booking for the same session")
	}

	// Begin transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update locker availability
	err := updateLockerAvailability(tx, req.LokerID, req.SessionID, false)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create booking
	booking := model.TrBooking{
		UserID:    req.UserID,
		SessionID: req.SessionID,
		LokerID:   req.LokerID,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Stsrc:     "A",
	}
	err = tx.Create(&booking).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

// Function to check if the user already has a booking for the same session
func userHasBooking(userID uuid.UUID, sessionID int) bool {
	db := database.GlobalDB
	var count int64
	db.Model(&model.TrBooking{}).Where("user_id = ? AND session_id = ?", userID, sessionID).Count(&count)
	return count > 0
}

// Function to update locker availability
func updateLockerAvailability(tx *gorm.DB, lockerID int, sessionID int, available bool) error {
	err := tx.Model(&model.MsLoker{}).Where("locker_id = ? AND session_id = ?", lockerID, sessionID).Update("available", available).Error
	return err
}

// Function to create a booking
// func createBooking(tx *sql.Tx, userID int, sessionID int, lockerID int) error {
// 	_, err := tx.Exec("INSERT INTO Booking (UserID, SessionID, LockerID) VALUES (?, ?, ?)", userID, sessionID, lockerID)
// 	return err
// }

// import (
// 	"Libeery/model"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func BookASessionForLater(c *gin.Context) {
// 	var req model.BookingRequestDTO

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Retrieve the locker
// 	locker := getLockerByID(req.LokerID) // You need to implement this function

// 	if locker == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Locker not found"})
// 		return
// 	}

// 	// Book the session
// 	locker.BookSession(&model.ForLaterSession{
// 		StartSession: req.StartSession,
// 		EndSession:   req.EndSession,
// 	}, &model.MsBookingStatus{
// 		BookingStatusID: req.BookingStatusID,
// 		// You may include other fields as needed
// 	})

// 	// Respond with success message
// 	c.JSON(http.StatusOK, gin.H{"message": "Session booked successfully"})
// }

// func getLockerByID(lockerID int) *model.MsLoker {
// 	// Assuming you have some mechanism to retrieve the locker from your data store (e.g., database)
// 	// Here is a simple example using a temporary in-memory slice of lockers
// 	// Replace this with your actual data retrieval mechanism

// 	// Temporary in-memory data (replace this with your actual data source)
// 	lockers := []*model.MsLoker{
// 		{LockerID: 1, RowNumber: 1, ColumnNumber: 1, Availability: make(map[*model.ForLaterSession]bool), BookingSessions: make(map[*model.ForLaterSession]*model.MsBookingStatus), Stsrc: "A"},
// 		{LockerID: 2, RowNumber: 2, ColumnNumber: 2, Availability: make(map[*model.ForLaterSession]bool), BookingSessions: make(map[*model.ForLaterSession]*model.MsBookingStatus), Stsrc: "A"},
// 		// Add more lockers as needed
// 	}

// 	// Search for the locker by ID
// 	for _, locker := range lockers {
// 		if locker.LockerID == lockerID {
// 			return locker
// 		}
// 	}

// 	// If the locker with the given ID is not found, return nil
// 	return nil
// }
