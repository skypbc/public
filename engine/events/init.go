package events

import "context"

// event_new.go
var NewEvent func(uuid UUID, kind Kind, category []Category, action Action) IEvent
var NewEvent2 func(opts NewEventArgs) IEvent
var NewEvent3 func(opts EventCloneOpts) IEvent

// meta.go
var NewMeta func(value map[string]any) IMeta
var NewMeta2 func(value map[string]any, size int64) IMeta

// utils.go
var SetCommandSignature func(ctx context.Context, rawEvent *RawEvent, cmdId int64) error

// dconfig.go
var MaxMetaSize func() int64

var SetEventDataFunc func(cfg IEngineConfig, fn DataProto) error
var SetEventSetDataFunc func(cfg IEngineConfig, fn SetDataProto) error

var SetEventSetDataRawFunc func(cfg IEngineConfig, fn SetDataProto) error
