package testdata

import "context"

func CtxWithBadName(myContext context.Context, a int, b string) { // want "the function has a context that is not named 'ctx'."

}
