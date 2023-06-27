package d4data

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slog"
)

const (
	keyFileName = "__fileName__"
	keySnoId    = "__snoID__"
	keyType     = "__type__"
	keyTypeHash = "__typeHash__"

	refKeyRaw       = "__raw__"
	refKeyGroup     = "__group__"
	refKeyGroupName = "groupName"
	refKeyName      = "name"
)

// TODO: pointer type

var (
	ErrNotFile    = errors.New("missing required fields for file")
	ErrNotObject  = errors.New("missing required fields for object")
	ErrNotType    = errors.New("missing required fields for type")
	ErrNotRef     = errors.New("missing required fields for ref")
	ErrNotVector2 = errors.New("missing required fields for 2d vector")
	ErrNotVector3 = errors.New("missing required fields for 3d vector")
)

type Type struct {
	Name string
	Hash int64
}

func (t Type) String() string {
	return t.Name
}

func HasTypeInfo(result gjson.Result) bool {
	return result.Get(keyType).Exists() &&
		result.Get(keyTypeHash).Exists()
}

func ParseType(result gjson.Result) (Type, error) {
	if !HasTypeInfo(result) {
		return Type{}, ErrNotType
	}

	return Type{
		Name: result.Get(keyType).Str,
		Hash: int64(result.Get(keyTypeHash).Num),
	}, nil
}

type Object struct {
	Type   Type
	Fields map[string]any
}

func IsObject(result gjson.Result) bool {
	return HasTypeInfo(result)
}

func ParseObject(result gjson.Result, skipFileMeta bool) (Object, error) {
	// Assure object
	if !IsObject(result) {
		return Object{}, ErrNotObject
	}

	// Get type info
	typeInfo, err := ParseType(result)
	if err != nil {
		return Object{}, err
	}

	fields := make(map[string]any)
	result.ForEach(func(fieldName, fieldValue gjson.Result) bool {
		switch fieldName.Str {
		case keyType, keyTypeHash:
			return true
		case keyFileName, keySnoId:
			if skipFileMeta {
				return true
			}
		case "":
			return true
		}

		// Don't use switch default so that  we can fallthrough from file meta
		fields[fieldName.Str] = ParseValue(fieldValue)
		return true
	})

	return Object{
		Type:   typeInfo,
		Fields: fields,
	}, nil
}

func ParseValue(value gjson.Result) any {
	if obj, err := ParseRef(value); err == nil {
		return obj
	}

	if obj, err := ParseObject(value, false); err == nil {
		return obj
	}

	if obj, err := ParseVector2(value); err == nil {
		return obj
	}

	if obj, err := ParseVector3(value); err == nil {
		return obj
	}

	if value.IsObject() {
		slog.Warn("Parsing generic object type", slog.Any("raw", value.Raw))
		fields := make(map[string]any)
		value.ForEach(func(key, value gjson.Result) bool {
			fields[key.String()] = ParseValue(value)
			return true
		})
		return fields
	}

	if value.IsArray() {
		values := make([]any, 0)
		value.ForEach(func(_, value gjson.Result) bool {
			values = append(values, ParseValue(value))
			return true
		})
		return values
	}

	return value.Value()
}

type Ref struct {
	Type      Type
	Raw       int64
	Group     int64
	GroupName string
	Extra     map[string]any
}

func (r Ref) String() string {
	return fmt.Sprintf("[%s] %s: %q", r.Type.Name, r.GroupName, r.Extra[refKeyName])
}

func IsRef(result gjson.Result) bool {
	return HasTypeInfo(result) &&
		result.Get(refKeyRaw).Exists() &&
		result.Get(refKeyGroup).Exists() &&
		result.Get(refKeyGroupName).Exists()
}

func ParseRef(value gjson.Result) (Ref, error) {
	if !IsRef(value) {
		return Ref{}, ErrNotRef
	}

	typeInfo, err := ParseType(value)
	if err != nil {
		return Ref{}, err
	}

	fields, ok := value.Value().(map[string]any)
	if !ok {
		return Ref{}, ErrNotRef
	}

	delete(fields, refKeyRaw)
	delete(fields, refKeyGroup)
	delete(fields, refKeyGroupName)

	return Ref{
		Type:      typeInfo,
		Raw:       value.Get(refKeyRaw).Int(),
		Group:     value.Get(refKeyGroup).Int(),
		GroupName: value.Get(refKeyGroupName).String(),
		Extra:     fields,
	}, nil
}

type Vector2 struct {
	X float64
	Y float64
}

func (v Vector2) String() string {
	return fmt.Sprintf("Vector(%f, %f)", v.X, v.Y)
}

func IsVector2(result gjson.Result) bool {
	var hasX, hasY, hasOther bool
	result.ForEach(func(key, value gjson.Result) bool {
		switch key.Str {
		case "x":
			hasX = true
		case "y":
			hasY = true
		default:
			hasOther = true
			return false
		}
		return true
	})
	return hasX && hasY && !hasOther
}

func ParseVector2(result gjson.Result) (Vector2, error) {
	if !IsVector2(result) {
		return Vector2{}, ErrNotVector2
	}

	return Vector2{
		X: result.Get("x").Float(),
		Y: result.Get("y").Float(),
	}, nil
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func IsVector3(result gjson.Result) bool {
	var hasX, hasY, hasZ, hasOther bool
	result.ForEach(func(key, value gjson.Result) bool {
		switch key.Str {
		case "x":
			hasX = true
		case "y":
			hasY = true
		case "z":
			hasZ = true
		default:
			hasOther = true
			return false
		}
		return true
	})
	return hasX && hasY && hasZ && !hasOther
}

func (v Vector3) String() string {
	return fmt.Sprintf("Vector(%f, %f, %f)", v.X, v.Y, v.Z)
}

func ParseVector3(result gjson.Result) (Vector3, error) {
	if !IsVector3(result) {
		return Vector3{}, ErrNotVector3
	}

	return Vector3{
		X: result.Get("x").Float(),
		Y: result.Get("y").Float(),
		Z: result.Get("z").Float(),
	}, nil
}

// File represents a file representing a SNO
type File struct {
	Object
	Name  string
	SnoId int64
	Raw__ gjson.Result
}

func IsFile(result gjson.Result) bool {
	return IsObject(result) && result.Get(keyFileName).Exists() &&
		result.Get(keySnoId).Exists()
}

func NewFile(result gjson.Result) (*File, error) {
	// Check if it's a file
	if !IsFile(result) {
		return nil, ErrNotFile
	}

	// Parse object
	obj, err := ParseObject(result, true)
	if err != nil {
		return nil, err
	}

	// Construct the file
	return &File{
		Object: obj,
		Name:   result.Get(keyFileName).String(),
		SnoId:  result.Get(keySnoId).Int(),
		Raw__:  result,
	}, nil
}
