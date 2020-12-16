package auth

// Credential 授权凭证
type Credential struct {
	AccessKey string
	SecretKey []byte
}

// NewCredential 构造凭证
func NewCredential(accessKey, secretKey string) *Credential {
	return &Credential{
		AccessKey: accessKey,
		SecretKey: []byte(secretKey),
	}
}
