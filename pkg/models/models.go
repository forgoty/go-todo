package models

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/forgoty/go-todo/pkg/web"
)

// Wrapper around echo context
type ReqContext struct {
	ctx web.Context
}

func NewContext(ctx web.Context) *ReqContext {
	return &ReqContext{
		ctx: ctx,
	}
}

func (c *ReqContext) Request() *http.Request {
	return c.ctx.Request()
}

func (c *ReqContext) SetRequest(r *http.Request) {
	c.ctx.SetRequest(r)
}

func (c *ReqContext) Response() *web.Response {
	return c.ctx.Response()
}

func (c *ReqContext) SetResponse(r *web.Response) {
	c.ctx.SetResponse(r)
}

func (c *ReqContext) IsTLS() bool {
	return c.ctx.IsTLS()
}

func (c *ReqContext) IsWebSocket() bool {
	return c.ctx.IsWebSocket()
}

func (c *ReqContext) Scheme() string {
	return c.ctx.Scheme()
}

func (c *ReqContext) RealIP() string {
	return c.ctx.RealIP()
}

func (c *ReqContext) Path() string {
	return c.ctx.Path()
}

func (c *ReqContext) SetPath(p string) {
	c.ctx.SetPath(p)
}

func (c *ReqContext) Param(name string) string {
	return c.ctx.Param(name)
}

func (c *ReqContext) ParamNames() []string {
	return c.ctx.ParamNames()
}

func (c *ReqContext) SetParamNames(names ...string) {
	c.ctx.SetParamNames(names...)
}

func (c *ReqContext) ParamValues() []string {
	return c.ctx.ParamValues()
}

func (c *ReqContext) SetParamValues(values ...string) {
	c.ctx.SetParamValues(values...)
}

func (c *ReqContext) QueryParam(name string) string {
	return c.ctx.QueryParam(name)
}

func (c *ReqContext) QueryParams() url.Values {
	return c.ctx.QueryParams()
}

func (c *ReqContext) QueryString() string {
	return c.ctx.QueryString()
}

func (c *ReqContext) FormValue(name string) string {
	return c.ctx.FormValue(name)
}

func (c *ReqContext) FormParams() (url.Values, error) {
	return c.ctx.FormParams()
}

func (c *ReqContext) FormFile(name string) (*multipart.FileHeader, error) {
	return c.ctx.FormFile(name)
}

func (c *ReqContext) MultipartForm() (*multipart.Form, error) {
	return c.ctx.MultipartForm()
}

func (c *ReqContext) Cookie(name string) (*http.Cookie, error) {
	return c.ctx.Cookie(name)
}

func (c *ReqContext) SetCookie(cookie *http.Cookie) {
	c.ctx.SetCookie(cookie)
}

func (c *ReqContext) Cookies() []*http.Cookie {
	return c.ctx.Cookies()
}

func (c *ReqContext) Get(key string) interface{} {
	return c.ctx.Get(key)
}

func (c *ReqContext) Set(key string, val interface{}) {
	c.ctx.Set(key, val)
}

func (c *ReqContext) Bind(i interface{}) error {
	return c.ctx.Bind(i)
}

func (c *ReqContext) Validate(i interface{}) error {
	return c.ctx.Validate(i)
}

func (c *ReqContext) Render(code int, name string, data interface{}) error {
	return c.ctx.Render(code, name, data)
}

func (c *ReqContext) HTML(code int, html string) error {
	return c.ctx.HTML(code, html)
}

func (c *ReqContext) HTMLBlob(code int, b []byte) error {
	return c.ctx.HTMLBlob(code, b)
}

func (c *ReqContext) String(code int, s string) error {
	return c.ctx.String(code, s)
}

func (c *ReqContext) JSON(code int, i interface{}) error {
	return c.ctx.JSON(code, i)
}

func (c *ReqContext) JSONPretty(code int, i interface{}, indent string) error {
	return c.ctx.JSONPretty(code, i, indent)
}

func (c *ReqContext) JSONBlob(code int, b []byte) error {
	return c.ctx.JSONBlob(code, b)
}

func (c *ReqContext) JSONP(code int, callback string, i interface{}) error {
	return c.ctx.JSONP(code, callback, i)
}

func (c *ReqContext) JSONPBlob(code int, callback string, b []byte) error {
	return c.ctx.JSONPBlob(code, callback, b)
}
func (c *ReqContext) XML(code int, i interface{}) error {
	return c.ctx.XML(code, i)
}

func (c *ReqContext) XMLPretty(code int, i interface{}, indent string) error {
	return c.ctx.XMLPretty(code, i, indent)
}

func (c *ReqContext) XMLBlob(code int, b []byte) error {
	return c.ctx.XMLBlob(code, b)
}

func (c *ReqContext) Blob(code int, contentType string, b []byte) error {
	return c.ctx.Blob(code, contentType, b)
}

func (c *ReqContext) Stream(code int, contentType string, r io.Reader) error {
	return c.ctx.Stream(code, contentType, r)
}

func (c *ReqContext) File(file string) error {
	return c.ctx.File(file)
}

func (c *ReqContext) Attachment(file, name string) error {
	return c.ctx.Attachment(file, name)
}

func (c *ReqContext) Inline(file, name string) error {
	return c.ctx.Inline(file, name)
}

func (c *ReqContext) NoContent(code int) error {
	return c.ctx.NoContent(code)
}

func (c *ReqContext) Redirect(code int, url string) error {
	return c.ctx.Redirect(code, url)
}

func (c *ReqContext) Error(err error) {
	c.ctx.Error(err)
}

func (c *ReqContext) Echo() *web.Handler {
	return c.ctx.Echo()
}

func (c *ReqContext) Handler() web.HandlerFunc {
	return c.ctx.Handler()
}

func (c *ReqContext) SetHandler(h web.HandlerFunc) {
	c.ctx.SetHandler(h)
}

func (c *ReqContext) Logger() web.Logger {
	return c.ctx.Logger()
}

func (c *ReqContext) SetLogger(l web.Logger) {
}

func (c *ReqContext) Reset(r *http.Request, w http.ResponseWriter) {
	c.ctx.Reset(r, w)
}
