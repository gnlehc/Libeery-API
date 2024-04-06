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

// Function to create a booking loker by session
func CreateBooking(c *gin.Context, reqBody model.BookingRequestDTO) error {
	// req initialization
	req := model.BookingRequestDTO{
		UserID:    reqBody.UserID,
		SessionID: reqBody.SessionID,
		LokerID:   reqBody.LokerID,
	}

	db := database.GlobalDB

	userIDexists := db.Where("user_id = ?", req.UserID).First(&model.MsUser{}).RowsAffected
	if userIDexists == 0 {
		return errors.New("user not found")
	}

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
	err := updateLockerAvailability(tx, req.LokerID, req.SessionID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create booking
	booking := model.TrBooking{
		UserID:          req.UserID,
		SessionID:       req.SessionID,
		LokerID:         req.LokerID,
		BookingStatusID: 1,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
		Stsrc:           "A",
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
// then when user check all the available loker in session_id = 2,
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

func GetAllBookingData(c *gin.Context) ([]*model.TrBooking, error) {
	db := database.GlobalDB
	var bookings []*model.TrBooking
	if err := db.Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func CheckInBooking(c *gin.Context) (model.TrBooking, error) {
	req := model.CheckInBookingRequestDTO{
		UserID:    c.Param("userID"),
		BookingID: c.Param("bookingID"),
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return model.TrBooking{}, err
	}

	bookingID, err := uuid.Parse(req.BookingID)
	if err != nil {
		return model.TrBooking{}, err
	}

	db := database.GlobalDB
	booking := model.TrBooking{}
	bookingIDexists := db.Where("booking_id = ? and user_id = ?", bookingID, userID).First(&booking).RowsAffected
	if bookingIDexists == 0 {
		return model.TrBooking{}, errors.New("booking not found")
	}

	booking.BookingStatusID = 2
	if err := db.Save(&booking).Error; err != nil {
		return model.TrBooking{}, err
	}
	return booking, nil
}

func GetUserBookingData(c *gin.Context) ([]*model.TrBooking, error) {
	db := database.GlobalDB
	var bookings []*model.TrBooking
	userID := c.Param("userID")

	if err := db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(400, gin.H{"message": "Error occurred while retrieving bookings"})
		return nil, err
	}
	return bookings, nil
}
