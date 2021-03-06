// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: multi/multi2.proto

package multi

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type Multi2_Color int

const (
	Multi2_BLUE  Multi2_Color = 0
	Multi2_GREEN Multi2_Color = 1
	Multi2_RED   Multi2_Color = 2
)

var Multi2_Color_name = map[int]string{
	0: "BLUE",
	1: "GREEN",
	2: "RED",
}
var Multi2_Color_value = map[string]int{
	"BLUE":  0,
	"GREEN": 1,
	"RED":   2,
}

func (x Multi2_Color) String() string {
	return Multi2_Color_name[int(x)]
}

type Multi2 struct {
	RequiredValue int32
	Color         Multi2_Color
}

// GetRequiredValue gets the RequiredValue of the Multi2.
func (m *Multi2) GetRequiredValue() (x int32) {
	if m == nil {
		return x
	}
	return m.RequiredValue
}

// GetColor gets the Color of the Multi2.
func (m *Multi2) GetColor() (x Multi2_Color) {
	if m == nil {
		return x
	}
	return m.Color
}

// MarshalToWriter marshals Multi2 to the provided writer.
func (m *Multi2) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if m.RequiredValue != 0 {
		writer.WriteInt32(1, m.RequiredValue)
	}

	if int(m.Color) != 0 {
		writer.WriteEnum(2, int(m.Color))
	}

	return
}

// Marshal marshals Multi2 to a slice of bytes.
func (m *Multi2) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Multi2 from the provided reader.
func (m *Multi2) UnmarshalFromReader(reader jspb.Reader) *Multi2 {
	for reader.Next() {
		if m == nil {
			m = &Multi2{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.RequiredValue = reader.ReadInt32()
		case 2:
			m.Color = Multi2_Color(reader.ReadEnum())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Multi2 from a slice of bytes.
func (m *Multi2) Unmarshal(rawBytes []byte) (*Multi2, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
