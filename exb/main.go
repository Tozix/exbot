package exb

func NewTrade(PublicKey string, PrivateKey string) (*Keys, error) {
	return &Keys{
		PublicKey:  PublicKey,
		PrivateKey: PrivateKey,
	}, nil
}
