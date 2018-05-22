package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, Kubernetes Seattle Folks! You said: %s", string(req))
}
