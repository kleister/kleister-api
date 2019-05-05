package user

type MetricsOptions struct {
	Service Service
	Metrics string
}

func NewMetrics(opts MetricsOptions) Service {
	return &metrics{
		service: opts.Service,
		metrics: opts.Metrics,
	}
}

type metrics struct {
	service Service
	metrics string
}

func (m *metrics) ListUser() {

}

func (m *metrics) CreateUser() {

}

func (m *metrics) UpdateUser() {

}

func (m *metrics) DeleteUser() {

}

func (m *metrics) ShowUser() {

}
