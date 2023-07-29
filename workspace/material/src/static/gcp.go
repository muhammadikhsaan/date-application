package static

type (
	uploadpayload struct {
		Target string
		Rules  uploadrules
	}

	uploadrules struct {
		MaxSize   uint
		Extention []string
	}
)

var (
	FILE_UPLOAD_GUIDE = map[string]uploadpayload{
		"CUgYnJKNdM": {
			Target: "user-profile",
			Rules: uploadrules{
				MaxSize:   2 * 1024 * 1024,
				Extention: []string{"jpg", "png", "jpeg"},
			},
		},
	}
)
