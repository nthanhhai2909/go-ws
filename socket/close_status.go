package socket

import "github.com/nthanhhai2909/go-commons-lang/stringutils"

type CloseStatus struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

// TODO HGA WILL ADD MORE STATUS HERE

// Normal - Code 1000 indicates a normal closure, meaning that the purpose for which the connection
// was established has been fulfilled.
var Normal = CloseStatus{Code: 1000, Reason: stringutils.EMPTY}

// GoingAway - Code 1001 indicates that an endpoint is "going away", such as a server going down
// or a browser having navigated away from a page.
var GoingAway = CloseStatus{Code: 1001, Reason: stringutils.EMPTY}

// ProtocolError - Code 1002 indicates that an endpoint is terminating the connection due to a protocol error.
var ProtocolError = CloseStatus{Code: 1002, Reason: stringutils.EMPTY}

// NotAcceptable - Code 1003 indicates that an endpoint is terminating the connection because it has received a type of data
// it cannot accept (e.g., an endpoint that understands only text data MAY send this if it receives a binary smsg).
var NotAcceptable = CloseStatus{Code: 1003, Reason: stringutils.EMPTY}

// NoStatusCode - Code 1004 is a reserved value and MUST NOT be set as a status code in a Close control frame by an endpoint.
// It is designated for use in applications expecting a status code to indicate that no status code was actually present.
var NoStatusCode = CloseStatus{Code: 1004, Reason: stringutils.EMPTY}
