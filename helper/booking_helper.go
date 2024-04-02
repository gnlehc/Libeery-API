package helper

import (
	"Libeery/database"
	model "Libeery/model"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Function to make a booking
func MakeBooking(c *gin.Context, reqBody model.BookingRequestDTO) error {
	// Mapping the fields from the request body struct to the BookingRequestDTO struct
	req := model.BookingRequestDTO{
		UserID:    reqBody.UserID,
		SessionID: reqBody.SessionID,
		LokerID:   reqBody.LokerID,
	}

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
	err := updateLockerAvailability(tx, req.LokerID, req.SessionID)
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

// MAIN LOGIC
// the availability should be updated along with the session,
// if user book with session_id = 1 and loker_id = 1,
// then when user check all the available loker in session_id = 1,
// loker_id = 1 should be available

// Function to update locker availability within a session
func updateLockerAvailability(tx *gorm.DB, lockerID int, sessionID int) error {
	// Check if the locker is booked for the specified session
	var count int64
	err := tx.Model(&model.TrBooking{}).
		Where("loker_id = ? AND session_id = ?", lockerID, sessionID).
		Count(&count).Error
	if err != nil {
		return err
	}

	// If the count is greater than 0, it means the locker is booked for the session
	if count > 0 {
		// Update the locker availability to Booked
		err = tx.Model(&model.MsLoker{}).
			Where("locker_id = ?", lockerID).
			Update("availability", "Booked").Error
		if err != nil {
			return err
		}
	} else {
		// If the count is 0, it means the locker is available for the session
		// Update the locker availability to Active
		err = tx.Model(&model.MsLoker{}).
			Where("locker_id = ?", lockerID).
			Update("availability", "Active").Error
		if err != nil {
			return err
		}
	}

	return nil
}

// func updateLockerAvailability(tx *gorm.DB, lockerID int, sessionID int) error {
// 	err := tx.Model(&model.MsLoker{}).
// 		Where("locker_id = ?", lockerID).
// 		Joins("JOIN (SELECT * FROM ms_sessions WHERE session_id = ?) AS sessions ON ms_lokers.session_id = sessions.session_id", sessionID).
// 		Update("Availability", "Booked").Error
// 	return err
// }
