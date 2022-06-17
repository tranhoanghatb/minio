package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *rebalStatus) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint8
		zb0001, err = dc.ReadUint8()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = rebalStatus(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z rebalStatus) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z rebalStatus) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rebalStatus) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint8
		zb0001, bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = rebalStatus(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z rebalStatus) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *rebalanceInfo) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "startTs":
			z.StartTime, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "stopTs":
			z.EndTime, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "EndTime")
				return
			}
		case "status":
			{
				var zb0002 uint8
				zb0002, err = dc.ReadUint8()
				if err != nil {
					err = msgp.WrapError(err, "Status")
					return
				}
				z.Status = rebalStatus(zb0002)
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
func (z rebalanceInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "startTs"
	err = en.Append(0x83, 0xa7, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73)
	if err != nil {
		return
	}
	err = en.WriteTime(z.StartTime)
	if err != nil {
		err = msgp.WrapError(err, "StartTime")
		return
	}
	// write "stopTs"
	err = en.Append(0xa6, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x73)
	if err != nil {
		return
	}
	err = en.WriteTime(z.EndTime)
	if err != nil {
		err = msgp.WrapError(err, "EndTime")
		return
	}
	// write "status"
	err = en.Append(0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint8(uint8(z.Status))
	if err != nil {
		err = msgp.WrapError(err, "Status")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z rebalanceInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "startTs"
	o = append(o, 0x83, 0xa7, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73)
	o = msgp.AppendTime(o, z.StartTime)
	// string "stopTs"
	o = append(o, 0xa6, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x73)
	o = msgp.AppendTime(o, z.EndTime)
	// string "status"
	o = append(o, 0xa6, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendUint8(o, uint8(z.Status))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rebalanceInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "startTs":
			z.StartTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StartTime")
				return
			}
		case "stopTs":
			z.EndTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "EndTime")
				return
			}
		case "status":
			{
				var zb0002 uint8
				zb0002, bts, err = msgp.ReadUint8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Status")
					return
				}
				z.Status = rebalStatus(zb0002)
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
func (z rebalanceInfo) Msgsize() (s int) {
	s = 1 + 8 + msgp.TimeSize + 7 + msgp.TimeSize + 7 + msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *rebalanceMeta) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "stopTs":
			z.StoppedAt, err = dc.ReadTime()
			if err != nil {
				err = msgp.WrapError(err, "StoppedAt")
				return
			}
		case "arn":
			err = dc.ReadExactBytes((z.ARN)[:])
			if err != nil {
				err = msgp.WrapError(err, "ARN")
				return
			}
		case "pf":
			z.PercentFreeGoal, err = dc.ReadFloat64()
			if err != nil {
				err = msgp.WrapError(err, "PercentFreeGoal")
				return
			}
		case "rss":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "PoolStats")
				return
			}
			if cap(z.PoolStats) >= int(zb0002) {
				z.PoolStats = (z.PoolStats)[:zb0002]
			} else {
				z.PoolStats = make([]*rebalanceStats, zb0002)
			}
			for za0002 := range z.PoolStats {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						err = msgp.WrapError(err, "PoolStats", za0002)
						return
					}
					z.PoolStats[za0002] = nil
				} else {
					if z.PoolStats[za0002] == nil {
						z.PoolStats[za0002] = new(rebalanceStats)
					}
					err = z.PoolStats[za0002].DecodeMsg(dc)
					if err != nil {
						err = msgp.WrapError(err, "PoolStats", za0002)
						return
					}
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

// EncodeMsg implements msgp.Encodable
func (z *rebalanceMeta) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "stopTs"
	err = en.Append(0x84, 0xa6, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x73)
	if err != nil {
		return
	}
	err = en.WriteTime(z.StoppedAt)
	if err != nil {
		err = msgp.WrapError(err, "StoppedAt")
		return
	}
	// write "arn"
	err = en.Append(0xa3, 0x61, 0x72, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteBytes((z.ARN)[:])
	if err != nil {
		err = msgp.WrapError(err, "ARN")
		return
	}
	// write "pf"
	err = en.Append(0xa2, 0x70, 0x66)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.PercentFreeGoal)
	if err != nil {
		err = msgp.WrapError(err, "PercentFreeGoal")
		return
	}
	// write "rss"
	err = en.Append(0xa3, 0x72, 0x73, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.PoolStats)))
	if err != nil {
		err = msgp.WrapError(err, "PoolStats")
		return
	}
	for za0002 := range z.PoolStats {
		if z.PoolStats[za0002] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.PoolStats[za0002].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, "PoolStats", za0002)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *rebalanceMeta) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "stopTs"
	o = append(o, 0x84, 0xa6, 0x73, 0x74, 0x6f, 0x70, 0x54, 0x73)
	o = msgp.AppendTime(o, z.StoppedAt)
	// string "arn"
	o = append(o, 0xa3, 0x61, 0x72, 0x6e)
	o = msgp.AppendBytes(o, (z.ARN)[:])
	// string "pf"
	o = append(o, 0xa2, 0x70, 0x66)
	o = msgp.AppendFloat64(o, z.PercentFreeGoal)
	// string "rss"
	o = append(o, 0xa3, 0x72, 0x73, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.PoolStats)))
	for za0002 := range z.PoolStats {
		if z.PoolStats[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.PoolStats[za0002].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, "PoolStats", za0002)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rebalanceMeta) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "stopTs":
			z.StoppedAt, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StoppedAt")
				return
			}
		case "arn":
			bts, err = msgp.ReadExactBytes(bts, (z.ARN)[:])
			if err != nil {
				err = msgp.WrapError(err, "ARN")
				return
			}
		case "pf":
			z.PercentFreeGoal, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PercentFreeGoal")
				return
			}
		case "rss":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PoolStats")
				return
			}
			if cap(z.PoolStats) >= int(zb0002) {
				z.PoolStats = (z.PoolStats)[:zb0002]
			} else {
				z.PoolStats = make([]*rebalanceStats, zb0002)
			}
			for za0002 := range z.PoolStats {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.PoolStats[za0002] = nil
				} else {
					if z.PoolStats[za0002] == nil {
						z.PoolStats[za0002] = new(rebalanceStats)
					}
					bts, err = z.PoolStats[za0002].UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "PoolStats", za0002)
						return
					}
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
func (z *rebalanceMeta) Msgsize() (s int) {
	s = 1 + 7 + msgp.TimeSize + 4 + msgp.ArrayHeaderSize + (16 * (msgp.ByteSize)) + 3 + msgp.Float64Size + 4 + msgp.ArrayHeaderSize
	for za0002 := range z.PoolStats {
		if z.PoolStats[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.PoolStats[za0002].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *rebalanceMetric) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint8
		zb0001, err = dc.ReadUint8()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = rebalanceMetric(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z rebalanceMetric) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint8(uint8(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z rebalanceMetric) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint8(o, uint8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rebalanceMetric) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint8
		zb0001, bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = rebalanceMetric(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z rebalanceMetric) Msgsize() (s int) {
	s = msgp.Uint8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *rebalanceMetrics) DecodeMsg(dc *msgp.Reader) (err error) {
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
func (z rebalanceMetrics) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z rebalanceMetrics) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rebalanceMetrics) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
func (z rebalanceMetrics) Msgsize() (s int) {
	s = 1
	return
}

// DecodeMsg implements msgp.Decodable
func (z *rebalanceStats) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "ifs":
			z.InitFreeSpace, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "InitFreeSpace")
				return
			}
		case "ic":
			z.InitCapacity, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "InitCapacity")
				return
			}
		case "bus":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Buckets")
				return
			}
			if cap(z.Buckets) >= int(zb0002) {
				z.Buckets = (z.Buckets)[:zb0002]
			} else {
				z.Buckets = make([]string, zb0002)
			}
			for za0001 := range z.Buckets {
				z.Buckets[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Buckets", za0001)
					return
				}
			}
		case "rbs":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "RebalancedBuckets")
				return
			}
			if cap(z.RebalancedBuckets) >= int(zb0003) {
				z.RebalancedBuckets = (z.RebalancedBuckets)[:zb0003]
			} else {
				z.RebalancedBuckets = make([]string, zb0003)
			}
			for za0002 := range z.RebalancedBuckets {
				z.RebalancedBuckets[za0002], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "RebalancedBuckets", za0002)
					return
				}
			}
		case "bu":
			z.Bucket, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "ob":
			z.Object, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Object")
				return
			}
		case "no":
			z.NumObjects, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "NumObjects")
				return
			}
		case "nv":
			z.NumVersions, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "NumVersions")
				return
			}
		case "bs":
			z.Bytes, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Bytes")
				return
			}
		case "par":
			z.Participating, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Participating")
				return
			}
		case "inf":
			err = z.Info.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Info")
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
func (z *rebalanceStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "ifs"
	err = en.Append(0x8b, 0xa3, 0x69, 0x66, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.InitFreeSpace)
	if err != nil {
		err = msgp.WrapError(err, "InitFreeSpace")
		return
	}
	// write "ic"
	err = en.Append(0xa2, 0x69, 0x63)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.InitCapacity)
	if err != nil {
		err = msgp.WrapError(err, "InitCapacity")
		return
	}
	// write "bus"
	err = en.Append(0xa3, 0x62, 0x75, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Buckets)))
	if err != nil {
		err = msgp.WrapError(err, "Buckets")
		return
	}
	for za0001 := range z.Buckets {
		err = en.WriteString(z.Buckets[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Buckets", za0001)
			return
		}
	}
	// write "rbs"
	err = en.Append(0xa3, 0x72, 0x62, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.RebalancedBuckets)))
	if err != nil {
		err = msgp.WrapError(err, "RebalancedBuckets")
		return
	}
	for za0002 := range z.RebalancedBuckets {
		err = en.WriteString(z.RebalancedBuckets[za0002])
		if err != nil {
			err = msgp.WrapError(err, "RebalancedBuckets", za0002)
			return
		}
	}
	// write "bu"
	err = en.Append(0xa2, 0x62, 0x75)
	if err != nil {
		return
	}
	err = en.WriteString(z.Bucket)
	if err != nil {
		err = msgp.WrapError(err, "Bucket")
		return
	}
	// write "ob"
	err = en.Append(0xa2, 0x6f, 0x62)
	if err != nil {
		return
	}
	err = en.WriteString(z.Object)
	if err != nil {
		err = msgp.WrapError(err, "Object")
		return
	}
	// write "no"
	err = en.Append(0xa2, 0x6e, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.NumObjects)
	if err != nil {
		err = msgp.WrapError(err, "NumObjects")
		return
	}
	// write "nv"
	err = en.Append(0xa2, 0x6e, 0x76)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.NumVersions)
	if err != nil {
		err = msgp.WrapError(err, "NumVersions")
		return
	}
	// write "bs"
	err = en.Append(0xa2, 0x62, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Bytes)
	if err != nil {
		err = msgp.WrapError(err, "Bytes")
		return
	}
	// write "par"
	err = en.Append(0xa3, 0x70, 0x61, 0x72)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Participating)
	if err != nil {
		err = msgp.WrapError(err, "Participating")
		return
	}
	// write "inf"
	err = en.Append(0xa3, 0x69, 0x6e, 0x66)
	if err != nil {
		return
	}
	err = z.Info.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *rebalanceStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "ifs"
	o = append(o, 0x8b, 0xa3, 0x69, 0x66, 0x73)
	o = msgp.AppendUint64(o, z.InitFreeSpace)
	// string "ic"
	o = append(o, 0xa2, 0x69, 0x63)
	o = msgp.AppendUint64(o, z.InitCapacity)
	// string "bus"
	o = append(o, 0xa3, 0x62, 0x75, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Buckets)))
	for za0001 := range z.Buckets {
		o = msgp.AppendString(o, z.Buckets[za0001])
	}
	// string "rbs"
	o = append(o, 0xa3, 0x72, 0x62, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.RebalancedBuckets)))
	for za0002 := range z.RebalancedBuckets {
		o = msgp.AppendString(o, z.RebalancedBuckets[za0002])
	}
	// string "bu"
	o = append(o, 0xa2, 0x62, 0x75)
	o = msgp.AppendString(o, z.Bucket)
	// string "ob"
	o = append(o, 0xa2, 0x6f, 0x62)
	o = msgp.AppendString(o, z.Object)
	// string "no"
	o = append(o, 0xa2, 0x6e, 0x6f)
	o = msgp.AppendUint64(o, z.NumObjects)
	// string "nv"
	o = append(o, 0xa2, 0x6e, 0x76)
	o = msgp.AppendUint64(o, z.NumVersions)
	// string "bs"
	o = append(o, 0xa2, 0x62, 0x73)
	o = msgp.AppendUint64(o, z.Bytes)
	// string "par"
	o = append(o, 0xa3, 0x70, 0x61, 0x72)
	o = msgp.AppendBool(o, z.Participating)
	// string "inf"
	o = append(o, 0xa3, 0x69, 0x6e, 0x66)
	o, err = z.Info.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Info")
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rebalanceStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ifs":
			z.InitFreeSpace, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "InitFreeSpace")
				return
			}
		case "ic":
			z.InitCapacity, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "InitCapacity")
				return
			}
		case "bus":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Buckets")
				return
			}
			if cap(z.Buckets) >= int(zb0002) {
				z.Buckets = (z.Buckets)[:zb0002]
			} else {
				z.Buckets = make([]string, zb0002)
			}
			for za0001 := range z.Buckets {
				z.Buckets[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Buckets", za0001)
					return
				}
			}
		case "rbs":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "RebalancedBuckets")
				return
			}
			if cap(z.RebalancedBuckets) >= int(zb0003) {
				z.RebalancedBuckets = (z.RebalancedBuckets)[:zb0003]
			} else {
				z.RebalancedBuckets = make([]string, zb0003)
			}
			for za0002 := range z.RebalancedBuckets {
				z.RebalancedBuckets[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RebalancedBuckets", za0002)
					return
				}
			}
		case "bu":
			z.Bucket, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bucket")
				return
			}
		case "ob":
			z.Object, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Object")
				return
			}
		case "no":
			z.NumObjects, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NumObjects")
				return
			}
		case "nv":
			z.NumVersions, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NumVersions")
				return
			}
		case "bs":
			z.Bytes, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Bytes")
				return
			}
		case "par":
			z.Participating, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Participating")
				return
			}
		case "inf":
			bts, err = z.Info.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Info")
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
func (z *rebalanceStats) Msgsize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 3 + msgp.Uint64Size + 4 + msgp.ArrayHeaderSize
	for za0001 := range z.Buckets {
		s += msgp.StringPrefixSize + len(z.Buckets[za0001])
	}
	s += 4 + msgp.ArrayHeaderSize
	for za0002 := range z.RebalancedBuckets {
		s += msgp.StringPrefixSize + len(z.RebalancedBuckets[za0002])
	}
	s += 3 + msgp.StringPrefixSize + len(z.Bucket) + 3 + msgp.StringPrefixSize + len(z.Object) + 3 + msgp.Uint64Size + 3 + msgp.Uint64Size + 3 + msgp.Uint64Size + 4 + msgp.BoolSize + 4 + z.Info.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *rstats) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(rstats, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(rebalanceStats)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z rstats) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				err = msgp.WrapError(err, zb0003)
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z rstats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, zb0003)
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *rstats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(rstats, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(rebalanceStats)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z rstats) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
