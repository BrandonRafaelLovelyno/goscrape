// This file is generated by "./lib/proto/generate"

package proto

/*

Tethering

The Tethering domain defines methods and events for browser port binding.

*/

// TetheringBind Request browser port binding.
type TetheringBind struct {
	// Port number to bind.
	Port int `json:"port"`
}

// ProtoReq name.
func (m TetheringBind) ProtoReq() string { return "Tethering.bind" }

// Call sends the request.
func (m TetheringBind) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TetheringUnbind Request browser port unbinding.
type TetheringUnbind struct {
	// Port number to unbind.
	Port int `json:"port"`
}

// ProtoReq name.
func (m TetheringUnbind) ProtoReq() string { return "Tethering.unbind" }

// Call sends the request.
func (m TetheringUnbind) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TetheringAccepted Informs that port was successfully bound and got a specified connection id.
type TetheringAccepted struct {
	// Port number that was successfully bound.
	Port int `json:"port"`

	// ConnectionID Connection id to be used.
	ConnectionID string `json:"connectionId"`
}

// ProtoEvent name.
func (evt TetheringAccepted) ProtoEvent() string {
	return "Tethering.accepted"
}
