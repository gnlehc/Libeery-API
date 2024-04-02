package model

// book loker by session time
type MsLoker struct {
	LockerID     int    `gorm:"primaryKey" json:"LockerID"`
	RowNumber    int    `json:"RowNumber"`
	ColumnNumber int    `json:"ColumnNumber"`
	Availability string `json:"Availability"`
	Stsrc        string `json:"Stsrc" gorm:"size:1"`
}

// func NewMsLoker(lockerID, rowNumber, columnNumber int) *MsLoker {
// 	return &MsLoker{
// 		LockerID:        lockerID,
// 		RowNumber:       rowNumber,
// 		ColumnNumber:    columnNumber,
// 		Availability:    make(map[*ForLaterSession]bool),
// 		BookingSessions: make(map[*ForLaterSession]*MsBookingStatus),
// 		Stsrc:           "A",
// 	}
// }

// // Method to book the locker for a given session
// func (l *MsLoker) BookSession(sessionID *ForLaterSession, booking *MsBookingStatus) {
// 	// Mark session as booked
// 	// default availability is true (available)
// 	l.Availability[sessionID] = false // after booked jadi false
// 	// Store booking session
// 	l.BookingSessions[sessionID] = booking
// }

// // Method to check if the locker is available for a given session
// func (l *MsLoker) IsAvailable(sessionID *ForLaterSession) bool {
// 	// Check availability for the session ID
// 	return l.Availability[sessionID]
// }
