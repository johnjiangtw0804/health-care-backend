package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewGormDatabase(dsn string, debug bool) (*GormDatabase, error) {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if debug {
		config.Logger = gormLogger.Default.LogMode(gormLogger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	return &GormDatabase{DB: db}, nil
}

func (d *GormDatabase) AutoMigrate() error {
	// here we don't actually need to use the gorm library. We can just use the raw sql
	if err := d.DB.Exec(`
	CREATE TABLE DOCTOR (
	DOCTOR_ID INT,
	FIRST_NAME VARCHAR(50) NOT NULL,
	LAST_NAME VARCHAR(50) NOT NULL,
	PRIMARY KEY (DOCTOR_ID));`).Error; err != nil {
		return err
	}

	if err := d.DB.Exec(`
	CREATE TABLE PATIENT (
	PATIENT_ID INT,
    FIRST_NAME VARCHAR(50) NOT NULL,
	LAST_NAME VARCHAR(50) NOT NULL,
    AGE INT NOT NULL,
    SEX CHAR NOT NULL,
    BLOOD_TYPE CHAR(2) NOT NULL,
    DOB DATE NOT NULL,
    DOCTOR_ID INT NOT NULL,
    PRIMARY KEY (PATIENT_ID),
    CONSTRAINT PATIENT_FK_DOCTOR_ID FOREIGN KEY (DOCTOR_ID) REFERENCES DOCTOR(DOCTOR_ID));`).Error; err != nil {
		return err
	}

	if err := d.DB.Exec(`
	CREATE TABLE VITAL_SIGN (
    PATIENT_ID INT,
    ISSUE_TIME TIMESTAMP,
    BODY_TEMPERTATURE FLOAT NOT NULL,
    PULSE_RATE INT NOT NULL,
    RESIPIRATE_RATE INT NOT NULL,
    SYSTOLIC_PRESSURE INT NOT NULL,
    DIASTOLIC_PRESSURE INT NOT NULL,
    PRIMARY KEY(PATIENT_ID, ISSUE_TIME),
    CONSTRAINT VITAL_SIGN_FK_PATIENT_ID FOREIGN KEY (PATIENT_ID) REFERENCES PATIENT(PATIENT_ID));`).Error; err != nil {
		return err
	}

	if err := d.DB.Exec(`
	CREATE TABLE PATIENT_MEDICATIONS (
	PATIENT_ID INT,
	PRESCRIBED_MEDICATIONS VARCHAR(50),
    PRIMARY KEY(PATIENT_ID, PRESCRIBED_MEDICATIONS),
    CONSTRAINT PATIENT_MEDICATIONS_FK_PATIENT_ID FOREIGN KEY (PATIENT_ID) REFERENCES PATIENT(PATIENT_ID));`).Error; err != nil {
		return err
	}

	if err := d.DB.Exec(`
	CREATE TABLE PATIENT_DISEASE (
	PATIENT_ID INT,
	DISEASE VARCHAR(50),
	PRIMARY KEY(PATIENT_ID, DISEASE),
    CONSTRAINT PATIENT_DISEASE_FK_PATIENT_ID FOREIGN KEY (PATIENT_ID) REFERENCES PATIENT(PATIENT_ID));`).Error; err != nil {
		return err
	}

	if err := d.DB.Exec(`
	CREATE TABLE NURSE (
	NURSE_ID INT,
    FIRST_NAME VARCHAR(50) NOT NULL,
	LAST_NAME VARCHAR(50) NOT NULL,
	PRIMARY KEY (NURSE_ID));`).Error; err != nil {
		return err
	}

	if err := d.DB.Exec(`
	CREATE TABLE PATIENT_NURSE (
	PATIENT_ID INT,
    NURSE_ID INT,
	PRIMARY KEY (PATIENT_ID, NURSE_ID),
    CONSTRAINT PATIENT_NURSE_FK_PATIENT_ID FOREIGN KEY (PATIENT_ID) REFERENCES PATIENT(PATIENT_ID),
    CONSTRAINT PATIENT_NURSE_FK_NURSE_ID FOREIGN KEY (NURSE_ID) REFERENCES NURSE(NURSE_ID));`).Error; err != nil {
		return err
	}

	// insert some doctors
	if err := d.DB.Exec(`
	INSERT INTO DOCTOR (DOCTOR_ID, FIRST_NAME, LAST_NAME)
		VALUES (1, 'John', 'Doe'),
       (2, 'Jane', 'Smith'),
       (3, 'Michael', 'Johnson');`).Error; err != nil {
		return err
	}

	// insert some patients
	if err := d.DB.Exec(`
	INSERT INTO PATIENT (PATIENT_ID, FIRST_NAME, LAST_NAME, AGE, SEX, BLOOD_TYPE, DOB, DOCTOR_ID)
	VALUES (1, 'Alice', 'Johnson', 35, 'F', 'A+', '1988-03-12', 1),
       (2, 'Bob', 'Smith', 45, 'M', 'B-', '1978-07-24', 2),
       (3, 'Carol', 'Davis', 28, 'F', 'O+', '1995-11-05', 1);`).Error; err != nil {
		return err
	}

	// insert some vital signs
	if err := d.DB.Exec(`
	INSERT INTO VITAL_SIGN (PATIENT_ID, ISSUE_TIME, BODY_TEMPERTATURE, PULSE_RATE, RESIPIRATE_RATE, SYSTOLIC_PRESSURE, DIASTOLIC_PRESSURE)
	VALUES (1, '2023-05-01 10:30:00', 98.6, 70, 18, 120, 80),
       (2, '2023-05-02 09:45:00', 99.2, 68, 16, 130, 85),
       (3, '2023-05-03 15:15:00', 98.8, 72, 20, 125, 82);`).Error; err != nil {
		return err
	}

	// insert some medications
	if err := d.DB.Exec(`
	INSERT INTO PATIENT_MEDICATIONS (PATIENT_ID, PRESCRIBED_MEDICATIONS)
	VALUES (1, 'Aspirin'),
       (1, 'Antibiotic'),
       (2, 'Painkiller'),
       (3, 'Antihistamine');`).Error; err != nil {
		return err
	}

	// insert some diseases
	if err := d.DB.Exec(`
	INSERT INTO PATIENT_DISEASE (PATIENT_ID, DISEASE)
	VALUES (1, 'Hypertension'),
       (2, 'Diabetes'),
       (3, 'Asthma');`).Error; err != nil {
		return err
	}

	// insert some nurses
	if err := d.DB.Exec(`
	INSERT INTO NURSE (NURSE_ID, FIRST_NAME, LAST_NAME)
	VALUES (1, 'Emily', 'Wilson'),
       (2, 'David', 'Brown'),
       (3, 'Sophia', 'Anderson');`).Error; err != nil {
		return err
	}

	// insert some patient-nurse relationships
	if err := d.DB.Exec(`
	INSERT INTO PATIENT_NURSE (PATIENT_ID, NURSE_ID)
	VALUES (1, 1),
       (2, 2),
       (3, 3);`).Error; err != nil {
		return err
	}

	// generate some views
	if err := d.DB.Exec(`
	CREATE VIEW PATIENT_DASHBOARD_VIEW AS (
    SELECT
        p.patient_id AS ID,
        p.first_name,
        p.last_name,
        p.age,
        p.sex,
        p.blood_type,
        p.dob AS DOB,
        p.doctor_id AS assigned_doctor_ID,
        v.body_tempertature,
        v.pulse_rate,
        v.resipirate_rate,
        v.systolic_pressure,
        v.diastolic_pressure,
        m.prescribed_medications AS current_prescribed_med,
        d.disease AS current_disease
    FROM
        PATIENT AS p
    JOIN
        vital_sign AS v ON p.patient_id = v.patient_id
    JOIN
        patient_medications AS m ON p.patient_id = m.patient_id
    JOIN
        patient_disease AS d ON p.patient_id = d.patient_id
	);`).Error; err != nil {
		return err
	}

	return nil
}
