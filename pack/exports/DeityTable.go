// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DeityTable struct {
	_tab flatbuffers.Table
}

func GetRootAsDeityTable(buf []byte, offset flatbuffers.UOffsetT) *DeityTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DeityTable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *DeityTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DeityTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DeityTable) Deities(obj *Deity, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DeityTable) DeitiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DeityTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DeityTableAddDeities(builder *flatbuffers.Builder, Deities flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Deities), 0)
}
func DeityTableStartDeitiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DeityTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Deity struct {
	_tab flatbuffers.Table
}

func GetRootAsDeity(buf []byte, offset flatbuffers.UOffsetT) *Deity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Deity{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Deity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Deity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Deity) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Deity) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Deity) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Deity) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Deity) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Deity) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DeityStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DeityAddId(builder *flatbuffers.Builder, Id uint32) {
	builder.PrependUint32Slot(0, Id, 0)
}
func DeityAddNameEn(builder *flatbuffers.Builder, NameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(NameEn), 0)
}
func DeityAddNameFr(builder *flatbuffers.Builder, NameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(NameFr), 0)
}
func DeityAddNameDe(builder *flatbuffers.Builder, NameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NameDe), 0)
}
func DeityAddNameJa(builder *flatbuffers.Builder, NameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(NameJa), 0)
}
func DeityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
