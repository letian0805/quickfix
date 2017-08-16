package internal

import "time"

//SessionSettings stores all of the configuration for a given session
type SessionSettings struct {
	ResetOnLogon                 bool
	RefreshOnLogon               bool
	ResetOnLogout                bool
	ResetOnDisconnect            bool
	HeartBtInt                   time.Duration
	SessionTime                  *TimeRange
	InitiateLogon                bool
	ResendRequestChunkSize       int
	EnableLastMsgSeqNumProcessed bool
	SkipCheckLatency             bool

	//required on logon for FIX.T.1 messages
	DefaultApplVerID string

	//specific to initiators
	ReconnectInterval    time.Duration
	SocketConnectAddress []string
}
