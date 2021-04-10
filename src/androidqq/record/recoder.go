package record

import "time"

type BotRecord struct {
	InitTime int

	// NoPicSig --> 0
	// SKey --> 1
	KeyMap map[int]*Key
}

func (r *BotRecord) SetKey(id, shelfLife int, key []byte) {
	t := time.Now().Second()
	k := Key{
		CreateTime: t,
		ExpireTime: t + shelfLife,
		Source:     key,
	}
	r.KeyMap[id] = &k
}

func (r *BotRecord) GetKey(id int) *Key {
	value, exits := r.KeyMap[id]
	if exits {
		return value
	}
	return nil
}
