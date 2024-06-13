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
	if !migrator.HasTable(&model.MsBook{}) {
		if err := db.AutoMigrate(&model.MsBook{}); err != nil {
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
	if err := seedDefaultBookData(db); err != nil {
		return err
	}

	return nil
}

func seedDefaultBookData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.MsBook{}).Count(&count).Error; err != nil {
		return err
	}
	if count <= 0 {
		defaultBookData := []model.MsBook{
			{
				BookID:    1,
				ISBN:      "9780143039001",
				Title:     "To Kill a Mockingbird",
				Author:    "Harper Lee",
				Publisher: "Harper Perennial Modern Classics",
				Edition:   "50th Anniversary",
				Year:      1960,
				Abstract:  "To Kill a Mockingbird is a novel by Harper Lee published in 1960. It explores themes of racial injustice, moral growth, and the loss of innocence in the American South during the 1930s. The story follows young Scout Finch and her brother Jem as their father, lawyer Atticus Finch, defends a black man falsely accused of raping a white woman. Through their experiences, the children learn about empathy, compassion, and the complexities of human nature.",
				Stock:     7,
				Photo:     "https://m.media-amazon.com/images/I/51IXWZzlgSL._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    2,
				ISBN:      "9780061120082",
				Title:     "Pride and Prejudice",
				Author:    "Jane Austen",
				Publisher: "HarperTeen",
				Edition:   "Reprint Edition",
				Year:      1813,
				Abstract:  "Pride and Prejudice is a romantic novel by Jane Austen, first published in 1813. It follows the story of Elizabeth Bennet, the spirited and independent protagonist, as she navigates the societal expectations and pressures of early 19th-century England. The novel explores themes of love, marriage, class, and reputation, and is known for its wit, irony, and memorable characters.",
				Stock:     2,
				Photo:     "https://m.media-amazon.com/images/I/51YxyvZCpXS._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    3,
				ISBN:      "9780743273563",
				Title:     "The Great Gatsby",
				Author:    "F. Scott Fitzgerald",
				Publisher: "Scribner",
				Edition:   "Reissue",
				Year:      1925,
				Abstract:  "The Great Gatsby is a novel by American author F. Scott Fitzgerald, first published in 1925. Set in the Roaring Twenties on Long Island, New York, the story follows the enigmatic Jay Gatsby and his pursuit of the elusive Daisy Buchanan. Through Gatsby's extravagant parties and romantic gestures, the novel explores themes of wealth, class, love, and the American Dream.",
				Stock:     9,
				Photo:     "https://m.media-amazon.com/images/I/515Ra7ttqIL._SX342_SY445_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    4,
				ISBN:      "9780451524935",
				Title:     "1984",
				Author:    "George Orwell",
				Publisher: "Signet Classic",
				Edition:   "Reissue",
				Year:      1949,
				Abstract:  "1984 is a dystopian novel by George Orwell published in 1949. It is set in a totalitarian society ruled by the oppressive Party led by Big Brother, where individualism and independent thought are suppressed. The novel follows protagonist Winston Smith as he rebels against the Party's control and seeks freedom and truth in a world of surveillance and propaganda.",
				Stock:     4,
				Photo:     "https://m.media-amazon.com/images/I/41aCCnK8p6L._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    5,
				ISBN:      "9780062561022",
				Title:     "The Catcher in the Rye",
				Author:    "J.D. Salinger",
				Publisher: "Little, Brown and Company",
				Edition:   "Reprint Edition",
				Year:      1951,
				Abstract:  "The Catcher in the Rye is a novel by J.D. Salinger, first published in 1951. It is narrated by Holden Caulfield, a disillusioned teenager who reflects on his experiences after being expelled from prep school. The novel explores themes of teenage angst, alienation, and the search for authenticity in a world perceived as phony by the protagonist.",
				Stock:     1,
				Photo:     "https://images-na.ssl-images-amazon.com/images/I/81OthjkJBuL.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    6,
				ISBN:      "9780143127550",
				Title:     "The Goldfinch",
				Author:    "Donna Tartt",
				Publisher: "Little, Brown and Company",
				Edition:   "Reprint Edition",
				Year:      2013,
				Abstract:  "The Goldfinch is a novel by Donna Tartt published in 2013. It follows the life of Theo Decker, a young boy who survives a terrorist attack at the Metropolitan Museum of Art in New York City. The novel explores themes of loss, trauma, art, and the search for meaning in the face of adversity.",
				Stock:     0,
				Photo:     "https://m.media-amazon.com/images/I/417FX4orgXL._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    7,
				ISBN:      "9780345803481",
				Title:     "The Bell Jar",
				Author:    "Sylvia Plath",
				Publisher: "Harper Perennial Modern Classics",
				Edition:   "Reprint Edition",
				Year:      1963,
				Abstract:  "The Bell Jar is a novel by Sylvia Plath, first published in 1963 under the pseudonym Victoria Lucas. It is a semi-autobiographical account of Plath's own struggles with mental illness, following protagonist Esther Greenwood's descent into depression and her experiences in a psychiatric hospital. The novel explores themes of identity, gender roles, and the pressures of societal expectations.",
				Stock:     0,
				Photo:     "https://m.media-amazon.com/images/I/415yzhIbkrL._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    8,
				ISBN:      "978140007998",
				Title:     "The Kite Runner",
				Author:    "Khaled Hosseini",
				Publisher: "Riverhead Books",
				Edition:   "Reprint Edition",
				Year:      2003,
				Abstract:  "The Kite Runner is a novel by Khaled Hosseini published in 2003. It follows the story of Amir, a young boy from Kabul, Afghanistan, and his journey to seek redemption for betraying his childhood friend Hassan. The novel explores themes of guilt, redemption, and the complex relationship between fathers and sons against the backdrop of the tumultuous events in Afghanistan from the fall of the monarchy to the rise of the Taliban regime.",
				Stock:     5,
				Photo:     "https://m.media-amazon.com/images/I/41fAt2RhwML._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    9,
				ISBN:      "9780679722762",
				Title:     "Beloved",
				Author:    "Toni Morrison",
				Publisher: "Vintage Books",
				Edition:   "Reprint Edition",
				Year:      1987,
				Abstract:  "Beloved is a novel by Toni Morrison published in 1987. Set after the American Civil War, it tells the story of Sethe, a former slave who escaped to Ohio, but is haunted by the memories of her past and the ghost of her deceased daughter. The novel explores themes of slavery, trauma, and the struggle for freedom, and employs a unique narrative style that blends realism with elements of magic and folklore.",
				Stock:     1,
				Photo:     "https://m.media-amazon.com/images/I/31xXfVRPC4L._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
			{
				BookID:    10,
				ISBN:      "9780446675536",
				Title:     "The Alchemist",
				Author:    "Paulo Coelho",
				Publisher: "HarperOne",
				Edition:   "Reissue",
				Year:      1988,
				Abstract:  "The Alchemist is a novel by Brazilian author Paulo Coelho, first published in 1988. It follows the journey of Santiago, a young Andalusian shepherd, as he travels from Spain to Egypt in search of a hidden treasure. Along the way, Santiago encounters various characters who impart wisdom and lessons about life, destiny, and the pursuit of one's dreams. The novel explores themes of self-discovery, personal legend, and the interconnectedness of all things.",
				Stock:     2,
				Photo:     "https://m.media-amazon.com/images/I/51YsnEoNr-L._SY445_SX342_.jpg",
				Stsrc:     "A",
			},
		}
		for _, data := range defaultBookData {
			if err := db.Create(&data).Error; err != nil {
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
			{LokerID: 1, RowNumber: 1, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LokerID: 2, RowNumber: 1, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LokerID: 3, RowNumber: 1, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LokerID: 4, RowNumber: 1, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LokerID: 5, RowNumber: 1, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LokerID: 6, RowNumber: 1, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LokerID: 7, RowNumber: 1, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LokerID: 8, RowNumber: 1, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LokerID: 9, RowNumber: 1, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LokerID: 10, RowNumber: 1, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LokerID: 11, RowNumber: 1, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LokerID: 12, RowNumber: 1, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LokerID: 13, RowNumber: 1, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LokerID: 14, RowNumber: 1, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LokerID: 15, RowNumber: 1, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LokerID: 16, RowNumber: 1, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LokerID: 17, RowNumber: 1, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LokerID: 18, RowNumber: 1, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LokerID: 19, RowNumber: 1, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LokerID: 20, RowNumber: 1, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LokerID: 21, RowNumber: 1, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LokerID: 22, RowNumber: 1, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LokerID: 23, RowNumber: 1, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LokerID: 24, RowNumber: 1, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LokerID: 25, RowNumber: 1, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LokerID: 26, RowNumber: 1, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LokerID: 27, RowNumber: 1, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LokerID: 28, RowNumber: 1, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LokerID: 29, RowNumber: 1, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LokerID: 30, RowNumber: 1, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LokerID: 31, RowNumber: 1, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LokerID: 32, RowNumber: 1, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LokerID: 33, RowNumber: 1, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LokerID: 34, RowNumber: 1, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LokerID: 35, RowNumber: 1, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LokerID: 36, RowNumber: 1, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LokerID: 37, RowNumber: 1, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LokerID: 38, RowNumber: 1, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LokerID: 39, RowNumber: 1, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LokerID: 40, RowNumber: 1, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LokerID: 41, RowNumber: 1, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LokerID: 42, RowNumber: 1, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LokerID: 43, RowNumber: 1, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LokerID: 44, RowNumber: 1, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LokerID: 45, RowNumber: 1, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LokerID: 46, RowNumber: 1, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LokerID: 47, RowNumber: 1, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LokerID: 48, RowNumber: 1, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LokerID: 49, RowNumber: 1, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LokerID: 50, RowNumber: 1, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LokerID: 51, RowNumber: 1, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LokerID: 52, RowNumber: 1, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LokerID: 53, RowNumber: 1, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LokerID: 54, RowNumber: 1, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LokerID: 55, RowNumber: 1, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LokerID: 56, RowNumber: 1, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LokerID: 57, RowNumber: 1, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LokerID: 58, RowNumber: 1, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LokerID: 59, RowNumber: 1, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LokerID: 60, RowNumber: 1, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LokerID: 61, RowNumber: 2, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LokerID: 62, RowNumber: 2, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LokerID: 63, RowNumber: 2, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LokerID: 64, RowNumber: 2, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LokerID: 65, RowNumber: 2, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LokerID: 66, RowNumber: 2, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LokerID: 67, RowNumber: 2, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LokerID: 68, RowNumber: 2, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LokerID: 69, RowNumber: 2, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LokerID: 70, RowNumber: 2, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LokerID: 71, RowNumber: 2, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LokerID: 72, RowNumber: 2, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LokerID: 73, RowNumber: 2, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LokerID: 74, RowNumber: 2, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LokerID: 75, RowNumber: 2, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LokerID: 76, RowNumber: 2, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LokerID: 77, RowNumber: 2, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LokerID: 78, RowNumber: 2, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LokerID: 79, RowNumber: 2, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LokerID: 80, RowNumber: 2, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LokerID: 81, RowNumber: 2, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LokerID: 82, RowNumber: 2, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LokerID: 83, RowNumber: 2, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LokerID: 84, RowNumber: 2, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LokerID: 85, RowNumber: 2, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LokerID: 86, RowNumber: 2, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LokerID: 87, RowNumber: 2, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LokerID: 88, RowNumber: 2, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LokerID: 89, RowNumber: 2, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LokerID: 90, RowNumber: 2, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LokerID: 91, RowNumber: 2, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LokerID: 92, RowNumber: 2, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LokerID: 93, RowNumber: 2, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LokerID: 94, RowNumber: 2, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LokerID: 95, RowNumber: 2, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LokerID: 96, RowNumber: 2, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LokerID: 97, RowNumber: 2, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LokerID: 98, RowNumber: 2, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LokerID: 99, RowNumber: 2, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LokerID: 100, RowNumber: 2, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LokerID: 101, RowNumber: 2, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LokerID: 102, RowNumber: 2, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LokerID: 103, RowNumber: 2, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LokerID: 104, RowNumber: 2, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LokerID: 105, RowNumber: 2, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LokerID: 106, RowNumber: 2, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LokerID: 107, RowNumber: 2, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LokerID: 108, RowNumber: 2, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LokerID: 109, RowNumber: 2, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LokerID: 110, RowNumber: 2, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LokerID: 111, RowNumber: 2, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LokerID: 112, RowNumber: 2, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LokerID: 113, RowNumber: 2, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LokerID: 114, RowNumber: 2, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LokerID: 115, RowNumber: 2, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LokerID: 116, RowNumber: 2, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LokerID: 117, RowNumber: 2, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LokerID: 118, RowNumber: 2, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LokerID: 119, RowNumber: 2, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LokerID: 120, RowNumber: 2, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LokerID: 121, RowNumber: 3, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LokerID: 122, RowNumber: 3, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LokerID: 123, RowNumber: 3, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LokerID: 124, RowNumber: 3, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LokerID: 125, RowNumber: 3, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LokerID: 126, RowNumber: 3, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LokerID: 127, RowNumber: 3, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LokerID: 128, RowNumber: 3, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LokerID: 129, RowNumber: 3, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LokerID: 130, RowNumber: 3, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LokerID: 131, RowNumber: 3, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LokerID: 132, RowNumber: 3, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LokerID: 133, RowNumber: 3, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LokerID: 134, RowNumber: 3, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LokerID: 135, RowNumber: 3, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LokerID: 136, RowNumber: 3, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LokerID: 137, RowNumber: 3, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LokerID: 138, RowNumber: 3, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LokerID: 139, RowNumber: 3, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LokerID: 140, RowNumber: 3, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LokerID: 141, RowNumber: 3, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LokerID: 142, RowNumber: 3, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LokerID: 143, RowNumber: 3, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LokerID: 144, RowNumber: 3, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LokerID: 145, RowNumber: 3, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LokerID: 146, RowNumber: 3, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LokerID: 147, RowNumber: 3, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LokerID: 148, RowNumber: 3, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LokerID: 149, RowNumber: 3, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LokerID: 150, RowNumber: 3, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LokerID: 151, RowNumber: 3, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LokerID: 152, RowNumber: 3, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LokerID: 153, RowNumber: 3, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LokerID: 154, RowNumber: 3, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LokerID: 155, RowNumber: 3, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LokerID: 156, RowNumber: 3, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LokerID: 157, RowNumber: 3, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LokerID: 158, RowNumber: 3, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LokerID: 159, RowNumber: 3, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LokerID: 160, RowNumber: 3, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LokerID: 161, RowNumber: 3, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LokerID: 162, RowNumber: 3, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LokerID: 163, RowNumber: 3, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LokerID: 164, RowNumber: 3, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LokerID: 165, RowNumber: 3, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LokerID: 166, RowNumber: 3, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LokerID: 167, RowNumber: 3, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LokerID: 168, RowNumber: 3, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LokerID: 169, RowNumber: 3, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LokerID: 170, RowNumber: 3, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LokerID: 171, RowNumber: 3, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LokerID: 172, RowNumber: 3, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LokerID: 173, RowNumber: 3, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LokerID: 174, RowNumber: 3, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LokerID: 175, RowNumber: 3, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LokerID: 176, RowNumber: 3, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LokerID: 177, RowNumber: 3, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LokerID: 178, RowNumber: 3, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LokerID: 179, RowNumber: 3, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LokerID: 180, RowNumber: 3, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LokerID: 181, RowNumber: 4, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LokerID: 182, RowNumber: 4, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LokerID: 183, RowNumber: 4, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LokerID: 184, RowNumber: 4, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LokerID: 185, RowNumber: 4, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LokerID: 186, RowNumber: 4, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LokerID: 187, RowNumber: 4, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LokerID: 188, RowNumber: 4, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LokerID: 189, RowNumber: 4, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LokerID: 190, RowNumber: 4, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LokerID: 191, RowNumber: 4, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LokerID: 192, RowNumber: 4, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LokerID: 193, RowNumber: 4, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LokerID: 194, RowNumber: 4, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LokerID: 195, RowNumber: 4, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LokerID: 196, RowNumber: 4, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LokerID: 197, RowNumber: 4, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LokerID: 198, RowNumber: 4, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LokerID: 199, RowNumber: 4, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LokerID: 200, RowNumber: 4, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LokerID: 201, RowNumber: 4, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LokerID: 202, RowNumber: 4, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LokerID: 203, RowNumber: 4, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LokerID: 204, RowNumber: 4, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LokerID: 205, RowNumber: 4, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LokerID: 206, RowNumber: 4, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LokerID: 207, RowNumber: 4, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LokerID: 208, RowNumber: 4, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LokerID: 209, RowNumber: 4, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LokerID: 210, RowNumber: 4, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LokerID: 211, RowNumber: 4, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LokerID: 212, RowNumber: 4, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LokerID: 213, RowNumber: 4, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LokerID: 214, RowNumber: 4, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LokerID: 215, RowNumber: 4, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LokerID: 216, RowNumber: 4, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LokerID: 217, RowNumber: 4, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LokerID: 218, RowNumber: 4, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LokerID: 219, RowNumber: 4, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LokerID: 220, RowNumber: 4, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LokerID: 221, RowNumber: 4, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LokerID: 222, RowNumber: 4, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LokerID: 223, RowNumber: 4, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LokerID: 224, RowNumber: 4, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LokerID: 225, RowNumber: 4, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LokerID: 226, RowNumber: 4, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LokerID: 227, RowNumber: 4, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LokerID: 228, RowNumber: 4, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LokerID: 229, RowNumber: 4, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LokerID: 230, RowNumber: 4, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LokerID: 231, RowNumber: 4, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LokerID: 232, RowNumber: 4, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LokerID: 233, RowNumber: 4, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LokerID: 234, RowNumber: 4, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
			{LokerID: 235, RowNumber: 4, ColumnNumber: 55, Availability: "Active", Stsrc: "A"},
			{LokerID: 236, RowNumber: 4, ColumnNumber: 56, Availability: "Active", Stsrc: "A"},
			{LokerID: 237, RowNumber: 4, ColumnNumber: 57, Availability: "Active", Stsrc: "A"},
			{LokerID: 238, RowNumber: 4, ColumnNumber: 58, Availability: "Active", Stsrc: "A"},
			{LokerID: 239, RowNumber: 4, ColumnNumber: 59, Availability: "Active", Stsrc: "A"},
			{LokerID: 240, RowNumber: 4, ColumnNumber: 60, Availability: "Active", Stsrc: "A"},
			{LokerID: 241, RowNumber: 5, ColumnNumber: 1, Availability: "Active", Stsrc: "A"},
			{LokerID: 242, RowNumber: 5, ColumnNumber: 2, Availability: "Active", Stsrc: "A"},
			{LokerID: 243, RowNumber: 5, ColumnNumber: 3, Availability: "Active", Stsrc: "A"},
			{LokerID: 244, RowNumber: 5, ColumnNumber: 4, Availability: "Active", Stsrc: "A"},
			{LokerID: 245, RowNumber: 5, ColumnNumber: 5, Availability: "Active", Stsrc: "A"},
			{LokerID: 246, RowNumber: 5, ColumnNumber: 6, Availability: "Active", Stsrc: "A"},
			{LokerID: 247, RowNumber: 5, ColumnNumber: 7, Availability: "Active", Stsrc: "A"},
			{LokerID: 248, RowNumber: 5, ColumnNumber: 8, Availability: "Active", Stsrc: "A"},
			{LokerID: 249, RowNumber: 5, ColumnNumber: 9, Availability: "Active", Stsrc: "A"},
			{LokerID: 250, RowNumber: 5, ColumnNumber: 10, Availability: "Active", Stsrc: "A"},
			{LokerID: 251, RowNumber: 5, ColumnNumber: 11, Availability: "Active", Stsrc: "A"},
			{LokerID: 252, RowNumber: 5, ColumnNumber: 12, Availability: "Active", Stsrc: "A"},
			{LokerID: 253, RowNumber: 5, ColumnNumber: 13, Availability: "Active", Stsrc: "A"},
			{LokerID: 254, RowNumber: 5, ColumnNumber: 14, Availability: "Active", Stsrc: "A"},
			{LokerID: 255, RowNumber: 5, ColumnNumber: 15, Availability: "Active", Stsrc: "A"},
			{LokerID: 256, RowNumber: 5, ColumnNumber: 16, Availability: "Active", Stsrc: "A"},
			{LokerID: 257, RowNumber: 5, ColumnNumber: 17, Availability: "Active", Stsrc: "A"},
			{LokerID: 258, RowNumber: 5, ColumnNumber: 18, Availability: "Active", Stsrc: "A"},
			{LokerID: 259, RowNumber: 5, ColumnNumber: 19, Availability: "Active", Stsrc: "A"},
			{LokerID: 260, RowNumber: 5, ColumnNumber: 20, Availability: "Active", Stsrc: "A"},
			{LokerID: 261, RowNumber: 5, ColumnNumber: 21, Availability: "Active", Stsrc: "A"},
			{LokerID: 262, RowNumber: 5, ColumnNumber: 22, Availability: "Active", Stsrc: "A"},
			{LokerID: 263, RowNumber: 5, ColumnNumber: 23, Availability: "Active", Stsrc: "A"},
			{LokerID: 264, RowNumber: 5, ColumnNumber: 24, Availability: "Active", Stsrc: "A"},
			{LokerID: 265, RowNumber: 5, ColumnNumber: 25, Availability: "Active", Stsrc: "A"},
			{LokerID: 266, RowNumber: 5, ColumnNumber: 26, Availability: "Active", Stsrc: "A"},
			{LokerID: 267, RowNumber: 5, ColumnNumber: 27, Availability: "Active", Stsrc: "A"},
			{LokerID: 268, RowNumber: 5, ColumnNumber: 28, Availability: "Active", Stsrc: "A"},
			{LokerID: 269, RowNumber: 5, ColumnNumber: 29, Availability: "Active", Stsrc: "A"},
			{LokerID: 270, RowNumber: 5, ColumnNumber: 30, Availability: "Active", Stsrc: "A"},
			{LokerID: 271, RowNumber: 5, ColumnNumber: 31, Availability: "Active", Stsrc: "A"},
			{LokerID: 272, RowNumber: 5, ColumnNumber: 32, Availability: "Active", Stsrc: "A"},
			{LokerID: 273, RowNumber: 5, ColumnNumber: 33, Availability: "Active", Stsrc: "A"},
			{LokerID: 274, RowNumber: 5, ColumnNumber: 34, Availability: "Active", Stsrc: "A"},
			{LokerID: 275, RowNumber: 5, ColumnNumber: 35, Availability: "Active", Stsrc: "A"},
			{LokerID: 276, RowNumber: 5, ColumnNumber: 36, Availability: "Active", Stsrc: "A"},
			{LokerID: 277, RowNumber: 5, ColumnNumber: 37, Availability: "Active", Stsrc: "A"},
			{LokerID: 278, RowNumber: 5, ColumnNumber: 38, Availability: "Active", Stsrc: "A"},
			{LokerID: 279, RowNumber: 5, ColumnNumber: 39, Availability: "Active", Stsrc: "A"},
			{LokerID: 280, RowNumber: 5, ColumnNumber: 40, Availability: "Active", Stsrc: "A"},
			{LokerID: 281, RowNumber: 5, ColumnNumber: 41, Availability: "Active", Stsrc: "A"},
			{LokerID: 282, RowNumber: 5, ColumnNumber: 42, Availability: "Active", Stsrc: "A"},
			{LokerID: 283, RowNumber: 5, ColumnNumber: 43, Availability: "Active", Stsrc: "A"},
			{LokerID: 284, RowNumber: 5, ColumnNumber: 44, Availability: "Active", Stsrc: "A"},
			{LokerID: 285, RowNumber: 5, ColumnNumber: 45, Availability: "Active", Stsrc: "A"},
			{LokerID: 286, RowNumber: 5, ColumnNumber: 46, Availability: "Active", Stsrc: "A"},
			{LokerID: 287, RowNumber: 5, ColumnNumber: 47, Availability: "Active", Stsrc: "A"},
			{LokerID: 288, RowNumber: 5, ColumnNumber: 48, Availability: "Active", Stsrc: "A"},
			{LokerID: 289, RowNumber: 5, ColumnNumber: 49, Availability: "Active", Stsrc: "A"},
			{LokerID: 290, RowNumber: 5, ColumnNumber: 50, Availability: "Active", Stsrc: "A"},
			{LokerID: 291, RowNumber: 5, ColumnNumber: 51, Availability: "Active", Stsrc: "A"},
			{LokerID: 292, RowNumber: 5, ColumnNumber: 52, Availability: "Active", Stsrc: "A"},
			{LokerID: 293, RowNumber: 5, ColumnNumber: 53, Availability: "Active", Stsrc: "A"},
			{LokerID: 294, RowNumber: 5, ColumnNumber: 54, Availability: "Active", Stsrc: "A"},
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
