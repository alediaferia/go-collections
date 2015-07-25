package collections

type InterfacesMapper func (value interface{}) (interface{})
type InterfacesFilter func (value interface{}) (bool)
type InterfacesReducer func (a, b interface{}) (interface{})

type Interfaces struct {
	values []interface{}
}

func NewFromSlice(values []interface{}) (*Interfaces) {
	this := new(Interfaces)
	this.values = values
	return this
}

func (this *Interfaces) Map(mapper InterfacesMapper) (*Interfaces) {
	new_ := make([]interface{}, 0, len(this.values))
	for _, v := range this.values {
		new_ = append(new_, mapper(v))
	}

	return &Interfaces{ values: new_ }
}

func (this *Interfaces) Filter(filter InterfacesFilter) (*Interfaces) {
	new_ := make([]interface{}, 0, len(this.values))
	for _, v := range this.values {
		if filter(v) {
			new_ = append(new_, v)
		}
	}
	return &Interfaces{ values: new_ }
}

func (this *Interfaces) Reduce(identity interface{}, reducer InterfacesReducer) (interface{}) {
	res := identity
	for _, v := range this.values {
		res = reducer(res, v)
	}

	return res
}
