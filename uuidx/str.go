package uuidx

import (
	"github.com/cnych/stardust/encodingx/base64x"
	"github.com/cnych/stardust/encodingx/hexx"
	"github.com/satori/go.uuid"
)

func Hex(o uuid.UUID) string {
	return hexx.EncodeString(o.Bytes())
}

func HexV1() string {
	return Hex(uuid.NewV1())
}

func HexV4() string {
	return Hex(uuid.NewV4())
}

func Base64(o uuid.UUID) string {
	return base64x.EncodeString(o.Bytes())
}

func Base64V1() string {
	return Base64(uuid.NewV1())
}

func Base64V4() string {
	return Base64(uuid.NewV4())
}
