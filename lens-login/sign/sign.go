package sign

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// sign message with sk
// return message and sig
func Sign(message, sk string) (sig string, err error) {
	//address := flag.String("address", "0xE7E9f12f99aD17d4786b9B1247C097e63ceaF8Db", "the login address")
	//secretKey := flag.String("sk", "", "the sk to signature") // 账户address的私钥

	// hash of message
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	ecdsaSK, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", err
	}
	// get sig
	signature, err := crypto.Sign(hash, ecdsaSK)
	if err != nil {
		return "", err
	}

	// to string
	sig = hexutil.Encode(signature)

	return sig, nil
}
