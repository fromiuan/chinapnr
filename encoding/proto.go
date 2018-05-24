package encoding

type CertFromPFX struct {
	Code             string `json:"Code"`
	Base64CertString string `json:"Base64CertString"`
}

type PrivateKeyIndexFromPFX struct {
	Code            string `json:"Code"`
	PrivateKeyIndex string `json:"privateKeyIndex"`
}

type P7SignMessageAttach struct {
	Code             string `json:"Code"`
	Base64CertString string `json:"Base64CertString"`
}

type P7VerifyMessageAttach struct {
	Base64Source        string `json:"Base64Source"`
	Base64SignatureCert string `json:"Base64SignatureCert"`
	Code                string `json:"Code"`
	Result              string `json:"Result"`
}
