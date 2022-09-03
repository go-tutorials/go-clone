package model

type User struct {
	Name   *string  `json:"name" yaml:"name"`
	Age    *int     `json:"age" yaml:"age"`
	Amount *float64 `json:"amount"  yaml:"amount"`
	Sex    *bool    `json:"sex" yaml:"sex"`
	Email  string   `json:"email" yaml:"email"`
	Phone  string   `json:"phone" yaml:"phone"`
	Ref1   string   `json:"ref1" yaml:"ref1"`
	Ref2   string   `json:"ref2" yaml:"ref2"`
	Ref3   string   `json:"ref3" yaml:"ref3"`
}

type RefConfig struct {
	Ref1 string `json:"ref1"`
	Ref2 string `json:"ref2"`
	Ref3 string `json:"ref3"`
}

type Request struct {
	Header Header `json:"header"`
	Body   Body   `json:"body"`
}

type Header struct {
	AppID          string `json:"appID"`
	AcceptLanguage string `json:"acceptLanguage"`
}

type Body struct {
	User        User    `json:"user"`
	AccountCode string  `json:"accountCode"`
	Amount      float64 `json:"amount"`
}
