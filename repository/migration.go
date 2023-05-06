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
    BLOOD_TYPE CHAR NOT NULL,
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

	// insert
	if err := d.DB.Exec(`
	INSERT INTO DOCTOR (DOCTOR_ID, FIRST_NAME, LAST_NAME)
		VALUES (1, 'John', 'Doe'),
       (2, 'Jane', 'Smith'),
       (3, 'Michael', 'Johnson');
	`).Error; err != nil {
		return err
	}

	// create views

	return nil
}
