module github.com/glevine/goexp/interfaces

go 1.21.5

replace (
	github.com/glevine/goexp/interfaces/v2 => ./v2
	github.com/glevine/goexp/interfaces/v3 => ./v3
)

require (
	github.com/glevine/goexp/interfaces/v2 v2.0.0-00010101000000-000000000000
	github.com/glevine/goexp/interfaces/v3 v3.0.0-00010101000000-000000000000
)
