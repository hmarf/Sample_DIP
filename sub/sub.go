package sub

type substruct struct {
}

type SubInterface interface {
	Sub() string
}

func NewSub() SubInterface {
	return &substruct{}
}

func (s *substruct) Sub() string {
	return "sub!!!"
}
