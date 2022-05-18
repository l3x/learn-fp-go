package maybe

type SuccessOrFailure interface {
	Success() bool
	Failure() bool
}
