package events

import "github.com/skypbc/public/include"

type IEngine = include.IEngine
type IEngineConfig = include.IEngineConfig
type UUID = include.UUID

type IData = include.IData
type IDataFromMem = include.IDataFromMem
type IDataFromFile = include.IDataFromFile

type DataProto = include.EventDataProto
type SetDataProto = include.EventSetDataProto

type Kind = include.EventKind
type Category = include.EventCategory
type Action = include.EventAction
type Type = include.EventType
type Status = include.EventStatus
type Priority = include.EventPriority
type DataType = include.EventDataType
type OwnerType = include.EventOwnerType
type Opt = include.EventOpt

type IEvent = include.IEvent
type IMeta = include.IMeta
type IMetaSource = include.IMetaSource

type Event2MapOpts = include.Event2MapOpts
type NewEventArgs = include.NewEventArgs
type EventCloneOpts = include.EventCloneOpts

type ICommandSignature = include.ICommandSignature
