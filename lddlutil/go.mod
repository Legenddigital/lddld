module github.com/Legenddigital/lddld/lddlutil

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/Legenddigital/base58 v1.0.0
	github.com/Legenddigital/lddld/chaincfg v1.0.1
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/lddlec v0.0.0-20180721005212-59fe2b293f69
	github.com/Legenddigital/lddld/lddlec/edwards v0.0.0-20180721005212-59fe2b293f69
	github.com/Legenddigital/lddld/lddlec/secp256k1 v1.0.0
	github.com/Legenddigital/lddld/wire v1.0.1
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
)

replace (
	github.com/Legenddigital/lddld/chaincfg => ../chaincfg
	github.com/Legenddigital/lddld/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/Legenddigital/lddld/lddlec => ../lddlec
	github.com/Legenddigital/lddld/lddlec/edwards => ../lddlec/edwards
	github.com/Legenddigital/lddld/lddlec/secp256k1 => ../lddlec/secp256k1
	github.com/Legenddigital/lddld/wire => ../wire
)
