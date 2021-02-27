package memory

type Provider struct{}

func New() *Provider {
	return &Provider{}
}

func (*Provider) Init() error {
	initUserStorage()
	initPassengerStorage()
	initDriverStorage()
	initMappingStorage()
	return nil
}
