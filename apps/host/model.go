package host

//数据模型
type Host struct {
	Id   string
	Name string
}

func NewHost() *Host {
	return &Host{}
}
