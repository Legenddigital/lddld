module github.com/Legenddigital/lddld/txscript

require (
	github.com/Legenddigital/lddld/chaincfg v1.0.1
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/lddlec v0.0.0-20180721031028-5369a485acf6
	github.com/Legenddigital/lddld/lddlec/edwards v0.0.0-20180721031028-5369a485acf6
	github.com/Legenddigital/lddld/lddlec/secp256k1 v1.0.0
	github.com/Legenddigital/lddld/lddlutil v1.0.0
	github.com/Legenddigital/lddld/wire v1.0.1
	github.com/Legenddigital/slog v1.0.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
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
