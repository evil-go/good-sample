package greet

import "net/http"

type Dao struct {
	DefaultMessage string
	BobMessage     string
	JuliaMessage   string
}

func (sdi Dao) GreetingForName(name string) (string, error) {
	switch name {
	case "Bob":
		return sdi.BobMessage, nil
	case "Julia":
		return sdi.JuliaMessage, nil
	default:
		return sdi.DefaultMessage, nil
	}
}

type Response struct {
	Message string
}

type GreetingFinder interface {
	GreetingForName(name string) (string, error)
}

type Service struct {
	GreetingFinder GreetingFinder
}

func (ssi Service) Greeting(name string) (Response, error) {
	msg, err := ssi.GreetingFinder.GreetingForName(name)
	if err != nil {
		return Response{}, err
	}
	return Response{Message: msg}, nil
}

type Greeter interface {
	Greeting(name string) (Response, error)
}

type Controller struct {
	Greeter Greeter
}

func (mc Controller) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	result, err := mc.Greeter.Greeting(req.URL.Query().Get("name"))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result.Message))
}
