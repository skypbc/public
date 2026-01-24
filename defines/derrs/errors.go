package derrs

import (
	"github.com/skypbc/goutils/gerrors"
)

type IError = gerrors.IError
type NewErrorArgs = gerrors.NewErrorArgs

var (
	CodeIncorrectParams = gerrors.CodeIncorrectParams
	CodeUnknown         = gerrors.CodeUnknown
	CodeNotFound        = gerrors.CodeNotFound

	CodeNotImplemented = gerrors.NewCode("10000000-0000-0000-0000-000000000000")

	CodeCommand         = gerrors.NewCode("10000000-0000-0000-0001-000000000000")
	CodeCommandNotFound = gerrors.NewCode("10000000-0000-0000-0001-000000000001")

	CodeAccess = gerrors.NewCode("10000000-0000-0000-0002-000000000000")
	CodeConfig = gerrors.NewCode("10000000-0000-0000-0003-000000000000")

	CodePlugin     = gerrors.NewCode("10000000-0000-0000-0004-000000000000")
	CodePluginLoad = gerrors.NewCode("10000000-0000-0000-0004-000000000001")
)

var Wrap = gerrors.Wrap
var NewIncorrectParamsError = gerrors.NewIncorrectParamsError
var NewUnknownError = gerrors.NewUnknownError
var NewNotFoundError = gerrors.NewNotFoundError

func NewCommandError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:       CodeCommand,
		Categories: []gerrors.Code{CodeCommand},
		Name:       "Command",
		Template:   "Command error",
		Parents:    err,
	})
}

func NewCommandNotFoundError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:       CodeCommandNotFound,
		Categories: []gerrors.Code{CodeCommand},
		Name:       "CommandNotFound",
		Template:   "Command not found",
		Parents:    err,
	})
}

func NewAccessError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:       CodeAccess,
		Categories: []gerrors.Code{CodeAccess},
		Name:       "Access",
		Template:   "Access error",
		Parents:    err,
	})
}

func NewConfigError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:       CodeConfig,
		Categories: []gerrors.Code{CodeConfig},
		Name:       "Config",
		Template:   "Configuration error",
		Parents:    err,
	})
}

func NewNotImplementedError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:     CodeNotImplemented,
		Name:     "NotImplemented",
		Template: "Not implemented error",
		Parents:  err,
	})
}

func NewPluginError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:       CodePlugin,
		Categories: []gerrors.Code{CodePlugin},
		Name:       "Plugin",
		Template:   "Plugin error",
		Parents:    err,
	})
}

func NewPluginLoadError(err ...error) IError {
	return gerrors.NewError(NewErrorArgs{
		Code:       CodePluginLoad,
		Categories: []gerrors.Code{CodePlugin},
		Name:       "PluginLoad",
		Template:   "Plugin load error",
		Parents:    err,
	})
}
