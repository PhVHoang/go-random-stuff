/*
A Context caries a deadline, cacelation signal and request-scoped values
across API boundary.
Its methods are safe for simultaneous use by miltiple goroutines
*/

type Context interface {
  // Done returns a channle that is closed when this Context
  // is canceled or times out.
  Done() <-chan struct{}
  // Err indicates why this context was canceled, after the Done
  // channel is closed.
  Err() error
  // Deadline returns the time when this Context will be canceled,
  // if any.
  Deadline() (deadline time.Time, ok bool)

  // Value returns the value associated with key or nil if none.
  Value(key interface{}) interface{}
}

/*
Derived contexts

The context package provides functions to derive new Context values from existing ones.
These values form a tree: when a Context is canceled, all Contexts derived from it are also
canceled.

Background is the root of any Context tree; it never canceled
*/

func Background() Context
