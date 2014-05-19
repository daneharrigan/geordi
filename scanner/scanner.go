package scanner

import (
	"errors"
	"github.com/daneharrigan/geordi/types"
)

var (
	ErrArgumentQuoting    = errors.New("incomplete argument quoting")
	ErrUnexpectedCR       = errors.New(`unexpected \r character`)
	ErrUnexpectedCmdChr   = errors.New("unexpected character in command name")
	ErrUnexpectedEscaping = errors.New("unexpected escaping in argument")
	ErrInvalidDecimal     = errors.New("invalid decimal placement")
	ErrInvalidDataType    = errors.New("invalid data type")
)

type Scanner struct {
	bytes []byte
	cmd   bool
	err   error
	i     int
	s     int
	f     int
	t     types.Type
}

func New(b []byte) *Scanner {
	return &Scanner{bytes: b, cmd: true}
}

func (s *Scanner) Scan() bool {
	if s.err != nil ||  s.i == len(s.bytes) {
		return false
	}

	if s.cmd {
		s.command()
		s.cmd = false
	} else {
		s.argument()
	}

	if s.err != nil ||  s.s < 0 {
		return false
	}

	return true
}

func (s *Scanner) Err() error {
	return s.err
}

func (s *Scanner) Type() types.Type {
	return s.t
}

func (s *Scanner) Bytes() []byte {
	return s.bytes[s.s:s.f]
}

func (s *Scanner) command() {
	s.t = types.String
	s.s = -1
	for _, c := range s.bytes {
		s.i++
		switch {
		case c == ' ' || c == '\t' || c == '\n':
			if s.s < 0 {
				continue
			}

			s.f = s.i - 1
			return
		case c == '\r':
			s.err = ErrUnexpectedCR
			return
		case c >= 'A' && c <= 'Z':
			if s.s < 0 {
				s.s = s.i - 1
			}
			continue
		default:
			s.err = ErrUnexpectedCmdChr
			return
		}
	}

	if s.f < s.i {
		s.f = s.i
	}
}

func (s *Scanner) argument() {
	s.s = -1
	s.t = -1

	var quoted bool
	var escaped bool

	for s.i < len(s.bytes) {
		c := s.bytes[s.i]
		s.i++
		if escaped {
			escaped = false
			continue
		}

		switch {
		case c == ' ' || c == '\t' || c == '\n':
			if !quoted {
				s.f = s.i - 1
				if s.s > 0 {
					return
				}
			}
		case c == '\r' && !quoted:
			s.err = ErrUnexpectedCR
			return
		case c == '"' && !quoted:
			quoted = true
			s.t = types.String
			s.s = s.i
			continue
		case c == '"' && quoted:
			quoted = false
			s.f = s.i - 1
			return
		case c == '\\':
			if !quoted {
				s.err = ErrUnexpectedEscaping
				return
			}

			escaped = true
			continue
		case (c >= '0' && c <= '9') || c == '.' || c == '-':
			if quoted {
				continue
			}

			if s.s < 0 {
				s.s = s.i - 1
			}

			if c == '.' {
				if s.t == types.Float {
					s.err = ErrInvalidDecimal
					return
				}

				s.t = types.Float
				continue
			}

			if s.t != types.Float {
				s.t = types.Int
			}
		default:
			if !quoted {
				s.err = ErrArgumentQuoting
				return
			}
		}
	}

	if quoted {
		s.err = ErrArgumentQuoting
		return
	}

	if s.f < s.i {
		s.f = s.i
	}
}
