package host

// HostPosition keeps track of a set of host configuration and the start and
// end lines within the HostsCollection
type HostPosition struct {
	Host
	startLine int
	endLine   int
}

// MakeHostPosition makes a new one of these
func MakeHostPosition() HostPosition {
	return HostPosition{Host: MakeHost()}
}

// SetStartLine sets start line
func (hostPosition *HostPosition) SetStartLine(startLine int) {
	hostPosition.startLine = startLine
}

// StartLine get startLine
func (hostPosition HostPosition) StartLine() int {
	return hostPosition.startLine
}

// SetEndLine sets the end line
func (hostPosition *HostPosition) SetEndLine(endLine int) {
	hostPosition.endLine = endLine
}

// EndLine gets the end line
func (hostPosition HostPosition) EndLine() int {
	return hostPosition.endLine
}
