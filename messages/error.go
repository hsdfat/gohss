package messages

import (
	"fmt"
	"gohss/hss_models"

	"gohss/modules/go-diameter/v4/diam"
	"gohss/modules/go-diameter/v4/diam/datatype"
)

type ErrorCode int32

const (
	ErrorCode_UNDEFINED ErrorCode = 0
	// Default success code
	ErrorCode_MULTI_ROUND_AUTH        ErrorCode = 1001
	ErrorCode_SUCCESS                 ErrorCode = 2001
	ErrorCode_LIMITED_SUCCESS         ErrorCode = 2002
	ErrorCode_COMMAND_UNSUPPORTED     ErrorCode = 3001
	ErrorCode_UNABLE_TO_DELIVER       ErrorCode = 3002
	ErrorCode_REALM_NOT_SERVED        ErrorCode = 3003
	ErrorCode_TOO_BUSY                ErrorCode = 3004
	ErrorCode_LOOP_DETECTED           ErrorCode = 3005
	ErrorCode_REDIRECT_INDICATION     ErrorCode = 3006
	ErrorCode_APPLICATION_UNSUPPORTED ErrorCode = 3007
	ErrorCode_INVALID_HDR_BITS        ErrorCode = 3008
	ErrorCode_INVALID_AVP_BITS        ErrorCode = 3009
	ErrorCode_UNKNOWN_PEER            ErrorCode = 3010
	ErrorCode_AUTHENTICATION_REJECTED ErrorCode = 4001
	ErrorCode_OUT_OF_SPACE            ErrorCode = 4002
	ErrorCode_ELECTION_LOST           ErrorCode = 4003
	ErrorCode_AUTHORIZATION_REJECTED  ErrorCode = 5003
	// Permanent Failures 7.4.3
	ErrorCode_USER_UNKNOWN             ErrorCode = 5001
	ErrorCode_UNKNOWN_SESSION_ID       ErrorCode = 5002
	ErrorCode_UNKNOWN_EPS_SUBSCRIPTION ErrorCode = 5420
	ErrorCode_RAT_NOT_ALLOWED          ErrorCode = 5421
	ErrorCode_ROAMING_NOT_ALLOWED      ErrorCode = 5004
	ErrorCode_EQUIPMENT_UNKNOWN        ErrorCode = 5422
	ErrorCode_UNKNOWN_SERVING_NODE     ErrorCode = 5423
	// Transient Failures 7.4.4
	ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE ErrorCode = 4181
)

// ConvertAuthErrorToFailureMessage creates a corresponding diameter failure message for an auth error.
func ConvertAuthErrorToFailureMessage(err error, msg *diam.Message, sessionID datatype.UTF8String, serverCfg *hss_models.DiameterConfig) *diam.Message {
	switch err.(type) {
	case AuthRejectedError:
		return ConstructFailureAnswer(msg, sessionID, serverCfg, uint32(ErrorCode_AUTHORIZATION_REJECTED))
	case AuthDataUnavailableError:
		return ConstructFailureAnswer(msg, sessionID, serverCfg, uint32(ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE))
	default:
		return ConstructFailureAnswer(msg, sessionID, serverCfg, uint32(diam.UnableToComply))
	}
}

type AuthDataUnavailableError struct {
	msg string
}

func (err AuthDataUnavailableError) Error() string {
	return fmt.Sprintf("Authentication data unavailable: %s", err.msg)
}

// NewAuthDataUnavailableError creates an AuthDataUnavailableError.
func NewAuthDataUnavailableError(msg string) AuthDataUnavailableError {
	return AuthDataUnavailableError{msg: msg}
}

// AuthRejectedError indicates that the HSS cannot return any authentication
// vectors due to unallowed attachment of the UE. See 3GPP TS 29.272 section 5.2.3.1.3.
type AuthRejectedError struct {
	msg string
}

func (err AuthRejectedError) Error() string {
	return fmt.Sprintf("Authentication rejected: %s", err.msg)
}

// NewAuthRejectedError creates an AuthRejectedError.
func NewAuthRejectedError(msg string) AuthRejectedError {
	return AuthRejectedError{msg: msg}
}
