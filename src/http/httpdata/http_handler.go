package httpdata

type HttpHandler interface {
	Method() string                               // Method of HTTP Request, Get/Post/Delete/etc
	URL() string                                  // URL of request
	PrepareRequestData() HttpDataWithValidator    // prepare the data holder for the request
	Handle(*RequestData) (interface{}, HttpError) // handler of the request
}
