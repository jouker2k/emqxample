module github.com/jwtea/emqxample/client

go 1.12

replace github.com/jwtea/emqxample/app/mqtt => ../mqtt

require (
	github.com/jwtea/emqxample/app/mqtt v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.2.0
)
