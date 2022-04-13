package models

import "time"

// Struct for `games`
type Games struct {
	gameid int32
	gametype string
	mapid int32
	name string
	datetime time.Time
	token string
	rounds int32
	teams int32
	roundlength int32
	mtime float64
	currentround int32
	distance_weight float64
	number_weight float64
	longest_weight float64
	longest_hq_weight int32
	multimove int8
	templateid int32
	cohortid int32
	active int8
	current int8
	opts string
	stealcount int8
	record int8
	ranked int8
}