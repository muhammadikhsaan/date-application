package interaction

type (
	ParamGetUserInteraction struct {
		UserID string
	}
)

type (
	ParamSaveInteraction struct {
		UserID   string
		TargetID string
		Action   string
	}
)
