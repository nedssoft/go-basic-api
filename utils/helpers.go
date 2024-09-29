package utils

import "reflect"

func DefaultValue[T interface{}](value T, defaultValue T) T {
    if reflect.ValueOf(value).IsZero() {
        return defaultValue
    }
    return value
}