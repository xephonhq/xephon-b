package serialize

import "github.com/xephonhq/xephon-b/pkg/common"

type Serializer interface {
	WriteInt(* common.IntPointWithSeries) ([]byte, error)
	WriteDouble(* common.DoublePointWithSeries) ([]byte, error)
}