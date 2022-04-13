package models

// Struct for `teams`
type Teams struct {
  teamid int32
  gameid int32
  color string
  token string
  hassteal int8
  longest int32
}