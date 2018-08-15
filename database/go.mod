module github.com/Legenddigital/lddld/database

require (
	github.com/btcsuite/goleveldb v1.0.0
	github.com/Legenddigital/lddld/chaincfg v1.0.1
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/lddlutil v1.0.0
	github.com/Legenddigital/lddld/wire v1.0.1
	github.com/Legenddigital/slog v1.0.0
	github.com/jessevdk/go-flags v1.4.0
)

replace (
	github.com/Legenddigital/lddld/chaincfg => ../chaincfg
	github.com/Legenddigital/lddld/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/Legenddigital/lddld/lddlec => ../lddlec
	github.com/Legenddigital/lddld/lddlec/edwards => ../lddlec/edwards
	github.com/Legenddigital/lddld/lddlec/secp256k1 => ../lddlec/secp256k1
	github.com/Legenddigital/lddld/lddlutil => ../lddlutil
	github.com/Legenddigital/lddld/wire => ../wire
)
