package errorhandler

type NotFoundErr struct {
	Message string
}

type UnauthorizedErr struct {
	Message string
}

type BadRequestErr struct {
	Message string
}
type InternalServerErr struct {
	Message string
}

func (e *NotFoundErr) Error() string {
	return e.Message
}
func (e *UnauthorizedErr) Error() string {
	return e.Message
}
func (e *BadRequestErr) Error() string {
	return e.Message
}
func (e *InternalServerErr) Error() string {
	return e.Message
}
