package e2e

type TestService struct {
	suites []*TestSuite
}

type TestServiceOption func(*TestService)

func NewTestService(opts ...TestServiceOption) *TestService {
	s := &TestService{
		suites: []*TestSuite{},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithTestSuite(ts *TestSuite) TestServiceOption {
	return func(s *TestService) {
		s.suites = append(s.suites, ts)
	}
}
