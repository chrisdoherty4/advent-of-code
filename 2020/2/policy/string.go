package policy

type Validator interface {
	Validate(s string) bool
}

type StringPolicy struct {
	requirements []Validator
}

func NewStringPolicy() StringPolicy {
	return StringPolicy{
		requirements: []Validator{},
	}
}

func (p *StringPolicy) AddRequirement(requirement Validator) {
	if requirement == nil {
		return
	}

	p.requirements = append(p.requirements, requirement)
}

func (p *StringPolicy) IsValid(stringToValidate string) bool {
	for _, requirement := range p.requirements {
		if !requirement.Validate(stringToValidate) {
			return false
		}
	}

	return true
}
