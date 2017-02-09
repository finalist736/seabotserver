package seabotserver

// FROM BOT = FB
type FBBvb struct {
	Place int   `json:"place"`
	Ships []int `json:"ships"`
}

type FBTurn struct {
	Field []int `json:"field"`
}

type FromBot struct {
	Auth string  `json:"auth"`
	Exit bool    `json:"exit"`
	Bvb  *FBBvb  `json:"bvb"`
	Turn *FBTurn `json:"turn"`
}

// TO BOT = TB

type ToBot struct {
	Auth *TBAuth `json:"auth,omitempty"`
}

type TBAuth struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	ID    int64  `json:"id"`
}
