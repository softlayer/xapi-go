package xapi

// OpaqueRefIsEmpty checks to see if a field that should be an OpaqueRef is actually empty.  Sometimes it's an empty string (rare)
// but other times it's OpaqueRef:NULL which is good to know.
func OpaqueRefIsEmpty(a string) bool {
	if a == "OpaqueRef:NULL" || a == "" {
		return true
	}

	return false
}
