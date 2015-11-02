// Code generated by protoc-gen-go.
// source: flow/flow.proto
// DO NOT EDIT!

/*
Package flow is a generated protocol buffer package.

It is generated from these files:
	flow/flow.proto

It has these top-level messages:
	FlowMessage
*/
package flow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FlowMessage struct {
	Uuid             *string            `protobuf:"bytes,1,req,name=Uuid" json:"Uuid,omitempty"`
	Host             *string            `protobuf:"bytes,2,req,name=Host" json:"Host,omitempty"`
	EtherSrc         *string            `protobuf:"bytes,3,req,name=EtherSrc" json:"EtherSrc,omitempty"`
	EtherDst         *string            `protobuf:"bytes,4,req,name=EtherDst" json:"EtherDst,omitempty"`
	EtherType        *string            `protobuf:"bytes,5,req,name=EtherType" json:"EtherType,omitempty"`
	Ipv4Src          *string            `protobuf:"bytes,6,req,name=Ipv4Src" json:"Ipv4Src,omitempty"`
	Ipv4Dst          *string            `protobuf:"bytes,7,req,name=Ipv4Dst" json:"Ipv4Dst,omitempty"`
	Path             *string            `protobuf:"bytes,8,req,name=Path" json:"Path,omitempty"`
	PortSrc          *uint32            `protobuf:"varint,9,req,name=PortSrc" json:"PortSrc,omitempty"`
	PortDst          *uint32            `protobuf:"varint,10,req,name=PortDst" json:"PortDst,omitempty"`
	Id               *uint64            `protobuf:"varint,11,req,name=Id" json:"Id,omitempty"`
	Timestamp        *uint64            `protobuf:"varint,12,req,name=Timestamp" json:"Timestamp,omitempty"`
	Attributes       *FlowMessage_Attrs `protobuf:"bytes,13,opt,name=Attributes" json:"Attributes,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *FlowMessage) Reset()         { *m = FlowMessage{} }
func (m *FlowMessage) String() string { return proto.CompactTextString(m) }
func (*FlowMessage) ProtoMessage()    {}

func (m *FlowMessage) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *FlowMessage) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *FlowMessage) GetEtherSrc() string {
	if m != nil && m.EtherSrc != nil {
		return *m.EtherSrc
	}
	return ""
}

func (m *FlowMessage) GetEtherDst() string {
	if m != nil && m.EtherDst != nil {
		return *m.EtherDst
	}
	return ""
}

func (m *FlowMessage) GetEtherType() string {
	if m != nil && m.EtherType != nil {
		return *m.EtherType
	}
	return ""
}

func (m *FlowMessage) GetIpv4Src() string {
	if m != nil && m.Ipv4Src != nil {
		return *m.Ipv4Src
	}
	return ""
}

func (m *FlowMessage) GetIpv4Dst() string {
	if m != nil && m.Ipv4Dst != nil {
		return *m.Ipv4Dst
	}
	return ""
}

func (m *FlowMessage) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *FlowMessage) GetPortSrc() uint32 {
	if m != nil && m.PortSrc != nil {
		return *m.PortSrc
	}
	return 0
}

func (m *FlowMessage) GetPortDst() uint32 {
	if m != nil && m.PortDst != nil {
		return *m.PortDst
	}
	return 0
}

func (m *FlowMessage) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *FlowMessage) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *FlowMessage) GetAttributes() *FlowMessage_Attrs {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type FlowMessage_InterfaceAttributes struct {
	TenantId         *string `protobuf:"bytes,1,opt,name=TenantId" json:"TenantId,omitempty"`
	VNI              *string `protobuf:"bytes,2,opt,name=VNI" json:"VNI,omitempty"`
	IfIndex          *uint32 `protobuf:"varint,3,opt,name=IfIndex" json:"IfIndex,omitempty"`
	IfName           *string `protobuf:"bytes,4,opt,name=IfName" json:"IfName,omitempty"`
	MTU              *uint32 `protobuf:"varint,5,opt,name=MTU" json:"MTU,omitempty"`
	BridgeName       *string `protobuf:"bytes,6,opt,name=BridgeName" json:"BridgeName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FlowMessage_InterfaceAttributes) Reset()         { *m = FlowMessage_InterfaceAttributes{} }
func (m *FlowMessage_InterfaceAttributes) String() string { return proto.CompactTextString(m) }
func (*FlowMessage_InterfaceAttributes) ProtoMessage()    {}

func (m *FlowMessage_InterfaceAttributes) GetTenantId() string {
	if m != nil && m.TenantId != nil {
		return *m.TenantId
	}
	return ""
}

func (m *FlowMessage_InterfaceAttributes) GetVNI() string {
	if m != nil && m.VNI != nil {
		return *m.VNI
	}
	return ""
}

func (m *FlowMessage_InterfaceAttributes) GetIfIndex() uint32 {
	if m != nil && m.IfIndex != nil {
		return *m.IfIndex
	}
	return 0
}

func (m *FlowMessage_InterfaceAttributes) GetIfName() string {
	if m != nil && m.IfName != nil {
		return *m.IfName
	}
	return ""
}

func (m *FlowMessage_InterfaceAttributes) GetMTU() uint32 {
	if m != nil && m.MTU != nil {
		return *m.MTU
	}
	return 0
}

func (m *FlowMessage_InterfaceAttributes) GetBridgeName() string {
	if m != nil && m.BridgeName != nil {
		return *m.BridgeName
	}
	return ""
}

type FlowMessage_Attrs struct {
	IntfAttrSrc      *FlowMessage_InterfaceAttributes `protobuf:"bytes,1,opt,name=IntfAttrSrc" json:"IntfAttrSrc,omitempty"`
	IntfAttrDst      *FlowMessage_InterfaceAttributes `protobuf:"bytes,2,opt,name=IntfAttrDst" json:"IntfAttrDst,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *FlowMessage_Attrs) Reset()         { *m = FlowMessage_Attrs{} }
func (m *FlowMessage_Attrs) String() string { return proto.CompactTextString(m) }
func (*FlowMessage_Attrs) ProtoMessage()    {}

func (m *FlowMessage_Attrs) GetIntfAttrSrc() *FlowMessage_InterfaceAttributes {
	if m != nil {
		return m.IntfAttrSrc
	}
	return nil
}

func (m *FlowMessage_Attrs) GetIntfAttrDst() *FlowMessage_InterfaceAttributes {
	if m != nil {
		return m.IntfAttrDst
	}
	return nil
}
