package constant

// ClientFailureStatusCode - status code
const (
	ClientFailureStatusCode = "ECHO-100400"
)

// statusText - map of status code to status text
var StatusText = map[string]string{
	"ECHO-100400": "Request parameters are invalid",
}

// HttpStatusCodes - map of status code to HTTP status code
var HttpStatusCodes = map[string]int{
	"ECHO-100400": 400,
}
