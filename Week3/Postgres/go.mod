module example.com/postgresDB

go 1.17

replace example.com/connection => ./connection

replace example.com/skelton => ./skelton

replace example.com/util => ./util

require (
	example.com/connection v0.0.0-00010101000000-000000000000
	example.com/skelton v0.0.0-00010101000000-000000000000
)

require (
	example.com/util v0.0.0-00010101000000-000000000000 // indirect
	github.com/lib/pq v1.10.4 // indirect
)
