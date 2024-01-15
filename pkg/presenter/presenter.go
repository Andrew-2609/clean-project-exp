package presenter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

const (
	JSON = "application/json"
	XML  = "application/xml"
)

type Presenter struct {
	Accept     string
	Payload    any
	formatters map[string]func(w http.ResponseWriter) error
}

func NewPresenter(accept string) (*Presenter, error) {
	switch accept {
	case JSON, XML:
		presenter := &Presenter{Accept: accept}
		presenter.initFormatters()
		return presenter, nil
	default:
		return nil, errors.New("invalid value for Accept header")
	}
}

func (p *Presenter) initFormatters() {
	p.formatters = map[string]func(w http.ResponseWriter) error{
		JSON: p.toJSON,
		XML:  p.toXML,
	}
}

func (p *Presenter) SetPayload(payload any) {
	p.Payload = payload
}

func (p *Presenter) FormatResponse(w http.ResponseWriter) error {
	if p.Payload == nil {
		return errors.New("payload of presenter is empty")
	}

	return p.formatters[p.Accept](w)
}

func (p *Presenter) toJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", JSON)

	if err := json.NewEncoder(w).Encode(p.Payload); err != nil {
		return errors.New("could not parse server response")
	}

	return nil
}

func (p *Presenter) toXML(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", XML)

	if err := xml.NewEncoder(w).Encode(p.Payload); err != nil {
		return errors.New("could not parse server response")
	}

	return nil
}
