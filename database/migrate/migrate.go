package migrate

import (
	"Libeery/model"
	"time"

	"gorm.io/gorm"
)

func DatabaseMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.MsMahasiswa{},
		&model.MsStaff{},
		&model.MsUser{},
		&model.TrBooking{},
		&model.MsBookingStatus{},
		&model.MsLoker{},
		&model.ForLaterSession{},
	); err != nil {
		return err
	}

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

	// Inserting data into MsUser from MsMahasiswa
	if err := db.Exec(`
		INSERT INTO Ms_Users ("user_id", "nim", "created_at", "updated_at", "stsrc")
		SELECT gen_random_uuid(), "nim", ?, ?, 'A'
		FROM Ms_Mahasiswas
		WHERE "nim" IS NOT NULL`, time.Now(), time.Now()).Error; err != nil {
		return err
	}

	// Inserting data into MsUser from MsStaff
	if err := db.Exec(`
		INSERT INTO Ms_Users ("user_id", "nis", "created_at", "updated_at", "stsrc")
		SELECT gen_random_uuid(), "nis", ?, ?, 'A'
		FROM Ms_Staffs
		WHERE "nis" IS NOT NULL`, time.Now(), time.Now()).Error; err != nil {
		return err
	}

	defaultForLaterSessions := []model.ForLaterSession{
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

	for _, session := range defaultForLaterSessions {
		if err := db.Create(&session).Error; err != nil {
			return err
		}
	}

	return nil
}
func parseTime(timeStr string) time.Time {
	layout := "15:04:05"
	t, _ := time.Parse(layout, timeStr)
	return t
}
