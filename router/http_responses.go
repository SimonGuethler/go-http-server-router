package router

func HttpResponseOK(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 200,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseCreated(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 201,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseAccepted(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 202,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNoContent(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 204,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseMovedPermanently(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 301,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseFound(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 302,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseSeeOther(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 303,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNotModified(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 304,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseTemporaryRedirect(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 307,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponsePermanentRedirect(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 308,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseBadRequest(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 400,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseUnauthorized(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 401,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseForbidden(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 403,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNotFound(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 404,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseMethodNotAllowed(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 405,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNotAcceptable(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 406,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseProxyAuthenticationRequired(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 407,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseRequestTimeout(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 408,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseConflict(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 409,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseGone(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 410,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseLengthRequired(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 411,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponsePreconditionFailed(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 412,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponsePayloadTooLarge(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 413,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseURITooLong(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 414,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseUnsupportedMediaType(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 415,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseRangeNotSatisfiable(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 416,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseExpectationFailed(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 417,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseTeapot(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 418,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseMisdirectedRequest(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 421,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseUnprocessableEntity(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 422,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseLocked(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 423,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseFailedDependency(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 424,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseUpgradeRequired(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 426,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponsePreconditionRequired(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 428,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseTooManyRequests(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 429,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseRequestHeaderFieldsTooLarge(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 431,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseUnavailableForLegalReasons(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 451,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseInternalServerError(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 500,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNotImplemented(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 501,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseBadGateway(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 502,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseServiceUnavailable(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 503,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseGatewayTimeout(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 504,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseHTTPVersionNotSupported(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 505,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseVariantAlsoNegotiates(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 506,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseInsufficientStorage(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 507,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseLoopDetected(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 508,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNotExtended(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 510,
		Headers:    nil,
		Body:       body,
	}
}

func HttpResponseNetworkAuthenticationRequired(body string) HttpResponse {
	return HttpResponse{
		StatusCode: 511,
		Headers:    nil,
		Body:       body,
	}
}
