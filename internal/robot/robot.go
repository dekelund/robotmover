package robot

type robotError string

func (e robotError) Error() string {
	return string(e)
}
