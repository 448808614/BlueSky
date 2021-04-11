package bluesky

const (
	// Fail 因为某种原因导致登录失败（看日志）
	Fail byte = iota
)

func (s *BlueSky) login() byte {

	return Fail
}
