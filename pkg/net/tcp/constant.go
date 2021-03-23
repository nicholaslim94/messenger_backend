package tcp

var (
	// TCP Request constant
	requestContacts            string = "requestContacts"
	updateContactLastRead      string = "updateContactLastRead"
	updateContactGroupLastRead string = "updateContactGroupLastRead"
	updateContactActive        string = "updateContactActive"
	addMessage                 string = "addMessage"
	requestMessages            string = "requestMessages"
	requestAccount             string = "requestAccount"
	addGroupMessage            string = "addGroupMessage"
	requestGroupMessages       string = "requestGroupMessages"
	requestGroup               string = "requestGroup"

	// TCP Response constant
	responseContacts      string = "responseContacts"
	responseMessages      string = "responseMessages"
	responseGroupMessages string = "responseGroupMessages"
	responseAccount       string = "responseAccount"
	responseGroup         string = "responseGroup"
)
