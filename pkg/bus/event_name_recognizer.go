package bus

import "reflect"

func RecognizeEventName(event Dispatchable) string {
	return reflect.TypeOf(event).Elem().Name()
}
