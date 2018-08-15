module github.com/Legenddigital/lddld/hdkeychain

require (
	github.com/Legenddigital/base58 v1.0.0
	github.com/Legenddigital/lddld/chaincfg v1.0.1
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/lddlec v0.0.0-20180721005914-d26200ec716b
	github.com/Legenddigital/lddld/lddlec/secp256k1 v1.0.0
	github.com/Legenddigital/lddld/lddlutil v1.0.0
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
