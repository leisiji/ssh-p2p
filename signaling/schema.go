package signaling

// URI default signaling server
const URI = "http://server.natappfree.cc:38230"

// ConnectInfo SDP by offer or answer
type ConnectInfo struct {
	Source string `json:"source"`
	SDP    string `json:"sdp"`
}
