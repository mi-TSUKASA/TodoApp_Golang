module config

go 1.17

require (
	gopkg.in/go-ini/ini.v1 v1.66.2
	utils v0.0.0-00010101000000-000000000000
)

require github.com/stretchr/testify v1.7.0 // indirect

replace utils => ../utils
