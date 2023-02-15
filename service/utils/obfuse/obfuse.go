package obfuse

import (
	"errors"

	hashids "github.com/speps/go-hashids/v2"
)

type Obfuscator struct {
	Salt      string
	MinLength int
}

var hashID *hashids.HashID

func NewObfuscator(salt string, minLength int) *Obfuscator {
	o := &Obfuscator{
		Salt:      salt,
		MinLength: minLength,
	}
	o.initHashID()

	return o
}

func (o *Obfuscator) initHashID() {
	hd := hashids.NewData()
	hd.Salt = o.Salt
	hd.MinLength = o.MinLength

	var err error
	hashID, err = hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
}

func (o *Obfuscator) Obfuscate(id uint) string {
	hid, err := hashID.Encode([]int{int(id)})
	if err != nil {
		return ""
	}
	return hid
}

func (o *Obfuscator) Deobfuscate(hid string) (uint, error) {
	if hid == "" {
		return 0, errors.New("hid is empty")
	}

	ids, err := hashID.DecodeWithError(hid)
	if err != nil {
		return 0, err
	}
	return uint(ids[0]), nil
}

func (o *Obfuscator) DeobfuscateHids(hids []string) ([]uint, error) {
	ids := make([]uint, len(hids))
	for key, hid := range hids {
		if hid == "" {
			return nil, errors.New("hid is an empty string")
		}
		id, err := o.Deobfuscate(hid)
		if err != nil {
			return nil, err
		}
		ids[key] = id
	}
	return ids, nil
}
