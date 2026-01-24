package events

import (
	"github.com/skypbc/goutils/gslice"
	"github.com/skypbc/public/defines/derrs"
	"github.com/skypbc/public/engine/cmds"
)

type RawEvent struct {
	// Уникальный идентификатор
	VirtualUuid UUID `json:"virtual_uuid" xml:"virtual_uuid" form:"virtual_uuid"`
	// Уникальный тег
	Uuid UUID `json:"uuid" xml:"uuid" form:"uuid"`
	// Вид (направление) ивента (number)
	Kind Kind `json:"kind" xml:"kind" form:"kind"`
	// Категория ивента (number)
	Category []Category `json:"category" xml:"category" form:"category"`
	// Операция (number)
	Action Action `json:"action" xml:"action" form:"action"`
	// Тип ивента
	Type Type `json:"type" xml:"type" form:"type"`
	// Статус
	Status Status `json:"status,omitempty" xml:"status" form:"status"`
	// Приоритет
	Priority Priority `json:"priority,omitempty" xml:"priority" form:"priority"`
	// Метаданные
	Meta map[string]any `json:"meta" xml:"meta" form:"meta"`
	// Размер метаданных
	MetaSize int64 `json:"meta_size,omitempty" xml:"meta_size" form:"meta_size"`
	// Размер данных привязанных к ивенту
	DataSize int64 `json:"data_size" xml:"data_size" form:"data_size"`
	// Максимальный размер данных, который может содержать ивент.
	// Если значение не задано, оно определяется из конфига
	DataMaxSize int64 `json:"data_max_size" xml:"data_max_size" form:"data_max_size"`
	// Тип данных
	DataType DataType `json:"data_type,omitempty" xml:"data_type" form:"data_type"`
	// Владелец
	OwnerUuid UUID `json:"owner_uuid" xml:"owner_uuid" form:"owner_uuid"`
	// Тип владелеца
	OwnerType OwnerType `json:"owner_type" xml:"owner_type" form:"owner_type"`
	// Время создания
	Created int64 `json:"created,omitempty" xml:"created" form:"created"`
	// Время когда ивент станет просроченным
	Expired int64 `json:"expired,omitempty" xml:"expired" form:"expired"`
	// Произвольные настройки. Могут отсутствовать.
	// Настраиваются кодом, в момент получения ивента.
	Opts map[Opt]any
}

func (w *RawEvent) CommandSignature() ICommandSignature {
	return cmds.NewCommandSignature(w.Kind, w.Category, w.Action)
}

func (w *RawEvent) SetCommandSignature(sig ICommandSignature) {
	w.Kind = sig.Kind()
	w.Category = gslice.Clone(sig.Category())
	w.Action = sig.Action()
}

func (w *RawEvent) Event() (event IEvent, err error) {
	if w.Meta == nil {
		w.Meta = map[string]any{}
	}
	if w.Opts == nil {
		w.Opts = map[Opt]any{}
	}

	meta := NewMeta(w.Meta)
	if meta.Size() > MaxMetaSize() {
		return nil, derrs.NewIncorrectParamsError().
			SetTemplate("meta.size > MaxMetaSize")
	}

	opts := NewEventArgs{
		VirtualUuid: w.VirtualUuid,
		Uuid:        w.Uuid,
		Kind:        w.Kind,
		Category:    w.Category,
		Action:      w.Action,
		Type:        w.Type,
		Status:      w.Status,
		Priority:    w.Priority,
		Meta:        meta,
		DataSize:    w.DataSize,
		DataMaxSize: w.DataMaxSize,
		DataType:    w.DataType,
		OwnerUuid:   w.OwnerUuid,
		OwnerType:   w.OwnerType,
		Created:     w.Created,
		Expired:     w.Expired,
		Opts:        w.Opts,
	}

	return NewEvent2(opts), nil
}
