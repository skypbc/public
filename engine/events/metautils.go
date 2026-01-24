package events

import "encoding/json"

func Meta[R any, T IMetaSource](metaSource T) (res R) {
	meta := meta(metaSource)
	bytes, err := json.Marshal(meta.Data())
	if err != nil {
		return res
	}
	json.Unmarshal(bytes, &res) //nolint:errcheck
	return res
}

func TryMeta[R any, T IMetaSource](metaSource T) (res R, ok bool) {
	meta := meta(metaSource)
	bytes, err := json.Marshal(meta.Data())
	if err != nil {
		return res, false
	}
	if err := json.Unmarshal(bytes, &res); err != nil {
		return res, false
	}
	return res, true
}

func MetaValue[R any, T IMetaSource](metaSource T, key string, d R) (res R) {
	meta := meta(metaSource)
	if err := meta.Value(key, &res); err == nil {
		return res
	}
	return d
}

func MetaTryValue[R any, T IMetaSource](metaSource T, key string, d R) (res R, ok bool) {
	meta := meta(metaSource)
	if err := meta.Value(key, &res); err == nil {
		return res, true
	}
	return d, false
}

func MetaTryValue2[R any, T IMetaSource](metaSource T, key string) (res R, ok bool) {
	meta := meta(metaSource)
	if err := meta.Value(key, &res); err == nil {
		return res, true
	}
	return res, false
}

func MetaValueOrErr[R any, T IMetaSource](metaSource T, key string) (res R, err error) {
	meta := meta(metaSource)
	if err := meta.Value(key, &res); err != nil {
		return res, err
	}
	return res, nil
}

func MetaSetValue[R any, T IMetaSource](metaSource T, key string, value R) error {
	meta := meta(metaSource)
	return meta.SetValue(key, value)
}

func meta[T IMetaSource](metaSource T) (res IMeta) {
	switch eventOrMeta := any(metaSource).(type) {
	case IEvent:
		res = eventOrMeta.Meta()
	case IMeta:
		res = eventOrMeta
	default:
		res = NewMeta(nil)
	}
	return res
}
