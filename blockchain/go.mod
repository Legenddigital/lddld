module github.com/Legenddigital/lddld/blockchain

require (
	github.com/Legenddigital/lddld/blockchain/stake v1.0.0
	github.com/Legenddigital/lddld/chaincfg v1.0.1
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/database v1.0.0
	github.com/Legenddigital/lddld/lddlec v0.0.0-20180801202239-0761de129164
	github.com/Legenddigital/lddld/lddlec/edwards v0.0.0-20180721031028-5369a485acf6
	github.com/Legenddigital/lddld/lddlec/secp256k1 v1.0.0
	github.com/Legenddigital/lddld/lddljson v1.0.0
	github.com/Legenddigital/lddld/lddlutil v1.0.0
	github.com/Legenddigital/lddld/gcs v1.0.0
	github.com/Legenddigital/lddld/txscript v1.0.0
	github.com/Legenddigital/lddld/wire v1.0.1
	github.com/Legenddigital/slog v1.0.0
)

replace (
	github.com/Legenddigital/lddld/blockchain/stake => ./stake
	github.com/Legenddigital/lddld/chaincfg => ../chaincfg
	github.com/Legenddigital/lddld/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/Legenddigital/lddld/database => ../database
	github.com/Legenddigital/lddld/lddlec => ../lddlec
	github.com/Legenddigital/lddld/lddlec/edwards => ../lddlec/edwards
	github.com/Legenddigital/lddld/lddlec/secp256k1 => ../lddlec/secp256k1
	github.com/Legenddigital/lddld/lddljson => ../lddljson
	github.com/Legenddigital/lddld/lddlutil => ../lddlutil
	github.com/Legenddigital/lddld/gcs => ../gcs
	github.com/Legenddigital/lddld/txscript => ../txscript
	github.com/Legenddigital/lddld/wire => ../wire
)
