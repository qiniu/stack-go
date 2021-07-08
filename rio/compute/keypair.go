package compute

import "github.com/qiniu/stack-go/components/client"

// KeyPair 密钥对类接口封装
type KeyPair struct {
	client *client.Client
}

// NewKeyPair 获得密钥对实例
func NewKeyPair(cli *client.Client) *KeyPair {
	return &KeyPair{client: cli}
}

// KeyPairInfo ..
type KeyPairInfo struct {
	KeyPairName string `json:"key_pair_name"`
	FingerPrint string `json:"finger_print"`
}
