package domain

type Existence struct {
	Exists    bool `json:"exists"`
}

type Outcome struct {
	Success    bool `json:"success"`
}

type OutcomeAndMsg struct {
	Success    bool `json:"success"`
	Message	string `json:"message"`
}

type MultiStatus struct {
	OutcomeAndMsgs []OutcomeAndMsg
}
