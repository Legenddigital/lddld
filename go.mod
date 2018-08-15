module github.com/Legenddigital/lddld

require (
	github.com/aead/siphash v0.0.0-20170329201724-e404fcfc8885
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/btcsuite/goleveldb v1.0.0
	github.com/btcsuite/snappy-go v1.0.0
	github.com/btcsuite/winsvc v1.0.0
	github.com/davecgh/go-spew v1.1.0
	github.com/dchest/blake256 v1.0.0
	github.com/Legenddigital/base58 v1.0.0
	github.com/Legenddigital/lddld/addrmgr v1.0.1
	github.com/Legenddigital/lddld/blockchain v1.0.0
	github.com/Legenddigital/lddld/blockchain/stake v1.0.0
	github.com/Legenddigital/lddld/certgen v1.0.1
	github.com/Legenddigital/lddld/chaincfg v1.0.1
	github.com/Legenddigital/lddld/chaincfg/chainhash v1.0.1
	github.com/Legenddigital/lddld/connmgr v1.0.0
	github.com/Legenddigital/lddld/database v1.0.0
	github.com/Legenddigital/lddld/lddlec v0.0.0-20180801202239-0761de129164
	github.com/Legenddigital/lddld/lddlec/secp256k1 v1.0.0
	github.com/Legenddigital/lddld/lddljson v1.0.0
	github.com/Legenddigital/lddld/lddlutil v1.0.0
	github.com/Legenddigital/lddld/gcs v1.0.0
	github.com/Legenddigital/lddld/hdkeychain v1.0.0
	github.com/Legenddigital/lddld/mempool v1.0.0
	github.com/Legenddigital/lddld/mining v1.0.0
	github.com/Legenddigital/lddld/peer v1.0.0
	github.com/Legenddigital/lddld/rpcclient v1.0.0
	github.com/Legenddigital/lddld/txscript v1.0.0
	github.com/Legenddigital/lddld/wire v1.0.1
	github.com/Legenddigital/slog v1.0.0
	github.com/gorilla/websocket v1.2.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
	golang.org/x/sys v0.0.0-20180522224204-88eb85aaee56
)

replace (
	github.com/Legenddigital/lddld/addrmgr => ./addrmgr
	github.com/Legenddigital/lddld/blockchain => ./blockchain
	github.com/Legenddigital/lddld/blockchain/stake => ./blockchain/stake
	github.com/Legenddigital/lddld/certgen => ./certgen
	github.com/Legenddigital/lddld/chaincfg => ./chaincfg
	github.com/Legenddigital/lddld/chaincfg/chainhash => ./chaincfg/chainhash
	github.com/Legenddigital/lddld/connmgr => ./connmgr
	github.com/Legenddigital/lddld/database => ./database
	github.com/Legenddigital/lddld/lddlec => ./lddlec
	github.com/Legenddigital/lddld/lddlec/edwards => ./lddlec/edwards
	github.com/Legenddigital/lddld/lddlec/secp256k1 => ./lddlec/secp256k1
	github.com/Legenddigital/lddld/lddljson => ./lddljson
	github.com/Legenddigital/lddld/lddlutil => ./lddlutil
	github.com/Legenddigital/lddld/gcs => ./gcs
	github.com/Legenddigital/lddld/hdkeychain => ./hdkeychain
	github.com/Legenddigital/lddld/limits => ./limits
	github.com/Legenddigital/lddld/mempool => ./mempool
	github.com/Legenddigital/lddld/mining => ./mining
	github.com/Legenddigital/lddld/peer => ./peer
	github.com/Legenddigital/lddld/rpcclient => ./rpcclient
	github.com/Legenddigital/lddld/txscript => ./txscript
	github.com/Legenddigital/lddld/wire => ./wire
)
