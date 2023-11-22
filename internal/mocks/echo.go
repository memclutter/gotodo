package mocks

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type EchoContext struct {
	mock.Mock
}

func (m *EchoContext) Request() *http.Request {
	args := m.Called()
	return args.Get(0).(*http.Request)
}

func (m *EchoContext) SetRequest(r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetResponse(r *echo.Response) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Response() *echo.Response {
	args := m.Called()
	return args.Get(0).(*echo.Response)
}

func (m *EchoContext) IsTLS() bool {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) IsWebSocket() bool {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Scheme() string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) RealIP() string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Path() string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetPath(p string) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Param(name string) string {
	args := m.Called(name)
	return args.String(0)
}

func (m *EchoContext) ParamNames() []string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetParamNames(names ...string) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) ParamValues() []string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetParamValues(values ...string) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) QueryParam(name string) string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) QueryParams() url.Values {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) QueryString() string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) FormValue(name string) string {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) FormParams() (url.Values, error) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) FormFile(name string) (*multipart.FileHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) MultipartForm() (*multipart.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Cookie(name string) (*http.Cookie, error) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetCookie(cookie *http.Cookie) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Cookies() []*http.Cookie {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Get(key string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Set(key string, val interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Bind(i interface{}) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *EchoContext) Validate(i interface{}) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *EchoContext) Render(code int, name string, data interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) HTML(code int, html string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) HTMLBlob(code int, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) String(code int, s string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) JSON(code int, i interface{}) error {
	args := m.Called(code, i)
	return args.Error(0)
}

func (m *EchoContext) JSONPretty(code int, i interface{}, indent string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) JSONBlob(code int, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) JSONP(code int, callback string, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) JSONPBlob(code int, callback string, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) XML(code int, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) XMLPretty(code int, i interface{}, indent string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) XMLBlob(code int, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Blob(code int, contentType string, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Stream(code int, contentType string, r io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) File(file string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Attachment(file string, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Inline(file string, name string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) NoContent(code int) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Redirect(code int, url string) error {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Error(err error) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Handler() echo.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetHandler(h echo.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Logger() echo.Logger {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) SetLogger(l echo.Logger) {
	//TODO implement me
	panic("implement me")
}

func (m *EchoContext) Echo() *echo.Echo {
	args := m.Called()
	return args.Get(0).(*echo.Echo)
}

func (m *EchoContext) Reset(r *http.Request, w http.ResponseWriter) {
	//TODO implement me
	panic("implement me")
}
