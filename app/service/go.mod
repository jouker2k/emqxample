module github.com/jwtea/emqxample/app/service

go 1.12

require (
	github.com/jwtea/emqxample/app/mqtt v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.2.0
	github.com/sirupsen/logrus v1.2.0
)

replace github.com/jwtea/emqxample/app/mqtt => ../mqtt
