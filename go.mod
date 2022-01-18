module todo_app

go 1.17

replace config => ./config
replace utils => ./utils

require config v0.0.0-00010101000000-000000000000

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/go-ini/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
)
