package sign

import (
	"crypto/ecdsa"
	"fmt"

	"golang.org/x/xerrors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rockiecn/middleware-test/upload/challenge"
)

// sign message with sk
// return message and sig
func Sign(sk string) (message string, sig string, err error) {
	//address := flag.String("address", "0xE7E9f12f99aD17d4786b9B1247C097e63ceaF8Db", "the login address")
	//secretKey := flag.String("sk", "", "the sk to signature") // 账户address的私钥

	// privateKey, err := crypto.GenerateKey()
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", xerrors.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// get address from pk
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// get message
	message, err = challenge.ChallengeFunc(address)
	if err != nil {
		return "", "", err
	}
	//fmt.Println("message with challenge: ", message)

	// hash of message
	hash := crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	ecdsaSK, err := crypto.HexToECDSA(sk)
	if err != nil {
		return "", "", err
	}
	// get sig
	signature, err := crypto.Sign(hash, ecdsaSK)
	if err != nil {
		return "", "", err
	}

	// to string
	sig = hexutil.Encode(signature)

	return message, sig, nil
}
