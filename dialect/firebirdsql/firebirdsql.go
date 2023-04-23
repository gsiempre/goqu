package firebirdsql

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func DialectOptions() *goqu.SQLDialectOptions {
	opts := goqu.DefaultDialectOptions()

	opts.TimeFormat = time.DateTime
	opts.BooleanOperatorLookup = map[exp.BooleanOperation][]byte{
		exp.EqOp:             []byte("="),
		exp.NeqOp:            []byte("<>"),
		exp.GtOp:             []byte(">"),
		exp.GteOp:            []byte(">="),
		exp.LtOp:             []byte("<"),
		exp.LteOp:            []byte("<="),
		exp.InOp:             []byte("IN"),
		exp.NotInOp:          []byte("NOT IN"),
		exp.IsOp:             []byte("IS"),
		exp.IsNotOp:          []byte("IS NOT"),
		exp.LikeOp:           []byte("LIKE"),
		exp.NotLikeOp:        []byte("NOT LIKE"),
		exp.ILikeOp:          []byte("LIKE"),
		exp.NotILikeOp:       []byte("NOT LIKE"),
		exp.RegexpLikeOp:     []byte("SIMILAR TO"),
		exp.RegexpNotLikeOp:  []byte("NOT SIMILAR TO"),
		exp.RegexpILikeOp:    []byte("SIMILAR TO"),
		exp.RegexpNotILikeOp: []byte("NOT SIMILAR TO"),
	}

	return opts
}

func init() {
	goqu.RegisterDialect("firebirdsql", DialectOptions())
}
