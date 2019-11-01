package stmutil

import (
	"context"

	"github.com/lukechampine/stm"
)

func ContextDoneVar(ctx context.Context) (*stm.Var, func()) {
	ctx, cancel := context.WithCancel(ctx)
	_var := stm.NewVar(false)
	go func() {
		<-ctx.Done()
		stm.AtomicSet(_var, true)
	}()
	return _var, cancel
}