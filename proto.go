package seabotserver

// структуры для обмена данными между ботами и сервером.

// FROM BOT = FB
type FBBvb struct {
	Place int   `json:"place"`
	Ships []int `json:"ships"`
}

type FBTurn struct {
	Shot [2]int `json:"shot"`
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
	Bvb  *TBBvb  `json:"bvb,omitempty"`
	Turn *TBTurn `json:"turn,omitempty"`
}

type TBTurn struct {
	ID     int64 `json:"id,omitempty"`
	Result int64 `json:"result,omitempty"`
}

type TBBvb struct {
	Wait  int       `json:"wait,omitempty"`
	ID    int64     `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Ships *[100]int `json:"ships,omitempty"`
}

//type ShipPlaces struct {
//	Palubs int
//	Places [][2]int
//}

type TBAuth struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	ID    int64  `json:"id"`
}
