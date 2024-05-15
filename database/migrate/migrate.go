package migrate

import (
	"Libeery/model"
	"time"

	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) error {
	migrator := db.Migrator()

	// Check if tables exist and migrate if necessary
	if !migrator.HasTable(&model.MsMahasiswa{}) {
		if err := db.AutoMigrate(&model.MsMahasiswa{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.MsStaff{}) {
		if err := db.AutoMigrate(&model.MsStaff{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.MsUser{}) {
		if err := db.AutoMigrate(&model.MsUser{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.TrBooking{}) {
		if err := db.AutoMigrate(&model.TrBooking{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.MsBookingStatus{}) {
		if err := db.AutoMigrate(&model.MsBookingStatus{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.MsLoker{}) {
		if err := db.AutoMigrate(&model.MsLoker{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.MsSession{}) {
		if err := db.AutoMigrate(&model.MsSession{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.MsAcara{}) {
		if err := db.AutoMigrate(&model.MsAcara{}); err != nil {
			return err
		}
	}

	// Data seeding
	if err := seedDefaultMahasiswaData(db); err != nil {
		return err
	}

	if err := seedDefaultStaffData(db); err != nil {
		return err
	}

	if err := seedDefaultBookingStatusData(db); err != nil {
		return err
	}

	if err := seedDefaultMsSession(db); err != nil {
		return err
	}
	if err := seedDefeaultMsLokerData(db); err != nil {
		return err
	}
	if err := seedDefaultMsUsersData(db); err != nil {
		return err
	}
	if err := seedDefaultAcara(db); err != nil {
		return err
	}

	return nil
}

// Seed default Mahasiswa
func seedDefaultMahasiswaData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsMahasiswa{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultMahasiswaData := []model.MsMahasiswa{
			{NIM: "2602057652", MhsName: "Chelsea Ng", MhsPassword: "isAdmin24", Stsrc: "A"},
			{NIM: "2602063831", MhsName: "Verena Vynne Sentosa", MhsPassword: "isMahasiswa24", Stsrc: "A"},
			{NIM: "2602063560", MhsName: "Christopher Verrell", MhsPassword: "verinthebuilding", Stsrc: "A"},
			{NIM: "2602053515", MhsName: "Jesslyn Amanda Mulyawan", MhsPassword: "isAdmin13", Stsrc: "A"},
			{NIM: "2602062652", MhsName: "Nicholas Owen Sentosa", MhsPassword: "owen123", Stsrc: "A"},
		}
		for _, data := range defaultMahasiswaData {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Seed default Staff
func seedDefaultStaffData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsStaff{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultStaffData := []model.MsStaff{
			{NIS: "D5416", StaffName: "Kanyadian Idhananta", StaffPassword: "isDosen", Stsrc: "A"},
			{NIS: "D6831", StaffName: "Islam Nur Alam", StaffPassword: "isAlam", Stsrc: "A"},
			{NIS: "D6657", StaffName: "Anderies Anderies", StaffPassword: "isAnderies", Stsrc: "A"},
			{NIS: "D6811", StaffName: "Faisal Asadi", StaffPassword: "isFaisal", Stsrc: "A"},
			{NIS: "D6835", StaffName: "Hanis Amalia Saputri", StaffPassword: "isHanis", Stsrc: "A"},
		}
		for _, data := range defaultStaffData {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Seed default BookingStatus
func seedDefaultBookingStatusData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsBookingStatus{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultBookingStatusData := []model.MsBookingStatus{
			{BookingStatusID: 1, StatusTitle: "Pending", Stsrc: "A"},
			{BookingStatusID: 2, StatusTitle: "Checked-In", Stsrc: "A"},
			{BookingStatusID: 3, StatusTitle: "Checked-Out", Stsrc: "A"},
		}
		for _, data := range defaultBookingStatusData {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// Seed default MsSession
func seedDefaultMsSession(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsSession{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultMsSession := []model.MsSession{
			{SessionID: 1, StartSession: parseTime("08:00:00"), EndSession: parseTime("09:00:00")},
			{SessionID: 2, StartSession: parseTime("09:00:00"), EndSession: parseTime("10:00:00")},
			{SessionID: 3, StartSession: parseTime("10:00:00"), EndSession: parseTime("11:00:00")},
			{SessionID: 4, StartSession: parseTime("11:00:00"), EndSession: parseTime("12:00:00")},
			{SessionID: 5, StartSession: parseTime("12:00:00"), EndSession: parseTime("13:00:00")},
			{SessionID: 6, StartSession: parseTime("13:00:00"), EndSession: parseTime("14:00:00")},
			{SessionID: 7, StartSession: parseTime("14:00:00"), EndSession: parseTime("15:00:00")},
			{SessionID: 8, StartSession: parseTime("15:00:00"), EndSession: parseTime("16:00:00")},
			{SessionID: 9, StartSession: parseTime("16:00:00"), EndSession: parseTime("17:00:00")},
			{SessionID: 10, StartSession: parseTime("17:00:00"), EndSession: parseTime("18:00:00")},
		}
		for _, session := range defaultMsSession {
			if err := db.Create(&session).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func parseTime(timeStr string) time.Time {
	layout := "15:04:05"
	t, _ := time.Parse(layout, timeStr)
	return t
}

// Seed default MsLoker
func seedDefeaultMsLokerData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsLoker{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultMsLokerData := []model.MsLoker{
			{LockerID: 1, RowNumber: 1, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LockerID: 2, RowNumber: 1, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LockerID: 3, RowNumber: 1, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LockerID: 4, RowNumber: 1, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LockerID: 5, RowNumber: 1, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LockerID: 6, RowNumber: 1, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LockerID: 7, RowNumber: 1, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LockerID: 8, RowNumber: 1, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LockerID: 9, RowNumber: 1, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LockerID: 10, RowNumber: 1, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LockerID: 11, RowNumber: 1, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LockerID: 12, RowNumber: 1, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LockerID: 13, RowNumber: 1, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LockerID: 14, RowNumber: 1, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LockerID: 15, RowNumber: 1, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LockerID: 16, RowNumber: 1, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LockerID: 17, RowNumber: 1, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LockerID: 18, RowNumber: 1, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LockerID: 19, RowNumber: 1, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LockerID: 20, RowNumber: 1, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LockerID: 21, RowNumber: 1, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LockerID: 22, RowNumber: 1, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LockerID: 23, RowNumber: 1, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LockerID: 24, RowNumber: 1, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LockerID: 25, RowNumber: 1, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LockerID: 26, RowNumber: 1, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LockerID: 27, RowNumber: 1, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LockerID: 28, RowNumber: 1, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LockerID: 29, RowNumber: 1, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LockerID: 30, RowNumber: 1, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LockerID: 31, RowNumber: 1, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LockerID: 32, RowNumber: 1, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LockerID: 33, RowNumber: 1, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LockerID: 34, RowNumber: 1, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LockerID: 35, RowNumber: 1, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LockerID: 36, RowNumber: 1, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LockerID: 37, RowNumber: 1, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LockerID: 38, RowNumber: 1, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LockerID: 39, RowNumber: 1, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LockerID: 40, RowNumber: 1, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LockerID: 41, RowNumber: 1, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LockerID: 42, RowNumber: 1, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LockerID: 43, RowNumber: 1, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LockerID: 44, RowNumber: 1, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LockerID: 45, RowNumber: 1, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LockerID: 46, RowNumber: 1, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LockerID: 47, RowNumber: 1, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LockerID: 48, RowNumber: 1, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LockerID: 49, RowNumber: 1, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LockerID: 50, RowNumber: 1, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LockerID: 51, RowNumber: 1, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LockerID: 52, RowNumber: 1, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LockerID: 53, RowNumber: 1, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LockerID: 54, RowNumber: 1, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LockerID: 55, RowNumber: 1, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LockerID: 56, RowNumber: 1, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LockerID: 57, RowNumber: 1, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LockerID: 58, RowNumber: 1, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LockerID: 59, RowNumber: 1, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LockerID: 60, RowNumber: 1, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LockerID: 61, RowNumber: 2, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LockerID: 62, RowNumber: 2, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LockerID: 63, RowNumber: 2, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LockerID: 64, RowNumber: 2, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LockerID: 65, RowNumber: 2, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LockerID: 66, RowNumber: 2, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LockerID: 67, RowNumber: 2, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LockerID: 68, RowNumber: 2, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LockerID: 69, RowNumber: 2, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LockerID: 70, RowNumber: 2, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LockerID: 71, RowNumber: 2, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LockerID: 72, RowNumber: 2, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LockerID: 73, RowNumber: 2, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LockerID: 74, RowNumber: 2, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LockerID: 75, RowNumber: 2, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LockerID: 76, RowNumber: 2, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LockerID: 77, RowNumber: 2, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LockerID: 78, RowNumber: 2, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LockerID: 79, RowNumber: 2, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LockerID: 80, RowNumber: 2, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LockerID: 81, RowNumber: 2, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LockerID: 82, RowNumber: 2, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LockerID: 83, RowNumber: 2, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LockerID: 84, RowNumber: 2, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LockerID: 85, RowNumber: 2, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LockerID: 86, RowNumber: 2, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LockerID: 87, RowNumber: 2, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LockerID: 88, RowNumber: 2, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LockerID: 89, RowNumber: 2, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LockerID: 90, RowNumber: 2, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LockerID: 91, RowNumber: 2, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LockerID: 92, RowNumber: 2, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LockerID: 93, RowNumber: 2, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LockerID: 94, RowNumber: 2, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LockerID: 95, RowNumber: 2, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LockerID: 96, RowNumber: 2, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LockerID: 97, RowNumber: 2, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LockerID: 98, RowNumber: 2, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LockerID: 99, RowNumber: 2, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LockerID: 100, RowNumber: 2, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LockerID: 101, RowNumber: 2, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LockerID: 102, RowNumber: 2, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LockerID: 103, RowNumber: 2, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LockerID: 104, RowNumber: 2, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LockerID: 105, RowNumber: 2, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LockerID: 106, RowNumber: 2, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LockerID: 107, RowNumber: 2, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LockerID: 108, RowNumber: 2, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LockerID: 109, RowNumber: 2, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LockerID: 110, RowNumber: 2, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LockerID: 111, RowNumber: 2, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LockerID: 112, RowNumber: 2, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LockerID: 113, RowNumber: 2, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LockerID: 114, RowNumber: 2, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LockerID: 115, RowNumber: 2, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LockerID: 116, RowNumber: 2, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LockerID: 117, RowNumber: 2, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LockerID: 118, RowNumber: 2, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LockerID: 119, RowNumber: 2, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LockerID: 120, RowNumber: 2, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LockerID: 121, RowNumber: 3, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LockerID: 122, RowNumber: 3, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LockerID: 123, RowNumber: 3, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LockerID: 124, RowNumber: 3, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LockerID: 125, RowNumber: 3, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LockerID: 126, RowNumber: 3, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LockerID: 127, RowNumber: 3, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LockerID: 128, RowNumber: 3, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LockerID: 129, RowNumber: 3, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LockerID: 130, RowNumber: 3, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LockerID: 131, RowNumber: 3, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LockerID: 132, RowNumber: 3, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LockerID: 133, RowNumber: 3, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LockerID: 134, RowNumber: 3, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LockerID: 135, RowNumber: 3, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LockerID: 136, RowNumber: 3, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LockerID: 137, RowNumber: 3, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LockerID: 138, RowNumber: 3, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LockerID: 139, RowNumber: 3, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LockerID: 140, RowNumber: 3, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LockerID: 141, RowNumber: 3, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LockerID: 142, RowNumber: 3, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LockerID: 143, RowNumber: 3, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LockerID: 144, RowNumber: 3, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LockerID: 145, RowNumber: 3, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LockerID: 146, RowNumber: 3, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LockerID: 147, RowNumber: 3, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LockerID: 148, RowNumber: 3, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LockerID: 149, RowNumber: 3, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LockerID: 150, RowNumber: 3, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LockerID: 151, RowNumber: 3, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LockerID: 152, RowNumber: 3, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LockerID: 153, RowNumber: 3, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LockerID: 154, RowNumber: 3, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LockerID: 155, RowNumber: 3, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LockerID: 156, RowNumber: 3, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LockerID: 157, RowNumber: 3, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LockerID: 158, RowNumber: 3, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LockerID: 159, RowNumber: 3, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LockerID: 160, RowNumber: 3, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LockerID: 161, RowNumber: 3, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LockerID: 162, RowNumber: 3, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LockerID: 163, RowNumber: 3, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LockerID: 164, RowNumber: 3, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LockerID: 165, RowNumber: 3, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LockerID: 166, RowNumber: 3, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LockerID: 167, RowNumber: 3, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LockerID: 168, RowNumber: 3, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LockerID: 169, RowNumber: 3, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LockerID: 170, RowNumber: 3, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LockerID: 171, RowNumber: 3, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LockerID: 172, RowNumber: 3, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LockerID: 173, RowNumber: 3, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LockerID: 174, RowNumber: 3, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LockerID: 175, RowNumber: 3, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LockerID: 176, RowNumber: 3, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LockerID: 177, RowNumber: 3, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LockerID: 178, RowNumber: 3, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LockerID: 179, RowNumber: 3, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LockerID: 180, RowNumber: 3, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LockerID: 181, RowNumber: 4, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LockerID: 182, RowNumber: 4, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LockerID: 183, RowNumber: 4, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LockerID: 184, RowNumber: 4, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LockerID: 185, RowNumber: 4, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LockerID: 186, RowNumber: 4, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LockerID: 187, RowNumber: 4, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LockerID: 188, RowNumber: 4, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LockerID: 189, RowNumber: 4, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LockerID: 190, RowNumber: 4, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LockerID: 191, RowNumber: 4, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LockerID: 192, RowNumber: 4, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LockerID: 193, RowNumber: 4, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LockerID: 194, RowNumber: 4, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LockerID: 195, RowNumber: 4, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LockerID: 196, RowNumber: 4, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LockerID: 197, RowNumber: 4, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LockerID: 198, RowNumber: 4, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LockerID: 199, RowNumber: 4, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LockerID: 200, RowNumber: 4, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LockerID: 201, RowNumber: 4, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LockerID: 202, RowNumber: 4, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LockerID: 203, RowNumber: 4, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LockerID: 204, RowNumber: 4, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LockerID: 205, RowNumber: 4, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LockerID: 206, RowNumber: 4, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LockerID: 207, RowNumber: 4, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LockerID: 208, RowNumber: 4, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LockerID: 209, RowNumber: 4, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LockerID: 210, RowNumber: 4, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LockerID: 211, RowNumber: 4, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LockerID: 212, RowNumber: 4, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LockerID: 213, RowNumber: 4, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LockerID: 214, RowNumber: 4, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LockerID: 215, RowNumber: 4, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LockerID: 216, RowNumber: 4, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LockerID: 217, RowNumber: 4, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LockerID: 218, RowNumber: 4, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LockerID: 219, RowNumber: 4, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LockerID: 220, RowNumber: 4, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LockerID: 221, RowNumber: 4, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LockerID: 222, RowNumber: 4, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LockerID: 223, RowNumber: 4, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LockerID: 224, RowNumber: 4, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LockerID: 225, RowNumber: 4, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LockerID: 226, RowNumber: 4, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LockerID: 227, RowNumber: 4, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LockerID: 228, RowNumber: 4, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LockerID: 229, RowNumber: 4, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LockerID: 230, RowNumber: 4, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LockerID: 231, RowNumber: 4, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LockerID: 232, RowNumber: 4, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LockerID: 233, RowNumber: 4, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LockerID: 234, RowNumber: 4, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LockerID: 235, RowNumber: 4, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LockerID: 236, RowNumber: 4, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LockerID: 237, RowNumber: 4, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LockerID: 238, RowNumber: 4, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LockerID: 239, RowNumber: 4, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LockerID: 240, RowNumber: 4, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LockerID: 241, RowNumber: 5, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LockerID: 242, RowNumber: 5, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LockerID: 243, RowNumber: 5, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LockerID: 244, RowNumber: 5, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LockerID: 245, RowNumber: 5, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LockerID: 246, RowNumber: 5, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LockerID: 247, RowNumber: 5, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LockerID: 248, RowNumber: 5, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LockerID: 249, RowNumber: 5, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LockerID: 250, RowNumber: 5, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LockerID: 251, RowNumber: 5, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LockerID: 252, RowNumber: 5, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LockerID: 253, RowNumber: 5, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LockerID: 254, RowNumber: 5, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LockerID: 255, RowNumber: 5, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LockerID: 256, RowNumber: 5, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LockerID: 257, RowNumber: 5, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LockerID: 258, RowNumber: 5, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LockerID: 259, RowNumber: 5, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LockerID: 260, RowNumber: 5, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LockerID: 261, RowNumber: 5, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LockerID: 262, RowNumber: 5, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LockerID: 263, RowNumber: 5, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LockerID: 264, RowNumber: 5, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LockerID: 265, RowNumber: 5, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LockerID: 266, RowNumber: 5, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LockerID: 267, RowNumber: 5, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LockerID: 268, RowNumber: 5, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LockerID: 269, RowNumber: 5, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LockerID: 270, RowNumber: 5, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LockerID: 271, RowNumber: 5, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LockerID: 272, RowNumber: 5, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LockerID: 273, RowNumber: 5, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LockerID: 274, RowNumber: 5, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LockerID: 275, RowNumber: 5, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LockerID: 276, RowNumber: 5, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LockerID: 277, RowNumber: 5, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LockerID: 278, RowNumber: 5, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LockerID: 279, RowNumber: 5, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LockerID: 280, RowNumber: 5, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LockerID: 281, RowNumber: 5, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LockerID: 282, RowNumber: 5, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LockerID: 283, RowNumber: 5, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LockerID: 284, RowNumber: 5, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LockerID: 285, RowNumber: 5, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LockerID: 286, RowNumber: 5, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LockerID: 287, RowNumber: 5, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LockerID: 288, RowNumber: 5, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LockerID: 289, RowNumber: 5, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LockerID: 290, RowNumber: 5, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LockerID: 291, RowNumber: 5, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LockerID: 292, RowNumber: 5, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LockerID: 293, RowNumber: 5, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LockerID: 294, RowNumber: 5, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
		}
		for _, data := range defaultMsLokerData {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func seedDefaultMsUsersData(db *gorm.DB) error {
	// Inserting data to MsUser for Mahasiswa
	var mahasiswas []model.MsMahasiswa
	if err := db.Find(&mahasiswas).Error; err != nil {
		return err
	}

	for _, mahasiswa := range mahasiswas {
		var count int64
		err := db.Table("ms_users").Where("nim = ?", mahasiswa.NIM).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			if err := db.Exec(`
            INSERT INTO ms_users ("user_id", "nim", "created_at", "updated_at", "stsrc")
            SELECT gen_random_uuid(), ?, ?, ?, 'A'`, mahasiswa.NIM, time.Now(), time.Now()).Error; err != nil {
				return err
			}
		}
	}

	// Inserting data to MsUser for Staff
	var staffs []model.MsStaff
	if err := db.Find(&staffs).Error; err != nil {
		return err
	}

	for _, staff := range staffs {
		var count int64
		err := db.Table("ms_users").Where("nis = ?", staff.NIS).Count(&count).Error
		if err != nil {
			return err
		}
		if count == 0 {
			if err := db.Exec(`
            INSERT INTO ms_users ("user_id", "nis", "created_at", "updated_at", "stsrc")
            SELECT gen_random_uuid(), ?, ?, ?, 'A'`, staff.NIS, time.Now(), time.Now()).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func seedDefaultAcara(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsAcara{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultAcaraData := []model.MsAcara{
			{
				AcaraID:        1,
				AcaraName:      "Sosialisasi BINUS MAYA Avatar",
				AcaraStartTime: parseTime("09:00:00"),
				AcaraEndTime:   parseTime("11:00:00"),
				AcaraDate:      time.Date(2024, time.May, 15, 0, 0, 0, 0, time.UTC),
				AcaraLocation:  "LKC Refreshment Room",
				AcaraDetails:   "Halo BINUSIANS! Dalam BINUS MAYA versi 3.0, kamu akan memiliki kesempatan untuk mengekspresikan diri dengan lebih unik dan personal melalui avatar kustom yang dapat kamu buat sendiri.",
				SpeakerName:    "Kanyadian Idananta, S.Kom., M.TI",
				RegisterLink:   "https://forms.gle/ZyXPmn8Pg6xccUFKA",
				AcaraImage:     "https://binus.ac.id/wp-content/uploads/2021/02/217-0-Binusacid.jpg",
				Stsrc:          "A",
			},
			{
				AcaraID:        2,
				AcaraName:      "Road to PILMAPRES BINUS University 2024",
				AcaraStartTime: parseTime("09:00:00"),
				AcaraEndTime:   parseTime("11:00:00"),
				AcaraDate:      time.Date(2024, time.May, 18, 0, 0, 0, 0, time.UTC),
				AcaraLocation:  "LKC Refreshment Room",
				AcaraDetails:   "Pemilihan Mahasiswa Berprestasi (PILMAPRES) merupakan kompetisi Mahasiswa yang diselenggarakan oleh Pusat Prestasi Nasional yang berada di bawah naungan Kementerian Riset, Teknologi, dan Pendidikan Tinggi (KEMENRISTEKDIKTI) setiap tahunnya.",
				SpeakerName:    "Herru Darmadi, S.Kom., M.TI",
				RegisterLink:   "https://forms.gle/LCJemRV7pXjUryuMA",
				AcaraImage:     "https://student.binus.ac.id/wp-content/uploads/2024/03/Student-Website-Web-Banner.jpg",
				Stsrc:          "A",
			},
			{
				AcaraID:        3,
				AcaraName:      "Merancang Produk untuk Sustainable Development Goals di BINUS University",
				AcaraStartTime: parseTime("09:00:00"),
				AcaraEndTime:   parseTime("11:00:00"),
				AcaraDate:      time.Date(2024, time.May, 19, 0, 0, 0, 0, time.UTC),
				AcaraLocation:  "LKC Refreshment Room",
				AcaraDetails:   "BINUS University ikut berpartisipasi dan mendukung penuh dalam 17 Tujuan Pembangunan Berkelanjutan (Sustainable Development Goals/SDGs), yang merupakan seruan mendesak bagi semua negara – baik maju maupun berkembang – untuk melakukan tindakan dalam kemitraan global.",
				SpeakerName:    "Alvina Aulia, S.Kom., M.TI.",
				RegisterLink:   "https://forms.gle/YXxK98wU7qjQyuXv6",
				AcaraImage:     "https://student.binus.ac.id/wp-content/uploads/2024/04/WhatsApp-Image-2024-04-02-at-13.37.43.jpeg",
				Stsrc:          "A",
			},
			{
				AcaraID:        4,
				AcaraName:      "Unicharm Goes To BINUS University",
				AcaraStartTime: parseTime("09:00:00"),
				AcaraEndTime:   parseTime("11:00:00"),
				AcaraDate:      time.Date(2024, time.May, 21, 0, 0, 0, 0, time.UTC),
				AcaraLocation:  "LKC Refreshment Room",
				AcaraDetails:   "Dalam acara ini, Unicharm memperkenalkan produk-produk unggulannya yang dapat digunakan dari berbagai usia, mulai dari usia balita, remaja, lanjut usia hingga hewan peliharaan.",
				SpeakerName:    "Anak Agung Ayu Mirah Krisnawati, S.Sos., M.I.Kom",
				RegisterLink:   "https://forms.gle/nS8DvwEC3iFLKhsd6",
				AcaraImage:     "https://student.binus.ac.id/wp-content/uploads/2024/03/WhatsApp-Image-2024-03-08-at-1.52.10-PM.jpeg",
				Stsrc:          "A",
			},
		}
		for _, data := range defaultAcaraData {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
