package memory

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (*Provider) Init() error {
	return nil
}
