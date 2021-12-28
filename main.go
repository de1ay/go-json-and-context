package main

import (
	jsonnames "github.com/de1ay/go-json-and-context/json/custom_names"
	jsonexclude "github.com/de1ay/go-json-and-context/json/exclude_field"
	jsonmarshal "github.com/de1ay/go-json-and-context/json/marshalling"
	jsonomitempty "github.com/de1ay/go-json-and-context/json/omitempty"

	// jsonserve "github.com/de1ay/go-json-and-context/json/serve"
	jsonbuiltintime "github.com/de1ay/go-json-and-context/json/built_in_time"
	jsoncustomtypes "github.com/de1ay/go-json-and-context/json/custom_types"
	jsontypes "github.com/de1ay/go-json-and-context/json/types"
	jsonunexported "github.com/de1ay/go-json-and-context/json/unexported"
	jsonunmarshal "github.com/de1ay/go-json-and-context/json/unmarshalling"
	jsonunstructed "github.com/de1ay/go-json-and-context/json/unstructed"

	// cancelchannel "github.com/de1ay/go-json-and-context/cancellation/channel"
	// cancelfunc "github.com/de1ay/go-json-and-context/cancellation/cancel_func"
	// canceltimeout "github.com/de1ay/go-json-and-context/cancellation/timeout"
	cancellationmulti "github.com/de1ay/go-json-and-context/cancellation/multiple_goroutines"
)

func main() {
	jsontypes.Slide32()
	jsonunexported.Slide33()
	jsonnames.Slide34()
	jsonexclude.Slide35()
	jsonomitempty.Slide36()
	jsonmarshal.Slide38()
	// next line runs a server in the main thread, this will block further execution
	// jsonserve.Slide39()
	jsonunmarshal.Slide40()
	jsonunstructed.Slide41()
	jsoncustomtypes.Slide45()
	jsonbuiltintime.Slide46()
	// cancelchannel.Slide49()
	// canceltimeout.Slide51()
	// cancelfunc.Slide52()
	cancellationmulti.Exec()
}
