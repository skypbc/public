package include

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

type EventDataProto func(ctx context.Context, event IEvent) (data IData, err error)
type EventSetDataProto func(ctx context.Context, event IEvent, data any, size ...int64) (err error)

type EventKind uint16

func (e EventKind) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventAction uint16

func (e EventAction) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventStatus uint16

func (e EventStatus) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventCategory uint16

func (e EventCategory) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventPriority uint16

func (e EventPriority) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventDataType uint16

func (e EventDataType) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventType uint8

func (e EventType) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventOwnerType uint8

func (e EventOwnerType) DbString() string {
	return strconv.FormatUint(uint64(e), 10)
}

type EventOpt int

func (e EventOpt) DbString() string {
	return strconv.FormatInt(int64(e), 10)
}

type IEvent interface {
	fmt.Stringer
	IMetaSource

	Uuid() UUID

	VirtualUuid() UUID
	SetVirtualUuid(uuid UUID)

	Kind() EventKind
	Category() []EventCategory
	Action() EventAction
	Type() EventType

	Status() EventStatus
	SetStatus(status EventStatus)

	Priority() EventPriority

	Meta() IMeta
	SetMeta(data map[string]any)
	MetaSize() int64

	IsData() bool
	Data(ctx context.Context) (IData, error)
	SetData(ctx context.Context, data any, size ...int64) error
	SetDataRaw(ctx context.Context, data any, size ...int64) error
	SetData2(ctx context.Context, data any, dataType EventDataType, size ...int64) error
	SetDataRaw2(ctx context.Context, data any, dataType EventDataType, size ...int64) error
	DataSize() int64
	SetDataSize(size int64)
	DataMaxSize() int64
	SetDataMaxSize(size int64)
	DataType() EventDataType
	SetDataType(dataType EventDataType)

	Opts() map[EventOpt]any
	OptVal(key EventOpt) (any, bool)
	SetOptVal(key EventOpt, val any)

	OwnerUuid() UUID
	OwnerType() EventOwnerType

	Created() int64
	Expired() int64

	ToMap() map[string]any

	Clone(opts EventCloneOpts) (IEvent, error)

	CloneWithType(type_ EventType) (IEvent, error)
	CloneWithMeta(type_ EventType, meta map[string]any) (IEvent, error)
	CloneWithData(ctx context.Context, type_ EventType, data any) (IEvent, error)
	CloneWithMetaAndData(ctx context.Context, type_ EventType, meta map[string]any, data any) (IEvent, error)
}

type IMeta interface {
	IMetaSource

	fmt.Stringer
	json.Marshaler
	json.Unmarshaler

	Value(key string, out any) (err error)
	RawValue(key string) (raw any, ok bool)

	SetValue(key string, value any) error

	Data() map[string]any
	SetData(data map[string]any)

	Size() int64
	Clone() (IMeta, error)
}

// IMetaSource используется для идентификации типов, которые могут предоставлять информацию о метаданных.
type IMetaSource interface {
	IsMetaSource()
}

type Event2MapOpts struct {
	MaxDataSize int64

	Status      bool
	Data        bool
	DataType    bool
	DataMaxSize bool
	ErrData     bool
	ErrSize     bool
	OwnerUuid   bool
	Created     bool
	Expired     bool
}

type NewEventArgs struct {
	VirtualUuid UUID
	Uuid        UUID
	Kind        EventKind
	Category    []EventCategory
	Action      EventAction
	Type        EventType
	Status      EventStatus
	Priority    EventPriority
	Meta        IMeta
	DataSize    int64
	DataMaxSize int64
	DataType    EventDataType
	OwnerUuid   UUID
	OwnerType   EventOwnerType
	Created     int64
	Expired     int64
	Opts        map[EventOpt]any
}

type EventCloneOpts struct {
	VirtualUuid *UUID
	Uuid        *UUID
	Kind        *EventKind
	Category    []EventCategory
	Action      *EventAction
	Type        *EventType
	Status      *EventStatus
	Priority    *EventPriority
	Meta        IMeta
	DataSize    *int64
	DataMaxSize *int64
	DataType    *EventDataType
	OwnerUuid   *UUID
	OwnerType   *EventOwnerType
	Created     *int64
	Expired     *int64
	Opts        map[EventOpt]any
}
