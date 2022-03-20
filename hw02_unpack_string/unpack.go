package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString      = errors.New("invalid string")
	ErrUnusableSymbol     = errors.New("unusable symbol")
	ErrUnusableSpecSymbol = errors.New("unusable special symbol")
	ErrInvalidTail        = errors.New("invalid tail")
	ErrInvalidDigit       = errors.New("invalid digit")
)

const EOT rune = 0
const (
	START = iota
	CHAR
	SLASH
	ERROR
)

type Parserstate struct {
	state int             // parser state
	last  rune            // last rune
	sb    strings.Builder // assembled string
}

func NewParser() *Parserstate {
	p := Parserstate{state: START, last: EOT}
	return &p
}

func (p *Parserstate) GetString() (string, error) {
	// emit EOT and get final string
	switch p.state {
	case START: // last lexemma was finalized
		return p.sb.String(), nil
	case CHAR: // complete lexemma and get out
		p.sb.WriteRune(p.last)
		return p.sb.String(), nil
	default: // something wrong
		return "", ErrInvalidTail
	}
}

func (p *Parserstate) Parser(r rune) error {
	// parse current symbol and switch state
	switch p.state {
	case START: // expect: digit, slash, letter
		switch {
		case unicode.IsDigit(r):
			return ErrInvalidString
		case r == '\\':
			if p.last != EOT {
				p.sb.WriteRune(p.last)
				p.last = EOT
			}
			p.state = SLASH
		case unicode.IsLetter(r):
			p.last = r
			p.state = CHAR
		default:
			return ErrUnusableSymbol
		}
	case CHAR: // one char was got, expect: digit, slash, letter
		switch {
		case unicode.IsLetter(r):
			if p.last != EOT {
				p.sb.WriteRune(p.last)
			}
			p.last = r
			p.state = CHAR
		case r == '\\':
			if p.last != EOT {
				p.sb.WriteRune(p.last)
			}
			p.state = SLASH
		case unicode.IsDigit(r):
			i, err := strconv.Atoi(string(r))
			if err != nil {
				return ErrInvalidDigit
			}
			p.sb.WriteString(strings.Repeat(string(p.last), i))
			p.state = START
		default:
			return ErrUnusableSymbol
		}
	case SLASH: // one slash was got, expect: digit, slash and a few special symbol's letter
		switch {
		case unicode.IsDigit(r), r == '\\':
			p.last = r
		case r == 's':
			p.last = ' '
		case r == 't':
			p.last = '\t'
		case r == 'n':
			p.last = '\n'
		case r == 'r':
			p.last = '\r'
		default:
			return ErrUnusableSpecSymbol
		}
		p.state = CHAR
	}
	return nil
}

func Unpack(s string) (string, error) {
	p := NewParser()
	for _, r := range s {
		err := p.Parser(r)
		if err != nil {
			return "", err
		}
	}
	s, err := p.GetString()
	if err != nil {
		return "", err
	}
	return s, nil
}
