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
	// Inserting data to MsUser
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
		}
		for _, data := range defaultMsLokerData {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
