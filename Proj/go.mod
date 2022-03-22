module TTK4145-Project

go 1.17

replace (
	elevator => ./elevator
	network => ./network
)

require elevator v0.0.0-00010101000000-000000000000
