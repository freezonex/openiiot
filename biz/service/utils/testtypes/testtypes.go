// testtypes is a package for describing test case structs
// using go struct embedding.
package testtypes

import (
	"context"
	"testing"

	"freezonex/openiiot/biz/dal/mysql"
)

// Label provides test struct with a descriptive name.
// This is mainly used for developer readability
// and for test names when using
// testing.T.Run(name, func ...)
type Label struct {
	// Name is the name of the test case.
	Name string
}

// Description provides more information about the test case.
type Description struct {
	// Description is the description of the test.
	Comment string
}

// Setup provides test with a setup function to perform
// some stateful operation(s) before the core logic of
// the test is executed.
type Setup struct {
	// SetupThunk is a thunk that takes in test type T.
	SetupThunk func(t *testing.T)
}

// SetupCtx provides same functionality as Setup but with access
// to context.
type SetupCtx struct {
	// SetupCtxThunk is a thunk that takes a test type T and context.
	SetupCtxThunk func(t *testing.T, ctx context.Context)
}

// SetupDB provides same functionality as Setup
// but with DB operations and context.
type SetupDB struct {
	// SetupDBThunk is a thunk that takes a test type T, context, and db handle.
	SetupDBThunk func(t *testing.T, ctx context.Context, db *mysql.MySQL)
}

// Teardown provides a test with a teardown function to perform
// some stateful operation(s) after the core logic of
// the test is executed.
type Teardown struct {
	// TeardownThunk is a thunk that takes in test type T.
	TeardownThunk func(t *testing.T)
}

// TeardownCtx provides same functionality as Teardown but with access
// to context.
type TeardownCtx struct {
	// TeardownCtxThunk is a thunk that takes in test type T and a context.
	TeardownCtxThunk func(t *testing.T, ctx context.Context)
}

// TeardownDB provides same functionality as Teardown
// but with DB operations and context.
type TeardownDB struct {
	// TeardownDBThunk is a thunk that takes in test type T , context, and db handle
	TeardownDBThunk func(t *testing.T, ctx context.Context, db *mysql.MySQL)
}
