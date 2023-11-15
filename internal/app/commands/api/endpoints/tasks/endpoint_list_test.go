package tasks

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ContextMock struct {
	mock.Mock
}

func (c *ContextMock) Request() *http.Request {
	args := c.Called()
	return args.Get(0).(*http.Request)
}

func (c *ContextMock) SetRequest(r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetResponse(r *echo.Response) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Response() *echo.Response {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) IsTLS() bool {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) IsWebSocket() bool {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Scheme() string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) RealIP() string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Path() string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetPath(p string) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Param(name string) string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) ParamNames() []string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetParamNames(names ...string) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) ParamValues() []string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetParamValues(values ...string) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) QueryParam(name string) string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) QueryParams() url.Values {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) QueryString() string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) FormValue(name string) string {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) FormParams() (url.Values, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) FormFile(name string) (*multipart.FileHeader, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) MultipartForm() (*multipart.Form, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Cookie(name string) (*http.Cookie, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetCookie(cookie *http.Cookie) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Cookies() []*http.Cookie {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Get(key string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Set(key string, val interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Bind(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Validate(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Render(code int, name string, data interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) HTML(code int, html string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) HTMLBlob(code int, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) String(code int, s string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) JSON(code int, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) JSONPretty(code int, i interface{}, indent string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) JSONBlob(code int, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) JSONP(code int, callback string, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) JSONPBlob(code int, callback string, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) XML(code int, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) XMLPretty(code int, i interface{}, indent string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) XMLBlob(code int, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Blob(code int, contentType string, b []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Stream(code int, contentType string, r io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) File(file string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Attachment(file string, name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Inline(file string, name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) NoContent(code int) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Redirect(code int, url string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Error(err error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Handler() echo.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetHandler(h echo.HandlerFunc) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Logger() echo.Logger {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) SetLogger(l echo.Logger) {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Echo() *echo.Echo {
	//TODO implement me
	panic("implement me")
}

func (c *ContextMock) Reset(r *http.Request, w http.ResponseWriter) {
	//TODO implement me
	panic("implement me")
}

func TestEndpoint_List(t *testing.T) {
	endpoint := NewEndpoint()
	c := &ContextMock{}
	require.NoError(t, endpoint.List(c))
}
