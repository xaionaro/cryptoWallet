package cryptoWalletInterfaces

import (
	"github.com/conejoninja/tesoro/pb/messages"
)

type Trezor interface {
	USBHIDWallet

	// CipherKeyValue is a function of symmetric encryption of key-value paris
	// using deterministic hierarchy
	//
	// See https://github.com/satoshilabs/slips/blob/master/slip-0011.md
	CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType)
}
