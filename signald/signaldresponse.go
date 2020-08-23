package signald

// Response is a response to a request to signald, or a new inbound message
type Response struct {
	ID   string
	Data ResponseData
	Type string
}

// ResponseData is where most of the data in the response is stored.
type ResponseData struct {
	Groups          []Group
	Accounts        []Account
	URI             string
	DataMessage     DataMessage
	Message         string
	Username        string
	Source          string
	SourceDevice    int
	Type            int
	IsReceipt       bool
	Timestamp       float64
	ServerTimestamp float64
}

// Group represents a group in signal
type Group struct {
	GroupID  string
	Members  []string
	Name     string
	AvatarID int
}

// Account represents a user account registered to signald
type Account struct {
	Username   string
	DeviceID   int
	Filename   string
	Registered bool
	HasKeys    bool `json:"has_keys"`
	Subscribed bool
}

// DataMessage is the main component of incoming text messages
type DataMessage struct {
	Timestamp        float64
	Message          string
	ExpiresInSeconds float64
	GroupInfo        IncomingGroupInfo
}

// IncomingGroupInfo is information about a particular group
type IncomingGroupInfo struct {
	GroupID string
	Type    string
}
