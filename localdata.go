package seabotserver

type QueueData struct {
	Bot  *TcpBot
	Bvb  *FBBvb
	Exit bool
}
