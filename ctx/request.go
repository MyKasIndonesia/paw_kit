package ctx

import "context"

var (
	contextKeyRequestID = contextKey("RequestID")
)

// RequestID return value of RequestID attached in the context
// Empty string will be returned if no RequestID attached in the context
func RequestID(c context.Context) string {
	RequestID := c.Value(contextKeyRequestID)
	if RequestID == nil {
		return ""
	}
	return RequestID.(string)
}

// SetRequestID return context with value requestID attached on them
// Accept empty string as the value of requestID
func SetRequestID(c context.Context, requestID string) context.Context {
	return context.WithValue(c, contextKeyRequestID, requestID)
}
