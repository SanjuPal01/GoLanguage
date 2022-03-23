module example.com/assignments

go 1.17

replace example.com/stack => ./stack

replace example.com/queue => ./queue

require (
	example.com/queue v0.0.0-00010101000000-000000000000
	example.com/stack v0.0.0-00010101000000-000000000000
)
