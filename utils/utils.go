package utils

import "strconv"

// Type Defs
type Methods map[string]bool
type ServiceToMethod map[string]Methods
type stmarg struct {
	service string
	methods []string
}

// Methods
func NewSTM(service string, methods ...string) stmarg {
	return stmarg{service: service, methods: methods}
}

func NewServiceToMethod(services ...stmarg) ServiceToMethod {
	stm := ServiceToMethod{}
	for _, c := range services {
		stm[c.service] = NewMethods(c.methods...)
	}
	return stm
}

func (s ServiceToMethod) Contains(content string) bool {
	return s[content] != nil
}

func (s ServiceToMethod) ContainsMethod(service string, method string) bool {
	return s[service].Contains(method)
}

func (s ServiceToMethod) AddMethod(service string, method string) {
	s[service].Add(method)
}

func NewMethods(methods ...string) Methods {
	m := Methods{}
	for _, c := range methods {
		m[c] = true
	}
	return m
}

func (s Methods) Contains(content string) bool {
	return s[content]
}

func (s Methods) Add(content string) {
	s[content] = true
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

func StringToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
