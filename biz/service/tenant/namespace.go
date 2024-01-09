package tenant

func NewNamespace(name string) *Namespace {
	return &Namespace{
		APIVersion: "v1",
		Kind:       "Namespace",
		Metadata: Metadata{
			Name: name,
		},
	}
}
