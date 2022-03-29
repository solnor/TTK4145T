module TTK4145-Project

go 1.17

replace (
	driver => ./driver
	elevator => ./elevator
	network => ./network
)

require driver v0.0.0-00010101000000-000000000000

require elevator v0.0.0-00010101000000-000000000000 // indirect
