package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	TIEMOUT    = 60
	METHOD_GET = "GET"
	EXIT_CODE  = -1
)

var client *HttpClient

type HttpClient struct {
	Client *http.Client
}

//创建http client
func NewHttpClient(timeout int) *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

//初始化
func InitHttp() error {
	client = NewHttpClient(TIEMOUT)
	return nil
}

//定义响应体
type HttpReq struct {
	ErrNo  int64 `json:"errno"`
	Errmsg int64 `json:"errmsg"`
	Data   int64 `json:"data"`
}

//Get请求

func (s *HttpClient) GetInformation(params map[string]string) (interface{}, error) {

	var result interface{}
	//此处可用配置库解析变量
	url := fmt.Sprintf("http://xxx.xxx.xxx.xxx/api/vi/allinfo")

	//创建http request请求方法,执行请求方法、url等
	//func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
	req, err := http.NewRequest(METHOD_GET, url, nil)
	if err != nil {
		return nil, err
	}

	//Reuqest 结构体信息
	//type Request struct {
	// Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
	//Method string
	// URL在服务端表示被请求的URI，在客户端表示要访问的URL。
	//
	// 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
	// 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串。
	// （参见RFC 2616, Section 5.1.2）
	//
	// 在客户端，URL的Host字段指定了要连接的服务器，
	// 而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值。
	//URL *url.URL
	// 接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
	//Proto      string // "HTTP/1.0"
	//ProtoMajor int    // 1
	//ProtoMinor int    // 0
	// Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	Connection: keep-alive
	// 则：
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Connection": {"keep-alive"},
	//	}
	// HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
	// 在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。
	//Header Header
	// Body是请求的主体。
	//
	// 在客户端，如果Body是nil表示该请求没有主体买入GET请求。
	// Client的Transport字段会负责调用Body的Close方法。
	//
	// 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF。
	// Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段。
	//Body io.ReadCloser
	// ContentLength记录相关内容的长度。
	// 如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据。
	// 在客户端，如果Body非nil而该字段为0，表示不知道Body的长度。
	//ContentLength int64
	// TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
	// 本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码。
	//TransferEncoding []string
	// Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。
	//Close bool
	// 在服务端，Host指定URL会在其上寻找资源的主机。
	// 根据RFC 2616，该值可以是Host头的值，或者URL自身提供的主机名。
	// Host的格式可以是"host:port"。
	//
	// 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
	// 如过该字段为""，Request.Write方法会使用URL字段的Host。
	//Host string
	// Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据。
	// 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
	//Form url.Values
	// PostForm是解析好的POST或PUT的表单数据。
	// 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
	//PostForm url.Values
	// MultipartForm是解析好的多部件表单，包括上传的文件。
	// 本字段只有在调用ParseMultipartForm后才有效。
	// 在客户端，会忽略请求中的本字段而使用Body替代。
	//MultipartForm *multipart.Form
	// Trailer指定了会在请求主体之后发送的额外的头域。
	//
	// 在服务端，Trailer字段必须初始化为只有trailer键，所有键都对应nil值。
	// （客户端会声明哪些trailer会发送）
	// 在处理器从Body读取时，不能使用本字段。
	// 在从Body的读取返回EOF后，Trailer字段会被更新完毕并包含非nil的值。
	// （如果客户端发送了这些键值对），此时才可以访问本字段。
	//
	// 在客户端，Trail必须初始化为一个包含将要发送的键值对的映射。（值可以是nil或其终值）
	// ContentLength字段必须是0或-1，以启用"chunked"传输编码发送请求。
	// 在开始发送请求后，Trailer可以在读取请求主体期间被修改，
	// 一旦请求主体返回EOF，调用者就不可再修改Trailer。
	//
	// 很少有HTTP客户端、服务端或代理支持HTTP trailer。
	//Trailer Header
	// RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志。
	// 本字段不是ReadRequest函数填写的，也没有定义格式。
	// 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址。
	// 客户端会忽略请求中的RemoteAddr字段。
	//RemoteAddr string
	// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
	// （参见RFC 2616, Section 5.1）
	// 一般应使用URI字段，在客户端设置请求的本字段会导致错误。
	//RequestURI string
	// TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
	// 本字段不是ReadRequest函数填写的。
	// 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
	// 客户端会忽略请求中的TLS字段。
	//TLS *tls.ConnectionState
	//获取请求参数
	//func (u *URL) Query() Values
	//Query方法解析RawQuery字段并返回其表示的Values类型键值对。

	q := req.URL.Query()
	for k, v := range params {
		//Add将value添加到key关联的值集里原有的值的后面。
		//type Values map[string][]string
		//Values将建映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，Values的键是大小写敏感的。
		//func (v Values) Add(key, value string)
		q.Add(k, v)
	}

	//type URL struct {
	//Scheme   string
	//Opaque   string    // 编码后的不透明数据
	//User     *Userinfo // 用户名和密码信息
	//Host     string    // host或host:port
	//Path     string
	//RawQuery string // 编码后的查询字符串，没有'?'
	//Fragment string // 引用的片段（文档位置），没有'#'
	//}

	//参数解析并排序
	//Encode方法将v编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	//func (v Values) Encode() string
	req.URL.RawQuery = q.Encode()

	req.Close = true

	//处理请求，返回响应体
	//func (c *Client) Do(req *Request) (resp *Response, err error)
	//Do方法发送请求，返回HTTP回复。它会遵守客户端c设置的策略（如重定向、cookie、认证）。

	//如果客户端的策略（如重定向）返回错误或存在HTTP协议错误时，本方法将返回该错误；如果回应的状态码不是2xx，本方法并不会返回错误。

	//如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它。如果返回值resp的主体未关闭，c下层的RoundTripper接口（一般为Transport类型）可能无法重用resp主体下层保持的TCP连接去执行之后的请求。

	//请求的主体，如果非nil，会在执行后被c.Transport关闭，即使出现错误。
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	//请求完成后延迟关闭连接
	defer resp.Body.Close()

	//读取响应体信息
	//func ReadAll(r io.Reader) ([]byte, error)
	//ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。成功的调用返回的err为nil而非EOF。因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误。
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var httpreq HttpReq

	//[]vyte转为struct
	err = json.Unmarshal(data, &httpreq)
	if err != nil {
		return nil, err
	}

	result = httpreq.Data

	return result, nil
}

//main函数
func main() {

	err := InitHttp()
	if err != nil {
		fmt.Printf("init http failed:%v\n", err)
		os.Exit(EXIT_CODE)
	}
	var params map[string]string
	params["params1"] = "xxxx"
	params["params2"] = "xxxx"
	params["params3"] = "xxxx"

	data, err := client.GetInformation(params)
	if err != nil {
		fmt.Printf("get info failed:%v\n", err)
		return
	}

	fmt.Println(data)
}
