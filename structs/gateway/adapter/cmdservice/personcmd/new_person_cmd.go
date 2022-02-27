package personcmd

import "github.com/booscaaa/rtdd-golang/gateway/core/domain"

type cmd struct {
	usecase domain.PersonUseCase
}

// NewPersonCMD .
func NewPersonCMD(
	usecase domain.PersonUseCase,
) domain.PersonCMD {
	return &cmd{
		usecase: usecase,
	}
}
