// Code generated by go-swagger; DO NOT EDIT.

package adminoperations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/electronic_order"
	"github.com/transcom/mymove/pkg/gen/adminapi/adminoperations/office"
)

// NewMymoveAPI creates a new Mymove instance
func NewMymoveAPI(spec *loads.Document) *MymoveAPI {
	return &MymoveAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		OfficeCreateOfficeUserHandler: office.CreateOfficeUserHandlerFunc(func(params office.CreateOfficeUserParams) middleware.Responder {
			return middleware.NotImplemented("operation OfficeCreateOfficeUser has not yet been implemented")
		}),
		ElectronicOrderGetElectronicOrdersTotalsHandler: electronic_order.GetElectronicOrdersTotalsHandlerFunc(func(params electronic_order.GetElectronicOrdersTotalsParams) middleware.Responder {
			return middleware.NotImplemented("operation ElectronicOrderGetElectronicOrdersTotals has not yet been implemented")
		}),
		OfficeGetOfficeUserHandler: office.GetOfficeUserHandlerFunc(func(params office.GetOfficeUserParams) middleware.Responder {
			return middleware.NotImplemented("operation OfficeGetOfficeUser has not yet been implemented")
		}),
		OfficeIndexAccessCodesHandler: office.IndexAccessCodesHandlerFunc(func(params office.IndexAccessCodesParams) middleware.Responder {
			return middleware.NotImplemented("operation OfficeIndexAccessCodes has not yet been implemented")
		}),
		ElectronicOrderIndexElectronicOrdersHandler: electronic_order.IndexElectronicOrdersHandlerFunc(func(params electronic_order.IndexElectronicOrdersParams) middleware.Responder {
			return middleware.NotImplemented("operation ElectronicOrderIndexElectronicOrders has not yet been implemented")
		}),
		OfficeIndexOfficeUsersHandler: office.IndexOfficeUsersHandlerFunc(func(params office.IndexOfficeUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation OfficeIndexOfficeUsers has not yet been implemented")
		}),
		OfficeIndexOfficesHandler: office.IndexOfficesHandlerFunc(func(params office.IndexOfficesParams) middleware.Responder {
			return middleware.NotImplemented("operation OfficeIndexOffices has not yet been implemented")
		}),
		OfficeUpdateOfficeUserHandler: office.UpdateOfficeUserHandlerFunc(func(params office.UpdateOfficeUserParams) middleware.Responder {
			return middleware.NotImplemented("operation OfficeUpdateOfficeUser has not yet been implemented")
		}),
	}
}

/*MymoveAPI The API for move.mil admin actions. */
type MymoveAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// OfficeCreateOfficeUserHandler sets the operation handler for the create office user operation
	OfficeCreateOfficeUserHandler office.CreateOfficeUserHandler
	// ElectronicOrderGetElectronicOrdersTotalsHandler sets the operation handler for the get electronic orders totals operation
	ElectronicOrderGetElectronicOrdersTotalsHandler electronic_order.GetElectronicOrdersTotalsHandler
	// OfficeGetOfficeUserHandler sets the operation handler for the get office user operation
	OfficeGetOfficeUserHandler office.GetOfficeUserHandler
	// OfficeIndexAccessCodesHandler sets the operation handler for the index access codes operation
	OfficeIndexAccessCodesHandler office.IndexAccessCodesHandler
	// ElectronicOrderIndexElectronicOrdersHandler sets the operation handler for the index electronic orders operation
	ElectronicOrderIndexElectronicOrdersHandler electronic_order.IndexElectronicOrdersHandler
	// OfficeIndexOfficeUsersHandler sets the operation handler for the index office users operation
	OfficeIndexOfficeUsersHandler office.IndexOfficeUsersHandler
	// OfficeIndexOfficesHandler sets the operation handler for the index offices operation
	OfficeIndexOfficesHandler office.IndexOfficesHandler
	// OfficeUpdateOfficeUserHandler sets the operation handler for the update office user operation
	OfficeUpdateOfficeUserHandler office.UpdateOfficeUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *MymoveAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MymoveAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MymoveAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MymoveAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MymoveAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MymoveAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MymoveAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MymoveAPI
func (o *MymoveAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.OfficeCreateOfficeUserHandler == nil {
		unregistered = append(unregistered, "office.CreateOfficeUserHandler")
	}

	if o.ElectronicOrderGetElectronicOrdersTotalsHandler == nil {
		unregistered = append(unregistered, "electronic_order.GetElectronicOrdersTotalsHandler")
	}

	if o.OfficeGetOfficeUserHandler == nil {
		unregistered = append(unregistered, "office.GetOfficeUserHandler")
	}

	if o.OfficeIndexAccessCodesHandler == nil {
		unregistered = append(unregistered, "office.IndexAccessCodesHandler")
	}

	if o.ElectronicOrderIndexElectronicOrdersHandler == nil {
		unregistered = append(unregistered, "electronic_order.IndexElectronicOrdersHandler")
	}

	if o.OfficeIndexOfficeUsersHandler == nil {
		unregistered = append(unregistered, "office.IndexOfficeUsersHandler")
	}

	if o.OfficeIndexOfficesHandler == nil {
		unregistered = append(unregistered, "office.IndexOfficesHandler")
	}

	if o.OfficeUpdateOfficeUserHandler == nil {
		unregistered = append(unregistered, "office.UpdateOfficeUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MymoveAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MymoveAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *MymoveAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *MymoveAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *MymoveAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MymoveAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the mymove API
func (o *MymoveAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MymoveAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/office_users"] = office.NewCreateOfficeUser(o.context, o.OfficeCreateOfficeUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/electronic_orders/totals"] = electronic_order.NewGetElectronicOrdersTotals(o.context, o.ElectronicOrderGetElectronicOrdersTotalsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/office_users/{officeUserId}"] = office.NewGetOfficeUser(o.context, o.OfficeGetOfficeUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/access_codes"] = office.NewIndexAccessCodes(o.context, o.OfficeIndexAccessCodesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/electronic_orders"] = electronic_order.NewIndexElectronicOrders(o.context, o.ElectronicOrderIndexElectronicOrdersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/office_users"] = office.NewIndexOfficeUsers(o.context, o.OfficeIndexOfficeUsersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/offices"] = office.NewIndexOffices(o.context, o.OfficeIndexOfficesHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/office_users/{officeUserId}"] = office.NewUpdateOfficeUser(o.context, o.OfficeUpdateOfficeUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MymoveAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MymoveAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MymoveAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MymoveAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
