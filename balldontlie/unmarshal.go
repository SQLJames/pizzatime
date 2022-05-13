package balldontlie

type unmarshalling interface {
	unmarshal() any
}

func unmarshal(data unmarshalling) any {
	return data.unmarshal()
}
