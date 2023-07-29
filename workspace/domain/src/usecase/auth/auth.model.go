package auth

type (
	ParamAuthRegister struct {
		Username string
		Email    string
		Password string
	}

	ParamAuthLogin struct {
		Username string
		Password string
	}

	ParamAuthConfirmation struct {
		SecondaryId string
	}

	ParamAuthMe struct {
		SecondaryId string
	}

	ParamAuthCheckAlredyUserCredential struct {
		Email    string
		Username string
	}
)
