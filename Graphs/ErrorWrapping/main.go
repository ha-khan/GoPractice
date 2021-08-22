package main

import (
	"errors"
	"fmt"
	"net/url"
)

/*
	type QueryError is meant to model an error wrapper defined in our apps layer

	Any errors returned by some "third party"l pkg/api would be wrapped in Err

	Effectively creating a Linked List of error nodes

	https://go.dev/blog/go1.13-errors

*/
type QueryError struct {
	Query string
	Err   error
}

// Implements the built-in error interface
func (e *QueryError) Error() string {

	// This combines state information recorded in QueryError struct with the embedded error
	// or next error list node
	return e.Query + ": " + e.Err.Error()

}

// Errors that want to create a chain of errors should implement the Unwrap method,
// so that they may be used in errors.As to traverse the errors LinkedList and get a specific errors node,
// the concrete type that implements the interface
func (e *QueryError) Unwrap() error {
	return e.Err
}

// Compares this QueryError's err, needed for errors.Is similar linked list traversal to errors.As
// except this one looks at each error nodes Error() value, rather than the concrete type that implements
// the errors interface
func (e *QueryError) Is(err error) bool {
	return e.Error() == err.Error()
}

func DoQuery(dns string) (err error) {

	doGet := func(dns string) error {
		return &url.Error{Op: "GET", URL: "google.com", Err: fmt.Errorf("unable to resolve host")}
	}

	err = doGet(dns)
	if err != nil {

		// Wrap "lower level" error and what ever state information that is associated with it to
		// another error generated at this "this" DoQuery layer
		// QueryError contains another error, url.Error, all associated state information is kept do to this "link list"
		err = &QueryError{Query: "Error with http driver", Err: err}
	}
	return
}

func main() {

	var err error

	err = DoQuery("google.com")
	if err != nil {
		// At this moment we don't know the extent of the error wrapping LL that has been constructed
		// thus far, we can traverse it to see if a specific errors node is some concrete type that
		// gives information on the "root" of the error, mainly by going to the tail node
		var urlErr = &url.Error{}
		if errors.As(err, &urlErr) {
			// Can dump more state information close to root of error
			fmt.Println(fmt.Sprintf("%+v", *urlErr))
		}

		var UnableToFindHost = errors.New("Error with http driver: GET \"google.com\": unable to resolve host")
		//var UnableToFindHost = errors.New("GET \"google.com\": unable to resolve host")
		if errors.Is(err, UnableToFindHost) {
			fmt.Println("FOUND")
		}
	}
}
