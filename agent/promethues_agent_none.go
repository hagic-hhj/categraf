//go:build no_prometheus

package agent

type PrometheusAgent struct {
}

func NewPrometheusAgent() AgentModule {
	return &PrometheusAgent{}
}

func (pa *PrometheusAgent) Start() error {
	return nil
}

func (pa *PrometheusAgent) Stop() error {
	return nil
}
