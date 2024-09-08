package msgutil

type Data map[string]interface{}

type Msg struct {
	Data Data
}

func NewMessage() Msg {
	return Msg{
		Data: make(Data),
	}
}

func (m Msg) Set(key string, value interface{}) Msg {
	m.Data[key] = value
	return m
}

func (m Msg) Done() Data {
	return m.Data
}

func RequestBodyParseErrorResponseMsg() Data {
	return NewMessage().Set("message", "Failed to parse request body").Done()
}

func JwtCreateErrorMsg() Data {
	return NewMessage().Set("message", "Failed to create JWT token").Done()
}

func SomethingWentWrongMsg() Data {
	return NewMessage().Set("message", "Something went wrong").Done()
}

func ExpectationFailedMsg() Data {
	return NewMessage().Set("message", "Expectation failed").Done()
}

func AccessForbiddenMsg() Data {
	return NewMessage().Set("message", "Access forbidden").Done()
}
