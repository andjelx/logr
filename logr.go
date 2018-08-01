// Package logr defines abstract interfaces for logging.  Packages can depend on
// these interfaces and callers can implement logging in whatever way is
// appropriate.
//
// This design derives from Dave Cheney's blog:
//     http://dave.cheney.net/2015/11/05/lets-talk-about-logging
//
// This is a BETA grade API.  Until there is a significant 2nd implementation,
// I don't really know how it will change.
<<<<<<< HEAD
<<<<<<< HEAD
=======
//
// The logging specifically makes it non-trivial to use format strings, to encourage
// attaching structured information instead of unstructured format strings.
//
// Usage
//
// Logging is done using a Logger.  Loggers can have name prefixes and tags attached,
// so that all log messages logged with that Logger have some base context associated.
//
// For instance, suppose we're trying to reconcile the state of an object, and we want
// to log that we've made some decision.
//
// With the traditional log package, we might write
//  log.Printf(
//      "decided to set field foo to value %q for object %s/%s",
//       targetValue, object.Namespace, object.Name)
//
// With logr's structured logging, we'd write
//  // elsewhere in the file, set up the logger to log with the prefix of "reconcilers",
//  // and the tag target-type=Foo, for extra context.
//  log := mainLogger.WithName("reconcilers").WithTag("target-type", "Foo")
//
//  // later on...
//  log.Info("setting field foo on object", "value", targetValue, "object", object)
//
// Depending on our logging implementation, we could then make logging decisions based on field values
// (like only logging such events for objects in a certain namespace), or copy the structured
// information into a structured log store.
//
// For logging errors, Logger has a convinience method called Error.  Suppose we wanted to log an
// error while reconciling.  With the traditional log package, we might write
//   log.Errorf("unable to reconcile object %s/%s: %v", object.Namespace, object.Name, err)
//
// With logr, we'd instead write
//   // assuming the above setup for log
//   log.Error(err, "unable to reconcile object", "object", object)
//
// This is mostly identical to
//   log.Info("unable to reconcile object", "error", err, "object", object)
//
// However, it's more convinient, and certain logging libraries may choose to attach additional
// information (such as stack traces) on calls to Error, so it's preferred to use Error to log errors.
//
// Parts of a log line
//
// Each log message from a Logger has four types of context:
// logger name, log verbosity, log message, and key-value pairs.
//
// The Logger name is constists of a series of name "segments" added by successive calls to WithName.
// These name segments may contain anything but periods.  Exactly how these are represented in the
// output is implementation-dependent.  A common format for implementations is to prefix log messages
// with the name segments, separated by periods.
//
// Log verbosity represents how little a log matters.  Level zero, the default, matters most.
// Increasing levels matter less and less.  Try to avoid lots of different verbosity levels,
// and instead provide useful keys, logger names, and log messages instead for users to filter on
// instead.
//
// The log message consists of a constant message attached to the the log line.  This
// should generally be a simple description of what's occuring, and should never be a format string.
//
// Variable information can then be attached using key/value pairs.  Keys are arbitrary strings,
// and values may be any Go object.
>>>>>>> parent of b8c64da... fixup! Structured Logging
=======
>>>>>>> parent of c62468b... Structured Logging
package logr

// TODO: consider structured logging, a la uber-go/zap
// TODO: consider other bits of glog functionality like Flush, InfoDepth, OutputStats

// InfoLogger represents the ability to log non-error messages.
type InfoLogger interface {
	// Info logs a non-error message.  This is behaviorally akin to fmt.Print.
	Info(args ...interface{})

	// Infof logs a formatted non-error message.
	Infof(format string, args ...interface{})

	// Enabled test whether this InfoLogger is enabled.  For example,
	// commandline flags might be used to set the logging verbosity and disable
	// some info logs.
	Enabled() bool
}

// Logger represents the ability to log messages, both errors and not.
type Logger interface {
	// All Loggers implement InfoLogger.  Calling InfoLogger methods directly on
	// a Logger value is equivalent to calling them on a V(0) InfoLogger.  For
	// example, logger.Info() produces the same result as logger.V(0).Info.
	InfoLogger

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> parent of c62468b... Structured Logging
	// Error logs a error message.  This is behaviorally akin to fmt.Print.
	Error(args ...interface{})

	// Errorf logs a formatted error message.
	Errorf(format string, args ...interface{})
<<<<<<< HEAD
=======
	// Error logs an error, with the given message and key/value pairs as context.
	// It functions as a convinience wrapper around Info, and generally behaves
	// equivalently to calling Info with the error attached as the "error" key,
	// but this method should be preferred for logging errors  (see the package
	// documentations for more information).
	//
	// The msg field should be used to add context to any underlying error,
	// while the err field should be used to attach the actual error that
	// triggered this log line, if present.
	Error(err error, msg string, keysAndValues ...interface{})
>>>>>>> parent of b8c64da... fixup! Structured Logging
=======
>>>>>>> parent of c62468b... Structured Logging

	// V returns an InfoLogger value for a specific verbosity level.  A higher
	// verbosity level means a log message is less important.
	V(level int) InfoLogger

<<<<<<< HEAD
<<<<<<< HEAD
	// NewWithPrefix returns a Logger which prefixes all messages.
	NewWithPrefix(prefix string) Logger
=======
	// WithTags adds some key-value pairs of context to a logger.
	// See Info for documentation on how key/value pairs work.
	WithTags(keysAndValues ...interface{}) Logger

	// WithName adds a new suffix to the logger's name.
	// Successive calls with WithName continue to append
	// suffixes to the logger's name.  Name segments should
	// not contain periods, but are otherwise freeform.
	WithName(name string) Logger
>>>>>>> parent of b8c64da... fixup! Structured Logging
=======
	// NewWithPrefix returns a Logger which prefixes all messages.
	NewWithPrefix(prefix string) Logger
>>>>>>> parent of c62468b... Structured Logging
}
