package docs

// Response model info
// @Description Response information
type Response struct {
	INN         string
	KPP         string
	CompanyName string
	Director    string
} //@name Response

// Request model info
// @Description Request information
type Request struct {
	INN string
} //@name Request
