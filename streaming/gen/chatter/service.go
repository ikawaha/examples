// Code generated by goa v3.2.3, DO NOT EDIT.
//
// chatter service
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package chatter

import (
	"context"

	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	"goa.design/goa/v3/security"
)

// The chatter service implements a simple client and server chat.
type Service interface {
	// Creates a valid JWT token for auth to chat.
	Login(context.Context, *LoginPayload) (res string, err error)
	// Echoes the message sent by the client.
	Echoer(context.Context, *EchoerPayload, EchoerServerStream) (err error)
	// Listens to the messages sent by the client.
	Listener(context.Context, *ListenerPayload, ListenerServerStream) (err error)
	// Summarizes the chat messages sent by the client.
	Summary(context.Context, *SummaryPayload, SummaryServerStream) (err error)
	// Subscribe to events sent when new chat messages are added.
	Subscribe(context.Context, *SubscribePayload, SubscribeServerStream) (err error)
	// Returns the chat messages sent to the server.
	// The "view" return value must have one of the following views
	//	- "tiny"
	//	- "default"
	History(context.Context, *HistoryPayload, HistoryServerStream) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "chatter"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"login", "echoer", "listener", "summary", "subscribe", "history"}

// EchoerServerStream is the interface a "echoer" endpoint server stream must
// satisfy.
type EchoerServerStream interface {
	// Send streams instances of "string".
	Send(string) error
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	// Close closes the stream.
	Close() error
}

// EchoerClientStream is the interface a "echoer" endpoint client stream must
// satisfy.
type EchoerClientStream interface {
	// Send streams instances of "string".
	Send(string) error
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	// Close closes the stream.
	Close() error
}

// ListenerServerStream is the interface a "listener" endpoint server stream
// must satisfy.
type ListenerServerStream interface {
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
	// Close closes the stream.
	Close() error
}

// ListenerClientStream is the interface a "listener" endpoint client stream
// must satisfy.
type ListenerClientStream interface {
	// Send streams instances of "string".
	Send(string) error
	// Close closes the stream.
	Close() error
}

// SummaryServerStream is the interface a "summary" endpoint server stream must
// satisfy.
type SummaryServerStream interface {
	// SendAndClose streams instances of "ChatSummaryCollection" and closes the
	// stream.
	SendAndClose(ChatSummaryCollection) error
	// Recv reads instances of "string" from the stream.
	Recv() (string, error)
}

// SummaryClientStream is the interface a "summary" endpoint client stream must
// satisfy.
type SummaryClientStream interface {
	// Send streams instances of "string".
	Send(string) error
	// CloseAndRecv stops sending messages to the stream and reads instances of
	// "ChatSummaryCollection" from the stream.
	CloseAndRecv() (ChatSummaryCollection, error)
}

// SubscribeServerStream is the interface a "subscribe" endpoint server stream
// must satisfy.
type SubscribeServerStream interface {
	// Send streams instances of "Event".
	Send(*Event) error
	// Close closes the stream.
	Close() error
}

// SubscribeClientStream is the interface a "subscribe" endpoint client stream
// must satisfy.
type SubscribeClientStream interface {
	// Recv reads instances of "Event" from the stream.
	Recv() (*Event, error)
}

// HistoryServerStream is the interface a "history" endpoint server stream must
// satisfy.
type HistoryServerStream interface {
	// Send streams instances of "ChatSummary".
	Send(*ChatSummary) error
	// Close closes the stream.
	Close() error
	// SetView sets the view used to render the result before streaming.
	SetView(view string)
}

// HistoryClientStream is the interface a "history" endpoint client stream must
// satisfy.
type HistoryClientStream interface {
	// Recv reads instances of "ChatSummary" from the stream.
	Recv() (*ChatSummary, error)
}

// Credentials used to authenticate to retrieve JWT token
type LoginPayload struct {
	User     string
	Password string
}

// EchoerPayload is the payload type of the chatter service echoer method.
type EchoerPayload struct {
	// JWT used for authentication
	Token string
}

// ListenerPayload is the payload type of the chatter service listener method.
type ListenerPayload struct {
	// JWT used for authentication
	Token string
}

// SummaryPayload is the payload type of the chatter service summary method.
type SummaryPayload struct {
	// JWT used for authentication
	Token string
}

// ChatSummaryCollection is the result type of the chatter service summary
// method.
type ChatSummaryCollection []*ChatSummary

// SubscribePayload is the payload type of the chatter service subscribe method.
type SubscribePayload struct {
	// JWT used for authentication
	Token string
}

// Event is the result type of the chatter service subscribe method.
type Event struct {
	// Message sent to the server
	Message string
	Action  string
	// Time at which the message was added
	AddedAt string
}

// HistoryPayload is the payload type of the chatter service history method.
type HistoryPayload struct {
	// JWT used for authentication
	Token string
	// View to use to render the result
	View *string
}

// ChatSummary is the result type of the chatter service history method.
type ChatSummary struct {
	// Message sent to the server
	Message string
	// Length of the message sent
	Length *int
	// Time at which the message was sent
	SentAt string
}

// Credentials are invalid
type Unauthorized string

type InvalidScopes string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return ""
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}

// NewChatSummaryCollection initializes result type ChatSummaryCollection from
// viewed result type ChatSummaryCollection.
func NewChatSummaryCollection(vres chatterviews.ChatSummaryCollection) ChatSummaryCollection {
	var res ChatSummaryCollection
	switch vres.View {
	case "tiny":
		res = newChatSummaryCollectionTiny(vres.Projected)
	case "default", "":
		res = newChatSummaryCollection(vres.Projected)
	}
	return res
}

// NewViewedChatSummaryCollection initializes viewed result type
// ChatSummaryCollection from result type ChatSummaryCollection using the given
// view.
func NewViewedChatSummaryCollection(res ChatSummaryCollection, view string) chatterviews.ChatSummaryCollection {
	var vres chatterviews.ChatSummaryCollection
	switch view {
	case "tiny":
		p := newChatSummaryCollectionViewTiny(res)
		vres = chatterviews.ChatSummaryCollection{Projected: p, View: "tiny"}
	case "default", "":
		p := newChatSummaryCollectionView(res)
		vres = chatterviews.ChatSummaryCollection{Projected: p, View: "default"}
	}
	return vres
}

// NewChatSummary initializes result type ChatSummary from viewed result type
// ChatSummary.
func NewChatSummary(vres *chatterviews.ChatSummary) *ChatSummary {
	var res *ChatSummary
	switch vres.View {
	case "tiny":
		res = newChatSummaryTiny(vres.Projected)
	case "default", "":
		res = newChatSummary(vres.Projected)
	}
	return res
}

// NewViewedChatSummary initializes viewed result type ChatSummary from result
// type ChatSummary using the given view.
func NewViewedChatSummary(res *ChatSummary, view string) *chatterviews.ChatSummary {
	var vres *chatterviews.ChatSummary
	switch view {
	case "tiny":
		p := newChatSummaryViewTiny(res)
		vres = &chatterviews.ChatSummary{Projected: p, View: "tiny"}
	case "default", "":
		p := newChatSummaryView(res)
		vres = &chatterviews.ChatSummary{Projected: p, View: "default"}
	}
	return vres
}

// newChatSummaryCollectionTiny converts projected type ChatSummaryCollection
// to service type ChatSummaryCollection.
func newChatSummaryCollectionTiny(vres chatterviews.ChatSummaryCollectionView) ChatSummaryCollection {
	res := make(ChatSummaryCollection, len(vres))
	for i, n := range vres {
		res[i] = newChatSummaryTiny(n)
	}
	return res
}

// newChatSummaryCollection converts projected type ChatSummaryCollection to
// service type ChatSummaryCollection.
func newChatSummaryCollection(vres chatterviews.ChatSummaryCollectionView) ChatSummaryCollection {
	res := make(ChatSummaryCollection, len(vres))
	for i, n := range vres {
		res[i] = newChatSummary(n)
	}
	return res
}

// newChatSummaryCollectionViewTiny projects result type ChatSummaryCollection
// to projected type ChatSummaryCollectionView using the "tiny" view.
func newChatSummaryCollectionViewTiny(res ChatSummaryCollection) chatterviews.ChatSummaryCollectionView {
	vres := make(chatterviews.ChatSummaryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newChatSummaryViewTiny(n)
	}
	return vres
}

// newChatSummaryCollectionView projects result type ChatSummaryCollection to
// projected type ChatSummaryCollectionView using the "default" view.
func newChatSummaryCollectionView(res ChatSummaryCollection) chatterviews.ChatSummaryCollectionView {
	vres := make(chatterviews.ChatSummaryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newChatSummaryView(n)
	}
	return vres
}

// newChatSummaryTiny converts projected type ChatSummary to service type
// ChatSummary.
func newChatSummaryTiny(vres *chatterviews.ChatSummaryView) *ChatSummary {
	res := &ChatSummary{}
	if vres.Message != nil {
		res.Message = *vres.Message
	}
	return res
}

// newChatSummary converts projected type ChatSummary to service type
// ChatSummary.
func newChatSummary(vres *chatterviews.ChatSummaryView) *ChatSummary {
	res := &ChatSummary{
		Length: vres.Length,
	}
	if vres.Message != nil {
		res.Message = *vres.Message
	}
	if vres.SentAt != nil {
		res.SentAt = *vres.SentAt
	}
	return res
}

// newChatSummaryViewTiny projects result type ChatSummary to projected type
// ChatSummaryView using the "tiny" view.
func newChatSummaryViewTiny(res *ChatSummary) *chatterviews.ChatSummaryView {
	vres := &chatterviews.ChatSummaryView{
		Message: &res.Message,
	}
	return vres
}

// newChatSummaryView projects result type ChatSummary to projected type
// ChatSummaryView using the "default" view.
func newChatSummaryView(res *ChatSummary) *chatterviews.ChatSummaryView {
	vres := &chatterviews.ChatSummaryView{
		Message: &res.Message,
		Length:  res.Length,
		SentAt:  &res.SentAt,
	}
	return vres
}
