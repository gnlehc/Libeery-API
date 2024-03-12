CREATE DATABASE LibeeryDB

USE LibeeryDB

CREATE TABLE MsMahasiswa(
	[NIM] CHAR(10) PRIMARY KEY,
	[MhsName] VARCHAR(50),
	[Password] VARCHAR(50)
);

INSERT INTO MsMahasiswa VALUES ('2602057652', 'Chelsea Ng', 'isAdmin24');
INSERT INTO MsMahasiswa VALUES ('2602063831', 'Verena Vynne Sentosa', 'isMahasiswa24');
INSERT INTO MsMahasiswa VALUES ('2602063560', 'Christopher Verrell', 'verinthebuilding');
INSERT INTO MsMahasiswa VALUES ('2602053515', 'Jesslyn Amanda Mulyawan', 'isAdmin13');
INSERT INTO MsMahasiswa VALUES ('2602062652', 'Nicholas Owen Sentosa', 'owen123');

SELECT * FROM MsMahasiswa;

CREATE TABLE MsStaff(
	[NIS] CHAR(10) PRIMARY KEY,
	[StaffName] VARCHAR(50),
	[Password] VARCHAR(50)
);

INSERT INTO MsStaff VALUES ('D5416', 'Kanyadian Idhananta', 'isDosen');
INSERT INTO MsStaff VALUES ('D6831', 'Islam Nur Alam', 'isAlam');
INSERT INTO MsStaff VALUES ('D6657', 'Anderies Anderies', 'isAnderies');
INSERT INTO MsStaff VALUES ('D6811', 'Faisal Asadi', 'isFaisal');
INSERT INTO MsStaff VALUES ('D6835', 'Hanis Amalia Saputri', 'isHanis');

SELECT * FROM MsStaff;

DROP TABLE MsMahasiswa;
DROP TABLE MsStaff;

