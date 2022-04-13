package models

//Struct for `simulations`
type Simulations struct {
	simulationid int32
	name string
	game_url string
	setup_url string
	env string
	configvar string
	field_mapping string
}