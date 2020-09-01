package signald

// Response is a response to a request to signald, or a new inbound message
type Response struct {
	ID    string
	Data  ResponseData
	Type  string
	Error error `json:"Error,omitempty"`
}

// RawResponse is a response to a request to signald, or a new inbound message
type RawResponse struct {
	JSON  map[string]interface{}
	Data  string
	Error error
	Done  bool
}

// ResponseData is where most of the data in the response is stored.
type ResponseData struct {
	Groups        []Group
	Accounts      []Account
	Contacts      []ContactInfo
	Identities    []Identity
	SendResults   []SendResult
	StatusMessage StatusMessage
	UserDetails   ContactTokenDetails
	Profile       Profile
	Timestamp     string
}

// Group represents a group in signal
type Group struct {
	GroupID  string
	Members  []RequestAddress
	Name     string
	Type     string
	AvatarID int
}

// Account represents a user account registered to signald
type Account struct {
	Username   string
	DeviceID   int
	Filename   string
	Registered bool
	Subscribed bool
	HasKeys    bool `json:"has_keys"`
}

// DataMessage is the main component of incoming text messages
type DataMessage struct {
	Timestamp        float64
	Message          string
	ExpiresInSeconds float64
	GroupInfo        struct {
		GroupID string
		Type    string
	}
}

// ContactInfo this is information about a contact
type ContactInfo struct {
	Name                  string
	Address               RequestAddress
	Color                 string
	ProfileKey            string
	MessageExpirationTime int
	InboxPosition         int64
}

// Identity this is information about an identity
type Identity struct {
	Address      RequestAddress
	Added        int64
	Fingerprint  string
	TrustLevel   string `json:"trust_level"`
	SafetyNumber string `json:"safety_number"`
}

// SendResult result of send command
type SendResult struct {
	Address RequestAddress
	Success struct {
		Unidentified bool
		NeedsSync    bool
	}
	NetworkFailure      bool
	UnregisteredFailure bool
}

// StatusMessage command status result
type StatusMessage struct {
	Message   string
	Error     bool
	Request   Request
	MsgNumber int `json:"msg_number"`
}

// ContactTokenDetails contact token details
type ContactTokenDetails struct {
	Token  string
	Relay  string
	Number string
	Voice  bool
	Video  bool
}

// Profile profile info
type Profile struct {
	Name                           string
	Avatar                         string
	IdentityKey                    string `json:"identity_key"`
	UnidentifiedAccess             string `json:"unidentified_access"`
	UnrestrictedUnidentifiedAccess bool   `json:"unrestricted_unidentified_access"`
}
