package ping

type Helper interface {
	Run() error
	CheckIP(string) (string, error)
	CheckPort(string, int) error
	CheckPorts(string, []int) error
}
