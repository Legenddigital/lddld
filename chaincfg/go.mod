module github.com/Legenddigital/lddld/chaincfg

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/wire v1.0.1
)

replace (
	github.com/Legenddigital/lddld/chaincfg/chainhash => ./chainhash
	github.com/Legenddigital/lddld/wire => ../wire
)
