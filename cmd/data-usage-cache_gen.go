package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *dataUsageCacheInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "LastUpdate":
			z.LastUpdate, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "NextCycle":
			z.NextCycle, err = dc.ReadUint8()
			if err != nil {
				err = msgp.WrapError(err, "NextCycle")
				return
			}
		case "NextBloomIdx":
			z.NextBloomIdx, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "NextBloomIdx")
				return
			}
		case "BloomFilter":
			z.BloomFilter, err = dc.ReadBytes(z.BloomFilter)
			if err != nil {
				err = msgp.WrapError(err, "BloomFilter")
				return
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

// EncodeMsg implements msgp.Encodable
func (z *dataUsageCacheInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// omitempty: check for empty values
	zb0001Len := uint32(5)
	var zb0001Mask uint8 /* 5 bits */
	if z.BloomFilter == nil {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	err = en.Append(0x80 | uint8(zb0001Len))
	if err != nil {
		return
	}
	if zb0001Len == 0 {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "LastUpdate"
	err = en.Append(0xaa, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LastUpdate)
	if err != nil {
		err = msgp.WrapError(err, "LastUpdate")
		return
	}
	// write "NextCycle"
	err = en.Append(0xa9, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint8(z.NextCycle)
	if err != nil {
		err = msgp.WrapError(err, "NextCycle")
		return
	}
	// write "NextBloomIdx"
	err = en.Append(0xac, 0x4e, 0x65, 0x78, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.NextBloomIdx)
	if err != nil {
		err = msgp.WrapError(err, "NextBloomIdx")
		return
	}
	if (zb0001Mask & 0x10) == 0 { // if not empty
		// write "BloomFilter"
		err = en.Append(0xab, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72)
		if err != nil {
			return
		}
		err = en.WriteBytes(z.BloomFilter)
		if err != nil {
			err = msgp.WrapError(err, "BloomFilter")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *dataUsageCacheInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(5)
	var zb0001Mask uint8 /* 5 bits */
	if z.BloomFilter == nil {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len == 0 {
		return
	}
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "LastUpdate"
	o = append(o, 0xaa, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendTime(o, z.LastUpdate)
	// string "NextCycle"
	o = append(o, 0xa9, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x79, 0x63, 0x6c, 0x65)
	o = msgp.AppendUint8(o, z.NextCycle)
	// string "NextBloomIdx"
	o = append(o, 0xac, 0x4e, 0x65, 0x78, 0x74, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x78)
	o = msgp.AppendUint64(o, z.NextBloomIdx)
	if (zb0001Mask & 0x10) == 0 { // if not empty
		// string "BloomFilter"
		o = append(o, 0xab, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72)
		o = msgp.AppendBytes(o, z.BloomFilter)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *dataUsageCacheInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "LastUpdate":
			z.LastUpdate, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "LastUpdate")
				return
			}
		case "NextCycle":
			z.NextCycle, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NextCycle")
				return
			}
		case "NextBloomIdx":
			z.NextBloomIdx, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NextBloomIdx")
				return
			}
		case "BloomFilter":
			z.BloomFilter, bts, err = msgp.ReadBytesBytes(bts, z.BloomFilter)
			if err != nil {
				err = msgp.WrapError(err, "BloomFilter")
				return
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
func (z *dataUsageCacheInfo) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 11 + msgp.TimeSize + 10 + msgp.Uint8Size + 13 + msgp.Uint64Size + 12 + msgp.BytesPrefixSize + len(z.BloomFilter)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *dataUsageEntry) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	z.Size, err = dc.ReadInt64()
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	z.Objects, err = dc.ReadUint64()
	if err != nil {
		err = msgp.WrapError(err, "Objects")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "ObjSizes")
		return
	}
	if zb0002 != uint32(dataUsageBucketLen) {
		err = msgp.ArrayError{Wanted: uint32(dataUsageBucketLen), Got: zb0002}
		return
	}
	for za0001 := range z.ObjSizes {
		z.ObjSizes[za0001], err = dc.ReadUint64()
		if err != nil {
			err = msgp.WrapError(err, "ObjSizes", za0001)
			return
		}
	}
	err = z.Children.DecodeMsg(dc)
	if err != nil {
		err = msgp.WrapError(err, "Children")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *dataUsageEntry) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 4
	err = en.Append(0x94)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Size)
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	err = en.WriteUint64(z.Objects)
	if err != nil {
		err = msgp.WrapError(err, "Objects")
		return
	}
	err = en.WriteArrayHeader(uint32(dataUsageBucketLen))
	if err != nil {
		err = msgp.WrapError(err, "ObjSizes")
		return
	}
	for za0001 := range z.ObjSizes {
		err = en.WriteUint64(z.ObjSizes[za0001])
		if err != nil {
			err = msgp.WrapError(err, "ObjSizes", za0001)
			return
		}
	}
	err = z.Children.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Children")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *dataUsageEntry) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 4
	o = append(o, 0x94)
	o = msgp.AppendInt64(o, z.Size)
	o = msgp.AppendUint64(o, z.Objects)
	o = msgp.AppendArrayHeader(o, uint32(dataUsageBucketLen))
	for za0001 := range z.ObjSizes {
		o = msgp.AppendUint64(o, z.ObjSizes[za0001])
	}
	o, err = z.Children.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Children")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *dataUsageEntry) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 4 {
		err = msgp.ArrayError{Wanted: 4, Got: zb0001}
		return
	}
	z.Size, bts, err = msgp.ReadInt64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	z.Objects, bts, err = msgp.ReadUint64Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Objects")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "ObjSizes")
		return
	}
	if zb0002 != uint32(dataUsageBucketLen) {
		err = msgp.ArrayError{Wanted: uint32(dataUsageBucketLen), Got: zb0002}
		return
	}
	for za0001 := range z.ObjSizes {
		z.ObjSizes[za0001], bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "ObjSizes", za0001)
			return
		}
	}
	bts, err = z.Children.UnmarshalMsg(bts)
	if err != nil {
		err = msgp.WrapError(err, "Children")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *dataUsageEntry) Msgsize() (s int) {
	s = 1 + msgp.Int64Size + msgp.Uint64Size + msgp.ArrayHeaderSize + (dataUsageBucketLen * (msgp.Uint64Size)) + z.Children.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *dataUsageHash) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint64
		zb0001, err = dc.ReadUint64()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = dataUsageHash(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z dataUsageHash) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint64(uint64(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z dataUsageHash) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint64(o, uint64(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *dataUsageHash) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint64
		zb0001, bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = dataUsageHash(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z dataUsageHash) Msgsize() (s int) {
	s = msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *sizeHistogram) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != uint32(dataUsageBucketLen) {
		err = msgp.ArrayError{Wanted: uint32(dataUsageBucketLen), Got: zb0001}
		return
	}
	for za0001 := range z {
		z[za0001], err = dc.ReadUint64()
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *sizeHistogram) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(dataUsageBucketLen))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for za0001 := range z {
		err = en.WriteUint64(z[za0001])
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *sizeHistogram) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(dataUsageBucketLen))
	for za0001 := range z {
		o = msgp.AppendUint64(o, z[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *sizeHistogram) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != uint32(dataUsageBucketLen) {
		err = msgp.ArrayError{Wanted: uint32(dataUsageBucketLen), Got: zb0001}
		return
	}
	for za0001 := range z {
		z[za0001], bts, err = msgp.ReadUint64Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *sizeHistogram) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (dataUsageBucketLen * (msgp.Uint64Size))
	return
}
