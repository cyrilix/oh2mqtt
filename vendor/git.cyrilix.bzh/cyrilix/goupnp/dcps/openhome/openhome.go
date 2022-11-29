// Client for UPnP Device Control Protocol OpenHome.
//
// This DCP is documented in detail at:
// - http://wiki.openhome.org/wiki/Av:Developer
//
// Typically, use one of the New* functions to create clients for services.
package openhome

// ***********************************************************
// GENERATED FILE - DO NOT EDIT BY HAND. See README.md
// ***********************************************************

import (
	"context"
	"net/url"
	"time"

	"git.cyrilix.bzh/cyrilix/goupnp"
	"git.cyrilix.bzh/cyrilix/goupnp/soap"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const ()

// Service URNs:
const (
	URN_Credentials_1          = "urn:av-openhome-org:service:Credentials:1"
	URN_Debug_1                = "urn:av-openhome-org:service:Debug:1"
	URN_Exakt_1                = "urn:av-openhome-org:service:Exakt:1"
	URN_Exakt_2                = "urn:av-openhome-org:service:Exakt:2"
	URN_Exakt_3                = "urn:av-openhome-org:service:Exakt:3"
	URN_Info_1                 = "urn:av-openhome-org:service:Info:1"
	URN_MediaServer_1          = "urn:av-openhome-org:service:MediaServer:1"
	URN_NetworkMonitor_1       = "urn:av-openhome-org:service:NetworkMonitor:1"
	URN_Pins_1                 = "urn:av-openhome-org:service:Pins:1"
	URN_Playlist_1             = "urn:av-openhome-org:service:Playlist:1"
	URN_PlaylistManager_1      = "urn:av-openhome-org:service:PlaylistManager:1"
	URN_Product_1              = "urn:av-openhome-org:service:Product:1"
	URN_Product_2              = "urn:av-openhome-org:service:Product:2"
	URN_Radio_1                = "urn:av-openhome-org:service:Radio:1"
	URN_Receiver_1             = "urn:av-openhome-org:service:Receiver:1"
	URN_Sender_1               = "urn:av-openhome-org:service:Sender:1"
	URN_Sender_2               = "urn:av-openhome-org:service:Sender:2"
	URN_SubscriptionLongPoll_1 = "urn:av-openhome-org:service:SubscriptionLongPoll:1"
	URN_Time_1                 = "urn:av-openhome-org:service:Time:1"
	URN_Transport_1            = "urn:av-openhome-org:service:Transport:1"
	URN_Volume_1               = "urn:av-openhome-org:service:Volume:1"
	URN_Volume_2               = "urn:av-openhome-org:service:Volume:2"
	URN_Volume_3               = "urn:av-openhome-org:service:Volume:3"
)

// Credentials1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Credentials:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Credentials1 struct {
	goupnp.ServiceClient
}

// NewCredentials1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewCredentials1ClientsCtx(ctx context.Context) (clients []*Credentials1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Credentials_1); err != nil {
		return
	}
	clients = newCredentials1ClientsFromGenericClients(genericClients)
	return
}

// NewCredentials1Clients is the legacy version of NewCredentials1ClientsCtx, but uses
// context.Background() as the context.
func NewCredentials1Clients() (clients []*Credentials1, errors []error, err error) {
	return NewCredentials1ClientsCtx(context.Background())
}

// NewCredentials1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewCredentials1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Credentials1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Credentials_1)
	if err != nil {
		return nil, err
	}
	return newCredentials1ClientsFromGenericClients(genericClients), nil
}

// NewCredentials1ClientsByURL is the legacy version of NewCredentials1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewCredentials1ClientsByURL(loc *url.URL) ([]*Credentials1, error) {
	return NewCredentials1ClientsByURLCtx(context.Background(), loc)
}

// NewCredentials1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewCredentials1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Credentials1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Credentials_1)
	if err != nil {
		return nil, err
	}
	return newCredentials1ClientsFromGenericClients(genericClients), nil
}

func newCredentials1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Credentials1 {
	clients := make([]*Credentials1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Credentials1{genericClients[i]}
	}
	return clients
}

func (client *Credentials1) ClearCtx(
	ctx context.Context,
	Id string,
) (err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalString(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "Clear", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Clear is the legacy version of ClearCtx, but uses
// context.Background() as the context.
func (client *Credentials1) Clear(Id string) (err error) {
	return client.ClearCtx(context.Background(),
		Id,
	)
}

func (client *Credentials1) GetCtx(
	ctx context.Context,
	Id string,
) (UserName string, Password []byte, Enabled bool, Status string, Data string, err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalString(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		UserName string
		Password string
		Enabled  string
		Status   string
		Data     string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "Get", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if UserName, err = soap.UnmarshalString(response.UserName); err != nil {
		return
	}
	if Password, err = soap.UnmarshalBinBase64(response.Password); err != nil {
		return
	}
	if Enabled, err = soap.UnmarshalBoolean(response.Enabled); err != nil {
		return
	}
	if Status, err = soap.UnmarshalString(response.Status); err != nil {
		return
	}
	if Data, err = soap.UnmarshalString(response.Data); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Get is the legacy version of GetCtx, but uses
// context.Background() as the context.
func (client *Credentials1) Get(Id string) (UserName string, Password []byte, Enabled bool, Status string, Data string, err error) {
	return client.GetCtx(context.Background(),
		Id,
	)
}

func (client *Credentials1) GetIdsCtx(
	ctx context.Context,
) (Ids string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Ids string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "GetIds", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Ids, err = soap.UnmarshalString(response.Ids); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetIds is the legacy version of GetIdsCtx, but uses
// context.Background() as the context.
func (client *Credentials1) GetIds() (Ids string, err error) {
	return client.GetIdsCtx(context.Background())
}

func (client *Credentials1) GetPublicKeyCtx(
	ctx context.Context,
) (PublicKey string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PublicKey string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "GetPublicKey", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PublicKey, err = soap.UnmarshalString(response.PublicKey); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPublicKey is the legacy version of GetPublicKeyCtx, but uses
// context.Background() as the context.
func (client *Credentials1) GetPublicKey() (PublicKey string, err error) {
	return client.GetPublicKeyCtx(context.Background())
}

func (client *Credentials1) GetSequenceNumberCtx(
	ctx context.Context,
) (SequenceNumber uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SequenceNumber string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "GetSequenceNumber", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SequenceNumber, err = soap.UnmarshalUi4(response.SequenceNumber); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetSequenceNumber is the legacy version of GetSequenceNumberCtx, but uses
// context.Background() as the context.
func (client *Credentials1) GetSequenceNumber() (SequenceNumber uint32, err error) {
	return client.GetSequenceNumberCtx(context.Background())
}

func (client *Credentials1) LoginCtx(
	ctx context.Context,
	Id string,
) (Token string, err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalString(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Token string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "Login", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Token, err = soap.UnmarshalString(response.Token); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Login is the legacy version of LoginCtx, but uses
// context.Background() as the context.
func (client *Credentials1) Login(Id string) (Token string, err error) {
	return client.LoginCtx(context.Background(),
		Id,
	)
}

func (client *Credentials1) ReLoginCtx(
	ctx context.Context,
	Id string,
	CurrentToken string,
) (NewToken string, err error) {
	// Request structure.
	request := &struct {
		Id           string
		CurrentToken string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalString(Id); err != nil {
		return
	}
	if request.CurrentToken, err = soap.MarshalString(CurrentToken); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NewToken string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "ReLogin", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewToken, err = soap.UnmarshalString(response.NewToken); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ReLogin is the legacy version of ReLoginCtx, but uses
// context.Background() as the context.
func (client *Credentials1) ReLogin(Id string, CurrentToken string) (NewToken string, err error) {
	return client.ReLoginCtx(context.Background(),
		Id,
		CurrentToken,
	)
}

func (client *Credentials1) SetCtx(
	ctx context.Context,
	Id string,
	UserName string,
	Password []byte,
) (err error) {
	// Request structure.
	request := &struct {
		Id       string
		UserName string
		Password string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalString(Id); err != nil {
		return
	}
	if request.UserName, err = soap.MarshalString(UserName); err != nil {
		return
	}
	if request.Password, err = soap.MarshalBinBase64(Password); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "Set", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Set is the legacy version of SetCtx, but uses
// context.Background() as the context.
func (client *Credentials1) Set(Id string, UserName string, Password []byte) (err error) {
	return client.SetCtx(context.Background(),
		Id,
		UserName,
		Password,
	)
}

func (client *Credentials1) SetEnabledCtx(
	ctx context.Context,
	Id string,
	Enabled bool,
) (err error) {
	// Request structure.
	request := &struct {
		Id      string
		Enabled string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalString(Id); err != nil {
		return
	}
	if request.Enabled, err = soap.MarshalBoolean(Enabled); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Credentials_1, "SetEnabled", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetEnabled is the legacy version of SetEnabledCtx, but uses
// context.Background() as the context.
func (client *Credentials1) SetEnabled(Id string, Enabled bool) (err error) {
	return client.SetEnabledCtx(context.Background(),
		Id,
		Enabled,
	)
}

// Debug1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Debug:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Debug1 struct {
	goupnp.ServiceClient
}

// NewDebug1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewDebug1ClientsCtx(ctx context.Context) (clients []*Debug1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Debug_1); err != nil {
		return
	}
	clients = newDebug1ClientsFromGenericClients(genericClients)
	return
}

// NewDebug1Clients is the legacy version of NewDebug1ClientsCtx, but uses
// context.Background() as the context.
func NewDebug1Clients() (clients []*Debug1, errors []error, err error) {
	return NewDebug1ClientsCtx(context.Background())
}

// NewDebug1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewDebug1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Debug1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Debug_1)
	if err != nil {
		return nil, err
	}
	return newDebug1ClientsFromGenericClients(genericClients), nil
}

// NewDebug1ClientsByURL is the legacy version of NewDebug1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewDebug1ClientsByURL(loc *url.URL) ([]*Debug1, error) {
	return NewDebug1ClientsByURLCtx(context.Background(), loc)
}

// NewDebug1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewDebug1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Debug1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Debug_1)
	if err != nil {
		return nil, err
	}
	return newDebug1ClientsFromGenericClients(genericClients), nil
}

func newDebug1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Debug1 {
	clients := make([]*Debug1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Debug1{genericClients[i]}
	}
	return clients
}

func (client *Debug1) GetLogCtx(
	ctx context.Context,
) (Log string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Log string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Debug_1, "GetLog", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Log, err = soap.UnmarshalString(response.Log); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetLog is the legacy version of GetLogCtx, but uses
// context.Background() as the context.
func (client *Debug1) GetLog() (Log string, err error) {
	return client.GetLogCtx(context.Background())
}

func (client *Debug1) SendLogCtx(
	ctx context.Context,
	Data string,
) (err error) {
	// Request structure.
	request := &struct {
		Data string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Data, err = soap.MarshalString(Data); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Debug_1, "SendLog", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SendLog is the legacy version of SendLogCtx, but uses
// context.Background() as the context.
func (client *Debug1) SendLog(Data string) (err error) {
	return client.SendLogCtx(context.Background(),
		Data,
	)
}

// Exakt1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Exakt:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Exakt1 struct {
	goupnp.ServiceClient
}

// NewExakt1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewExakt1ClientsCtx(ctx context.Context) (clients []*Exakt1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Exakt_1); err != nil {
		return
	}
	clients = newExakt1ClientsFromGenericClients(genericClients)
	return
}

// NewExakt1Clients is the legacy version of NewExakt1ClientsCtx, but uses
// context.Background() as the context.
func NewExakt1Clients() (clients []*Exakt1, errors []error, err error) {
	return NewExakt1ClientsCtx(context.Background())
}

// NewExakt1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewExakt1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Exakt1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Exakt_1)
	if err != nil {
		return nil, err
	}
	return newExakt1ClientsFromGenericClients(genericClients), nil
}

// NewExakt1ClientsByURL is the legacy version of NewExakt1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewExakt1ClientsByURL(loc *url.URL) ([]*Exakt1, error) {
	return NewExakt1ClientsByURLCtx(context.Background(), loc)
}

// NewExakt1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewExakt1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Exakt1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Exakt_1)
	if err != nil {
		return nil, err
	}
	return newExakt1ClientsFromGenericClients(genericClients), nil
}

func newExakt1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Exakt1 {
	clients := make([]*Exakt1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Exakt1{genericClients[i]}
	}
	return clients
}

func (client *Exakt1) ConnectionStatusCtx(
	ctx context.Context,
) (ConnectionStatus string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ConnectionStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_1, "ConnectionStatus", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ConnectionStatus, err = soap.UnmarshalString(response.ConnectionStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ConnectionStatus is the legacy version of ConnectionStatusCtx, but uses
// context.Background() as the context.
func (client *Exakt1) ConnectionStatus() (ConnectionStatus string, err error) {
	return client.ConnectionStatusCtx(context.Background())
}

func (client *Exakt1) DeviceListCtx(
	ctx context.Context,
) (List string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		List string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_1, "DeviceList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if List, err = soap.UnmarshalString(response.List); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// DeviceList is the legacy version of DeviceListCtx, but uses
// context.Background() as the context.
func (client *Exakt1) DeviceList() (List string, err error) {
	return client.DeviceListCtx(context.Background())
}

func (client *Exakt1) DeviceSettingsCtx(
	ctx context.Context,
	DeviceId string,
) (Settings string, err error) {
	// Request structure.
	request := &struct {
		DeviceId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Settings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_1, "DeviceSettings", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Settings, err = soap.UnmarshalString(response.Settings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// DeviceSettings is the legacy version of DeviceSettingsCtx, but uses
// context.Background() as the context.
func (client *Exakt1) DeviceSettings(DeviceId string) (Settings string, err error) {
	return client.DeviceSettingsCtx(context.Background(),
		DeviceId,
	)
}

func (client *Exakt1) ReprogramCtx(
	ctx context.Context,
	DeviceId string,
	FileUri *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		FileUri  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_1, "Reprogram", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Reprogram is the legacy version of ReprogramCtx, but uses
// context.Background() as the context.
func (client *Exakt1) Reprogram(DeviceId string, FileUri *url.URL) (err error) {
	return client.ReprogramCtx(context.Background(),
		DeviceId,
		FileUri,
	)
}

func (client *Exakt1) ReprogramFallbackCtx(
	ctx context.Context,
	DeviceId string,
	FileUri *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		FileUri  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_1, "ReprogramFallback", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ReprogramFallback is the legacy version of ReprogramFallbackCtx, but uses
// context.Background() as the context.
func (client *Exakt1) ReprogramFallback(DeviceId string, FileUri *url.URL) (err error) {
	return client.ReprogramFallbackCtx(context.Background(),
		DeviceId,
		FileUri,
	)
}

func (client *Exakt1) SetCtx(
	ctx context.Context,
	DeviceId string,
	BankId uint32,
	FileUri *url.URL,
	Mute bool,
	Persist bool,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		BankId   string
		FileUri  string
		Mute     string
		Persist  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.BankId, err = soap.MarshalUi4(BankId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	if request.Mute, err = soap.MarshalBoolean(Mute); err != nil {
		return
	}
	if request.Persist, err = soap.MarshalBoolean(Persist); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_1, "Set", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Set is the legacy version of SetCtx, but uses
// context.Background() as the context.
func (client *Exakt1) Set(DeviceId string, BankId uint32, FileUri *url.URL, Mute bool, Persist bool) (err error) {
	return client.SetCtx(context.Background(),
		DeviceId,
		BankId,
		FileUri,
		Mute,
		Persist,
	)
}

// Exakt2 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Exakt:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Exakt2 struct {
	goupnp.ServiceClient
}

// NewExakt2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewExakt2ClientsCtx(ctx context.Context) (clients []*Exakt2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Exakt_2); err != nil {
		return
	}
	clients = newExakt2ClientsFromGenericClients(genericClients)
	return
}

// NewExakt2Clients is the legacy version of NewExakt2ClientsCtx, but uses
// context.Background() as the context.
func NewExakt2Clients() (clients []*Exakt2, errors []error, err error) {
	return NewExakt2ClientsCtx(context.Background())
}

// NewExakt2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewExakt2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Exakt2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Exakt_2)
	if err != nil {
		return nil, err
	}
	return newExakt2ClientsFromGenericClients(genericClients), nil
}

// NewExakt2ClientsByURL is the legacy version of NewExakt2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewExakt2ClientsByURL(loc *url.URL) ([]*Exakt2, error) {
	return NewExakt2ClientsByURLCtx(context.Background(), loc)
}

// NewExakt2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewExakt2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Exakt2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Exakt_2)
	if err != nil {
		return nil, err
	}
	return newExakt2ClientsFromGenericClients(genericClients), nil
}

func newExakt2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Exakt2 {
	clients := make([]*Exakt2, len(genericClients))
	for i := range genericClients {
		clients[i] = &Exakt2{genericClients[i]}
	}
	return clients
}

func (client *Exakt2) ConnectionStatusCtx(
	ctx context.Context,
) (ConnectionStatus string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ConnectionStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "ConnectionStatus", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ConnectionStatus, err = soap.UnmarshalString(response.ConnectionStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ConnectionStatus is the legacy version of ConnectionStatusCtx, but uses
// context.Background() as the context.
func (client *Exakt2) ConnectionStatus() (ConnectionStatus string, err error) {
	return client.ConnectionStatusCtx(context.Background())
}

func (client *Exakt2) DeviceListCtx(
	ctx context.Context,
) (List string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		List string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "DeviceList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if List, err = soap.UnmarshalString(response.List); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// DeviceList is the legacy version of DeviceListCtx, but uses
// context.Background() as the context.
func (client *Exakt2) DeviceList() (List string, err error) {
	return client.DeviceListCtx(context.Background())
}

func (client *Exakt2) DeviceSettingsCtx(
	ctx context.Context,
	DeviceId string,
) (Settings string, err error) {
	// Request structure.
	request := &struct {
		DeviceId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Settings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "DeviceSettings", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Settings, err = soap.UnmarshalString(response.Settings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// DeviceSettings is the legacy version of DeviceSettingsCtx, but uses
// context.Background() as the context.
func (client *Exakt2) DeviceSettings(DeviceId string) (Settings string, err error) {
	return client.DeviceSettingsCtx(context.Background(),
		DeviceId,
	)
}

func (client *Exakt2) ReprogramCtx(
	ctx context.Context,
	DeviceId string,
	FileUri *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		FileUri  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "Reprogram", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Reprogram is the legacy version of ReprogramCtx, but uses
// context.Background() as the context.
func (client *Exakt2) Reprogram(DeviceId string, FileUri *url.URL) (err error) {
	return client.ReprogramCtx(context.Background(),
		DeviceId,
		FileUri,
	)
}

func (client *Exakt2) ReprogramFallbackCtx(
	ctx context.Context,
	DeviceId string,
	FileUri *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		FileUri  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "ReprogramFallback", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ReprogramFallback is the legacy version of ReprogramFallbackCtx, but uses
// context.Background() as the context.
func (client *Exakt2) ReprogramFallback(DeviceId string, FileUri *url.URL) (err error) {
	return client.ReprogramFallbackCtx(context.Background(),
		DeviceId,
		FileUri,
	)
}

func (client *Exakt2) SetCtx(
	ctx context.Context,
	DeviceId string,
	BankId uint32,
	FileUri *url.URL,
	Mute bool,
	Persist bool,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		BankId   string
		FileUri  string
		Mute     string
		Persist  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.BankId, err = soap.MarshalUi4(BankId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	if request.Mute, err = soap.MarshalBoolean(Mute); err != nil {
		return
	}
	if request.Persist, err = soap.MarshalBoolean(Persist); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "Set", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Set is the legacy version of SetCtx, but uses
// context.Background() as the context.
func (client *Exakt2) Set(DeviceId string, BankId uint32, FileUri *url.URL, Mute bool, Persist bool) (err error) {
	return client.SetCtx(context.Background(),
		DeviceId,
		BankId,
		FileUri,
		Mute,
		Persist,
	)
}

func (client *Exakt2) VersionCtx(
	ctx context.Context,
) (Version string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Version string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_2, "Version", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Version, err = soap.UnmarshalString(response.Version); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Version is the legacy version of VersionCtx, but uses
// context.Background() as the context.
func (client *Exakt2) Version() (Version string, err error) {
	return client.VersionCtx(context.Background())
}

// Exakt3 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Exakt:3". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Exakt3 struct {
	goupnp.ServiceClient
}

// NewExakt3ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewExakt3ClientsCtx(ctx context.Context) (clients []*Exakt3, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Exakt_3); err != nil {
		return
	}
	clients = newExakt3ClientsFromGenericClients(genericClients)
	return
}

// NewExakt3Clients is the legacy version of NewExakt3ClientsCtx, but uses
// context.Background() as the context.
func NewExakt3Clients() (clients []*Exakt3, errors []error, err error) {
	return NewExakt3ClientsCtx(context.Background())
}

// NewExakt3ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewExakt3ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Exakt3, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Exakt_3)
	if err != nil {
		return nil, err
	}
	return newExakt3ClientsFromGenericClients(genericClients), nil
}

// NewExakt3ClientsByURL is the legacy version of NewExakt3ClientsByURLCtx, but uses
// context.Background() as the context.
func NewExakt3ClientsByURL(loc *url.URL) ([]*Exakt3, error) {
	return NewExakt3ClientsByURLCtx(context.Background(), loc)
}

// NewExakt3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewExakt3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Exakt3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Exakt_3)
	if err != nil {
		return nil, err
	}
	return newExakt3ClientsFromGenericClients(genericClients), nil
}

func newExakt3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Exakt3 {
	clients := make([]*Exakt3, len(genericClients))
	for i := range genericClients {
		clients[i] = &Exakt3{genericClients[i]}
	}
	return clients
}

func (client *Exakt3) AudioChannelsCtx(
	ctx context.Context,
) (AudioChannels string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		AudioChannels string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "AudioChannels", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if AudioChannels, err = soap.UnmarshalString(response.AudioChannels); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// AudioChannels is the legacy version of AudioChannelsCtx, but uses
// context.Background() as the context.
func (client *Exakt3) AudioChannels() (AudioChannels string, err error) {
	return client.AudioChannelsCtx(context.Background())
}

func (client *Exakt3) ChannelMapCtx(
	ctx context.Context,
) (ChannelMap string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ChannelMap string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "ChannelMap", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ChannelMap, err = soap.UnmarshalString(response.ChannelMap); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ChannelMap is the legacy version of ChannelMapCtx, but uses
// context.Background() as the context.
func (client *Exakt3) ChannelMap() (ChannelMap string, err error) {
	return client.ChannelMapCtx(context.Background())
}

func (client *Exakt3) ConnectionStatusCtx(
	ctx context.Context,
) (ConnectionStatus string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ConnectionStatus string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "ConnectionStatus", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ConnectionStatus, err = soap.UnmarshalString(response.ConnectionStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ConnectionStatus is the legacy version of ConnectionStatusCtx, but uses
// context.Background() as the context.
func (client *Exakt3) ConnectionStatus() (ConnectionStatus string, err error) {
	return client.ConnectionStatusCtx(context.Background())
}

func (client *Exakt3) DeviceListCtx(
	ctx context.Context,
) (List string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		List string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "DeviceList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if List, err = soap.UnmarshalString(response.List); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// DeviceList is the legacy version of DeviceListCtx, but uses
// context.Background() as the context.
func (client *Exakt3) DeviceList() (List string, err error) {
	return client.DeviceListCtx(context.Background())
}

func (client *Exakt3) DeviceSettingsCtx(
	ctx context.Context,
	DeviceId string,
) (Settings string, err error) {
	// Request structure.
	request := &struct {
		DeviceId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Settings string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "DeviceSettings", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Settings, err = soap.UnmarshalString(response.Settings); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// DeviceSettings is the legacy version of DeviceSettingsCtx, but uses
// context.Background() as the context.
func (client *Exakt3) DeviceSettings(DeviceId string) (Settings string, err error) {
	return client.DeviceSettingsCtx(context.Background(),
		DeviceId,
	)
}

func (client *Exakt3) ReprogramCtx(
	ctx context.Context,
	DeviceId string,
	FileUri *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		FileUri  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "Reprogram", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Reprogram is the legacy version of ReprogramCtx, but uses
// context.Background() as the context.
func (client *Exakt3) Reprogram(DeviceId string, FileUri *url.URL) (err error) {
	return client.ReprogramCtx(context.Background(),
		DeviceId,
		FileUri,
	)
}

func (client *Exakt3) ReprogramFallbackCtx(
	ctx context.Context,
	DeviceId string,
	FileUri *url.URL,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		FileUri  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "ReprogramFallback", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// ReprogramFallback is the legacy version of ReprogramFallbackCtx, but uses
// context.Background() as the context.
func (client *Exakt3) ReprogramFallback(DeviceId string, FileUri *url.URL) (err error) {
	return client.ReprogramFallbackCtx(context.Background(),
		DeviceId,
		FileUri,
	)
}

func (client *Exakt3) SetCtx(
	ctx context.Context,
	DeviceId string,
	BankId uint32,
	FileUri *url.URL,
	Mute bool,
	Persist bool,
) (err error) {
	// Request structure.
	request := &struct {
		DeviceId string
		BankId   string
		FileUri  string
		Mute     string
		Persist  string
	}{}
	// BEGIN Marshal arguments into request.

	if request.DeviceId, err = soap.MarshalString(DeviceId); err != nil {
		return
	}
	if request.BankId, err = soap.MarshalUi4(BankId); err != nil {
		return
	}
	if request.FileUri, err = soap.MarshalURI(FileUri); err != nil {
		return
	}
	if request.Mute, err = soap.MarshalBoolean(Mute); err != nil {
		return
	}
	if request.Persist, err = soap.MarshalBoolean(Persist); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "Set", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Set is the legacy version of SetCtx, but uses
// context.Background() as the context.
func (client *Exakt3) Set(DeviceId string, BankId uint32, FileUri *url.URL, Mute bool, Persist bool) (err error) {
	return client.SetCtx(context.Background(),
		DeviceId,
		BankId,
		FileUri,
		Mute,
		Persist,
	)
}

func (client *Exakt3) SetAudioChannelsCtx(
	ctx context.Context,
	AudioChannels string,
) (err error) {
	// Request structure.
	request := &struct {
		AudioChannels string
	}{}
	// BEGIN Marshal arguments into request.

	if request.AudioChannels, err = soap.MarshalString(AudioChannels); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "SetAudioChannels", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAudioChannels is the legacy version of SetAudioChannelsCtx, but uses
// context.Background() as the context.
func (client *Exakt3) SetAudioChannels(AudioChannels string) (err error) {
	return client.SetAudioChannelsCtx(context.Background(),
		AudioChannels,
	)
}

func (client *Exakt3) SetChannelMapCtx(
	ctx context.Context,
	ChannelMap string,
) (err error) {
	// Request structure.
	request := &struct {
		ChannelMap string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ChannelMap, err = soap.MarshalString(ChannelMap); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "SetChannelMap", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetChannelMap is the legacy version of SetChannelMapCtx, but uses
// context.Background() as the context.
func (client *Exakt3) SetChannelMap(ChannelMap string) (err error) {
	return client.SetChannelMapCtx(context.Background(),
		ChannelMap,
	)
}

func (client *Exakt3) VersionCtx(
	ctx context.Context,
) (Version string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Version string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Exakt_3, "Version", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Version, err = soap.UnmarshalString(response.Version); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Version is the legacy version of VersionCtx, but uses
// context.Background() as the context.
func (client *Exakt3) Version() (Version string, err error) {
	return client.VersionCtx(context.Background())
}

// Info1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Info:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Info1 struct {
	goupnp.ServiceClient
}

// NewInfo1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewInfo1ClientsCtx(ctx context.Context) (clients []*Info1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Info_1); err != nil {
		return
	}
	clients = newInfo1ClientsFromGenericClients(genericClients)
	return
}

// NewInfo1Clients is the legacy version of NewInfo1ClientsCtx, but uses
// context.Background() as the context.
func NewInfo1Clients() (clients []*Info1, errors []error, err error) {
	return NewInfo1ClientsCtx(context.Background())
}

// NewInfo1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewInfo1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Info1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Info_1)
	if err != nil {
		return nil, err
	}
	return newInfo1ClientsFromGenericClients(genericClients), nil
}

// NewInfo1ClientsByURL is the legacy version of NewInfo1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewInfo1ClientsByURL(loc *url.URL) ([]*Info1, error) {
	return NewInfo1ClientsByURLCtx(context.Background(), loc)
}

// NewInfo1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewInfo1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Info1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Info_1)
	if err != nil {
		return nil, err
	}
	return newInfo1ClientsFromGenericClients(genericClients), nil
}

func newInfo1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Info1 {
	clients := make([]*Info1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Info1{genericClients[i]}
	}
	return clients
}

func (client *Info1) CountersCtx(
	ctx context.Context,
) (TrackCount uint32, DetailsCount uint32, MetatextCount uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TrackCount    string
		DetailsCount  string
		MetatextCount string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Info_1, "Counters", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TrackCount, err = soap.UnmarshalUi4(response.TrackCount); err != nil {
		return
	}
	if DetailsCount, err = soap.UnmarshalUi4(response.DetailsCount); err != nil {
		return
	}
	if MetatextCount, err = soap.UnmarshalUi4(response.MetatextCount); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Counters is the legacy version of CountersCtx, but uses
// context.Background() as the context.
func (client *Info1) Counters() (TrackCount uint32, DetailsCount uint32, MetatextCount uint32, err error) {
	return client.CountersCtx(context.Background())
}

func (client *Info1) DetailsCtx(
	ctx context.Context,
) (Duration uint32, BitRate uint32, BitDepth uint32, SampleRate uint32, Lossless bool, CodecName string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Duration   string
		BitRate    string
		BitDepth   string
		SampleRate string
		Lossless   string
		CodecName  string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Info_1, "Details", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Duration, err = soap.UnmarshalUi4(response.Duration); err != nil {
		return
	}
	if BitRate, err = soap.UnmarshalUi4(response.BitRate); err != nil {
		return
	}
	if BitDepth, err = soap.UnmarshalUi4(response.BitDepth); err != nil {
		return
	}
	if SampleRate, err = soap.UnmarshalUi4(response.SampleRate); err != nil {
		return
	}
	if Lossless, err = soap.UnmarshalBoolean(response.Lossless); err != nil {
		return
	}
	if CodecName, err = soap.UnmarshalString(response.CodecName); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Details is the legacy version of DetailsCtx, but uses
// context.Background() as the context.
func (client *Info1) Details() (Duration uint32, BitRate uint32, BitDepth uint32, SampleRate uint32, Lossless bool, CodecName string, err error) {
	return client.DetailsCtx(context.Background())
}

func (client *Info1) MetatextCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Info_1, "Metatext", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Metatext is the legacy version of MetatextCtx, but uses
// context.Background() as the context.
func (client *Info1) Metatext() (Value string, err error) {
	return client.MetatextCtx(context.Background())
}

func (client *Info1) TrackCtx(
	ctx context.Context,
) (Uri string, Metadata string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Uri      string
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Info_1, "Track", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Uri, err = soap.UnmarshalString(response.Uri); err != nil {
		return
	}
	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Track is the legacy version of TrackCtx, but uses
// context.Background() as the context.
func (client *Info1) Track() (Uri string, Metadata string, err error) {
	return client.TrackCtx(context.Background())
}

// MediaServer1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:MediaServer:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type MediaServer1 struct {
	goupnp.ServiceClient
}

// NewMediaServer1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewMediaServer1ClientsCtx(ctx context.Context) (clients []*MediaServer1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_MediaServer_1); err != nil {
		return
	}
	clients = newMediaServer1ClientsFromGenericClients(genericClients)
	return
}

// NewMediaServer1Clients is the legacy version of NewMediaServer1ClientsCtx, but uses
// context.Background() as the context.
func NewMediaServer1Clients() (clients []*MediaServer1, errors []error, err error) {
	return NewMediaServer1ClientsCtx(context.Background())
}

// NewMediaServer1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewMediaServer1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*MediaServer1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_MediaServer_1)
	if err != nil {
		return nil, err
	}
	return newMediaServer1ClientsFromGenericClients(genericClients), nil
}

// NewMediaServer1ClientsByURL is the legacy version of NewMediaServer1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewMediaServer1ClientsByURL(loc *url.URL) ([]*MediaServer1, error) {
	return NewMediaServer1ClientsByURLCtx(context.Background(), loc)
}

// NewMediaServer1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewMediaServer1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*MediaServer1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_MediaServer_1)
	if err != nil {
		return nil, err
	}
	return newMediaServer1ClientsFromGenericClients(genericClients), nil
}

func newMediaServer1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*MediaServer1 {
	clients := make([]*MediaServer1, len(genericClients))
	for i := range genericClients {
		clients[i] = &MediaServer1{genericClients[i]}
	}
	return clients
}

func (client *MediaServer1) AttributesCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "Attributes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Attributes is the legacy version of AttributesCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) Attributes() (Value string, err error) {
	return client.AttributesCtx(context.Background())
}

func (client *MediaServer1) BrowsePortCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "BrowsePort", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// BrowsePort is the legacy version of BrowsePortCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) BrowsePort() (Value uint32, err error) {
	return client.BrowsePortCtx(context.Background())
}

func (client *MediaServer1) ManufacturerCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "Manufacturer", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Manufacturer is the legacy version of ManufacturerCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) Manufacturer() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ManufacturerCtx(context.Background())
}

func (client *MediaServer1) ModelCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "Model", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Model is the legacy version of ModelCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) Model() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ModelCtx(context.Background())
}

func (client *MediaServer1) ProductCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "Product", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Product is the legacy version of ProductCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) Product() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ProductCtx(context.Background())
}

func (client *MediaServer1) QueryPortCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "QueryPort", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// QueryPort is the legacy version of QueryPortCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) QueryPort() (Value uint32, err error) {
	return client.QueryPortCtx(context.Background())
}

func (client *MediaServer1) UpdateCountCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_MediaServer_1, "UpdateCount", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// UpdateCount is the legacy version of UpdateCountCtx, but uses
// context.Background() as the context.
func (client *MediaServer1) UpdateCount() (Value uint32, err error) {
	return client.UpdateCountCtx(context.Background())
}

// NetworkMonitor1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:NetworkMonitor:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type NetworkMonitor1 struct {
	goupnp.ServiceClient
}

// NewNetworkMonitor1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewNetworkMonitor1ClientsCtx(ctx context.Context) (clients []*NetworkMonitor1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_NetworkMonitor_1); err != nil {
		return
	}
	clients = newNetworkMonitor1ClientsFromGenericClients(genericClients)
	return
}

// NewNetworkMonitor1Clients is the legacy version of NewNetworkMonitor1ClientsCtx, but uses
// context.Background() as the context.
func NewNetworkMonitor1Clients() (clients []*NetworkMonitor1, errors []error, err error) {
	return NewNetworkMonitor1ClientsCtx(context.Background())
}

// NewNetworkMonitor1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewNetworkMonitor1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*NetworkMonitor1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_NetworkMonitor_1)
	if err != nil {
		return nil, err
	}
	return newNetworkMonitor1ClientsFromGenericClients(genericClients), nil
}

// NewNetworkMonitor1ClientsByURL is the legacy version of NewNetworkMonitor1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewNetworkMonitor1ClientsByURL(loc *url.URL) ([]*NetworkMonitor1, error) {
	return NewNetworkMonitor1ClientsByURLCtx(context.Background(), loc)
}

// NewNetworkMonitor1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewNetworkMonitor1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*NetworkMonitor1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_NetworkMonitor_1)
	if err != nil {
		return nil, err
	}
	return newNetworkMonitor1ClientsFromGenericClients(genericClients), nil
}

func newNetworkMonitor1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*NetworkMonitor1 {
	clients := make([]*NetworkMonitor1, len(genericClients))
	for i := range genericClients {
		clients[i] = &NetworkMonitor1{genericClients[i]}
	}
	return clients
}

func (client *NetworkMonitor1) NameCtx(
	ctx context.Context,
) (Name string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_NetworkMonitor_1, "Name", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Name is the legacy version of NameCtx, but uses
// context.Background() as the context.
func (client *NetworkMonitor1) Name() (Name string, err error) {
	return client.NameCtx(context.Background())
}

func (client *NetworkMonitor1) PortsCtx(
	ctx context.Context,
) (Sender uint32, Receiver uint32, Results uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Sender   string
		Receiver string
		Results  string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_NetworkMonitor_1, "Ports", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Sender, err = soap.UnmarshalUi4(response.Sender); err != nil {
		return
	}
	if Receiver, err = soap.UnmarshalUi4(response.Receiver); err != nil {
		return
	}
	if Results, err = soap.UnmarshalUi4(response.Results); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Ports is the legacy version of PortsCtx, but uses
// context.Background() as the context.
func (client *NetworkMonitor1) Ports() (Sender uint32, Receiver uint32, Results uint32, err error) {
	return client.PortsCtx(context.Background())
}

// Pins1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Pins:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Pins1 struct {
	goupnp.ServiceClient
}

// NewPins1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewPins1ClientsCtx(ctx context.Context) (clients []*Pins1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Pins_1); err != nil {
		return
	}
	clients = newPins1ClientsFromGenericClients(genericClients)
	return
}

// NewPins1Clients is the legacy version of NewPins1ClientsCtx, but uses
// context.Background() as the context.
func NewPins1Clients() (clients []*Pins1, errors []error, err error) {
	return NewPins1ClientsCtx(context.Background())
}

// NewPins1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewPins1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Pins1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Pins_1)
	if err != nil {
		return nil, err
	}
	return newPins1ClientsFromGenericClients(genericClients), nil
}

// NewPins1ClientsByURL is the legacy version of NewPins1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewPins1ClientsByURL(loc *url.URL) ([]*Pins1, error) {
	return NewPins1ClientsByURLCtx(context.Background(), loc)
}

// NewPins1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewPins1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Pins1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Pins_1)
	if err != nil {
		return nil, err
	}
	return newPins1ClientsFromGenericClients(genericClients), nil
}

func newPins1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Pins1 {
	clients := make([]*Pins1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Pins1{genericClients[i]}
	}
	return clients
}

func (client *Pins1) ClearCtx(
	ctx context.Context,
	Id uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "Clear", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Clear is the legacy version of ClearCtx, but uses
// context.Background() as the context.
func (client *Pins1) Clear(Id uint32) (err error) {
	return client.ClearCtx(context.Background(),
		Id,
	)
}

func (client *Pins1) GetDeviceAccountMaxCtx(
	ctx context.Context,
) (DeviceMax uint32, AccountMax uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		DeviceMax  string
		AccountMax string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "GetDeviceAccountMax", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if DeviceMax, err = soap.UnmarshalUi4(response.DeviceMax); err != nil {
		return
	}
	if AccountMax, err = soap.UnmarshalUi4(response.AccountMax); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetDeviceAccountMax is the legacy version of GetDeviceAccountMaxCtx, but uses
// context.Background() as the context.
func (client *Pins1) GetDeviceAccountMax() (DeviceMax uint32, AccountMax uint32, err error) {
	return client.GetDeviceAccountMaxCtx(context.Background())
}

func (client *Pins1) GetIdArrayCtx(
	ctx context.Context,
) (IdArray string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		IdArray string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "GetIdArray", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if IdArray, err = soap.UnmarshalString(response.IdArray); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetIdArray is the legacy version of GetIdArrayCtx, but uses
// context.Background() as the context.
func (client *Pins1) GetIdArray() (IdArray string, err error) {
	return client.GetIdArrayCtx(context.Background())
}

func (client *Pins1) GetModesCtx(
	ctx context.Context,
) (Modes string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Modes string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "GetModes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Modes, err = soap.UnmarshalString(response.Modes); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetModes is the legacy version of GetModesCtx, but uses
// context.Background() as the context.
func (client *Pins1) GetModes() (Modes string, err error) {
	return client.GetModesCtx(context.Background())
}

func (client *Pins1) InvokeIdCtx(
	ctx context.Context,
	Id uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "InvokeId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// InvokeId is the legacy version of InvokeIdCtx, but uses
// context.Background() as the context.
func (client *Pins1) InvokeId(Id uint32) (err error) {
	return client.InvokeIdCtx(context.Background(),
		Id,
	)
}

func (client *Pins1) InvokeIndexCtx(
	ctx context.Context,
	Index uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Index string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Index, err = soap.MarshalUi4(Index); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "InvokeIndex", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// InvokeIndex is the legacy version of InvokeIndexCtx, but uses
// context.Background() as the context.
func (client *Pins1) InvokeIndex(Index uint32) (err error) {
	return client.InvokeIndexCtx(context.Background(),
		Index,
	)
}

func (client *Pins1) ReadListCtx(
	ctx context.Context,
	Ids string,
) (List string, err error) {
	// Request structure.
	request := &struct {
		Ids string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Ids, err = soap.MarshalString(Ids); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		List string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "ReadList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if List, err = soap.UnmarshalString(response.List); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ReadList is the legacy version of ReadListCtx, but uses
// context.Background() as the context.
func (client *Pins1) ReadList(Ids string) (List string, err error) {
	return client.ReadListCtx(context.Background(),
		Ids,
	)
}

func (client *Pins1) SetAccountCtx(
	ctx context.Context,
	Index uint32,
	Mode string,
	Type string,
	Uri string,
	Title string,
	Description string,
	ArtworkUri string,
	Shuffle bool,
) (err error) {
	// Request structure.
	request := &struct {
		Index       string
		Mode        string
		Type        string
		Uri         string
		Title       string
		Description string
		ArtworkUri  string
		Shuffle     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Index, err = soap.MarshalUi4(Index); err != nil {
		return
	}
	if request.Mode, err = soap.MarshalString(Mode); err != nil {
		return
	}
	if request.Type, err = soap.MarshalString(Type); err != nil {
		return
	}
	if request.Uri, err = soap.MarshalString(Uri); err != nil {
		return
	}
	if request.Title, err = soap.MarshalString(Title); err != nil {
		return
	}
	if request.Description, err = soap.MarshalString(Description); err != nil {
		return
	}
	if request.ArtworkUri, err = soap.MarshalString(ArtworkUri); err != nil {
		return
	}
	if request.Shuffle, err = soap.MarshalBoolean(Shuffle); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "SetAccount", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetAccount is the legacy version of SetAccountCtx, but uses
// context.Background() as the context.
func (client *Pins1) SetAccount(Index uint32, Mode string, Type string, Uri string, Title string, Description string, ArtworkUri string, Shuffle bool) (err error) {
	return client.SetAccountCtx(context.Background(),
		Index,
		Mode,
		Type,
		Uri,
		Title,
		Description,
		ArtworkUri,
		Shuffle,
	)
}

func (client *Pins1) SetDeviceCtx(
	ctx context.Context,
	Index uint32,
	Mode string,
	Type string,
	Uri string,
	Title string,
	Description string,
	ArtworkUri string,
	Shuffle bool,
) (err error) {
	// Request structure.
	request := &struct {
		Index       string
		Mode        string
		Type        string
		Uri         string
		Title       string
		Description string
		ArtworkUri  string
		Shuffle     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Index, err = soap.MarshalUi4(Index); err != nil {
		return
	}
	if request.Mode, err = soap.MarshalString(Mode); err != nil {
		return
	}
	if request.Type, err = soap.MarshalString(Type); err != nil {
		return
	}
	if request.Uri, err = soap.MarshalString(Uri); err != nil {
		return
	}
	if request.Title, err = soap.MarshalString(Title); err != nil {
		return
	}
	if request.Description, err = soap.MarshalString(Description); err != nil {
		return
	}
	if request.ArtworkUri, err = soap.MarshalString(ArtworkUri); err != nil {
		return
	}
	if request.Shuffle, err = soap.MarshalBoolean(Shuffle); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "SetDevice", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetDevice is the legacy version of SetDeviceCtx, but uses
// context.Background() as the context.
func (client *Pins1) SetDevice(Index uint32, Mode string, Type string, Uri string, Title string, Description string, ArtworkUri string, Shuffle bool) (err error) {
	return client.SetDeviceCtx(context.Background(),
		Index,
		Mode,
		Type,
		Uri,
		Title,
		Description,
		ArtworkUri,
		Shuffle,
	)
}

func (client *Pins1) SwapCtx(
	ctx context.Context,
	Index1 uint32,
	Index2 uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Index1 string
		Index2 string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Index1, err = soap.MarshalUi4(Index1); err != nil {
		return
	}
	if request.Index2, err = soap.MarshalUi4(Index2); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Pins_1, "Swap", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Swap is the legacy version of SwapCtx, but uses
// context.Background() as the context.
func (client *Pins1) Swap(Index1 uint32, Index2 uint32) (err error) {
	return client.SwapCtx(context.Background(),
		Index1,
		Index2,
	)
}

// Playlist1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Playlist:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Playlist1 struct {
	goupnp.ServiceClient
}

// NewPlaylist1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewPlaylist1ClientsCtx(ctx context.Context) (clients []*Playlist1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Playlist_1); err != nil {
		return
	}
	clients = newPlaylist1ClientsFromGenericClients(genericClients)
	return
}

// NewPlaylist1Clients is the legacy version of NewPlaylist1ClientsCtx, but uses
// context.Background() as the context.
func NewPlaylist1Clients() (clients []*Playlist1, errors []error, err error) {
	return NewPlaylist1ClientsCtx(context.Background())
}

// NewPlaylist1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewPlaylist1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Playlist1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Playlist_1)
	if err != nil {
		return nil, err
	}
	return newPlaylist1ClientsFromGenericClients(genericClients), nil
}

// NewPlaylist1ClientsByURL is the legacy version of NewPlaylist1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewPlaylist1ClientsByURL(loc *url.URL) ([]*Playlist1, error) {
	return NewPlaylist1ClientsByURLCtx(context.Background(), loc)
}

// NewPlaylist1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewPlaylist1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Playlist1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Playlist_1)
	if err != nil {
		return nil, err
	}
	return newPlaylist1ClientsFromGenericClients(genericClients), nil
}

func newPlaylist1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Playlist1 {
	clients := make([]*Playlist1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Playlist1{genericClients[i]}
	}
	return clients
}

func (client *Playlist1) DeleteAllCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "DeleteAll", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteAll is the legacy version of DeleteAllCtx, but uses
// context.Background() as the context.
func (client *Playlist1) DeleteAll() (err error) {
	return client.DeleteAllCtx(context.Background())
}

func (client *Playlist1) DeleteIdCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "DeleteId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteId is the legacy version of DeleteIdCtx, but uses
// context.Background() as the context.
func (client *Playlist1) DeleteId(Value uint32) (err error) {
	return client.DeleteIdCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) IdCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Id", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Id is the legacy version of IdCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Id() (Value uint32, err error) {
	return client.IdCtx(context.Background())
}

func (client *Playlist1) IdArrayCtx(
	ctx context.Context,
) (Token uint32, Array []byte, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Token string
		Array string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "IdArray", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Token, err = soap.UnmarshalUi4(response.Token); err != nil {
		return
	}
	if Array, err = soap.UnmarshalBinBase64(response.Array); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// IdArray is the legacy version of IdArrayCtx, but uses
// context.Background() as the context.
func (client *Playlist1) IdArray() (Token uint32, Array []byte, err error) {
	return client.IdArrayCtx(context.Background())
}

func (client *Playlist1) IdArrayChangedCtx(
	ctx context.Context,
	Token uint32,
) (Value bool, err error) {
	// Request structure.
	request := &struct {
		Token string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Token, err = soap.MarshalUi4(Token); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "IdArrayChanged", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// IdArrayChanged is the legacy version of IdArrayChangedCtx, but uses
// context.Background() as the context.
func (client *Playlist1) IdArrayChanged(Token uint32) (Value bool, err error) {
	return client.IdArrayChangedCtx(context.Background(),
		Token,
	)
}

func (client *Playlist1) InsertCtx(
	ctx context.Context,
	AfterId uint32,
	Uri string,
	Metadata string,
) (NewId uint32, err error) {
	// Request structure.
	request := &struct {
		AfterId  string
		Uri      string
		Metadata string
	}{}
	// BEGIN Marshal arguments into request.

	if request.AfterId, err = soap.MarshalUi4(AfterId); err != nil {
		return
	}
	if request.Uri, err = soap.MarshalString(Uri); err != nil {
		return
	}
	if request.Metadata, err = soap.MarshalString(Metadata); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NewId string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Insert", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewId, err = soap.UnmarshalUi4(response.NewId); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Insert is the legacy version of InsertCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Insert(AfterId uint32, Uri string, Metadata string) (NewId uint32, err error) {
	return client.InsertCtx(context.Background(),
		AfterId,
		Uri,
		Metadata,
	)
}

func (client *Playlist1) NextCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Next", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Next is the legacy version of NextCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Next() (err error) {
	return client.NextCtx(context.Background())
}

func (client *Playlist1) PauseCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Pause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Pause() (err error) {
	return client.PauseCtx(context.Background())
}

func (client *Playlist1) PlayCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Play() (err error) {
	return client.PlayCtx(context.Background())
}

func (client *Playlist1) PreviousCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Previous", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Previous is the legacy version of PreviousCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Previous() (err error) {
	return client.PreviousCtx(context.Background())
}

func (client *Playlist1) ProtocolInfoCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "ProtocolInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ProtocolInfo is the legacy version of ProtocolInfoCtx, but uses
// context.Background() as the context.
func (client *Playlist1) ProtocolInfo() (Value string, err error) {
	return client.ProtocolInfoCtx(context.Background())
}

func (client *Playlist1) ReadCtx(
	ctx context.Context,
	Id uint32,
) (Uri string, Metadata string, err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Uri      string
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Read", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Uri, err = soap.UnmarshalString(response.Uri); err != nil {
		return
	}
	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Read is the legacy version of ReadCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Read(Id uint32) (Uri string, Metadata string, err error) {
	return client.ReadCtx(context.Background(),
		Id,
	)
}

func (client *Playlist1) ReadListCtx(
	ctx context.Context,
	IdList string,
) (TrackList string, err error) {
	// Request structure.
	request := &struct {
		IdList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.IdList, err = soap.MarshalString(IdList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TrackList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "ReadList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TrackList, err = soap.UnmarshalString(response.TrackList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ReadList is the legacy version of ReadListCtx, but uses
// context.Background() as the context.
func (client *Playlist1) ReadList(IdList string) (TrackList string, err error) {
	return client.ReadListCtx(context.Background(),
		IdList,
	)
}

func (client *Playlist1) RepeatCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Repeat", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Repeat is the legacy version of RepeatCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Repeat() (Value bool, err error) {
	return client.RepeatCtx(context.Background())
}

func (client *Playlist1) SeekIdCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "SeekId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekId is the legacy version of SeekIdCtx, but uses
// context.Background() as the context.
func (client *Playlist1) SeekId(Value uint32) (err error) {
	return client.SeekIdCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) SeekIndexCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "SeekIndex", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekIndex is the legacy version of SeekIndexCtx, but uses
// context.Background() as the context.
func (client *Playlist1) SeekIndex(Value uint32) (err error) {
	return client.SeekIndexCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) SeekSecondAbsoluteCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "SeekSecondAbsolute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekSecondAbsolute is the legacy version of SeekSecondAbsoluteCtx, but uses
// context.Background() as the context.
func (client *Playlist1) SeekSecondAbsolute(Value uint32) (err error) {
	return client.SeekSecondAbsoluteCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) SeekSecondRelativeCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "SeekSecondRelative", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekSecondRelative is the legacy version of SeekSecondRelativeCtx, but uses
// context.Background() as the context.
func (client *Playlist1) SeekSecondRelative(Value int32) (err error) {
	return client.SeekSecondRelativeCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) SetRepeatCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "SetRepeat", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRepeat is the legacy version of SetRepeatCtx, but uses
// context.Background() as the context.
func (client *Playlist1) SetRepeat(Value bool) (err error) {
	return client.SetRepeatCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) SetShuffleCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "SetShuffle", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetShuffle is the legacy version of SetShuffleCtx, but uses
// context.Background() as the context.
func (client *Playlist1) SetShuffle(Value bool) (err error) {
	return client.SetShuffleCtx(context.Background(),
		Value,
	)
}

func (client *Playlist1) ShuffleCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Shuffle", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Shuffle is the legacy version of ShuffleCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Shuffle() (Value bool, err error) {
	return client.ShuffleCtx(context.Background())
}

func (client *Playlist1) StopCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *Playlist1) Stop() (err error) {
	return client.StopCtx(context.Background())
}

func (client *Playlist1) TracksMaxCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "TracksMax", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// TracksMax is the legacy version of TracksMaxCtx, but uses
// context.Background() as the context.
func (client *Playlist1) TracksMax() (Value uint32, err error) {
	return client.TracksMaxCtx(context.Background())
}

// Return values:
//
// * Value: allowed values: Playing, Paused, Stopped, Buffering
func (client *Playlist1) TransportStateCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Playlist_1, "TransportState", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// TransportState is the legacy version of TransportStateCtx, but uses
// context.Background() as the context.
func (client *Playlist1) TransportState() (Value string, err error) {
	return client.TransportStateCtx(context.Background())
}

// PlaylistManager1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:PlaylistManager:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type PlaylistManager1 struct {
	goupnp.ServiceClient
}

// NewPlaylistManager1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewPlaylistManager1ClientsCtx(ctx context.Context) (clients []*PlaylistManager1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_PlaylistManager_1); err != nil {
		return
	}
	clients = newPlaylistManager1ClientsFromGenericClients(genericClients)
	return
}

// NewPlaylistManager1Clients is the legacy version of NewPlaylistManager1ClientsCtx, but uses
// context.Background() as the context.
func NewPlaylistManager1Clients() (clients []*PlaylistManager1, errors []error, err error) {
	return NewPlaylistManager1ClientsCtx(context.Background())
}

// NewPlaylistManager1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewPlaylistManager1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*PlaylistManager1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_PlaylistManager_1)
	if err != nil {
		return nil, err
	}
	return newPlaylistManager1ClientsFromGenericClients(genericClients), nil
}

// NewPlaylistManager1ClientsByURL is the legacy version of NewPlaylistManager1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewPlaylistManager1ClientsByURL(loc *url.URL) ([]*PlaylistManager1, error) {
	return NewPlaylistManager1ClientsByURLCtx(context.Background(), loc)
}

// NewPlaylistManager1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewPlaylistManager1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*PlaylistManager1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_PlaylistManager_1)
	if err != nil {
		return nil, err
	}
	return newPlaylistManager1ClientsFromGenericClients(genericClients), nil
}

func newPlaylistManager1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*PlaylistManager1 {
	clients := make([]*PlaylistManager1, len(genericClients))
	for i := range genericClients {
		clients[i] = &PlaylistManager1{genericClients[i]}
	}
	return clients
}

func (client *PlaylistManager1) DeleteAllCtx(
	ctx context.Context,
	Id uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "DeleteAll", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteAll is the legacy version of DeleteAllCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) DeleteAll(Id uint32) (err error) {
	return client.DeleteAllCtx(context.Background(),
		Id,
	)
}

func (client *PlaylistManager1) DeleteIdCtx(
	ctx context.Context,
	Id uint32,
	TrackId uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Id      string
		TrackId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.TrackId, err = soap.MarshalUi4(TrackId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "DeleteId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// DeleteId is the legacy version of DeleteIdCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) DeleteId(Id uint32, TrackId uint32) (err error) {
	return client.DeleteIdCtx(context.Background(),
		Id,
		TrackId,
	)
}

func (client *PlaylistManager1) ImagesXmlCtx(
	ctx context.Context,
) (ImagesXml string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ImagesXml string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "ImagesXml", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ImagesXml, err = soap.UnmarshalString(response.ImagesXml); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ImagesXml is the legacy version of ImagesXmlCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) ImagesXml() (ImagesXml string, err error) {
	return client.ImagesXmlCtx(context.Background())
}

func (client *PlaylistManager1) InsertCtx(
	ctx context.Context,
	Id uint32,
	AfterTrackId uint32,
	Metadata string,
) (NewTrackId uint32, err error) {
	// Request structure.
	request := &struct {
		Id           string
		AfterTrackId string
		Metadata     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.AfterTrackId, err = soap.MarshalUi4(AfterTrackId); err != nil {
		return
	}
	if request.Metadata, err = soap.MarshalString(Metadata); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NewTrackId string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "Insert", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTrackId, err = soap.UnmarshalUi4(response.NewTrackId); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Insert is the legacy version of InsertCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) Insert(Id uint32, AfterTrackId uint32, Metadata string) (NewTrackId uint32, err error) {
	return client.InsertCtx(context.Background(),
		Id,
		AfterTrackId,
		Metadata,
	)
}

func (client *PlaylistManager1) MetadataCtx(
	ctx context.Context,
) (Metadata string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "Metadata", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Metadata is the legacy version of MetadataCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) Metadata() (Metadata string, err error) {
	return client.MetadataCtx(context.Background())
}

func (client *PlaylistManager1) PlaylistArraysCtx(
	ctx context.Context,
) (Token uint32, IdArray []byte, TokenArray []byte, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Token      string
		IdArray    string
		TokenArray string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistArrays", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Token, err = soap.UnmarshalUi4(response.Token); err != nil {
		return
	}
	if IdArray, err = soap.UnmarshalBinBase64(response.IdArray); err != nil {
		return
	}
	if TokenArray, err = soap.UnmarshalBinBase64(response.TokenArray); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistArrays is the legacy version of PlaylistArraysCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistArrays() (Token uint32, IdArray []byte, TokenArray []byte, err error) {
	return client.PlaylistArraysCtx(context.Background())
}

func (client *PlaylistManager1) PlaylistArraysChangedCtx(
	ctx context.Context,
	Token uint32,
) (Value bool, err error) {
	// Request structure.
	request := &struct {
		Token string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Token, err = soap.MarshalUi4(Token); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistArraysChanged", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistArraysChanged is the legacy version of PlaylistArraysChangedCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistArraysChanged(Token uint32) (Value bool, err error) {
	return client.PlaylistArraysChangedCtx(context.Background(),
		Token,
	)
}

func (client *PlaylistManager1) PlaylistDeleteIdCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistDeleteId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PlaylistDeleteId is the legacy version of PlaylistDeleteIdCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistDeleteId(Value uint32) (err error) {
	return client.PlaylistDeleteIdCtx(context.Background(),
		Value,
	)
}

func (client *PlaylistManager1) PlaylistInsertCtx(
	ctx context.Context,
	AfterId uint32,
	Name string,
	Description string,
	ImageId uint32,
) (NewId uint32, err error) {
	// Request structure.
	request := &struct {
		AfterId     string
		Name        string
		Description string
		ImageId     string
	}{}
	// BEGIN Marshal arguments into request.

	if request.AfterId, err = soap.MarshalUi4(AfterId); err != nil {
		return
	}
	if request.Name, err = soap.MarshalString(Name); err != nil {
		return
	}
	if request.Description, err = soap.MarshalString(Description); err != nil {
		return
	}
	if request.ImageId, err = soap.MarshalUi4(ImageId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		NewId string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistInsert", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewId, err = soap.UnmarshalUi4(response.NewId); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistInsert is the legacy version of PlaylistInsertCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistInsert(AfterId uint32, Name string, Description string, ImageId uint32) (NewId uint32, err error) {
	return client.PlaylistInsertCtx(context.Background(),
		AfterId,
		Name,
		Description,
		ImageId,
	)
}

func (client *PlaylistManager1) PlaylistMoveCtx(
	ctx context.Context,
	Id uint32,
	AfterId uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Id      string
		AfterId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.AfterId, err = soap.MarshalUi4(AfterId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistMove", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PlaylistMove is the legacy version of PlaylistMoveCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistMove(Id uint32, AfterId uint32) (err error) {
	return client.PlaylistMoveCtx(context.Background(),
		Id,
		AfterId,
	)
}

func (client *PlaylistManager1) PlaylistReadCtx(
	ctx context.Context,
	Id uint32,
) (Name string, Description string, ImageId uint32, err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name        string
		Description string
		ImageId     string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistRead", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Description, err = soap.UnmarshalString(response.Description); err != nil {
		return
	}
	if ImageId, err = soap.UnmarshalUi4(response.ImageId); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistRead is the legacy version of PlaylistReadCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistRead(Id uint32) (Name string, Description string, ImageId uint32, err error) {
	return client.PlaylistReadCtx(context.Background(),
		Id,
	)
}

func (client *PlaylistManager1) PlaylistReadArrayCtx(
	ctx context.Context,
	Id uint32,
) (Array []byte, err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Array string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistReadArray", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Array, err = soap.UnmarshalBinBase64(response.Array); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistReadArray is the legacy version of PlaylistReadArrayCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistReadArray(Id uint32) (Array []byte, err error) {
	return client.PlaylistReadArrayCtx(context.Background(),
		Id,
	)
}

func (client *PlaylistManager1) PlaylistReadListCtx(
	ctx context.Context,
	IdList string,
) (PlaylistList string, err error) {
	// Request structure.
	request := &struct {
		IdList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.IdList, err = soap.MarshalString(IdList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		PlaylistList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistReadList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if PlaylistList, err = soap.UnmarshalString(response.PlaylistList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistReadList is the legacy version of PlaylistReadListCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistReadList(IdList string) (PlaylistList string, err error) {
	return client.PlaylistReadListCtx(context.Background(),
		IdList,
	)
}

func (client *PlaylistManager1) PlaylistSetDescriptionCtx(
	ctx context.Context,
	Id uint32,
	Description string,
) (err error) {
	// Request structure.
	request := &struct {
		Id          string
		Description string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.Description, err = soap.MarshalString(Description); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistSetDescription", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PlaylistSetDescription is the legacy version of PlaylistSetDescriptionCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistSetDescription(Id uint32, Description string) (err error) {
	return client.PlaylistSetDescriptionCtx(context.Background(),
		Id,
		Description,
	)
}

func (client *PlaylistManager1) PlaylistSetImageIdCtx(
	ctx context.Context,
	Id uint32,
	ImageId uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Id      string
		ImageId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.ImageId, err = soap.MarshalUi4(ImageId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistSetImageId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PlaylistSetImageId is the legacy version of PlaylistSetImageIdCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistSetImageId(Id uint32, ImageId uint32) (err error) {
	return client.PlaylistSetImageIdCtx(context.Background(),
		Id,
		ImageId,
	)
}

func (client *PlaylistManager1) PlaylistSetNameCtx(
	ctx context.Context,
	Id uint32,
	Name string,
) (err error) {
	// Request structure.
	request := &struct {
		Id   string
		Name string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.Name, err = soap.MarshalString(Name); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistSetName", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PlaylistSetName is the legacy version of PlaylistSetNameCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistSetName(Id uint32, Name string) (err error) {
	return client.PlaylistSetNameCtx(context.Background(),
		Id,
		Name,
	)
}

func (client *PlaylistManager1) PlaylistsMaxCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "PlaylistsMax", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PlaylistsMax is the legacy version of PlaylistsMaxCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) PlaylistsMax() (Value uint32, err error) {
	return client.PlaylistsMaxCtx(context.Background())
}

func (client *PlaylistManager1) ReadCtx(
	ctx context.Context,
	Id uint32,
	TrackId uint32,
) (Metadata string, err error) {
	// Request structure.
	request := &struct {
		Id      string
		TrackId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.TrackId, err = soap.MarshalUi4(TrackId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "Read", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Read is the legacy version of ReadCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) Read(Id uint32, TrackId uint32) (Metadata string, err error) {
	return client.ReadCtx(context.Background(),
		Id,
		TrackId,
	)
}

func (client *PlaylistManager1) ReadListCtx(
	ctx context.Context,
	Id uint32,
	TrackIdList string,
) (TrackList string, err error) {
	// Request structure.
	request := &struct {
		Id          string
		TrackIdList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	if request.TrackIdList, err = soap.MarshalString(TrackIdList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TrackList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "ReadList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TrackList, err = soap.UnmarshalString(response.TrackList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ReadList is the legacy version of ReadListCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) ReadList(Id uint32, TrackIdList string) (TrackList string, err error) {
	return client.ReadListCtx(context.Background(),
		Id,
		TrackIdList,
	)
}

func (client *PlaylistManager1) TracksMaxCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlaylistManager_1, "TracksMax", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// TracksMax is the legacy version of TracksMaxCtx, but uses
// context.Background() as the context.
func (client *PlaylistManager1) TracksMax() (Value uint32, err error) {
	return client.TracksMaxCtx(context.Background())
}

// Product1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Product:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Product1 struct {
	goupnp.ServiceClient
}

// NewProduct1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewProduct1ClientsCtx(ctx context.Context) (clients []*Product1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Product_1); err != nil {
		return
	}
	clients = newProduct1ClientsFromGenericClients(genericClients)
	return
}

// NewProduct1Clients is the legacy version of NewProduct1ClientsCtx, but uses
// context.Background() as the context.
func NewProduct1Clients() (clients []*Product1, errors []error, err error) {
	return NewProduct1ClientsCtx(context.Background())
}

// NewProduct1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewProduct1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Product1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Product_1)
	if err != nil {
		return nil, err
	}
	return newProduct1ClientsFromGenericClients(genericClients), nil
}

// NewProduct1ClientsByURL is the legacy version of NewProduct1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewProduct1ClientsByURL(loc *url.URL) ([]*Product1, error) {
	return NewProduct1ClientsByURLCtx(context.Background(), loc)
}

// NewProduct1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewProduct1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Product1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Product_1)
	if err != nil {
		return nil, err
	}
	return newProduct1ClientsFromGenericClients(genericClients), nil
}

func newProduct1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Product1 {
	clients := make([]*Product1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Product1{genericClients[i]}
	}
	return clients
}

func (client *Product1) AttributesCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "Attributes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Attributes is the legacy version of AttributesCtx, but uses
// context.Background() as the context.
func (client *Product1) Attributes() (Value string, err error) {
	return client.AttributesCtx(context.Background())
}

func (client *Product1) ManufacturerCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "Manufacturer", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Manufacturer is the legacy version of ManufacturerCtx, but uses
// context.Background() as the context.
func (client *Product1) Manufacturer() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ManufacturerCtx(context.Background())
}

func (client *Product1) ModelCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "Model", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Model is the legacy version of ModelCtx, but uses
// context.Background() as the context.
func (client *Product1) Model() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ModelCtx(context.Background())
}

func (client *Product1) ProductCtx(
	ctx context.Context,
) (Room string, Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Room     string
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "Product", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Room, err = soap.UnmarshalString(response.Room); err != nil {
		return
	}
	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Product is the legacy version of ProductCtx, but uses
// context.Background() as the context.
func (client *Product1) Product() (Room string, Name string, Info string, Url string, ImageUri string, err error) {
	return client.ProductCtx(context.Background())
}

func (client *Product1) SetSourceIndexCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SetSourceIndex", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSourceIndex is the legacy version of SetSourceIndexCtx, but uses
// context.Background() as the context.
func (client *Product1) SetSourceIndex(Value uint32) (err error) {
	return client.SetSourceIndexCtx(context.Background(),
		Value,
	)
}

func (client *Product1) SetSourceIndexByNameCtx(
	ctx context.Context,
	Value string,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalString(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SetSourceIndexByName", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSourceIndexByName is the legacy version of SetSourceIndexByNameCtx, but uses
// context.Background() as the context.
func (client *Product1) SetSourceIndexByName(Value string) (err error) {
	return client.SetSourceIndexByNameCtx(context.Background(),
		Value,
	)
}

func (client *Product1) SetStandbyCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SetStandby", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetStandby is the legacy version of SetStandbyCtx, but uses
// context.Background() as the context.
func (client *Product1) SetStandby(Value bool) (err error) {
	return client.SetStandbyCtx(context.Background(),
		Value,
	)
}

func (client *Product1) SourceCtx(
	ctx context.Context,
	Index uint32,
) (SystemName string, Type string, Name string, Visible bool, err error) {
	// Request structure.
	request := &struct {
		Index string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Index, err = soap.MarshalUi4(Index); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SystemName string
		Type       string
		Name       string
		Visible    string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "Source", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SystemName, err = soap.UnmarshalString(response.SystemName); err != nil {
		return
	}
	if Type, err = soap.UnmarshalString(response.Type); err != nil {
		return
	}
	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Visible, err = soap.UnmarshalBoolean(response.Visible); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Source is the legacy version of SourceCtx, but uses
// context.Background() as the context.
func (client *Product1) Source(Index uint32) (SystemName string, Type string, Name string, Visible bool, err error) {
	return client.SourceCtx(context.Background(),
		Index,
	)
}

func (client *Product1) SourceCountCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SourceCount", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceCount is the legacy version of SourceCountCtx, but uses
// context.Background() as the context.
func (client *Product1) SourceCount() (Value uint32, err error) {
	return client.SourceCountCtx(context.Background())
}

func (client *Product1) SourceIndexCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SourceIndex", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceIndex is the legacy version of SourceIndexCtx, but uses
// context.Background() as the context.
func (client *Product1) SourceIndex() (Value uint32, err error) {
	return client.SourceIndexCtx(context.Background())
}

func (client *Product1) SourceXmlCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SourceXml", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceXml is the legacy version of SourceXmlCtx, but uses
// context.Background() as the context.
func (client *Product1) SourceXml() (Value string, err error) {
	return client.SourceXmlCtx(context.Background())
}

func (client *Product1) SourceXmlChangeCountCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "SourceXmlChangeCount", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceXmlChangeCount is the legacy version of SourceXmlChangeCountCtx, but uses
// context.Background() as the context.
func (client *Product1) SourceXmlChangeCount() (Value uint32, err error) {
	return client.SourceXmlChangeCountCtx(context.Background())
}

func (client *Product1) StandbyCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_1, "Standby", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Standby is the legacy version of StandbyCtx, but uses
// context.Background() as the context.
func (client *Product1) Standby() (Value bool, err error) {
	return client.StandbyCtx(context.Background())
}

// Product2 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Product:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Product2 struct {
	goupnp.ServiceClient
}

// NewProduct2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewProduct2ClientsCtx(ctx context.Context) (clients []*Product2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Product_2); err != nil {
		return
	}
	clients = newProduct2ClientsFromGenericClients(genericClients)
	return
}

// NewProduct2Clients is the legacy version of NewProduct2ClientsCtx, but uses
// context.Background() as the context.
func NewProduct2Clients() (clients []*Product2, errors []error, err error) {
	return NewProduct2ClientsCtx(context.Background())
}

// NewProduct2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewProduct2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Product2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Product_2)
	if err != nil {
		return nil, err
	}
	return newProduct2ClientsFromGenericClients(genericClients), nil
}

// NewProduct2ClientsByURL is the legacy version of NewProduct2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewProduct2ClientsByURL(loc *url.URL) ([]*Product2, error) {
	return NewProduct2ClientsByURLCtx(context.Background(), loc)
}

// NewProduct2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewProduct2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Product2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Product_2)
	if err != nil {
		return nil, err
	}
	return newProduct2ClientsFromGenericClients(genericClients), nil
}

func newProduct2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Product2 {
	clients := make([]*Product2, len(genericClients))
	for i := range genericClients {
		clients[i] = &Product2{genericClients[i]}
	}
	return clients
}

func (client *Product2) AttributesCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "Attributes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Attributes is the legacy version of AttributesCtx, but uses
// context.Background() as the context.
func (client *Product2) Attributes() (Value string, err error) {
	return client.AttributesCtx(context.Background())
}

func (client *Product2) ManufacturerCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "Manufacturer", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Manufacturer is the legacy version of ManufacturerCtx, but uses
// context.Background() as the context.
func (client *Product2) Manufacturer() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ManufacturerCtx(context.Background())
}

func (client *Product2) ModelCtx(
	ctx context.Context,
) (Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "Model", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Model is the legacy version of ModelCtx, but uses
// context.Background() as the context.
func (client *Product2) Model() (Name string, Info string, Url string, ImageUri string, err error) {
	return client.ModelCtx(context.Background())
}

func (client *Product2) ProductCtx(
	ctx context.Context,
) (Room string, Name string, Info string, Url string, ImageUri string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Room     string
		Name     string
		Info     string
		Url      string
		ImageUri string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "Product", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Room, err = soap.UnmarshalString(response.Room); err != nil {
		return
	}
	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Info, err = soap.UnmarshalString(response.Info); err != nil {
		return
	}
	if Url, err = soap.UnmarshalString(response.Url); err != nil {
		return
	}
	if ImageUri, err = soap.UnmarshalString(response.ImageUri); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Product is the legacy version of ProductCtx, but uses
// context.Background() as the context.
func (client *Product2) Product() (Room string, Name string, Info string, Url string, ImageUri string, err error) {
	return client.ProductCtx(context.Background())
}

func (client *Product2) SetSourceBySystemNameCtx(
	ctx context.Context,
	Value string,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalString(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SetSourceBySystemName", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSourceBySystemName is the legacy version of SetSourceBySystemNameCtx, but uses
// context.Background() as the context.
func (client *Product2) SetSourceBySystemName(Value string) (err error) {
	return client.SetSourceBySystemNameCtx(context.Background(),
		Value,
	)
}

func (client *Product2) SetSourceIndexCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SetSourceIndex", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSourceIndex is the legacy version of SetSourceIndexCtx, but uses
// context.Background() as the context.
func (client *Product2) SetSourceIndex(Value uint32) (err error) {
	return client.SetSourceIndexCtx(context.Background(),
		Value,
	)
}

func (client *Product2) SetSourceIndexByNameCtx(
	ctx context.Context,
	Value string,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalString(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SetSourceIndexByName", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSourceIndexByName is the legacy version of SetSourceIndexByNameCtx, but uses
// context.Background() as the context.
func (client *Product2) SetSourceIndexByName(Value string) (err error) {
	return client.SetSourceIndexByNameCtx(context.Background(),
		Value,
	)
}

func (client *Product2) SetStandbyCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SetStandby", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetStandby is the legacy version of SetStandbyCtx, but uses
// context.Background() as the context.
func (client *Product2) SetStandby(Value bool) (err error) {
	return client.SetStandbyCtx(context.Background(),
		Value,
	)
}

func (client *Product2) SourceCtx(
	ctx context.Context,
	Index uint32,
) (SystemName string, Type string, Name string, Visible bool, err error) {
	// Request structure.
	request := &struct {
		Index string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Index, err = soap.MarshalUi4(Index); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		SystemName string
		Type       string
		Name       string
		Visible    string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "Source", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if SystemName, err = soap.UnmarshalString(response.SystemName); err != nil {
		return
	}
	if Type, err = soap.UnmarshalString(response.Type); err != nil {
		return
	}
	if Name, err = soap.UnmarshalString(response.Name); err != nil {
		return
	}
	if Visible, err = soap.UnmarshalBoolean(response.Visible); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Source is the legacy version of SourceCtx, but uses
// context.Background() as the context.
func (client *Product2) Source(Index uint32) (SystemName string, Type string, Name string, Visible bool, err error) {
	return client.SourceCtx(context.Background(),
		Index,
	)
}

func (client *Product2) SourceCountCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SourceCount", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceCount is the legacy version of SourceCountCtx, but uses
// context.Background() as the context.
func (client *Product2) SourceCount() (Value uint32, err error) {
	return client.SourceCountCtx(context.Background())
}

func (client *Product2) SourceIndexCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SourceIndex", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceIndex is the legacy version of SourceIndexCtx, but uses
// context.Background() as the context.
func (client *Product2) SourceIndex() (Value uint32, err error) {
	return client.SourceIndexCtx(context.Background())
}

func (client *Product2) SourceXmlCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SourceXml", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceXml is the legacy version of SourceXmlCtx, but uses
// context.Background() as the context.
func (client *Product2) SourceXml() (Value string, err error) {
	return client.SourceXmlCtx(context.Background())
}

func (client *Product2) SourceXmlChangeCountCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "SourceXmlChangeCount", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// SourceXmlChangeCount is the legacy version of SourceXmlChangeCountCtx, but uses
// context.Background() as the context.
func (client *Product2) SourceXmlChangeCount() (Value uint32, err error) {
	return client.SourceXmlChangeCountCtx(context.Background())
}

func (client *Product2) StandbyCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Product_2, "Standby", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Standby is the legacy version of StandbyCtx, but uses
// context.Background() as the context.
func (client *Product2) Standby() (Value bool, err error) {
	return client.StandbyCtx(context.Background())
}

// Radio1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Radio:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Radio1 struct {
	goupnp.ServiceClient
}

// NewRadio1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRadio1ClientsCtx(ctx context.Context) (clients []*Radio1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Radio_1); err != nil {
		return
	}
	clients = newRadio1ClientsFromGenericClients(genericClients)
	return
}

// NewRadio1Clients is the legacy version of NewRadio1ClientsCtx, but uses
// context.Background() as the context.
func NewRadio1Clients() (clients []*Radio1, errors []error, err error) {
	return NewRadio1ClientsCtx(context.Background())
}

// NewRadio1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRadio1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Radio1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Radio_1)
	if err != nil {
		return nil, err
	}
	return newRadio1ClientsFromGenericClients(genericClients), nil
}

// NewRadio1ClientsByURL is the legacy version of NewRadio1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewRadio1ClientsByURL(loc *url.URL) ([]*Radio1, error) {
	return NewRadio1ClientsByURLCtx(context.Background(), loc)
}

// NewRadio1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRadio1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Radio1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Radio_1)
	if err != nil {
		return nil, err
	}
	return newRadio1ClientsFromGenericClients(genericClients), nil
}

func newRadio1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Radio1 {
	clients := make([]*Radio1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Radio1{genericClients[i]}
	}
	return clients
}

func (client *Radio1) ChannelCtx(
	ctx context.Context,
) (Uri string, Metadata string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Uri      string
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "Channel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Uri, err = soap.UnmarshalString(response.Uri); err != nil {
		return
	}
	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Channel is the legacy version of ChannelCtx, but uses
// context.Background() as the context.
func (client *Radio1) Channel() (Uri string, Metadata string, err error) {
	return client.ChannelCtx(context.Background())
}

func (client *Radio1) ChannelsMaxCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "ChannelsMax", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ChannelsMax is the legacy version of ChannelsMaxCtx, but uses
// context.Background() as the context.
func (client *Radio1) ChannelsMax() (Value uint32, err error) {
	return client.ChannelsMaxCtx(context.Background())
}

func (client *Radio1) IdCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "Id", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Id is the legacy version of IdCtx, but uses
// context.Background() as the context.
func (client *Radio1) Id() (Value uint32, err error) {
	return client.IdCtx(context.Background())
}

func (client *Radio1) IdArrayCtx(
	ctx context.Context,
) (Token uint32, Array []byte, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Token string
		Array string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "IdArray", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Token, err = soap.UnmarshalUi4(response.Token); err != nil {
		return
	}
	if Array, err = soap.UnmarshalBinBase64(response.Array); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// IdArray is the legacy version of IdArrayCtx, but uses
// context.Background() as the context.
func (client *Radio1) IdArray() (Token uint32, Array []byte, err error) {
	return client.IdArrayCtx(context.Background())
}

func (client *Radio1) IdArrayChangedCtx(
	ctx context.Context,
	Token uint32,
) (Value bool, err error) {
	// Request structure.
	request := &struct {
		Token string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Token, err = soap.MarshalUi4(Token); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "IdArrayChanged", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// IdArrayChanged is the legacy version of IdArrayChangedCtx, but uses
// context.Background() as the context.
func (client *Radio1) IdArrayChanged(Token uint32) (Value bool, err error) {
	return client.IdArrayChangedCtx(context.Background(),
		Token,
	)
}

func (client *Radio1) PauseCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "Pause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *Radio1) Pause() (err error) {
	return client.PauseCtx(context.Background())
}

func (client *Radio1) PlayCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *Radio1) Play() (err error) {
	return client.PlayCtx(context.Background())
}

func (client *Radio1) ProtocolInfoCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "ProtocolInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ProtocolInfo is the legacy version of ProtocolInfoCtx, but uses
// context.Background() as the context.
func (client *Radio1) ProtocolInfo() (Value string, err error) {
	return client.ProtocolInfoCtx(context.Background())
}

func (client *Radio1) ReadCtx(
	ctx context.Context,
	Id uint32,
) (Metadata string, err error) {
	// Request structure.
	request := &struct {
		Id string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Id, err = soap.MarshalUi4(Id); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "Read", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Read is the legacy version of ReadCtx, but uses
// context.Background() as the context.
func (client *Radio1) Read(Id uint32) (Metadata string, err error) {
	return client.ReadCtx(context.Background(),
		Id,
	)
}

func (client *Radio1) ReadListCtx(
	ctx context.Context,
	IdList string,
) (ChannelList string, err error) {
	// Request structure.
	request := &struct {
		IdList string
	}{}
	// BEGIN Marshal arguments into request.

	if request.IdList, err = soap.MarshalString(IdList); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		ChannelList string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "ReadList", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if ChannelList, err = soap.UnmarshalString(response.ChannelList); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ReadList is the legacy version of ReadListCtx, but uses
// context.Background() as the context.
func (client *Radio1) ReadList(IdList string) (ChannelList string, err error) {
	return client.ReadListCtx(context.Background(),
		IdList,
	)
}

func (client *Radio1) SeekSecondAbsoluteCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "SeekSecondAbsolute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekSecondAbsolute is the legacy version of SeekSecondAbsoluteCtx, but uses
// context.Background() as the context.
func (client *Radio1) SeekSecondAbsolute(Value uint32) (err error) {
	return client.SeekSecondAbsoluteCtx(context.Background(),
		Value,
	)
}

func (client *Radio1) SeekSecondRelativeCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "SeekSecondRelative", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekSecondRelative is the legacy version of SeekSecondRelativeCtx, but uses
// context.Background() as the context.
func (client *Radio1) SeekSecondRelative(Value int32) (err error) {
	return client.SeekSecondRelativeCtx(context.Background(),
		Value,
	)
}

func (client *Radio1) SetChannelCtx(
	ctx context.Context,
	Uri string,
	Metadata string,
) (err error) {
	// Request structure.
	request := &struct {
		Uri      string
		Metadata string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Uri, err = soap.MarshalString(Uri); err != nil {
		return
	}
	if request.Metadata, err = soap.MarshalString(Metadata); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "SetChannel", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetChannel is the legacy version of SetChannelCtx, but uses
// context.Background() as the context.
func (client *Radio1) SetChannel(Uri string, Metadata string) (err error) {
	return client.SetChannelCtx(context.Background(),
		Uri,
		Metadata,
	)
}

func (client *Radio1) SetIdCtx(
	ctx context.Context,
	Value uint32,
	Uri string,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
		Uri   string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	if request.Uri, err = soap.MarshalString(Uri); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "SetId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetId is the legacy version of SetIdCtx, but uses
// context.Background() as the context.
func (client *Radio1) SetId(Value uint32, Uri string) (err error) {
	return client.SetIdCtx(context.Background(),
		Value,
		Uri,
	)
}

func (client *Radio1) StopCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *Radio1) Stop() (err error) {
	return client.StopCtx(context.Background())
}

// Return values:
//
// * Value: allowed values: Stopped, Playing, Paused, Buffering
func (client *Radio1) TransportStateCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Radio_1, "TransportState", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// TransportState is the legacy version of TransportStateCtx, but uses
// context.Background() as the context.
func (client *Radio1) TransportState() (Value string, err error) {
	return client.TransportStateCtx(context.Background())
}

// Receiver1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Receiver:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Receiver1 struct {
	goupnp.ServiceClient
}

// NewReceiver1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewReceiver1ClientsCtx(ctx context.Context) (clients []*Receiver1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Receiver_1); err != nil {
		return
	}
	clients = newReceiver1ClientsFromGenericClients(genericClients)
	return
}

// NewReceiver1Clients is the legacy version of NewReceiver1ClientsCtx, but uses
// context.Background() as the context.
func NewReceiver1Clients() (clients []*Receiver1, errors []error, err error) {
	return NewReceiver1ClientsCtx(context.Background())
}

// NewReceiver1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewReceiver1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Receiver1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Receiver_1)
	if err != nil {
		return nil, err
	}
	return newReceiver1ClientsFromGenericClients(genericClients), nil
}

// NewReceiver1ClientsByURL is the legacy version of NewReceiver1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewReceiver1ClientsByURL(loc *url.URL) ([]*Receiver1, error) {
	return NewReceiver1ClientsByURLCtx(context.Background(), loc)
}

// NewReceiver1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewReceiver1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Receiver1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Receiver_1)
	if err != nil {
		return nil, err
	}
	return newReceiver1ClientsFromGenericClients(genericClients), nil
}

func newReceiver1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Receiver1 {
	clients := make([]*Receiver1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Receiver1{genericClients[i]}
	}
	return clients
}

func (client *Receiver1) PlayCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Receiver_1, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *Receiver1) Play() (err error) {
	return client.PlayCtx(context.Background())
}

func (client *Receiver1) ProtocolInfoCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Receiver_1, "ProtocolInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ProtocolInfo is the legacy version of ProtocolInfoCtx, but uses
// context.Background() as the context.
func (client *Receiver1) ProtocolInfo() (Value string, err error) {
	return client.ProtocolInfoCtx(context.Background())
}

func (client *Receiver1) SenderCtx(
	ctx context.Context,
) (Uri string, Metadata string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Uri      string
		Metadata string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Receiver_1, "Sender", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Uri, err = soap.UnmarshalString(response.Uri); err != nil {
		return
	}
	if Metadata, err = soap.UnmarshalString(response.Metadata); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Sender is the legacy version of SenderCtx, but uses
// context.Background() as the context.
func (client *Receiver1) Sender() (Uri string, Metadata string, err error) {
	return client.SenderCtx(context.Background())
}

func (client *Receiver1) SetSenderCtx(
	ctx context.Context,
	Uri string,
	Metadata string,
) (err error) {
	// Request structure.
	request := &struct {
		Uri      string
		Metadata string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Uri, err = soap.MarshalString(Uri); err != nil {
		return
	}
	if request.Metadata, err = soap.MarshalString(Metadata); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Receiver_1, "SetSender", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetSender is the legacy version of SetSenderCtx, but uses
// context.Background() as the context.
func (client *Receiver1) SetSender(Uri string, Metadata string) (err error) {
	return client.SetSenderCtx(context.Background(),
		Uri,
		Metadata,
	)
}

func (client *Receiver1) StopCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Receiver_1, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *Receiver1) Stop() (err error) {
	return client.StopCtx(context.Background())
}

// Return values:
//
// * Value: allowed values: Stopped, Playing, Waiting, Buffering
func (client *Receiver1) TransportStateCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Receiver_1, "TransportState", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// TransportState is the legacy version of TransportStateCtx, but uses
// context.Background() as the context.
func (client *Receiver1) TransportState() (Value string, err error) {
	return client.TransportStateCtx(context.Background())
}

// Sender1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Sender:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Sender1 struct {
	goupnp.ServiceClient
}

// NewSender1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewSender1ClientsCtx(ctx context.Context) (clients []*Sender1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Sender_1); err != nil {
		return
	}
	clients = newSender1ClientsFromGenericClients(genericClients)
	return
}

// NewSender1Clients is the legacy version of NewSender1ClientsCtx, but uses
// context.Background() as the context.
func NewSender1Clients() (clients []*Sender1, errors []error, err error) {
	return NewSender1ClientsCtx(context.Background())
}

// NewSender1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewSender1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Sender1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Sender_1)
	if err != nil {
		return nil, err
	}
	return newSender1ClientsFromGenericClients(genericClients), nil
}

// NewSender1ClientsByURL is the legacy version of NewSender1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewSender1ClientsByURL(loc *url.URL) ([]*Sender1, error) {
	return NewSender1ClientsByURLCtx(context.Background(), loc)
}

// NewSender1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewSender1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Sender1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Sender_1)
	if err != nil {
		return nil, err
	}
	return newSender1ClientsFromGenericClients(genericClients), nil
}

func newSender1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Sender1 {
	clients := make([]*Sender1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Sender1{genericClients[i]}
	}
	return clients
}

func (client *Sender1) AttributesCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_1, "Attributes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Attributes is the legacy version of AttributesCtx, but uses
// context.Background() as the context.
func (client *Sender1) Attributes() (Value string, err error) {
	return client.AttributesCtx(context.Background())
}

func (client *Sender1) AudioCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_1, "Audio", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Audio is the legacy version of AudioCtx, but uses
// context.Background() as the context.
func (client *Sender1) Audio() (Value bool, err error) {
	return client.AudioCtx(context.Background())
}

func (client *Sender1) MetadataCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_1, "Metadata", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Metadata is the legacy version of MetadataCtx, but uses
// context.Background() as the context.
func (client *Sender1) Metadata() (Value string, err error) {
	return client.MetadataCtx(context.Background())
}

func (client *Sender1) PresentationUrlCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_1, "PresentationUrl", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PresentationUrl is the legacy version of PresentationUrlCtx, but uses
// context.Background() as the context.
func (client *Sender1) PresentationUrl() (Value string, err error) {
	return client.PresentationUrlCtx(context.Background())
}

// Return values:
//
// * Value: allowed values: Enabled, Disabled, Blocked
func (client *Sender1) StatusCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_1, "Status", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Status is the legacy version of StatusCtx, but uses
// context.Background() as the context.
func (client *Sender1) Status() (Value string, err error) {
	return client.StatusCtx(context.Background())
}

// Sender2 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Sender:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Sender2 struct {
	goupnp.ServiceClient
}

// NewSender2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewSender2ClientsCtx(ctx context.Context) (clients []*Sender2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Sender_2); err != nil {
		return
	}
	clients = newSender2ClientsFromGenericClients(genericClients)
	return
}

// NewSender2Clients is the legacy version of NewSender2ClientsCtx, but uses
// context.Background() as the context.
func NewSender2Clients() (clients []*Sender2, errors []error, err error) {
	return NewSender2ClientsCtx(context.Background())
}

// NewSender2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewSender2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Sender2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Sender_2)
	if err != nil {
		return nil, err
	}
	return newSender2ClientsFromGenericClients(genericClients), nil
}

// NewSender2ClientsByURL is the legacy version of NewSender2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewSender2ClientsByURL(loc *url.URL) ([]*Sender2, error) {
	return NewSender2ClientsByURLCtx(context.Background(), loc)
}

// NewSender2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewSender2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Sender2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Sender_2)
	if err != nil {
		return nil, err
	}
	return newSender2ClientsFromGenericClients(genericClients), nil
}

func newSender2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Sender2 {
	clients := make([]*Sender2, len(genericClients))
	for i := range genericClients {
		clients[i] = &Sender2{genericClients[i]}
	}
	return clients
}

func (client *Sender2) AttributesCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "Attributes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Attributes is the legacy version of AttributesCtx, but uses
// context.Background() as the context.
func (client *Sender2) Attributes() (Value string, err error) {
	return client.AttributesCtx(context.Background())
}

func (client *Sender2) AudioCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "Audio", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Audio is the legacy version of AudioCtx, but uses
// context.Background() as the context.
func (client *Sender2) Audio() (Value bool, err error) {
	return client.AudioCtx(context.Background())
}

func (client *Sender2) EnabledCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "Enabled", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Enabled is the legacy version of EnabledCtx, but uses
// context.Background() as the context.
func (client *Sender2) Enabled() (Value bool, err error) {
	return client.EnabledCtx(context.Background())
}

func (client *Sender2) MetadataCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "Metadata", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Metadata is the legacy version of MetadataCtx, but uses
// context.Background() as the context.
func (client *Sender2) Metadata() (Value string, err error) {
	return client.MetadataCtx(context.Background())
}

func (client *Sender2) PresentationUrlCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "PresentationUrl", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// PresentationUrl is the legacy version of PresentationUrlCtx, but uses
// context.Background() as the context.
func (client *Sender2) PresentationUrl() (Value string, err error) {
	return client.PresentationUrlCtx(context.Background())
}

// Return values:
//
// * Value: allowed values: Enabled, Disabled, Blocked
func (client *Sender2) StatusCtx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "Status", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Status is the legacy version of StatusCtx, but uses
// context.Background() as the context.
func (client *Sender2) Status() (Value string, err error) {
	return client.StatusCtx(context.Background())
}

// Return values:
//
// * Value: allowed values: Sending, Ready, Blocked, Inactive, Disabled
func (client *Sender2) Status2Ctx(
	ctx context.Context,
) (Value string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Sender_2, "Status2", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalString(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Status2 is the legacy version of Status2Ctx, but uses
// context.Background() as the context.
func (client *Sender2) Status2() (Value string, err error) {
	return client.Status2Ctx(context.Background())
}

// SubscriptionLongPoll1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:SubscriptionLongPoll:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type SubscriptionLongPoll1 struct {
	goupnp.ServiceClient
}

// NewSubscriptionLongPoll1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewSubscriptionLongPoll1ClientsCtx(ctx context.Context) (clients []*SubscriptionLongPoll1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_SubscriptionLongPoll_1); err != nil {
		return
	}
	clients = newSubscriptionLongPoll1ClientsFromGenericClients(genericClients)
	return
}

// NewSubscriptionLongPoll1Clients is the legacy version of NewSubscriptionLongPoll1ClientsCtx, but uses
// context.Background() as the context.
func NewSubscriptionLongPoll1Clients() (clients []*SubscriptionLongPoll1, errors []error, err error) {
	return NewSubscriptionLongPoll1ClientsCtx(context.Background())
}

// NewSubscriptionLongPoll1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewSubscriptionLongPoll1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*SubscriptionLongPoll1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_SubscriptionLongPoll_1)
	if err != nil {
		return nil, err
	}
	return newSubscriptionLongPoll1ClientsFromGenericClients(genericClients), nil
}

// NewSubscriptionLongPoll1ClientsByURL is the legacy version of NewSubscriptionLongPoll1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewSubscriptionLongPoll1ClientsByURL(loc *url.URL) ([]*SubscriptionLongPoll1, error) {
	return NewSubscriptionLongPoll1ClientsByURLCtx(context.Background(), loc)
}

// NewSubscriptionLongPoll1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewSubscriptionLongPoll1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*SubscriptionLongPoll1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_SubscriptionLongPoll_1)
	if err != nil {
		return nil, err
	}
	return newSubscriptionLongPoll1ClientsFromGenericClients(genericClients), nil
}

func newSubscriptionLongPoll1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*SubscriptionLongPoll1 {
	clients := make([]*SubscriptionLongPoll1, len(genericClients))
	for i := range genericClients {
		clients[i] = &SubscriptionLongPoll1{genericClients[i]}
	}
	return clients
}

func (client *SubscriptionLongPoll1) GetPropertyUpdatesCtx(
	ctx context.Context,
	ClientId string,
) (Updates string, err error) {
	// Request structure.
	request := &struct {
		ClientId string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ClientId, err = soap.MarshalString(ClientId); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Updates string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_SubscriptionLongPoll_1, "GetPropertyUpdates", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Updates, err = soap.UnmarshalString(response.Updates); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// GetPropertyUpdates is the legacy version of GetPropertyUpdatesCtx, but uses
// context.Background() as the context.
func (client *SubscriptionLongPoll1) GetPropertyUpdates(ClientId string) (Updates string, err error) {
	return client.GetPropertyUpdatesCtx(context.Background(),
		ClientId,
	)
}

func (client *SubscriptionLongPoll1) RenewCtx(
	ctx context.Context,
	Sid string,
	RequestedDuration uint32,
) (Duration uint32, err error) {
	// Request structure.
	request := &struct {
		Sid               string
		RequestedDuration string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Sid, err = soap.MarshalString(Sid); err != nil {
		return
	}
	if request.RequestedDuration, err = soap.MarshalUi4(RequestedDuration); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Duration string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_SubscriptionLongPoll_1, "Renew", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Duration, err = soap.UnmarshalUi4(response.Duration); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Renew is the legacy version of RenewCtx, but uses
// context.Background() as the context.
func (client *SubscriptionLongPoll1) Renew(Sid string, RequestedDuration uint32) (Duration uint32, err error) {
	return client.RenewCtx(context.Background(),
		Sid,
		RequestedDuration,
	)
}

func (client *SubscriptionLongPoll1) SubscribeCtx(
	ctx context.Context,
	ClientId string,
	Udn string,
	Service string,
	RequestedDuration uint32,
) (Sid string, Duration uint32, err error) {
	// Request structure.
	request := &struct {
		ClientId          string
		Udn               string
		Service           string
		RequestedDuration string
	}{}
	// BEGIN Marshal arguments into request.

	if request.ClientId, err = soap.MarshalString(ClientId); err != nil {
		return
	}
	if request.Udn, err = soap.MarshalString(Udn); err != nil {
		return
	}
	if request.Service, err = soap.MarshalString(Service); err != nil {
		return
	}
	if request.RequestedDuration, err = soap.MarshalUi4(RequestedDuration); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Sid      string
		Duration string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_SubscriptionLongPoll_1, "Subscribe", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Sid, err = soap.UnmarshalString(response.Sid); err != nil {
		return
	}
	if Duration, err = soap.UnmarshalUi4(response.Duration); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Subscribe is the legacy version of SubscribeCtx, but uses
// context.Background() as the context.
func (client *SubscriptionLongPoll1) Subscribe(ClientId string, Udn string, Service string, RequestedDuration uint32) (Sid string, Duration uint32, err error) {
	return client.SubscribeCtx(context.Background(),
		ClientId,
		Udn,
		Service,
		RequestedDuration,
	)
}

func (client *SubscriptionLongPoll1) UnsubscribeCtx(
	ctx context.Context,
	Sid string,
) (err error) {
	// Request structure.
	request := &struct {
		Sid string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Sid, err = soap.MarshalString(Sid); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_SubscriptionLongPoll_1, "Unsubscribe", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Unsubscribe is the legacy version of UnsubscribeCtx, but uses
// context.Background() as the context.
func (client *SubscriptionLongPoll1) Unsubscribe(Sid string) (err error) {
	return client.UnsubscribeCtx(context.Background(),
		Sid,
	)
}

// Time1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Time:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Time1 struct {
	goupnp.ServiceClient
}

// NewTime1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewTime1ClientsCtx(ctx context.Context) (clients []*Time1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Time_1); err != nil {
		return
	}
	clients = newTime1ClientsFromGenericClients(genericClients)
	return
}

// NewTime1Clients is the legacy version of NewTime1ClientsCtx, but uses
// context.Background() as the context.
func NewTime1Clients() (clients []*Time1, errors []error, err error) {
	return NewTime1ClientsCtx(context.Background())
}

// NewTime1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewTime1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Time1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Time_1)
	if err != nil {
		return nil, err
	}
	return newTime1ClientsFromGenericClients(genericClients), nil
}

// NewTime1ClientsByURL is the legacy version of NewTime1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewTime1ClientsByURL(loc *url.URL) ([]*Time1, error) {
	return NewTime1ClientsByURLCtx(context.Background(), loc)
}

// NewTime1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewTime1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Time1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Time_1)
	if err != nil {
		return nil, err
	}
	return newTime1ClientsFromGenericClients(genericClients), nil
}

func newTime1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Time1 {
	clients := make([]*Time1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Time1{genericClients[i]}
	}
	return clients
}

func (client *Time1) TimeCtx(
	ctx context.Context,
) (TrackCount uint32, Duration uint32, Seconds uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TrackCount string
		Duration   string
		Seconds    string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Time_1, "Time", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TrackCount, err = soap.UnmarshalUi4(response.TrackCount); err != nil {
		return
	}
	if Duration, err = soap.UnmarshalUi4(response.Duration); err != nil {
		return
	}
	if Seconds, err = soap.UnmarshalUi4(response.Seconds); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Time is the legacy version of TimeCtx, but uses
// context.Background() as the context.
func (client *Time1) Time() (TrackCount uint32, Duration uint32, Seconds uint32, err error) {
	return client.TimeCtx(context.Background())
}

// Transport1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Transport:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Transport1 struct {
	goupnp.ServiceClient
}

// NewTransport1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewTransport1ClientsCtx(ctx context.Context) (clients []*Transport1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Transport_1); err != nil {
		return
	}
	clients = newTransport1ClientsFromGenericClients(genericClients)
	return
}

// NewTransport1Clients is the legacy version of NewTransport1ClientsCtx, but uses
// context.Background() as the context.
func NewTransport1Clients() (clients []*Transport1, errors []error, err error) {
	return NewTransport1ClientsCtx(context.Background())
}

// NewTransport1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewTransport1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Transport1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Transport_1)
	if err != nil {
		return nil, err
	}
	return newTransport1ClientsFromGenericClients(genericClients), nil
}

// NewTransport1ClientsByURL is the legacy version of NewTransport1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewTransport1ClientsByURL(loc *url.URL) ([]*Transport1, error) {
	return NewTransport1ClientsByURLCtx(context.Background(), loc)
}

// NewTransport1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewTransport1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Transport1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Transport_1)
	if err != nil {
		return nil, err
	}
	return newTransport1ClientsFromGenericClients(genericClients), nil
}

func newTransport1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Transport1 {
	clients := make([]*Transport1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Transport1{genericClients[i]}
	}
	return clients
}

func (client *Transport1) ModeInfoCtx(
	ctx context.Context,
) (CanSkipNext bool, CanSkipPrevious bool, CanRepeat bool, CanShuffle bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		CanSkipNext     string
		CanSkipPrevious string
		CanRepeat       string
		CanShuffle      string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "ModeInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if CanSkipNext, err = soap.UnmarshalBoolean(response.CanSkipNext); err != nil {
		return
	}
	if CanSkipPrevious, err = soap.UnmarshalBoolean(response.CanSkipPrevious); err != nil {
		return
	}
	if CanRepeat, err = soap.UnmarshalBoolean(response.CanRepeat); err != nil {
		return
	}
	if CanShuffle, err = soap.UnmarshalBoolean(response.CanShuffle); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// ModeInfo is the legacy version of ModeInfoCtx, but uses
// context.Background() as the context.
func (client *Transport1) ModeInfo() (CanSkipNext bool, CanSkipPrevious bool, CanRepeat bool, CanShuffle bool, err error) {
	return client.ModeInfoCtx(context.Background())
}

func (client *Transport1) ModesCtx(
	ctx context.Context,
) (Modes string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Modes string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "Modes", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Modes, err = soap.UnmarshalString(response.Modes); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Modes is the legacy version of ModesCtx, but uses
// context.Background() as the context.
func (client *Transport1) Modes() (Modes string, err error) {
	return client.ModesCtx(context.Background())
}

func (client *Transport1) PauseCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "Pause", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *Transport1) Pause() (err error) {
	return client.PauseCtx(context.Background())
}

func (client *Transport1) PlayCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "Play", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *Transport1) Play() (err error) {
	return client.PlayCtx(context.Background())
}

func (client *Transport1) PlayAsCtx(
	ctx context.Context,
	Mode string,
	Command string,
) (err error) {
	// Request structure.
	request := &struct {
		Mode    string
		Command string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Mode, err = soap.MarshalString(Mode); err != nil {
		return
	}
	if request.Command, err = soap.MarshalString(Command); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "PlayAs", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// PlayAs is the legacy version of PlayAsCtx, but uses
// context.Background() as the context.
func (client *Transport1) PlayAs(Mode string, Command string) (err error) {
	return client.PlayAsCtx(context.Background(),
		Mode,
		Command,
	)
}

func (client *Transport1) RepeatCtx(
	ctx context.Context,
) (Repeat bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Repeat string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "Repeat", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Repeat, err = soap.UnmarshalBoolean(response.Repeat); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Repeat is the legacy version of RepeatCtx, but uses
// context.Background() as the context.
func (client *Transport1) Repeat() (Repeat bool, err error) {
	return client.RepeatCtx(context.Background())
}

func (client *Transport1) SeekSecondAbsoluteCtx(
	ctx context.Context,
	StreamId uint32,
	SecondAbsolute uint32,
) (err error) {
	// Request structure.
	request := &struct {
		StreamId       string
		SecondAbsolute string
	}{}
	// BEGIN Marshal arguments into request.

	if request.StreamId, err = soap.MarshalUi4(StreamId); err != nil {
		return
	}
	if request.SecondAbsolute, err = soap.MarshalUi4(SecondAbsolute); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "SeekSecondAbsolute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekSecondAbsolute is the legacy version of SeekSecondAbsoluteCtx, but uses
// context.Background() as the context.
func (client *Transport1) SeekSecondAbsolute(StreamId uint32, SecondAbsolute uint32) (err error) {
	return client.SeekSecondAbsoluteCtx(context.Background(),
		StreamId,
		SecondAbsolute,
	)
}

func (client *Transport1) SeekSecondRelativeCtx(
	ctx context.Context,
	StreamId uint32,
	SecondRelative int32,
) (err error) {
	// Request structure.
	request := &struct {
		StreamId       string
		SecondRelative string
	}{}
	// BEGIN Marshal arguments into request.

	if request.StreamId, err = soap.MarshalUi4(StreamId); err != nil {
		return
	}
	if request.SecondRelative, err = soap.MarshalI4(SecondRelative); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "SeekSecondRelative", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SeekSecondRelative is the legacy version of SeekSecondRelativeCtx, but uses
// context.Background() as the context.
func (client *Transport1) SeekSecondRelative(StreamId uint32, SecondRelative int32) (err error) {
	return client.SeekSecondRelativeCtx(context.Background(),
		StreamId,
		SecondRelative,
	)
}

func (client *Transport1) SetRepeatCtx(
	ctx context.Context,
	Repeat bool,
) (err error) {
	// Request structure.
	request := &struct {
		Repeat string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Repeat, err = soap.MarshalBoolean(Repeat); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "SetRepeat", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetRepeat is the legacy version of SetRepeatCtx, but uses
// context.Background() as the context.
func (client *Transport1) SetRepeat(Repeat bool) (err error) {
	return client.SetRepeatCtx(context.Background(),
		Repeat,
	)
}

func (client *Transport1) SetShuffleCtx(
	ctx context.Context,
	Shuffle bool,
) (err error) {
	// Request structure.
	request := &struct {
		Shuffle string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Shuffle, err = soap.MarshalBoolean(Shuffle); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "SetShuffle", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetShuffle is the legacy version of SetShuffleCtx, but uses
// context.Background() as the context.
func (client *Transport1) SetShuffle(Shuffle bool) (err error) {
	return client.SetShuffleCtx(context.Background(),
		Shuffle,
	)
}

func (client *Transport1) ShuffleCtx(
	ctx context.Context,
) (Shuffle bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Shuffle string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "Shuffle", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Shuffle, err = soap.UnmarshalBoolean(response.Shuffle); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Shuffle is the legacy version of ShuffleCtx, but uses
// context.Background() as the context.
func (client *Transport1) Shuffle() (Shuffle bool, err error) {
	return client.ShuffleCtx(context.Background())
}

func (client *Transport1) SkipNextCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "SkipNext", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SkipNext is the legacy version of SkipNextCtx, but uses
// context.Background() as the context.
func (client *Transport1) SkipNext() (err error) {
	return client.SkipNextCtx(context.Background())
}

func (client *Transport1) SkipPreviousCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "SkipPrevious", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SkipPrevious is the legacy version of SkipPreviousCtx, but uses
// context.Background() as the context.
func (client *Transport1) SkipPrevious() (err error) {
	return client.SkipPreviousCtx(context.Background())
}

func (client *Transport1) StopCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "Stop", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *Transport1) Stop() (err error) {
	return client.StopCtx(context.Background())
}

func (client *Transport1) StreamIdCtx(
	ctx context.Context,
) (StreamId uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		StreamId string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "StreamId", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if StreamId, err = soap.UnmarshalUi4(response.StreamId); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// StreamId is the legacy version of StreamIdCtx, but uses
// context.Background() as the context.
func (client *Transport1) StreamId() (StreamId uint32, err error) {
	return client.StreamIdCtx(context.Background())
}

func (client *Transport1) StreamInfoCtx(
	ctx context.Context,
) (StreamId uint32, CanSeek bool, CanPause bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		StreamId string
		CanSeek  string
		CanPause string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "StreamInfo", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if StreamId, err = soap.UnmarshalUi4(response.StreamId); err != nil {
		return
	}
	if CanSeek, err = soap.UnmarshalBoolean(response.CanSeek); err != nil {
		return
	}
	if CanPause, err = soap.UnmarshalBoolean(response.CanPause); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// StreamInfo is the legacy version of StreamInfoCtx, but uses
// context.Background() as the context.
func (client *Transport1) StreamInfo() (StreamId uint32, CanSeek bool, CanPause bool, err error) {
	return client.StreamInfoCtx(context.Background())
}

// Return values:
//
// * State: allowed values: Playing, Paused, Stopped, Buffering, Waiting
func (client *Transport1) TransportStateCtx(
	ctx context.Context,
) (State string, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		State string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Transport_1, "TransportState", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if State, err = soap.UnmarshalString(response.State); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// TransportState is the legacy version of TransportStateCtx, but uses
// context.Background() as the context.
func (client *Transport1) TransportState() (State string, err error) {
	return client.TransportStateCtx(context.Background())
}

// Volume1 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Volume:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Volume1 struct {
	goupnp.ServiceClient
}

// NewVolume1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewVolume1ClientsCtx(ctx context.Context) (clients []*Volume1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Volume_1); err != nil {
		return
	}
	clients = newVolume1ClientsFromGenericClients(genericClients)
	return
}

// NewVolume1Clients is the legacy version of NewVolume1ClientsCtx, but uses
// context.Background() as the context.
func NewVolume1Clients() (clients []*Volume1, errors []error, err error) {
	return NewVolume1ClientsCtx(context.Background())
}

// NewVolume1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewVolume1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Volume1, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Volume_1)
	if err != nil {
		return nil, err
	}
	return newVolume1ClientsFromGenericClients(genericClients), nil
}

// NewVolume1ClientsByURL is the legacy version of NewVolume1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewVolume1ClientsByURL(loc *url.URL) ([]*Volume1, error) {
	return NewVolume1ClientsByURLCtx(context.Background(), loc)
}

// NewVolume1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewVolume1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Volume1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Volume_1)
	if err != nil {
		return nil, err
	}
	return newVolume1ClientsFromGenericClients(genericClients), nil
}

func newVolume1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Volume1 {
	clients := make([]*Volume1, len(genericClients))
	for i := range genericClients {
		clients[i] = &Volume1{genericClients[i]}
	}
	return clients
}

func (client *Volume1) BalanceCtx(
	ctx context.Context,
) (Value int32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "Balance", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalI4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Balance is the legacy version of BalanceCtx, but uses
// context.Background() as the context.
func (client *Volume1) Balance() (Value int32, err error) {
	return client.BalanceCtx(context.Background())
}

func (client *Volume1) BalanceDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "BalanceDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// BalanceDec is the legacy version of BalanceDecCtx, but uses
// context.Background() as the context.
func (client *Volume1) BalanceDec() (err error) {
	return client.BalanceDecCtx(context.Background())
}

func (client *Volume1) BalanceIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "BalanceInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// BalanceInc is the legacy version of BalanceIncCtx, but uses
// context.Background() as the context.
func (client *Volume1) BalanceInc() (err error) {
	return client.BalanceIncCtx(context.Background())
}

func (client *Volume1) CharacteristicsCtx(
	ctx context.Context,
) (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32, VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		VolumeMax            string
		VolumeUnity          string
		VolumeSteps          string
		VolumeMilliDbPerStep string
		BalanceMax           string
		FadeMax              string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "Characteristics", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if VolumeMax, err = soap.UnmarshalUi4(response.VolumeMax); err != nil {
		return
	}
	if VolumeUnity, err = soap.UnmarshalUi4(response.VolumeUnity); err != nil {
		return
	}
	if VolumeSteps, err = soap.UnmarshalUi4(response.VolumeSteps); err != nil {
		return
	}
	if VolumeMilliDbPerStep, err = soap.UnmarshalUi4(response.VolumeMilliDbPerStep); err != nil {
		return
	}
	if BalanceMax, err = soap.UnmarshalUi4(response.BalanceMax); err != nil {
		return
	}
	if FadeMax, err = soap.UnmarshalUi4(response.FadeMax); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Characteristics is the legacy version of CharacteristicsCtx, but uses
// context.Background() as the context.
func (client *Volume1) Characteristics() (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32, VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	return client.CharacteristicsCtx(context.Background())
}

func (client *Volume1) FadeCtx(
	ctx context.Context,
) (Value int32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "Fade", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalI4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Fade is the legacy version of FadeCtx, but uses
// context.Background() as the context.
func (client *Volume1) Fade() (Value int32, err error) {
	return client.FadeCtx(context.Background())
}

func (client *Volume1) FadeDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "FadeDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// FadeDec is the legacy version of FadeDecCtx, but uses
// context.Background() as the context.
func (client *Volume1) FadeDec() (err error) {
	return client.FadeDecCtx(context.Background())
}

func (client *Volume1) FadeIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "FadeInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// FadeInc is the legacy version of FadeIncCtx, but uses
// context.Background() as the context.
func (client *Volume1) FadeInc() (err error) {
	return client.FadeIncCtx(context.Background())
}

func (client *Volume1) MuteCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "Mute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Mute is the legacy version of MuteCtx, but uses
// context.Background() as the context.
func (client *Volume1) Mute() (Value bool, err error) {
	return client.MuteCtx(context.Background())
}

func (client *Volume1) SetBalanceCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "SetBalance", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBalance is the legacy version of SetBalanceCtx, but uses
// context.Background() as the context.
func (client *Volume1) SetBalance(Value int32) (err error) {
	return client.SetBalanceCtx(context.Background(),
		Value,
	)
}

func (client *Volume1) SetFadeCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "SetFade", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetFade is the legacy version of SetFadeCtx, but uses
// context.Background() as the context.
func (client *Volume1) SetFade(Value int32) (err error) {
	return client.SetFadeCtx(context.Background(),
		Value,
	)
}

func (client *Volume1) SetMuteCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "SetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *Volume1) SetMute(Value bool) (err error) {
	return client.SetMuteCtx(context.Background(),
		Value,
	)
}

func (client *Volume1) SetVolumeCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "SetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *Volume1) SetVolume(Value uint32) (err error) {
	return client.SetVolumeCtx(context.Background(),
		Value,
	)
}

func (client *Volume1) VolumeCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "Volume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Volume is the legacy version of VolumeCtx, but uses
// context.Background() as the context.
func (client *Volume1) Volume() (Value uint32, err error) {
	return client.VolumeCtx(context.Background())
}

func (client *Volume1) VolumeDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "VolumeDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// VolumeDec is the legacy version of VolumeDecCtx, but uses
// context.Background() as the context.
func (client *Volume1) VolumeDec() (err error) {
	return client.VolumeDecCtx(context.Background())
}

func (client *Volume1) VolumeIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "VolumeInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// VolumeInc is the legacy version of VolumeIncCtx, but uses
// context.Background() as the context.
func (client *Volume1) VolumeInc() (err error) {
	return client.VolumeIncCtx(context.Background())
}

func (client *Volume1) VolumeLimitCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_1, "VolumeLimit", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// VolumeLimit is the legacy version of VolumeLimitCtx, but uses
// context.Background() as the context.
func (client *Volume1) VolumeLimit() (Value uint32, err error) {
	return client.VolumeLimitCtx(context.Background())
}

// Volume2 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Volume:2". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Volume2 struct {
	goupnp.ServiceClient
}

// NewVolume2ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewVolume2ClientsCtx(ctx context.Context) (clients []*Volume2, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Volume_2); err != nil {
		return
	}
	clients = newVolume2ClientsFromGenericClients(genericClients)
	return
}

// NewVolume2Clients is the legacy version of NewVolume2ClientsCtx, but uses
// context.Background() as the context.
func NewVolume2Clients() (clients []*Volume2, errors []error, err error) {
	return NewVolume2ClientsCtx(context.Background())
}

// NewVolume2ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewVolume2ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Volume2, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Volume_2)
	if err != nil {
		return nil, err
	}
	return newVolume2ClientsFromGenericClients(genericClients), nil
}

// NewVolume2ClientsByURL is the legacy version of NewVolume2ClientsByURLCtx, but uses
// context.Background() as the context.
func NewVolume2ClientsByURL(loc *url.URL) ([]*Volume2, error) {
	return NewVolume2ClientsByURLCtx(context.Background(), loc)
}

// NewVolume2ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewVolume2ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Volume2, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Volume_2)
	if err != nil {
		return nil, err
	}
	return newVolume2ClientsFromGenericClients(genericClients), nil
}

func newVolume2ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Volume2 {
	clients := make([]*Volume2, len(genericClients))
	for i := range genericClients {
		clients[i] = &Volume2{genericClients[i]}
	}
	return clients
}

func (client *Volume2) BalanceCtx(
	ctx context.Context,
) (Value int32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "Balance", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalI4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Balance is the legacy version of BalanceCtx, but uses
// context.Background() as the context.
func (client *Volume2) Balance() (Value int32, err error) {
	return client.BalanceCtx(context.Background())
}

func (client *Volume2) BalanceDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "BalanceDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// BalanceDec is the legacy version of BalanceDecCtx, but uses
// context.Background() as the context.
func (client *Volume2) BalanceDec() (err error) {
	return client.BalanceDecCtx(context.Background())
}

func (client *Volume2) BalanceIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "BalanceInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// BalanceInc is the legacy version of BalanceIncCtx, but uses
// context.Background() as the context.
func (client *Volume2) BalanceInc() (err error) {
	return client.BalanceIncCtx(context.Background())
}

func (client *Volume2) CharacteristicsCtx(
	ctx context.Context,
) (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32, VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		VolumeMax            string
		VolumeUnity          string
		VolumeSteps          string
		VolumeMilliDbPerStep string
		BalanceMax           string
		FadeMax              string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "Characteristics", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if VolumeMax, err = soap.UnmarshalUi4(response.VolumeMax); err != nil {
		return
	}
	if VolumeUnity, err = soap.UnmarshalUi4(response.VolumeUnity); err != nil {
		return
	}
	if VolumeSteps, err = soap.UnmarshalUi4(response.VolumeSteps); err != nil {
		return
	}
	if VolumeMilliDbPerStep, err = soap.UnmarshalUi4(response.VolumeMilliDbPerStep); err != nil {
		return
	}
	if BalanceMax, err = soap.UnmarshalUi4(response.BalanceMax); err != nil {
		return
	}
	if FadeMax, err = soap.UnmarshalUi4(response.FadeMax); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Characteristics is the legacy version of CharacteristicsCtx, but uses
// context.Background() as the context.
func (client *Volume2) Characteristics() (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32, VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	return client.CharacteristicsCtx(context.Background())
}

func (client *Volume2) FadeCtx(
	ctx context.Context,
) (Value int32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "Fade", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalI4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Fade is the legacy version of FadeCtx, but uses
// context.Background() as the context.
func (client *Volume2) Fade() (Value int32, err error) {
	return client.FadeCtx(context.Background())
}

func (client *Volume2) FadeDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "FadeDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// FadeDec is the legacy version of FadeDecCtx, but uses
// context.Background() as the context.
func (client *Volume2) FadeDec() (err error) {
	return client.FadeDecCtx(context.Background())
}

func (client *Volume2) FadeIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "FadeInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// FadeInc is the legacy version of FadeIncCtx, but uses
// context.Background() as the context.
func (client *Volume2) FadeInc() (err error) {
	return client.FadeIncCtx(context.Background())
}

func (client *Volume2) MuteCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "Mute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Mute is the legacy version of MuteCtx, but uses
// context.Background() as the context.
func (client *Volume2) Mute() (Value bool, err error) {
	return client.MuteCtx(context.Background())
}

func (client *Volume2) SetBalanceCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "SetBalance", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBalance is the legacy version of SetBalanceCtx, but uses
// context.Background() as the context.
func (client *Volume2) SetBalance(Value int32) (err error) {
	return client.SetBalanceCtx(context.Background(),
		Value,
	)
}

func (client *Volume2) SetFadeCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "SetFade", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetFade is the legacy version of SetFadeCtx, but uses
// context.Background() as the context.
func (client *Volume2) SetFade(Value int32) (err error) {
	return client.SetFadeCtx(context.Background(),
		Value,
	)
}

func (client *Volume2) SetMuteCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "SetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *Volume2) SetMute(Value bool) (err error) {
	return client.SetMuteCtx(context.Background(),
		Value,
	)
}

func (client *Volume2) SetVolumeCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "SetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *Volume2) SetVolume(Value uint32) (err error) {
	return client.SetVolumeCtx(context.Background(),
		Value,
	)
}

func (client *Volume2) UnityGainCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "UnityGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// UnityGain is the legacy version of UnityGainCtx, but uses
// context.Background() as the context.
func (client *Volume2) UnityGain() (Value bool, err error) {
	return client.UnityGainCtx(context.Background())
}

func (client *Volume2) VolumeCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "Volume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Volume is the legacy version of VolumeCtx, but uses
// context.Background() as the context.
func (client *Volume2) Volume() (Value uint32, err error) {
	return client.VolumeCtx(context.Background())
}

func (client *Volume2) VolumeDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "VolumeDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// VolumeDec is the legacy version of VolumeDecCtx, but uses
// context.Background() as the context.
func (client *Volume2) VolumeDec() (err error) {
	return client.VolumeDecCtx(context.Background())
}

func (client *Volume2) VolumeIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "VolumeInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// VolumeInc is the legacy version of VolumeIncCtx, but uses
// context.Background() as the context.
func (client *Volume2) VolumeInc() (err error) {
	return client.VolumeIncCtx(context.Background())
}

func (client *Volume2) VolumeLimitCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_2, "VolumeLimit", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// VolumeLimit is the legacy version of VolumeLimitCtx, but uses
// context.Background() as the context.
func (client *Volume2) VolumeLimit() (Value uint32, err error) {
	return client.VolumeLimitCtx(context.Background())
}

// Volume3 is a client for UPnP SOAP service with URN "urn:av-openhome-org:service:Volume:3". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type Volume3 struct {
	goupnp.ServiceClient
}

// NewVolume3ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewVolume3ClientsCtx(ctx context.Context) (clients []*Volume3, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_Volume_3); err != nil {
		return
	}
	clients = newVolume3ClientsFromGenericClients(genericClients)
	return
}

// NewVolume3Clients is the legacy version of NewVolume3ClientsCtx, but uses
// context.Background() as the context.
func NewVolume3Clients() (clients []*Volume3, errors []error, err error) {
	return NewVolume3ClientsCtx(context.Background())
}

// NewVolume3ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewVolume3ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*Volume3, error) {
	genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_Volume_3)
	if err != nil {
		return nil, err
	}
	return newVolume3ClientsFromGenericClients(genericClients), nil
}

// NewVolume3ClientsByURL is the legacy version of NewVolume3ClientsByURLCtx, but uses
// context.Background() as the context.
func NewVolume3ClientsByURL(loc *url.URL) ([]*Volume3, error) {
	return NewVolume3ClientsByURLCtx(context.Background(), loc)
}

// NewVolume3ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewVolume3ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*Volume3, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_Volume_3)
	if err != nil {
		return nil, err
	}
	return newVolume3ClientsFromGenericClients(genericClients), nil
}

func newVolume3ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*Volume3 {
	clients := make([]*Volume3, len(genericClients))
	for i := range genericClients {
		clients[i] = &Volume3{genericClients[i]}
	}
	return clients
}

func (client *Volume3) BalanceCtx(
	ctx context.Context,
) (Value int32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "Balance", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalI4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Balance is the legacy version of BalanceCtx, but uses
// context.Background() as the context.
func (client *Volume3) Balance() (Value int32, err error) {
	return client.BalanceCtx(context.Background())
}

func (client *Volume3) BalanceDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "BalanceDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// BalanceDec is the legacy version of BalanceDecCtx, but uses
// context.Background() as the context.
func (client *Volume3) BalanceDec() (err error) {
	return client.BalanceDecCtx(context.Background())
}

func (client *Volume3) BalanceIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "BalanceInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// BalanceInc is the legacy version of BalanceIncCtx, but uses
// context.Background() as the context.
func (client *Volume3) BalanceInc() (err error) {
	return client.BalanceIncCtx(context.Background())
}

func (client *Volume3) CharacteristicsCtx(
	ctx context.Context,
) (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32, VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		VolumeMax            string
		VolumeUnity          string
		VolumeSteps          string
		VolumeMilliDbPerStep string
		BalanceMax           string
		FadeMax              string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "Characteristics", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if VolumeMax, err = soap.UnmarshalUi4(response.VolumeMax); err != nil {
		return
	}
	if VolumeUnity, err = soap.UnmarshalUi4(response.VolumeUnity); err != nil {
		return
	}
	if VolumeSteps, err = soap.UnmarshalUi4(response.VolumeSteps); err != nil {
		return
	}
	if VolumeMilliDbPerStep, err = soap.UnmarshalUi4(response.VolumeMilliDbPerStep); err != nil {
		return
	}
	if BalanceMax, err = soap.UnmarshalUi4(response.BalanceMax); err != nil {
		return
	}
	if FadeMax, err = soap.UnmarshalUi4(response.FadeMax); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Characteristics is the legacy version of CharacteristicsCtx, but uses
// context.Background() as the context.
func (client *Volume3) Characteristics() (VolumeMax uint32, VolumeUnity uint32, VolumeSteps uint32, VolumeMilliDbPerStep uint32, BalanceMax uint32, FadeMax uint32, err error) {
	return client.CharacteristicsCtx(context.Background())
}

func (client *Volume3) FadeCtx(
	ctx context.Context,
) (Value int32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "Fade", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalI4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Fade is the legacy version of FadeCtx, but uses
// context.Background() as the context.
func (client *Volume3) Fade() (Value int32, err error) {
	return client.FadeCtx(context.Background())
}

func (client *Volume3) FadeDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "FadeDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// FadeDec is the legacy version of FadeDecCtx, but uses
// context.Background() as the context.
func (client *Volume3) FadeDec() (err error) {
	return client.FadeDecCtx(context.Background())
}

func (client *Volume3) FadeIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "FadeInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// FadeInc is the legacy version of FadeIncCtx, but uses
// context.Background() as the context.
func (client *Volume3) FadeInc() (err error) {
	return client.FadeIncCtx(context.Background())
}

func (client *Volume3) MuteCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "Mute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Mute is the legacy version of MuteCtx, but uses
// context.Background() as the context.
func (client *Volume3) Mute() (Value bool, err error) {
	return client.MuteCtx(context.Background())
}

func (client *Volume3) SetBalanceCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "SetBalance", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetBalance is the legacy version of SetBalanceCtx, but uses
// context.Background() as the context.
func (client *Volume3) SetBalance(Value int32) (err error) {
	return client.SetBalanceCtx(context.Background(),
		Value,
	)
}

func (client *Volume3) SetFadeCtx(
	ctx context.Context,
	Value int32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalI4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "SetFade", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetFade is the legacy version of SetFadeCtx, but uses
// context.Background() as the context.
func (client *Volume3) SetFade(Value int32) (err error) {
	return client.SetFadeCtx(context.Background(),
		Value,
	)
}

func (client *Volume3) SetMuteCtx(
	ctx context.Context,
	Value bool,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalBoolean(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "SetMute", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *Volume3) SetMute(Value bool) (err error) {
	return client.SetMuteCtx(context.Background(),
		Value,
	)
}

func (client *Volume3) SetTrimCtx(
	ctx context.Context,
	Channel string,
	TrimBinaryMilliDb int32,
) (err error) {
	// Request structure.
	request := &struct {
		Channel           string
		TrimBinaryMilliDb string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	if request.TrimBinaryMilliDb, err = soap.MarshalI4(TrimBinaryMilliDb); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "SetTrim", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetTrim is the legacy version of SetTrimCtx, but uses
// context.Background() as the context.
func (client *Volume3) SetTrim(Channel string, TrimBinaryMilliDb int32) (err error) {
	return client.SetTrimCtx(context.Background(),
		Channel,
		TrimBinaryMilliDb,
	)
}

func (client *Volume3) SetVolumeCtx(
	ctx context.Context,
	Value uint32,
) (err error) {
	// Request structure.
	request := &struct {
		Value string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Value, err = soap.MarshalUi4(Value); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "SetVolume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *Volume3) SetVolume(Value uint32) (err error) {
	return client.SetVolumeCtx(context.Background(),
		Value,
	)
}

func (client *Volume3) SetVolumeOffsetCtx(
	ctx context.Context,
	Channel string,
	VolumeOffsetBinaryMilliDb int32,
) (err error) {
	// Request structure.
	request := &struct {
		Channel                   string
		VolumeOffsetBinaryMilliDb string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	if request.VolumeOffsetBinaryMilliDb, err = soap.MarshalI4(VolumeOffsetBinaryMilliDb); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "SetVolumeOffset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// SetVolumeOffset is the legacy version of SetVolumeOffsetCtx, but uses
// context.Background() as the context.
func (client *Volume3) SetVolumeOffset(Channel string, VolumeOffsetBinaryMilliDb int32) (err error) {
	return client.SetVolumeOffsetCtx(context.Background(),
		Channel,
		VolumeOffsetBinaryMilliDb,
	)
}

func (client *Volume3) TrimCtx(
	ctx context.Context,
	Channel string,
) (TrimBinaryMilliDb int32, err error) {
	// Request structure.
	request := &struct {
		Channel string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		TrimBinaryMilliDb string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "Trim", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if TrimBinaryMilliDb, err = soap.UnmarshalI4(response.TrimBinaryMilliDb); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Trim is the legacy version of TrimCtx, but uses
// context.Background() as the context.
func (client *Volume3) Trim(Channel string) (TrimBinaryMilliDb int32, err error) {
	return client.TrimCtx(context.Background(),
		Channel,
	)
}

func (client *Volume3) UnityGainCtx(
	ctx context.Context,
) (Value bool, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "UnityGain", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalBoolean(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// UnityGain is the legacy version of UnityGainCtx, but uses
// context.Background() as the context.
func (client *Volume3) UnityGain() (Value bool, err error) {
	return client.UnityGainCtx(context.Background())
}

func (client *Volume3) VolumeCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "Volume", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Volume is the legacy version of VolumeCtx, but uses
// context.Background() as the context.
func (client *Volume3) Volume() (Value uint32, err error) {
	return client.VolumeCtx(context.Background())
}

func (client *Volume3) VolumeDecCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "VolumeDec", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// VolumeDec is the legacy version of VolumeDecCtx, but uses
// context.Background() as the context.
func (client *Volume3) VolumeDec() (err error) {
	return client.VolumeDecCtx(context.Background())
}

func (client *Volume3) VolumeIncCtx(
	ctx context.Context,
) (err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := interface{}(nil)

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "VolumeInc", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// VolumeInc is the legacy version of VolumeIncCtx, but uses
// context.Background() as the context.
func (client *Volume3) VolumeInc() (err error) {
	return client.VolumeIncCtx(context.Background())
}

func (client *Volume3) VolumeLimitCtx(
	ctx context.Context,
) (Value uint32, err error) {
	// Request structure.
	request := interface{}(nil)
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		Value string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "VolumeLimit", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if Value, err = soap.UnmarshalUi4(response.Value); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// VolumeLimit is the legacy version of VolumeLimitCtx, but uses
// context.Background() as the context.
func (client *Volume3) VolumeLimit() (Value uint32, err error) {
	return client.VolumeLimitCtx(context.Background())
}

func (client *Volume3) VolumeOffsetCtx(
	ctx context.Context,
	Channel string,
) (VolumeOffsetBinaryMilliDb int32, err error) {
	// Request structure.
	request := &struct {
		Channel string
	}{}
	// BEGIN Marshal arguments into request.

	if request.Channel, err = soap.MarshalString(Channel); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	response := &struct {
		VolumeOffsetBinaryMilliDb string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformActionCtx(ctx, URN_Volume_3, "VolumeOffset", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if VolumeOffsetBinaryMilliDb, err = soap.UnmarshalI4(response.VolumeOffsetBinaryMilliDb); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// VolumeOffset is the legacy version of VolumeOffsetCtx, but uses
// context.Background() as the context.
func (client *Volume3) VolumeOffset(Channel string) (VolumeOffsetBinaryMilliDb int32, err error) {
	return client.VolumeOffsetCtx(context.Background(),
		Channel,
	)
}
