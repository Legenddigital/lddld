module github.com/Legenddigital/lddld/mining

require (
	github.com/Legenddigital/lddld/blockchain v1.0.0
	github.com/Legenddigital/lddld/blockchain/stake v1.0.0
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/lddljson v1.0.0
	github.com/Legenddigital/lddld/lddlutil v1.0.0
	github.com/Legenddigital/lddld/gcs v1.0.0
	github.com/Legenddigital/lddld/wire v1.0.1
)

replace (
	github.com/Legenddigital/lddld/blockchain => ../blockchain
	github.com/Legenddigital/lddld/blockchain/stake => ../blockchain/stake
	github.com/Legenddigital/lddld/chaincfg => ../chaincfg
	github.com/Legenddigital/lddld/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/Legenddigital/lddld/lddlec => ../lddlec
	github.com/Legenddigital/lddld/lddlec/edwards => ../lddlec/edwards
	github.com/Legenddigital/lddld/lddlec/secp256k1 => ../lddlec/secp256k1
	github.com/Legenddigital/lddld/lddljson => ../lddljson
	github.com/Legenddigital/lddld/lddlutil => ../lddlutil
	github.com/Legenddigital/lddld/gcs => ../gcs
	github.com/Legenddigital/lddld/wire => ../wire
)
