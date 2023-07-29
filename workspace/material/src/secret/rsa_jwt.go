package secret

import (
	"fmt"
	"os"

	"pensiel.com/material/src/static"
)

var (
	JWT_PUBLICKEY  []byte
	JWT_PRIVATEKEY []byte
)

func LoadSecretKeyJWT() {
	private, err := os.ReadFile(fmt.Sprintf("%s/rsa_jwt", static.CREDENTIAL_PATH))

	if err != nil {
		panic("can't load jwt private key")
	}

	JWT_PRIVATEKEY = private

	public, err := os.ReadFile(fmt.Sprintf("%s/rsa_jwt.pub", static.CREDENTIAL_PATH))

	if err != nil {
		panic("can't load jwt public key")
	}

	JWT_PUBLICKEY = public

}
