package testdata

import "context"

func CtxDefinedTwice(ctx1 context.Context, ctx2 context.Context) { // want "the function has more than one context defined in the parameters."

}
