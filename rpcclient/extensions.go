// Copyright (c) 2014-2015 The btcsuite developers
// Copyright (c) 2015-2017 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpcclient

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Legenddigital/lddld/chaincfg/chainhash"
	"github.com/Legenddigital/lddld/lddljson"
	"github.com/Legenddigital/lddld/lddlutil"
	"github.com/Legenddigital/lddld/wire"
)

var (
	// zeroUint32 is the zero value for a uint32.
	zeroUint32 = uint32(0)
)

// FutureCreateEncryptedWalletResult is a future promise to deliver the error
// result of a CreateEncryptedWalletAsync RPC invocation.
type FutureCreateEncryptedWalletResult chan *response

// Receive waits for and returns the error response promised by the future.
func (r FutureCreateEncryptedWalletResult) Receive() error {
	_, err := receiveFuture(r)
	return err
}

// CreateEncryptedWalletAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See CreateEncryptedWallet for the blocking version and more details.
//
// NOTE: This is a lddlwallet extension.
func (c *Client) CreateEncryptedWalletAsync(passphrase string) FutureCreateEncryptedWalletResult {
	cmd := lddljson.NewCreateEncryptedWalletCmd(passphrase)
	return c.sendCmd(cmd)
}

// CreateEncryptedWallet requests the creation of an encrypted wallet.  Wallets
// managed by lddlwallet are only written to disk with encrypted private keys,
// and generating wallets on the fly is impossible as it requires user input for
// the encryption passphrase.  This RPC specifies the passphrase and instructs
// the wallet creation.  This may error if a wallet is already opened, or the
// new wallet cannot be written to disk.
//
// NOTE: This is a lddlwallet extension.
func (c *Client) CreateEncryptedWallet(passphrase string) error {
	return c.CreateEncryptedWalletAsync(passphrase).Receive()
}

// FutureDebugLevelResult is a future promise to deliver the result of a
// DebugLevelAsync RPC invocation (or an applicable error).
type FutureDebugLevelResult chan *response

// Receive waits for the response promised by the future and returns the result
// of setting the debug logging level to the passed level specification or the
// list of of the available subsystems for the special keyword 'show'.
func (r FutureDebugLevelResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmashal the result as a string.
	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

// DebugLevelAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See DebugLevel for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) DebugLevelAsync(levelSpec string) FutureDebugLevelResult {
	cmd := lddljson.NewDebugLevelCmd(levelSpec)
	return c.sendCmd(cmd)
}

// DebugLevel dynamically sets the debug logging level to the passed level
// specification.
//
// The levelspec can be either a debug level or of the form:
// 	<subsystem>=<level>,<subsystem2>=<level2>,...
//
// Additionally, the special keyword 'show' can be used to get a list of the
// available subsystems.
//
// NOTE: This is a lddld extension.
func (c *Client) DebugLevel(levelSpec string) (string, error) {
	return c.DebugLevelAsync(levelSpec).Receive()
}

// FutureEstimateStakeDiffResult is a future promise to deliver the result of a
// EstimateStakeDiffAsync RPC invocation (or an applicable error).
type FutureEstimateStakeDiffResult chan *response

// Receive waits for the response promised by the future and returns the hash
// and height of the block in the longest (best) chain.
func (r FutureEstimateStakeDiffResult) Receive() (*lddljson.EstimateStakeDiffResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarsal result as a estimatestakediff result object.
	var est lddljson.EstimateStakeDiffResult
	err = json.Unmarshal(res, &est)
	if err != nil {
		return nil, err
	}

	return &est, nil
}

// EstimateStakeDiffAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See EstimateStakeDiff for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) EstimateStakeDiffAsync(tickets *uint32) FutureEstimateStakeDiffResult {
	cmd := lddljson.NewEstimateStakeDiffCmd(tickets)
	return c.sendCmd(cmd)
}

// EstimateStakeDiff returns the minimum, maximum, and expected next stake
// difficulty.
//
// NOTE: This is a lddld extension.
func (c *Client) EstimateStakeDiff(tickets *uint32) (*lddljson.EstimateStakeDiffResult, error) {
	return c.EstimateStakeDiffAsync(tickets).Receive()
}

// FutureExistsAddressResult is a future promise to deliver the result
// of a FutureExistsAddressResultAsync RPC invocation (or an applicable error).
type FutureExistsAddressResult chan *response

// Receive waits for the response promised by the future and returns whether or
// not an address exists in the blockchain or mempool.
func (r FutureExistsAddressResult) Receive() (bool, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return false, err
	}

	// Unmarshal the result as a bool.
	var exists bool
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// ExistsAddressAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) ExistsAddressAsync(address lddlutil.Address) FutureExistsAddressResult {
	cmd := lddljson.NewExistsAddressCmd(address.EncodeAddress())
	return c.sendCmd(cmd)
}

// ExistsAddress returns information about whether or not an address has been
// used on the main chain or in mempool.
//
// NOTE: This is a lddld extension.
func (c *Client) ExistsAddress(address lddlutil.Address) (bool, error) {
	return c.ExistsAddressAsync(address).Receive()
}

// FutureExistsAddressesResult is a future promise to deliver the result
// of a FutureExistsAddressesResultAsync RPC invocation (or an
// applicable error).
type FutureExistsAddressesResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the addresses exist.
func (r FutureExistsAddressesResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal the result as a string.
	var exists string
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return "", err
	}
	return exists, nil
}

// ExistsAddressesAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) ExistsAddressesAsync(addresses []lddlutil.Address) FutureExistsAddressesResult {
	addrsStr := make([]string, len(addresses))
	for i := range addresses {
		addrsStr[i] = addresses[i].EncodeAddress()
	}

	cmd := lddljson.NewExistsAddressesCmd(addrsStr)
	return c.sendCmd(cmd)
}

// ExistsAddresses returns information about whether or not an address exists
// in the blockchain or memory pool.
//
// NOTE: This is a lddld extension.
func (c *Client) ExistsAddresses(addresses []lddlutil.Address) (string, error) {
	return c.ExistsAddressesAsync(addresses).Receive()
}

// FutureExistsMissedTicketsResult is a future promise to deliver the result of
// an ExistsMissedTicketsAsync RPC invocation (or an applicable error).
type FutureExistsMissedTicketsResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the ticket exists in the missed ticket database.
func (r FutureExistsMissedTicketsResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal the result as a string.
	var exists string
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return "", err
	}
	return exists, nil
}

// ExistsMissedTicketsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
func (c *Client) ExistsMissedTicketsAsync(hashes []*chainhash.Hash) FutureExistsMissedTicketsResult {
	hashBlob := make([]byte, len(hashes)*chainhash.HashSize)
	for i, hash := range hashes {
		copy(hashBlob[i*chainhash.HashSize:(i+1)*chainhash.HashSize],
			hash[:])
	}
	cmd := lddljson.NewExistsMissedTicketsCmd(hex.EncodeToString(hashBlob))
	return c.sendCmd(cmd)
}

// ExistsMissedTickets returns a hex-encoded bitset describing whether or not
// ticket hashes exists in the missed ticket database.
func (c *Client) ExistsMissedTickets(hashes []*chainhash.Hash) (string, error) {
	return c.ExistsMissedTicketsAsync(hashes).Receive()
}

// FutureExistsExpiredTicketsResult is a future promise to deliver the result
// of a FutureExistsExpiredTicketsResultAsync RPC invocation (or an
// applicable error).
type FutureExistsExpiredTicketsResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the ticket exists in the live ticket database.
func (r FutureExistsExpiredTicketsResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal the result as a string.
	var exists string
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return "", err
	}
	return exists, nil
}

// ExistsExpiredTicketsAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) ExistsExpiredTicketsAsync(hashes []*chainhash.Hash) FutureExistsExpiredTicketsResult {
	hashBlob := make([]byte, len(hashes)*chainhash.HashSize)
	for i, hash := range hashes {
		copy(hashBlob[i*chainhash.HashSize:(i+1)*chainhash.HashSize],
			hash[:])
	}
	cmd := lddljson.NewExistsExpiredTicketsCmd(hex.EncodeToString(hashBlob))
	return c.sendCmd(cmd)
}

// ExistsExpiredTickets returns information about whether or not a ticket hash exists
// in the expired ticket database.
//
// NOTE: This is a lddld extension.
func (c *Client) ExistsExpiredTickets(hashes []*chainhash.Hash) (string, error) {
	return c.ExistsExpiredTicketsAsync(hashes).Receive()
}

// FutureExistsLiveTicketResult is a future promise to deliver the result
// of a FutureExistsLiveTicketResultAsync RPC invocation (or an
// applicable error).
type FutureExistsLiveTicketResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the ticket exists in the live ticket database.
func (r FutureExistsLiveTicketResult) Receive() (bool, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return false, err
	}

	// Unmarshal the result as a bool.
	var exists bool
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// ExistsLiveTicketAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) ExistsLiveTicketAsync(hash *chainhash.Hash) FutureExistsLiveTicketResult {
	cmd := lddljson.NewExistsLiveTicketCmd(hash.String())
	return c.sendCmd(cmd)
}

// ExistsLiveTicket returns information about whether or not a ticket hash exists
// in the live ticket database.
//
// NOTE: This is a lddld extension.
func (c *Client) ExistsLiveTicket(hash *chainhash.Hash) (bool, error) {
	return c.ExistsLiveTicketAsync(hash).Receive()
}

// FutureExistsLiveTicketsResult is a future promise to deliver the result
// of a FutureExistsLiveTicketsResultAsync RPC invocation (or an
// applicable error).
type FutureExistsLiveTicketsResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the ticket exists in the live ticket database.
func (r FutureExistsLiveTicketsResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal the result as a string.
	var exists string
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return "", err
	}
	return exists, nil
}

// ExistsLiveTicketsAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) ExistsLiveTicketsAsync(hashes []*chainhash.Hash) FutureExistsLiveTicketsResult {
	hashBlob := make([]byte, len(hashes)*chainhash.HashSize)
	for i, hash := range hashes {
		copy(hashBlob[i*chainhash.HashSize:(i+1)*chainhash.HashSize],
			hash[:])
	}
	cmd := lddljson.NewExistsLiveTicketsCmd(hex.EncodeToString(hashBlob))
	return c.sendCmd(cmd)
}

// ExistsLiveTickets returns information about whether or not a ticket hash exists
// in the live ticket database.
//
// NOTE: This is a lddld extension.
func (c *Client) ExistsLiveTickets(hashes []*chainhash.Hash) (string, error) {
	return c.ExistsLiveTicketsAsync(hashes).Receive()
}

// FutureExistsMempoolTxsResult is a future promise to deliver the result
// of a FutureExistsMempoolTxsResultAsync RPC invocation (or an
// applicable error).
type FutureExistsMempoolTxsResult chan *response

// Receive waits for the response promised by the future and returns whether
// or not the ticket exists in the mempool.
func (r FutureExistsMempoolTxsResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal the result as a string.
	var exists string
	err = json.Unmarshal(res, &exists)
	if err != nil {
		return "", err
	}
	return exists, nil
}

// ExistsMempoolTxsAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) ExistsMempoolTxsAsync(hashes []*chainhash.Hash) FutureExistsMempoolTxsResult {
	hashBlob := make([]byte, len(hashes)*chainhash.HashSize)
	for i, hash := range hashes {
		copy(hashBlob[i*chainhash.HashSize:(i+1)*chainhash.HashSize],
			hash[:])
	}
	cmd := lddljson.NewExistsMempoolTxsCmd(hex.EncodeToString(hashBlob))
	return c.sendCmd(cmd)
}

// ExistsMempoolTxs returns information about whether or not a ticket hash exists
// in the live ticket database.
//
// NOTE: This is a lddld extension.
func (c *Client) ExistsMempoolTxs(hashes []*chainhash.Hash) (string, error) {
	return c.ExistsMempoolTxsAsync(hashes).Receive()
}

// FutureExportWatchingWalletResult is a future promise to deliver the result of
// an ExportWatchingWalletAsync RPC invocation (or an applicable error).
type FutureExportWatchingWalletResult chan *response

// Receive waits for the response promised by the future and returns the
// exported wallet.
func (r FutureExportWatchingWalletResult) Receive() ([]byte, []byte, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, nil, err
	}

	// Unmarshal result as a JSON object.
	var obj map[string]interface{}
	err = json.Unmarshal(res, &obj)
	if err != nil {
		return nil, nil, err
	}

	// Check for the wallet and tx string fields in the object.
	base64Wallet, ok := obj["wallet"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected response type for "+
			"exportwatchingwallet 'wallet' field: %T\n",
			obj["wallet"])
	}
	base64TxStore, ok := obj["tx"].(string)
	if !ok {
		return nil, nil, fmt.Errorf("unexpected response type for "+
			"exportwatchingwallet 'tx' field: %T\n",
			obj["tx"])
	}

	walletBytes, err := base64.StdEncoding.DecodeString(base64Wallet)
	if err != nil {
		return nil, nil, err
	}

	txStoreBytes, err := base64.StdEncoding.DecodeString(base64TxStore)
	if err != nil {
		return nil, nil, err
	}

	return walletBytes, txStoreBytes, nil

}

// ExportWatchingWalletAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ExportWatchingWallet for the blocking version and more details.
//
// NOTE: This is a lddlwallet extension.
func (c *Client) ExportWatchingWalletAsync(account string) FutureExportWatchingWalletResult {
	cmd := lddljson.NewExportWatchingWalletCmd(&account, lddljson.Bool(true))
	return c.sendCmd(cmd)
}

// ExportWatchingWallet returns the raw bytes for a watching-only version of
// wallet.bin and tx.bin, respectively, for the specified account that can be
// used by lddlwallet to enable a wallet which does not have the private keys
// necessary to spend funds.
//
// NOTE: This is a lddlwallet extension.
func (c *Client) ExportWatchingWallet(account string) ([]byte, []byte, error) {
	return c.ExportWatchingWalletAsync(account).Receive()
}

// FutureGetBestBlockResult is a future promise to deliver the result of a
// GetBestBlockAsync RPC invocation (or an applicable error).
type FutureGetBestBlockResult chan *response

// Receive waits for the response promised by the future and returns the hash
// and height of the block in the longest (best) chain.
func (r FutureGetBestBlockResult) Receive() (*chainhash.Hash, int64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, 0, err
	}

	// Unmarshal result as a getbestblock result object.
	var bestBlock lddljson.GetBestBlockResult
	err = json.Unmarshal(res, &bestBlock)
	if err != nil {
		return nil, 0, err
	}

	// Convert to hash from string.
	hash, err := chainhash.NewHashFromStr(bestBlock.Hash)
	if err != nil {
		return nil, 0, err
	}

	return hash, bestBlock.Height, nil
}

// GetBestBlockAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBestBlock for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetBestBlockAsync() FutureGetBestBlockResult {
	cmd := lddljson.NewGetBestBlockCmd()
	return c.sendCmd(cmd)
}

// GetBestBlock returns the hash and height of the block in the longest (best)
// chain.
//
// NOTE: This is a lddld extension.
func (c *Client) GetBestBlock() (*chainhash.Hash, int64, error) {
	return c.GetBestBlockAsync().Receive()
}

// FutureGetCurrentNetResult is a future promise to deliver the result of a
// GetCurrentNetAsync RPC invocation (or an applicable error).
type FutureGetCurrentNetResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetCurrentNetResult) Receive() (wire.CurrencyNet, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal result as an int64.
	var net int64
	err = json.Unmarshal(res, &net)
	if err != nil {
		return 0, err
	}

	return wire.CurrencyNet(net), nil
}

// GetCurrentNetAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetCurrentNet for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetCurrentNetAsync() FutureGetCurrentNetResult {
	cmd := lddljson.NewGetCurrentNetCmd()
	return c.sendCmd(cmd)
}

// GetCurrentNet returns the network the server is running on.
//
// NOTE: This is a lddld extension.
func (c *Client) GetCurrentNet() (wire.CurrencyNet, error) {
	return c.GetCurrentNetAsync().Receive()
}

// FutureGetHeadersResult is a future promise to deliver the result of a
// getheaders RPC invocation (or an applicable error).
type FutureGetHeadersResult chan *response

// Receive waits for the response promised by the future and returns the
// getheaders result.
func (r FutureGetHeadersResult) Receive() (*lddljson.GetHeadersResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarsal result as a getheaders result object.
	var vr lddljson.GetHeadersResult
	err = json.Unmarshal(res, &vr)
	if err != nil {
		return nil, err
	}

	return &vr, nil
}

// GetHeadersAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the returned instance.
//
// See GetHeaders for the blocking version and more details.
func (c *Client) GetHeadersAsync(blockLocators []*chainhash.Hash, hashStop *chainhash.Hash) FutureGetHeadersResult {
	concatenatedLocators := make([]byte, chainhash.HashSize*len(blockLocators))
	for i := range blockLocators {
		copy(concatenatedLocators[i*chainhash.HashSize:], blockLocators[i][:])
	}
	cmd := lddljson.NewGetHeadersCmd(hex.EncodeToString(concatenatedLocators),
		hashStop.String())
	return c.sendCmd(cmd)
}

// GetHeaders mimics the wire protocol getheaders and headers messages by
// returning all headers on the main chain after the first known block in the
// locators, up until a block hash matches hashStop.
func (c *Client) GetHeaders(blockLocators []*chainhash.Hash, hashStop *chainhash.Hash) (*lddljson.GetHeadersResult, error) {
	return c.GetHeadersAsync(blockLocators, hashStop).Receive()
}

// FutureGetStakeDifficultyResult is a future promise to deliver the result of a
// GetStakeDifficultyAsync RPC invocation (or an applicable error).
type FutureGetStakeDifficultyResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetStakeDifficultyResult) Receive() (*lddljson.GetStakeDifficultyResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a lddljson.GetStakeDifficultyResult.
	var gsdr lddljson.GetStakeDifficultyResult
	err = json.Unmarshal(res, &gsdr)
	if err != nil {
		return nil, err
	}

	return &gsdr, nil
}

// GetStakeDifficultyAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetStakeDifficulty for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetStakeDifficultyAsync() FutureGetStakeDifficultyResult {
	cmd := lddljson.NewGetStakeDifficultyCmd()
	return c.sendCmd(cmd)
}

// GetStakeDifficulty returns the current and next stake difficulty.
//
// NOTE: This is a lddld extension.
func (c *Client) GetStakeDifficulty() (*lddljson.GetStakeDifficultyResult, error) {
	return c.GetStakeDifficultyAsync().Receive()
}

// FutureGetStakeVersionsResult is a future promise to deliver the result of a
// GetStakeVersionsAsync RPC invocation (or an applicable error).
type FutureGetStakeVersionsResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetStakeVersionsResult) Receive() (*lddljson.GetStakeVersionsResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a lddljson.GetStakeVersionsResult.
	var gsvr lddljson.GetStakeVersionsResult
	err = json.Unmarshal(res, &gsvr)
	if err != nil {
		return nil, err
	}

	return &gsvr, nil
}

// GetStakeVersionInfoAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetStakeVersionInfo for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetStakeVersionInfoAsync(count int32) FutureGetStakeVersionInfoResult {
	cmd := lddljson.NewGetStakeVersionInfoCmd(count)
	return c.sendCmd(cmd)
}

// GetStakeVersionInfo returns the stake versions results for past requested intervals (count).
//
// NOTE: This is a lddld extension.
func (c *Client) GetStakeVersionInfo(count int32) (*lddljson.GetStakeVersionInfoResult, error) {
	return c.GetStakeVersionInfoAsync(count).Receive()
}

// FutureGetStakeVersionInfoResult is a future promise to deliver the result of a
// GetStakeVersionInfoAsync RPC invocation (or an applicable error).
type FutureGetStakeVersionInfoResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetStakeVersionInfoResult) Receive() (*lddljson.GetStakeVersionInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a lddljson.GetStakeVersionInfoResult.
	var gsvr lddljson.GetStakeVersionInfoResult
	err = json.Unmarshal(res, &gsvr)
	if err != nil {
		return nil, err
	}

	return &gsvr, nil
}

// GetStakeVersionsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetStakeVersions for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetStakeVersionsAsync(hash string, count int32) FutureGetStakeVersionsResult {
	cmd := lddljson.NewGetStakeVersionsCmd(hash, count)
	return c.sendCmd(cmd)
}

// GetStakeVersions returns the stake versions and vote versions of past requested blocks.
//
// NOTE: This is a lddld extension.
func (c *Client) GetStakeVersions(hash string, count int32) (*lddljson.GetStakeVersionsResult, error) {
	return c.GetStakeVersionsAsync(hash, count).Receive()
}

// FutureGetTicketPoolValueResult is a future promise to deliver the result of a
// GetTicketPoolValueAsync RPC invocation (or an applicable error).
type FutureGetTicketPoolValueResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetTicketPoolValueResult) Receive() (lddlutil.Amount, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal result as a float64.
	var val float64
	err = json.Unmarshal(res, &val)
	if err != nil {
		return 0, err
	}

	// Convert to an amount.
	amt, err := lddlutil.NewAmount(val)
	if err != nil {
		return 0, err
	}

	return amt, nil
}

// GetTicketPoolValueAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetTicketPoolValue for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetTicketPoolValueAsync() FutureGetTicketPoolValueResult {
	cmd := lddljson.NewGetTicketPoolValueCmd()
	return c.sendCmd(cmd)
}

// GetTicketPoolValue returns the value of the live ticket pool.
//
// NOTE: This is a lddld extension.
func (c *Client) GetTicketPoolValue() (lddlutil.Amount, error) {
	return c.GetTicketPoolValueAsync().Receive()
}

// FutureGetVoteInfoResult is a future promise to deliver the result of a
// GetVoteInfoAsync RPC invocation (or an applicable error).
type FutureGetVoteInfoResult chan *response

// Receive waits for the response promised by the future and returns the network
// the server is running on.
func (r FutureGetVoteInfoResult) Receive() (*lddljson.GetVoteInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a lddljson.GetVoteInfoResult.
	var gsvr lddljson.GetVoteInfoResult
	err = json.Unmarshal(res, &gsvr)
	if err != nil {
		return nil, err
	}

	return &gsvr, nil
}

// GetVoteInfoAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See GetVoteInfo for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) GetVoteInfoAsync(version uint32) FutureGetVoteInfoResult {
	cmd := lddljson.NewGetVoteInfoCmd(version)
	return c.sendCmd(cmd)
}

// GetVoteInfo returns voting information for the specified stake version. This
// includes current voting window, quorum, total votes and agendas.
//
// NOTE: This is a lddld extension.
func (c *Client) GetVoteInfo(version uint32) (*lddljson.GetVoteInfoResult, error) {
	return c.GetVoteInfoAsync(version).Receive()
}

// FutureListAddressTransactionsResult is a future promise to deliver the result
// of a ListAddressTransactionsAsync RPC invocation (or an applicable error).
type FutureListAddressTransactionsResult chan *response

// Receive waits for the response promised by the future and returns information
// about all transactions associated with the provided addresses.
func (r FutureListAddressTransactionsResult) Receive() ([]lddljson.ListTransactionsResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as an array of listtransactions objects.
	var transactions []lddljson.ListTransactionsResult
	err = json.Unmarshal(res, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// ListAddressTransactionsAsync returns an instance of a type that can be used
// to get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ListAddressTransactions for the blocking version and more details.
//
// NOTE: This is a lddld extension.
func (c *Client) ListAddressTransactionsAsync(addresses []lddlutil.Address, account string) FutureListAddressTransactionsResult {
	// Convert addresses to strings.
	addrs := make([]string, 0, len(addresses))
	for _, addr := range addresses {
		addrs = append(addrs, addr.EncodeAddress())
	}
	cmd := lddljson.NewListAddressTransactionsCmd(addrs, &account)
	return c.sendCmd(cmd)
}

// ListAddressTransactions returns information about all transactions associated
// with the provided addresses.
//
// NOTE: This is a lddlwallet extension.
func (c *Client) ListAddressTransactions(addresses []lddlutil.Address, account string) ([]lddljson.ListTransactionsResult, error) {
	return c.ListAddressTransactionsAsync(addresses, account).Receive()
}

// FutureLiveTicketsResult is a future promise to deliver the result
// of a FutureLiveTicketsResultAsync RPC invocation (or an applicable error).
type FutureLiveTicketsResult chan *response

// Receive waits for the response promised by the future and returns all
// currently missed tickets from the missed ticket database.
func (r FutureLiveTicketsResult) Receive() ([]*chainhash.Hash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a lddljson.LiveTicketsResult.
	var container lddljson.LiveTicketsResult
	err = json.Unmarshal(res, &container)
	if err != nil {
		return nil, err
	}

	liveTickets := make([]*chainhash.Hash, 0, len(container.Tickets))
	for _, ticketStr := range container.Tickets {
		h, err := chainhash.NewHashFromStr(ticketStr)
		if err != nil {
			return nil, err
		}
		liveTickets = append(liveTickets, h)
	}

	return liveTickets, nil
}

// LiveTicketsAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) LiveTicketsAsync() FutureLiveTicketsResult {
	cmd := lddljson.NewLiveTicketsCmd()
	return c.sendCmd(cmd)
}

// LiveTickets returns all currently missed tickets from the missed
// ticket database in the daemon.
//
// NOTE: This is a lddld extension.
func (c *Client) LiveTickets() ([]*chainhash.Hash, error) {
	return c.LiveTicketsAsync().Receive()
}

// FutureMissedTicketsResult is a future promise to deliver the result
// of a FutureMissedTicketsResultAsync RPC invocation (or an applicable error).
type FutureMissedTicketsResult chan *response

// Receive waits for the response promised by the future and returns all
// currently missed tickets from the missed ticket database.
func (r FutureMissedTicketsResult) Receive() ([]*chainhash.Hash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a lddljson.MissedTicketsResult.
	var container lddljson.MissedTicketsResult
	err = json.Unmarshal(res, &container)
	if err != nil {
		return nil, err
	}

	missedTickets := make([]*chainhash.Hash, 0, len(container.Tickets))
	for _, ticketStr := range container.Tickets {
		h, err := chainhash.NewHashFromStr(ticketStr)
		if err != nil {
			return nil, err
		}
		missedTickets = append(missedTickets, h)
	}

	return missedTickets, nil
}

// MissedTicketsAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
func (c *Client) MissedTicketsAsync() FutureMissedTicketsResult {
	cmd := lddljson.NewMissedTicketsCmd()
	return c.sendCmd(cmd)
}

// MissedTickets returns all currently missed tickets from the missed
// ticket database in the daemon.
//
// NOTE: This is a lddld extension.
func (c *Client) MissedTickets() ([]*chainhash.Hash, error) {
	return c.MissedTicketsAsync().Receive()
}

// FutureSessionResult is a future promise to deliver the result of a
// SessionAsync RPC invocation (or an applicable error).
type FutureSessionResult chan *response

// Receive waits for the response promised by the future and returns the
// session result.
func (r FutureSessionResult) Receive() (*lddljson.SessionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a session result object.
	var session lddljson.SessionResult
	err = json.Unmarshal(res, &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

// SessionAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See Session for the blocking version and more details.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) SessionAsync() FutureSessionResult {
	// Not supported in HTTP POST mode.
	if c.config.HTTPPostMode {
		return newFutureError(ErrWebsocketsRequired)
	}

	cmd := lddljson.NewSessionCmd()
	return c.sendCmd(cmd)
}

// Session returns details regarding a websocket client's current connection.
//
// This RPC requires the client to be running in websocket mode.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) Session() (*lddljson.SessionResult, error) {
	return c.SessionAsync().Receive()
}

// FutureTicketFeeInfoResult is a future promise to deliver the result of a
// TicketFeeInfoAsync RPC invocation (or an applicable error).
type FutureTicketFeeInfoResult chan *response

// Receive waits for the response promised by the future and returns the
// ticketfeeinfo result.
func (r FutureTicketFeeInfoResult) Receive() (*lddljson.TicketFeeInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarsal result as a ticketfeeinfo result object.
	var tfir lddljson.TicketFeeInfoResult
	err = json.Unmarshal(res, &tfir)
	if err != nil {
		return nil, err
	}

	return &tfir, nil
}

// TicketFeeInfoAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See TicketFeeInfo for the blocking version and more details.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) TicketFeeInfoAsync(blocks *uint32, windows *uint32) FutureTicketFeeInfoResult {
	// Not supported in HTTP POST mode.
	if c.config.HTTPPostMode {
		return newFutureError(ErrWebsocketsRequired)
	}

	// Avoid passing actual nil values, since they can cause arguments
	// not to pass. Pass zero values instead.
	if blocks == nil {
		blocks = &zeroUint32
	}
	if windows == nil {
		windows = &zeroUint32
	}

	cmd := lddljson.NewTicketFeeInfoCmd(blocks, windows)
	return c.sendCmd(cmd)
}

// TicketFeeInfo returns information about ticket fees.
//
// This RPC requires the client to be running in websocket mode.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) TicketFeeInfo(blocks *uint32, windows *uint32) (*lddljson.TicketFeeInfoResult, error) {
	return c.TicketFeeInfoAsync(blocks, windows).Receive()
}

// FutureTicketVWAPResult is a future promise to deliver the result of a
// TicketVWAPAsync RPC invocation (or an applicable error).
type FutureTicketVWAPResult chan *response

// Receive waits for the response promised by the future and returns the
// ticketvwap result.
func (r FutureTicketVWAPResult) Receive() (lddlutil.Amount, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarsal result as a ticketvwap result object.
	var vwap float64
	err = json.Unmarshal(res, &vwap)
	if err != nil {
		return 0, err
	}

	amt, err := lddlutil.NewAmount(vwap)
	if err != nil {
		return 0, err
	}

	return amt, nil
}

// TicketVWAPAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See TicketVWAP for the blocking version and more details.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) TicketVWAPAsync(start *uint32, end *uint32) FutureTicketVWAPResult {
	// Not supported in HTTP POST mode.
	if c.config.HTTPPostMode {
		return newFutureError(ErrWebsocketsRequired)
	}

	cmd := lddljson.NewTicketVWAPCmd(start, end)
	return c.sendCmd(cmd)
}

// TicketVWAP returns the vwap weighted average price of tickets.
//
// This RPC requires the client to be running in websocket mode.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) TicketVWAP(start *uint32, end *uint32) (lddlutil.Amount, error) {
	return c.TicketVWAPAsync(start, end).Receive()
}

// FutureTxFeeInfoResult is a future promise to deliver the result of a
// TxFeeInfoAsync RPC invocation (or an applicable error).
type FutureTxFeeInfoResult chan *response

// Receive waits for the response promised by the future and returns the
// txfeeinfo result.
func (r FutureTxFeeInfoResult) Receive() (*lddljson.TxFeeInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarsal result as a txfeeinfo result object.
	var tfir lddljson.TxFeeInfoResult
	err = json.Unmarshal(res, &tfir)
	if err != nil {
		return nil, err
	}

	return &tfir, nil
}

// TxFeeInfoAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See TxFeeInfo for the blocking version and more details.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) TxFeeInfoAsync(blocks *uint32, start *uint32, end *uint32) FutureTxFeeInfoResult {
	// Not supported in HTTP POST mode.
	if c.config.HTTPPostMode {
		return newFutureError(ErrWebsocketsRequired)
	}

	cmd := lddljson.NewTxFeeInfoCmd(blocks, start, end)
	return c.sendCmd(cmd)
}

// TxFeeInfo returns information about tx fees.
//
// This RPC requires the client to be running in websocket mode.
//
// NOTE: This is a Legenddigital extension.
func (c *Client) TxFeeInfo(blocks *uint32, start *uint32, end *uint32) (*lddljson.TxFeeInfoResult, error) {
	return c.TxFeeInfoAsync(blocks, start, end).Receive()
}

// FutureVersionResult is a future promise to deliver the result of a version
// RPC invocation (or an applicable error).
type FutureVersionResult chan *response

// Receive waits for the response promised by the future and returns the version
// result.
func (r FutureVersionResult) Receive() (map[string]lddljson.VersionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarsal result as a version result object.
	var vr map[string]lddljson.VersionResult
	err = json.Unmarshal(res, &vr)
	if err != nil {
		return nil, err
	}

	return vr, nil
}

// VersionAsync returns an instance of a type that can be used to get the result
// of the RPC at some future time by invoking the Receive function on the returned instance.
//
// See Version for the blocking version and more details.
func (c *Client) VersionAsync() FutureVersionResult {
	cmd := lddljson.NewVersionCmd()
	return c.sendCmd(cmd)
}

// Version returns information about the server's JSON-RPC API versions.
func (c *Client) Version() (map[string]lddljson.VersionResult, error) {
	return c.VersionAsync().Receive()
}
