package scanner

// https://gist.github.com/daneharrigan/dd51c0e02d156c0259ba

type Scanner struct {
	b []byte
	e error
	s int
	f int
	
}

func New(b []byte) *Scanner {
	return &Scanner{b, nil, -1, -1}
}

func (s *Scanner) Err() error {
	return s.e
}

func (s *Scanner) Bytes() []byte {
	return s.b[s.s:s.f]
}

func (s *Scanner) Scan() bool {
	return false
}
