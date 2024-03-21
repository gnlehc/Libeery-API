-- psql -U admin -d Libeery-DB

-- CTRL + L to clear terminal in docker

CREATE TABLE MsMahasiswa(
    "NIM" CHAR(10) PRIMARY KEY,
    "MhsName" VARCHAR(50),
    "MhsPassword" VARCHAR(50),
    "Stsrc" CHAR(1)
);

INSERT INTO MsMahasiswa VALUES ('2602057652', 'Chelsea Ng', 'isAdmin24', 'A');
INSERT INTO MsMahasiswa VALUES ('2602063831', 'Verena Vynne Sentosa', 'isMahasiswa24', 'A');
INSERT INTO MsMahasiswa VALUES ('2602063560', 'Christopher Verrell', 'verinthebuilding', 'A');
INSERT INTO MsMahasiswa VALUES ('2602053515', 'Jesslyn Amanda Mulyawan', 'isAdmin13', 'A');
INSERT INTO MsMahasiswa VALUES ('2602062652', 'Nicholas Owen Sentosa', 'owen123', 'A');

SELECT * FROM MsMahasiswa;

CREATE TABLE MsStaff(
    "NIS" CHAR(5) PRIMARY KEY,
    "StaffName" VARCHAR(50),
    "StaffPassword" VARCHAR(50),
    "Stsrc" CHAR(1)
);

INSERT INTO MsStaff VALUES ('D5416', 'Kanyadian Idhananta', 'isDosen', 'A');
INSERT INTO MsStaff VALUES ('D6831', 'Islam Nur Alam', 'isAlam', 'A');
INSERT INTO MsStaff VALUES ('D6657', 'Anderies Anderies', 'isAnderies', 'A');
INSERT INTO MsStaff VALUES ('D6811', 'Faisal Asadi', 'isFaisal', 'A');
INSERT INTO MsStaff VALUES ('D6835', 'Hanis Amalia Saputri', 'isHanis', 'A');

SELECT * FROM MsStaff;

CREATE TABLE MsUser(
    "UserID" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- aggregation
    "NIM" CHAR(10),
    "NIS" CHAR(5),
    "CreatedAt" TIMESTAMP,
    "UpdatedAt" TIMESTAMP,
    "Stsrc" CHAR(1),
);

-- Inserting all data from MsMahasiswa into MsUser
INSERT INTO MsUser ("UserID", "NIM", "CreatedAt", "UpdatedAt", "Stsrc")
SELECT gen_random_uuid(), "NIM", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'A'
FROM MsMahasiswa;

-- Inserting all data from MsStaff into MsUser
INSERT INTO MsUser ("UserID", "NIS", "CreatedAt", "UpdatedAt", "Stsrc")
SELECT gen_random_uuid(), "NIS", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'A'
FROM MsStaff;

SELECT * FROM MsUser;

CREATE TABLE MsBookingStatus(
    BookingStatusID INT PRIMARY KEY,
    StatusTitle VARCHAR(255),
    Stsrc CHAR(1)
);

INSERT INTO MsBookingStatus (BookingStatusID, StatusTitle, Stsrc) VALUES (1, "Pending", "A");
INSERT INTO MsBookingStatus (BookingStatusID, StatusTitle, Stsrc) VALUES (2, "Checked-In", "A");
INSERT INTO MsBookingStatus (BookingStatusID, StatusTitle, Stsrc) VALUES (3, "Checked-Out", "A");

CREATE TABLE MsLoker(
    LockerID INT PRIMARY KEY,
    RowNumber INT,
    ColumnNumber INT
);

-- INSERT INTO MsLoker()

CREATE TABLE TrBooking(
    BookingID UUID PRIMARY KEY,
    UserID INT,
    BookingStatusID INT,
    LokerID INT,
    StartSession TIME,
    EndSession TIME,
    CheckInTime TIME, 
    CheckOutTime TIME, 
    CreatedAt TIME,
    UpdatedAt TIME,
    Stsrc CHAR(1),
    FOREIGN KEY (UserID) REFERENCES MsUser(UserID), 
    FOREIGN KEY (BookingStatusID) REFERENCES MsBookingStatus(BookingStatusID),
    FOREIGN KEY (LokerID) REFERENCES MsLoker(LockerID)
);

CREATE TABLE MsForLaterSession(
    SessionID INT PRIMARY KEY,
    StartSession TIME,
    EndSession TIME
);

INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (1, '08:00:00', '09:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (2, '09:00:00', '10:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (3, '10:00:00', '11:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (4, '11:00:00', '12:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (5, '12:00:00', '13:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (6, '13:00:00', '14:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (7, '14:00:00', '15:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (8, '15:00:00', '16:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (9, '16:00:00', '17:00:00');
INSERT INTO MsForLaterSession (SessionID, StartSession, EndSession) VALUES (10, '17:00:00', '18:00:00');

DROP TABLE MsMahasiswa;
DROP TABLE MsStaff;
DROP TABLE MsUser;
