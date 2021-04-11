package record

const (
	NoPicSig int = iota
	SKey
	TLV174
	TLV402
	TLV403
)

type Key struct {
	Source     []byte
	CreateTime int
	ExpireTime int
}
