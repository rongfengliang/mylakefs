// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/treeverse/lakefs/api/gen/models"
)

// NewLakefsAPI creates a new Lakefs instance
func NewLakefsAPI(spec *loads.Document) *LakefsAPI {
	return &LakefsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		CommitHandler: CommitHandlerFunc(func(params CommitParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.Commit has not yet been implemented")
		}),
		CreateBranchHandler: CreateBranchHandlerFunc(func(params CreateBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateBranch has not yet been implemented")
		}),
		CreateRepositoryHandler: CreateRepositoryHandlerFunc(func(params CreateRepositoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateRepository has not yet been implemented")
		}),
		DeleteBranchHandler: DeleteBranchHandlerFunc(func(params DeleteBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteBranch has not yet been implemented")
		}),
		DeleteRepositoryHandler: DeleteRepositoryHandlerFunc(func(params DeleteRepositoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteRepository has not yet been implemented")
		}),
		GetBranchHandler: GetBranchHandlerFunc(func(params GetBranchParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBranch has not yet been implemented")
		}),
		GetCommitHandler: GetCommitHandlerFunc(func(params GetCommitParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCommit has not yet been implemented")
		}),
		GetRepositoryHandler: GetRepositoryHandlerFunc(func(params GetRepositoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRepository has not yet been implemented")
		}),
		ListBranchesHandler: ListBranchesHandlerFunc(func(params ListBranchesParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListBranches has not yet been implemented")
		}),
		ListRepositoriesHandler: ListRepositoriesHandlerFunc(func(params ListRepositoriesParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListRepositories has not yet been implemented")
		}), // Applies when the Authorization header is set with the Basic scheme
		BasicAuthAuth: func(user string, pass string) (*models.User, error) {
			return nil, errors.NotImplemented("basic auth  (basic_auth) has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*LakefsAPI lakeFS HTTP API */
type LakefsAPI struct {
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
	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BasicAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	BasicAuthAuth func(string, string) (*models.User, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CommitHandler sets the operation handler for the commit operation
	CommitHandler CommitHandler
	// CreateBranchHandler sets the operation handler for the create branch operation
	CreateBranchHandler CreateBranchHandler
	// CreateRepositoryHandler sets the operation handler for the create repository operation
	CreateRepositoryHandler CreateRepositoryHandler
	// DeleteBranchHandler sets the operation handler for the delete branch operation
	DeleteBranchHandler DeleteBranchHandler
	// DeleteRepositoryHandler sets the operation handler for the delete repository operation
	DeleteRepositoryHandler DeleteRepositoryHandler
	// GetBranchHandler sets the operation handler for the get branch operation
	GetBranchHandler GetBranchHandler
	// GetCommitHandler sets the operation handler for the get commit operation
	GetCommitHandler GetCommitHandler
	// GetRepositoryHandler sets the operation handler for the get repository operation
	GetRepositoryHandler GetRepositoryHandler
	// ListBranchesHandler sets the operation handler for the list branches operation
	ListBranchesHandler ListBranchesHandler
	// ListRepositoriesHandler sets the operation handler for the list repositories operation
	ListRepositoriesHandler ListRepositoriesHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *LakefsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *LakefsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *LakefsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *LakefsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *LakefsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *LakefsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *LakefsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the LakefsAPI
func (o *LakefsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BasicAuthAuth == nil {
		unregistered = append(unregistered, "BasicAuthAuth")
	}

	if o.CommitHandler == nil {
		unregistered = append(unregistered, "Operations.CommitHandler")
	}

	if o.CreateBranchHandler == nil {
		unregistered = append(unregistered, "Operations.CreateBranchHandler")
	}

	if o.CreateRepositoryHandler == nil {
		unregistered = append(unregistered, "Operations.CreateRepositoryHandler")
	}

	if o.DeleteBranchHandler == nil {
		unregistered = append(unregistered, "Operations.DeleteBranchHandler")
	}

	if o.DeleteRepositoryHandler == nil {
		unregistered = append(unregistered, "Operations.DeleteRepositoryHandler")
	}

	if o.GetBranchHandler == nil {
		unregistered = append(unregistered, "Operations.GetBranchHandler")
	}

	if o.GetCommitHandler == nil {
		unregistered = append(unregistered, "Operations.GetCommitHandler")
	}

	if o.GetRepositoryHandler == nil {
		unregistered = append(unregistered, "Operations.GetRepositoryHandler")
	}

	if o.ListBranchesHandler == nil {
		unregistered = append(unregistered, "Operations.ListBranchesHandler")
	}

	if o.ListRepositoriesHandler == nil {
		unregistered = append(unregistered, "Operations.ListRepositoriesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *LakefsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *LakefsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {

		case "basic_auth":
			result[name] = o.BasicAuthenticator(func(username, password string) (interface{}, error) {
				return o.BasicAuthAuth(username, password)
			})

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *LakefsAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *LakefsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
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

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *LakefsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
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
func (o *LakefsAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the lakefs API
func (o *LakefsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *LakefsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repositoryId}/branches/{branchId}/commits"] = NewCommit(o.context, o.CommitHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repositoryId}/branches"] = NewCreateBranch(o.context, o.CreateBranchHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories"] = NewCreateRepository(o.context, o.CreateRepositoryHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/repositories/{repositoryId}/branches/{branchId}"] = NewDeleteBranch(o.context, o.DeleteBranchHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/repositories/{repositoryId}"] = NewDeleteRepository(o.context, o.DeleteRepositoryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches/{branchId}"] = NewGetBranch(o.context, o.GetBranchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/commits/{commitId}"] = NewGetCommit(o.context, o.GetCommitHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}"] = NewGetRepository(o.context, o.GetRepositoryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repositoryId}/branches"] = NewListBranches(o.context, o.ListBranchesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories"] = NewListRepositories(o.context, o.ListRepositoriesHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *LakefsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *LakefsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *LakefsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *LakefsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}