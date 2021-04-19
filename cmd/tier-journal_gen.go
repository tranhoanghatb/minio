// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *jentry) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "obj":
			z.ObjName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ObjName")
				return
			}
		case "tier":
			z.TierName, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "TierName")
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
func (z jentry) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "obj"
	err = en.Append(0x82, 0xa3, 0x6f, 0x62, 0x6a)
	if err != nil {
		return
	}
	err = en.WriteString(z.ObjName)
	if err != nil {
		err = msgp.WrapError(err, "ObjName")
		return
	}
	// write "tier"
	err = en.Append(0xa4, 0x74, 0x69, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.TierName)
	if err != nil {
		err = msgp.WrapError(err, "TierName")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z jentry) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "obj"
	o = append(o, 0x82, 0xa3, 0x6f, 0x62, 0x6a)
	o = msgp.AppendString(o, z.ObjName)
	// string "tier"
	o = append(o, 0xa4, 0x74, 0x69, 0x65, 0x72)
	o = msgp.AppendString(o, z.TierName)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *jentry) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "obj":
			z.ObjName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ObjName")
				return
			}
		case "tier":
			z.TierName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TierName")
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
func (z jentry) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.ObjName) + 5 + msgp.StringPrefixSize + len(z.TierName)
	return
}
