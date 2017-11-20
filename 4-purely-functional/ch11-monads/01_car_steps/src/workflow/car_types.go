package workflow

type CarData struct {
	Car       *Car   `json:"car"`
	Timestamp string `json:"timestamp,omitempty"`
}

type Options struct {
	Option1 string `json:"option_1"`
	Option2 string `json:"option_2,omitempty"`
	Option3 string `json:"option_3,omitempty"`
}

type Car struct {
	VIN string `json:"vin"`
	Make string `json:"make"`
	Model string `json:"model"`
	Options *Options `json:"options,omitempty"`
}

