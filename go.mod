module go.nhat.io/fontforge

go 1.22

toolchain go1.22.0

require (
	github.com/Masterminds/semver/v3 v3.3.1
	github.com/stretchr/testify v1.10.0
	go.nhat.io/python/v3 v3.11.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.nhat.io/cpy/v3 v3.11.2 // indirect
	go.nhat.io/once v0.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// v0.2.0 is incompatible with github.com/Masterminds/semver/v3 v3.3.1.
retract v0.2.0
