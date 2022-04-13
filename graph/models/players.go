package models

import "time"

// Struct for `players`
type Players struct {
	playerid int32
	cohortid int32
	name string
	email string
	age int32
	dob time.Time
	username string
	password string
	admin int8
	order int8
	reset_token string
	password_sent time.Time
	last_login time.Time
	timezone string
	at_id string
	hashed_password string
	start_date time.Time
	user_id string
}