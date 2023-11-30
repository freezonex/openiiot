package binding

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/bytedance/go-tagexpr/v2"
	"github.com/henrylee2cn/goutil"
	"github.com/henrylee2cn/goutil/tpack"
)

func wrapRespond(ctx *app.RequestContext) http.ResponseWriter {
	r := &bindRespond{
		ctx:    ctx,
		header: make(http.Header),
	}
	return r
}

type bindRespond struct {
	ctx    *app.RequestContext
	header http.Header
}

func (r *bindRespond) Header() http.Header {
	return r.header
}

func (r *bindRespond) Write([]byte) (int, error) {
	panic("implement me")
}

func (r *bindRespond) WriteHeader(int) {
	panic("implement me")
}

var defaultResponding = NewResponding(&TagNames{
	SetHeader: "header",
	SetCookie: "cookie",
})

// TagNames struct tag naming
type TagNames struct {
	// SetHeader use 'setheader' by default when empty
	SetHeader string
	// SetCookie use 'setcookie' by default when empty
	SetCookie string
}

// Responding class for writing header and cookie
type Responding struct {
	responses map[uintptr]*response
	lock      sync.RWMutex
	tagNames  TagNames
	vm        *tagexpr.VM
}

// NewResponding creates *Responding object.
func NewResponding(tagNames *TagNames) *Responding {
	if tagNames == nil {
		tagNames = new(TagNames)
	}
	goutil.InitAndGetString(&tagNames.SetHeader, "setheader")
	goutil.InitAndGetString(&tagNames.SetCookie, "setcookie")
	return &Responding{
		tagNames:  *tagNames,
		vm:        tagexpr.New(),
		responses: make(map[uintptr]*response, 1024),
	}
}

// WriteHeader writes header and cookie from result to w.
func (r *Responding) WriteHeader(w http.ResponseWriter, result interface{}) error {
	if result == nil {
		return nil
	}
	v, err := r.structValueOf(result)
	if err != nil || !v.IsValid() {
		return nil
	}
	resp, err := r.getOrPrepareResponse(v)
	if err != nil {
		return err
	}
	expr, err := r.vm.Run(result)
	if err != nil {
		return err
	}
	for _, c := range resp.cookies {
		if err = c.SetCookie(w, expr); err != nil {
			return err
		}
	}
	for _, h := range resp.headers {
		h.AddHeader(w, expr)
	}
	return nil
}

func (r *Responding) getOrPrepareResponse(value reflect.Value) (*response, error) {
	runtimeTypeID := tpack.From(value).RuntimeTypeID()
	r.lock.RLock()
	resp, ok := r.responses[runtimeTypeID]
	r.lock.RUnlock()
	if ok {
		return resp, nil
	}
	expr, err := r.vm.Run(reflect.New(value.Type()).Elem())
	if err != nil {
		return nil, err
	}
	resp = newResponse()
	expr.RangeFields(func(fh *tagexpr.FieldHandler) bool {
		field := fh.StructField()
		fs := fh.StringSelector()
		err = resp.setHeader(field, fs, r.tagNames.SetHeader)
		if err != nil {
			return false
		}
		err = resp.setCookie(field, fs, r.tagNames.SetCookie)
		if err != nil {
			return false
		}
		return true
	})
	r.lock.Lock()
	r.responses[runtimeTypeID] = resp
	r.lock.Unlock()
	return resp, err
}

func cleanTagValue(tagVal string) string {
	tagVal = strings.TrimSpace(tagVal)
	tagVal = strings.Replace(tagVal, " ", "", -1)
	tagVal = strings.Replace(tagVal, "\t", "", -1)
	return tagVal
}

func splitRequired(tagVal string) (string, bool) {
	name := strings.TrimSuffix(tagVal, ",required")
	return name, name != tagVal
}

func (r *response) setHeader(field reflect.StructField, fs, tagName string) error {
	tagVal, ok := field.Tag.Lookup(tagName)
	if !ok {
		return nil
	}
	if err := checkString(field.Name, field.Type); err != nil {
		return err
	}
	name, required := splitRequired(cleanTagValue(tagVal))
	if name == "" {
		name = field.Name
	}
	r.headers[name] = &header{
		name:          name,
		valueFS:       fs,
		valueRequired: required,
	}
	return nil
}

func (r *response) setCookie(field reflect.StructField, fs, tagName string) error {
	tagVal, ok := field.Tag.Lookup(tagName)
	if !ok {
		return nil
	}
	name, required := splitRequired(cleanTagValue(tagVal))
	if name == "" {
		name = field.Name
	}
	a := strings.SplitN(name, ",", 2)
	name = a[0]
	c, ok := r.cookies[name]
	if !ok {
		c = &cookie{
			name: name,
		}
		r.cookies[name] = c
	}
	var pos string
	if len(a) == 2 {
		pos = a[1]
	}
	switch strings.ToLower(pos) {
	case "path":
		if err := checkString(field.Name, field.Type); err != nil {
			return err
		}
		c.pathFS = fs
		c.pathRequired = required
	case "domain":
		if err := checkString(field.Name, field.Type); err != nil {
			return err
		}
		c.domainFS = fs
		c.domainRequired = required
	case "expires":
		if err := checkString(field.Name, field.Type); err != nil {
			if err = checkTime(field.Name, field.Type); err != nil {
				return err
			}
		}
		c.expiresFS = fs
		c.expiresRequired = required
	case "maxage":
		if err := checkInt(field.Name, field.Type); err != nil {
			return err
		}
		c.maxAgeFS = fs
		c.maxAgeRequired = required
	case "secure":
		if err := checkBool(field.Name, field.Type); err != nil {
			return err
		}
		c.secureFS = fs
		c.secureRequired = required
	case "httponly":
		if err := checkBool(field.Name, field.Type); err != nil {
			return err
		}
		c.httpOnlyFS = fs
		c.httpOnlyRequired = required
	case "samesite":
		if err := checkInt(field.Name, field.Type); err != nil {
			return err
		}
		c.sameSiteFS = fs
		c.sameSiteRequired = required
	case "value":
		fallthrough
	default: // Value
		if err := checkString(field.Name, field.Type); err != nil {
			return err
		}
		c.valueFS = fs
		c.valueRequired = required
	}
	return nil
}

func (r *Responding) structValueOf(value interface{}) (reflect.Value, error) {
	v, ok := value.(reflect.Value)
	if !ok {
		v = reflect.ValueOf(value)
	}
	v = goutil.DereferenceValue(v)
	if v.Kind() != reflect.Struct {
		return v, fmt.Errorf("type %T is not a non-nil struct", value)
	}
	return v, nil
}

type response struct {
	cookies map[string]*cookie // <name,cookie>
	headers map[string]*header // <name,header>
}

func newResponse() *response {
	return &response{
		cookies: make(map[string]*cookie, 32),
		headers: make(map[string]*header, 32),
	}
}

type header struct {
	name          string
	valueFS       string
	valueRequired bool
}

func (h *header) AddHeader(w http.ResponseWriter, tagExpr *tagexpr.TagExpr) error {
	v := getString(tagExpr, h.valueFS)
	if v == "" && h.valueRequired {
		return fmt.Errorf("the header %s missing required value", h.name)
	}
	w.Header().Add(h.name, v)
	return nil
}

type cookie struct {
	name string

	valueFS    string
	pathFS     string
	domainFS   string
	expiresFS  string
	maxAgeFS   string
	secureFS   string
	httpOnlyFS string
	sameSiteFS string

	valueRequired    bool
	pathRequired     bool
	domainRequired   bool
	expiresRequired  bool
	maxAgeRequired   bool
	secureRequired   bool
	httpOnlyRequired bool
	sameSiteRequired bool
}

func (c *cookie) SetCookie(w http.ResponseWriter, tagExpr *tagexpr.TagExpr) error {
	ck, err := c.Cookie(tagExpr)
	if err != nil {
		return err
	}
	http.SetCookie(w, ck)
	return nil
}

func (c *cookie) Cookie(tagExpr *tagexpr.TagExpr) (*http.Cookie, error) {
	ck := &http.Cookie{
		Name: c.Name(),
	}
	var err error
	if ck.Value, err = c.Value(tagExpr); err == nil {
		if ck.Path, err = c.Path(tagExpr); err == nil {
			if ck.Domain, err = c.Domain(tagExpr); err == nil {
				if ck.Expires, err = c.Expires(tagExpr); err == nil {
					if ck.MaxAge, err = c.MaxAge(tagExpr); err == nil {
						if ck.Secure, err = c.Secure(tagExpr); err == nil {
							if ck.HttpOnly, err = c.HttpOnly(tagExpr); err == nil {
								ck.SameSite, err = c.SameSite(tagExpr)
							}
						}
					}
				}
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return ck, nil
}

func (c *cookie) Name() string {
	return c.name
}

func (c *cookie) Value(tagExpr *tagexpr.TagExpr) (string, error) {
	v := url.QueryEscape(getString(tagExpr, c.valueFS))
	if v == "" && c.valueRequired {
		return v, fmt.Errorf("the cookie %s missing required Value field", c.name)
	}
	return v, nil
}

func (c *cookie) Path(tagExpr *tagexpr.TagExpr) (string, error) {
	v := getString(tagExpr, c.pathFS)
	if v == "" {
		if c.pathRequired {
			return v, fmt.Errorf("the cookie %s missing required Path field", c.name)
		}
		v = "/"
	}
	return v, nil
}

func (c *cookie) Domain(tagExpr *tagexpr.TagExpr) (string, error) {
	v := getString(tagExpr, c.domainFS)
	if v == "" && c.domainRequired {
		return v, fmt.Errorf("the cookie %s missing required Domain field", c.name)
	}
	return v, nil
}

func (c *cookie) Expires(tagExpr *tagexpr.TagExpr) (time.Time, error) {
	v, err := getTime(tagExpr, c.expiresFS)
	if err != nil {
		return v, err
	}
	if v == (time.Time{}) && c.expiresRequired {
		return v, fmt.Errorf("the cookie %s missing required Expires field", c.name)
	}
	return v, nil
}

func (c *cookie) MaxAge(tagExpr *tagexpr.TagExpr) (int, error) {
	v := getInt(tagExpr, c.maxAgeFS)
	if v == 0 && c.maxAgeRequired {
		return v, fmt.Errorf("the cookie %s missing required MaxAge field", c.name)
	}
	return v, nil
}

func (c *cookie) Secure(tagExpr *tagexpr.TagExpr) (bool, error) {
	v, valid := getBool(tagExpr, c.secureFS)
	if !valid && c.secureRequired {
		return v, fmt.Errorf("the cookie %s missing required Secure field", c.name)
	}
	return v, nil
}

func (c *cookie) HttpOnly(tagExpr *tagexpr.TagExpr) (bool, error) {
	v, valid := getBool(tagExpr, c.httpOnlyFS)
	if !valid && c.httpOnlyRequired {
		return v, fmt.Errorf("the cookie %s missing required HttpOnly field", c.name)
	}
	return v, nil
}

func (c *cookie) SameSite(tagExpr *tagexpr.TagExpr) (http.SameSite, error) {
	v := http.SameSite(getInt(tagExpr, c.sameSiteFS))
	if v == 0 && c.sameSiteRequired {
		return v, fmt.Errorf("missing required header parameter: %s", c.sameSiteFS)
	}
	return v, nil
}

func getInt(tagExpr *tagexpr.TagExpr, fs string) int {
	v := getElem(tagExpr, fs)
	if !v.IsValid() {
		return 0
	}
	return int(v.Int())
}

func getBool(tagExpr *tagexpr.TagExpr, fs string) (ret, valid bool) {
	v := getElem(tagExpr, fs)
	if !v.IsValid() {
		return false, false
	}
	return v.Bool(), true
}

func getTime(tagExpr *tagexpr.TagExpr, fs string) (time.Time, error) {
	v := getElem(tagExpr, fs)
	if !v.IsValid() {
		return time.Time{}, nil
	}
	if v.Kind() == reflect.String {
		return time.Parse(time.RFC1123, v.String())
	}
	if v.CanInterface() {
		t, ok := v.Interface().(time.Time)
		if ok {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("type %s cannot be converted to time.Time", v.Type().String())
}

func getString(tagExpr *tagexpr.TagExpr, fs string) string {
	v := getElem(tagExpr, fs)
	if !v.IsValid() {
		return ""
	}
	return v.String()
}

func getElem(tagExpr *tagexpr.TagExpr, fs string) reflect.Value {
	if fs == "" {
		return reflect.Value{}
	}
	fh, _ := tagExpr.Field(fs)
	v := fh.Value(false)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func checkString(fieldName string, t reflect.Type) error {
	t = goutil.DereferenceType(t)
	if t.Kind() != reflect.String {
		return fmt.Errorf("header field %s must be string or *string, but got: %s", fieldName, t.String())
	}
	return nil
}

func checkBool(fieldName string, t reflect.Type) error {
	t = goutil.DereferenceType(t)
	if t.Kind() != reflect.Bool {
		return fmt.Errorf("header field %s must be bool or *bool, but got: %s", fieldName, t.String())
	}
	return nil
}

func checkTime(fieldName string, t reflect.Type) error {
	t = goutil.DereferenceType(t)
	if t.String() != "time.Time" {
		return fmt.Errorf("header field %s must be time.Time or *time.Time, but got: %s", fieldName, t.String())
	}
	return nil
}

func checkInt(fieldName string, t reflect.Type) error {
	t = goutil.DereferenceType(t)
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	default:
		return fmt.Errorf("header field %s must be signed integer or signed integer pointer, but got: %s", fieldName, t.String())
	}
	return nil
}
