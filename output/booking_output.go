package output

import "Libeery/model"

type BookingOutput struct {
	Data       []*model.TrBooking
	BaseOutput BaseOutput
}

type CheckInBookingOutput struct {
	Data       model.TrBooking
	BaseOutput BaseOutput
}
