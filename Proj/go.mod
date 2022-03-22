module TTK4145-Project

go 1.17

require(
    elevator v0.0.0
    network v0.0.0
)

replace(
    elevator => ./elevator
    network => ./network
)