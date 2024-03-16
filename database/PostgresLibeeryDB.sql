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

DROP TABLE MsMahasiswa;
DROP TABLE MsStaff;
DROP TABLE MsUser;
