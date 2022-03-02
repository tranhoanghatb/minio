package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/minio/madmin-go"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *tierConfigMgrV1) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tiers":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Tiers")
				return
			}
			if z.Tiers == nil {
				z.Tiers = make(map[string]tierConfigV1, zb0002)
			} else if len(z.Tiers) > 0 {
				for key := range z.Tiers {
					delete(z.Tiers, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 tierConfigV1
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Tiers")
					return
				}
				err = za0002.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Tiers", za0001)
					return
				}
				z.Tiers[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *tierConfigMgrV1) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Tiers":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tiers")
				return
			}
			if z.Tiers == nil {
				z.Tiers = make(map[string]tierConfigV1, zb0002)
			} else if len(z.Tiers) > 0 {
				for key := range z.Tiers {
					delete(z.Tiers, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 tierConfigV1
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tiers")
					return
				}
				bts, err = za0002.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tiers", za0001)
					return
				}
				z.Tiers[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *tierConfigMgrV1) Msgsize() (s int) {
	s = 1 + 6 + msgp.MapHeaderSize
	if z.Tiers != nil {
		for za0001, za0002 := range z.Tiers {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + za0002.Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *tierConfigV1) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "Type":
			{
				var zb0002 int
				zb0002, err = dc.ReadInt()
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = tierTypeV1(zb0002)
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "S3":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "S3")
					return
				}
				z.S3 = nil
			} else {
				if z.S3 == nil {
					z.S3 = new(madmin.TierS3)
				}
				err = z.S3.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "S3")
					return
				}
			}
		case "Azure":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Azure")
					return
				}
				z.Azure = nil
			} else {
				if z.Azure == nil {
					z.Azure = new(madmin.TierAzure)
				}
				err = z.Azure.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Azure")
					return
				}
			}
		case "GCS":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "GCS")
					return
				}
				z.GCS = nil
			} else {
				if z.GCS == nil {
					z.GCS = new(madmin.TierGCS)
				}
				err = z.GCS.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "GCS")
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *tierConfigV1) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Version":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "Type":
			{
				var zb0002 int
				zb0002, bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = tierTypeV1(zb0002)
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "S3":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.S3 = nil
			} else {
				if z.S3 == nil {
					z.S3 = new(madmin.TierS3)
				}
				bts, err = z.S3.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "S3")
					return
				}
			}
		case "Azure":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Azure = nil
			} else {
				if z.Azure == nil {
					z.Azure = new(madmin.TierAzure)
				}
				bts, err = z.Azure.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Azure")
					return
				}
			}
		case "GCS":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.GCS = nil
			} else {
				if z.GCS == nil {
					z.GCS = new(madmin.TierGCS)
				}
				bts, err = z.GCS.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "GCS")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *tierConfigV1) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Version) + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 3
	if z.S3 == nil {
		s += msgp.NilSize
	} else {
		s += z.S3.Msgsize()
	}
	s += 6
	if z.Azure == nil {
		s += msgp.NilSize
	} else {
		s += z.Azure.Msgsize()
	}
	s += 4
	if z.GCS == nil {
		s += msgp.NilSize
	} else {
		s += z.GCS.Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *tierTypeV1) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int
		zb0001, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = tierTypeV1(zb0001)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *tierTypeV1) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int
		zb0001, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = tierTypeV1(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z tierTypeV1) Msgsize() (s int) {
	s = msgp.IntSize
	return
}
