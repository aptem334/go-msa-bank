// Package restapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.2 DO NOT EDIT.
package restapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add a new account to the store
	// (POST /account)
	AddAccount(w http.ResponseWriter, r *http.Request)
	// Update an existing account
	// (PUT /account)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	// Delete account from the store
	// (DELETE /account/{id})
	DeleteAccount(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)
	// Get account from the store
	// (GET /account/{id})
	GetAccount(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)
	// update balance for account from the store
	// (POST /account/{id}/update-balance)
	UpdateAccountBalance(w http.ResponseWriter, r *http.Request, id openapi_types.UUID)
	// Get all accounts from the store
	// (GET /accounts)
	GetAccounts(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// AddAccount operation middleware
func (siw *ServerInterfaceWrapper) AddAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddAccount(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateAccount operation middleware
func (siw *ServerInterfaceWrapper) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateAccount(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteAccount operation middleware
func (siw *ServerInterfaceWrapper) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteAccount(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAccount operation middleware
func (siw *ServerInterfaceWrapper) GetAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAccount(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateAccountBalance operation middleware
func (siw *ServerInterfaceWrapper) UpdateAccountBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", mux.Vars(r)["id"], &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateAccountBalance(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetAccounts operation middleware
func (siw *ServerInterfaceWrapper) GetAccounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAccounts(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/account", wrapper.AddAccount).Methods("POST")

	r.HandleFunc(options.BaseURL+"/account", wrapper.UpdateAccount).Methods("PUT")

	r.HandleFunc(options.BaseURL+"/account/{id}", wrapper.DeleteAccount).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/account/{id}", wrapper.GetAccount).Methods("GET")

	r.HandleFunc(options.BaseURL+"/account/{id}/update-balance", wrapper.UpdateAccountBalance).Methods("POST")

	r.HandleFunc(options.BaseURL+"/accounts", wrapper.GetAccounts).Methods("GET")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xXy24jtxL9FYL3LvshO5ONVtGMJwMBgRPESTaBFyWyJHHSTTJktWTB0L8Hxe6W1XrA",
	"csYTB1mJaJL1OOdUFfUolau9s2gpyvGjjGqJNaTlRCnXWOKlD85jIINpA9qN26aeYeAPtPEoxzJSMHYh",
	"t5mEur/Zbdn27DaTqjJoaap5c+5CDSTHsmmMltmxHbT6BghP+jCXmYgEgc4Y2WYy4J+NCajl+PeDvHZZ",
	"7MW8b+4puvudWzf7jIrY7Ycl2AV2EL6HCqzCYyTZO0ZCPTkH2EGIhxeOXfMNfCAMFqobp5IbjVEF48k4",
	"K8fye2O1cA2J2gUUMOPl3RoWi5R0Eyo5lksiPy7L2H4ujEuI27ljc8pZApWixRoMnwdvCKH+bnhh6PeX",
	"pYnCRAEiQu0rFBHDCoOYuyA6oGIhM1kZhTYmtCzUnNvEg1qiuC5GhwGu1+sC0m7hwqLsrsbyh+mHj7d3",
	"H/PrYlQsqa44GsJQxx/ndxhWhsk4TrJMR0qWkaGKj3S4iDpCPgP7R95pJFdRZnKFIba5XRWjYsROnEcL",
	"3six/KYYFVcykx5omVgoYa+gXEy/LAZgfLgg5ETrvuiynur3Tm960LG9DN5XRqVr5efI/vu65dX/A87l",
	"WP6vfCrssqvqsreeVDJk5yck0WpI0BJIWEQdBTkxQwFao+Y1LVFEcgHlviwpNJh0Gr1j9DmK69HoWHmx",
	"UQpjnDeV2CXOoL0bfXt8eGpXUBktjPUNpUKITV1D2LRACRAW16LD9DA4gkXcK+mcKbnP5EOunMYF2rxD",
	"N585vck7mfGaw2F/R9z86jUQvhk93dYbUjQ6T9H0RsSGs0bdnn13fLZPwDoSc9dYfZb439hmcizwQWH7",
	"ech/S4YAPmEiGbvohfBl3G+zXZWWj0Zv2+AqbKfHMMyb9H3P71Av7faTXjwEqJEwcGyPr8Gu4ZvcXWTW",
	"N8o0AoesZ3vae2ZYbu//0SIeAijmwdXP1/A2kws8UZ6fkP4TWH/tPvJq7H1Ceil1h9VVNqmM89ne8+jk",
	"XBz03v4xdSHNU/2vYPr1h8XJF+YJxn9uffco8GtLpatitkPyokHxdtI8MU5u+zGya8B/R8StAHsgEjYv",
	"0/SLx0oC9JkOFuUXEmAI63gxEzvFQgiw+fpNo6p6kOPFnYOtpH8LbakP/gOceJvzW351Jbn0OoPn2sPE",
	"m6dC33e6vd/+FQAA//+M0SWfHA8AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}