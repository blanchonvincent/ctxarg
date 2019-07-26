package testdata

import "context"

func CtxAsSecondArg(a int, ctx context.Context, b string) { // want "the function has a context that is not the first argument"

}
