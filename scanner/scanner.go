package scanner

import "errors"

import "fmt"

// https://gist.github.com/daneharrigan/dd51c0e02d156c0259ba

type Scanner struct {
	b []byte
	e error
	c bool
	s int
	f int
	i int
}

var (
	ErrArgumentQuoting    = errors.New("incomplete argument quoting")
	ErrUnexpectedCR       = errors.New(`unexpected \r character`)
	ErrUnexpectedCmdChr   = errors.New("unexpected character in command name")
	ErrUnexpectedEscaping = errors.New("unexpected escaping in argument")
	ErrInvalidDataType    = errors.New("invalid data type")
)

func New(b []byte) *Scanner {
	return &Scanner{b: b, c: true}
}

func (s *Scanner) Scan() bool {
	if s.c {
		s.s, s.f, s.i, s.e = scanCommand(s.b)
	} else {
		s.s, s.f, s.i, s.e = scanArgument(s.b, s.i)
	}

	if s.e != nil {
		return false
	}

	fmt.Printf("i=%d len=%d\n", s.i, len(s.b))
	return s.i+1 == len(s.b)
}

func (s *Scanner) Err() error {
	return s.e
}

func (s *Scanner) Bytes() []byte {
	return s.b[s.s:s.f]
}

func scanCommand(b []byte) (int, int, int, error) {
	s := -1
	f := -1
	i := 0

	for n, c := range b {
		i = n
		switch {
		case c == ' ' || c == '\t' || c == '\n':
			fmt.Printf("s=%d f=%d\n", s, f)
			if f > 0 {
				return s, f, i, nil
			}

			continue
		case c >= 65 && c <= 90:
			if s < 0 {
				s = i
			}

			f = i+1
		case c == '\r':
			return 0, 0, i, ErrUnexpectedCR
		default:
			return 0, 0, i, ErrUnexpectedCmdChr
		}
	}

	return s, f, i, nil
}

func scanArgument(b []byte, i int) (int, int, int, error) {
	var isQuoted, isEscaped bool
	s := -1
	f := -1

	for _, c := range b {
		i++
		if isEscaped {
			isEscaped = false
			continue
		}

		switch {
		case c == '"' && !isQuoted:
			s = i
			isQuoted = true
		case c == '"' && isQuoted:
			f = i
			return s, f, i, nil
		case c == '\\':
			if !isQuoted {
				return 0, 0, 0, ErrUnexpectedEscaping
			}

			isEscaped = true
			continue
		}
	}

	if isQuoted {
		return 0, 0, 0, ErrArgumentQuoting
	}

	return s, f, i, nil
}

/*

func (s *Scanner) Scan() bool {
	var c byte
	s.s = -1
	s.f = -1
whitespace:
	if len(s.b) == s.i {
		if s.q {
			s.e = ErrArgumentQuoting
			return false
		}

		if s.f < len(s.b) {
			s.f = len(s.b)
		fmt.Printf("bytes=%q\n", s.Bytes())
			return false
		}

		fmt.Printf("bytes=%q\n", s.Bytes())
		return false
	} else {
		c = s.b[s.i]
	}

	switch c {
	case '\n', '\t', ' ':
		if !s.q && s.f < 0 {
			s.f = s.i
		}

		if s.f > s.s {
			return true
		}

		s.i++
		goto whitespace
	case '\r':
		s.e = ErrUnexpectedCR
		return false
	default:
		if s.c {
			goto command
		}

		goto argument
	}  

	return false
command:
	c = s.b[s.i]
	switch {
	case c == '\n' || c == '\t' || c == ' ':
		s.f = s.i
		s.c = false
		goto whitespace
	case c >= 65 && c <= 90:
		if s.s < 0 {
			s.s = s.i
		}
		s.i++
		goto command
	default:
		s.e = ErrUnexpectedCmdChr
		return false
	}
argument:
	c = s.b[s.i]
	switch {
	case c == '"' && !s.q:
		s.s = s.i+1
		s.q = true
	case c == '"' && s.q:
		s.q = false
		s.f = s.i
	case c == '\\':
		if !s.q {
			s.e = ErrUnexpectedEscaping
			return false
		}

		s.i+=2 // skip over escaped char
		goto argument
	default:
		if s.s < 0 {
			s.s = s.i
		}
	}

	s.i++
	goto whitespace
}
*/
