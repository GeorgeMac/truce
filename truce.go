// Code generated by gocode.Generate; DO NOT EDIT.

package truce

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = Runtime
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.LookupField(name)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 652 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\x84Sak\xdbH\x10\xddu|pZr\xf7\x03\x0e\x0e\x86\xf5q$\x97\x8bs\u07ce\nB0\x8e\xd3\x04\x8am\x12\xf5KU\x15\xd6\xca\xca^*i\xd5\xd5\xca\xc4M\\h\x9b\xa6\xfd\xdc\x1f\x1c\x95]\u02ceM\xed\xe4\xd3\x0eo\xe6\xbd7\xf3d\xffV~\xad\xe1Z\xf9\r\xe1\xf2\x13B\xff\x97\x1f\xb70\xde\x16i\xaeY\x1a\xf2c\xa6\x99\xc1\xf1\x16\xae\x9fK\xa9q\r\xe1z\x9f\xe9\x11\xdeF\xf8\x97\x13\x11\xf3\x1c\x97w\b\xa1?\xcb/5\x8c\u007f\xf7\x83\xb0\xe0\xcdH\xc4\x15\xf3\x0e\xe1\xf2\x16\xa1\x9d\xf2\xf3\x16\u01bf>\xe0\xb7\b\xd7p\xbd\xcb\x12n\x84\xea\x16$\b\xa1\xfb\xfaw\xb3\t\xc6\xf8\x8f\xa1\u0423b\xd0\fer0\xe4R\ry\xc2\xc2\x03\xad\x8a\x90c\x8c\x1d[4\u00c2\xe3\xfbz\x90\xb1\xf0-\x1br\xb0 !y\xc6C\x11\x89\x90i!\xd3\u0705k\xe2\xf8\xe3\xc3\\+\x91\x0e\x03\x17\x1a\xad\xfe\x19\xfcmPg\xccU.d\xea\u00988S2%\xa6e\xe7\x17\x8d\xc3\x0f\xf4\x8d\xff\xdf\xfe\xb3`\xef/J\x1c\xadX\x9agR\xe9\xfc\u020e9#\xad\xb3#\x17\x1a\xa7\x9e\xd77\x12NT\xa4\u10ed\xe3\xa7K\xbe'U\xaf2wR\x96p\x17R\xe2\x18\xe2\x948z\x92\xf1u<o\x92\xf1\r\x9c)\xb1\xd6\xcb;\xe7.\xf8\xcdf\xb3\xa2\x13'S<\x12W.\xcc\x00\u00d8\xefaY3A{fk\xffU\xe0\xb3\xfd\xf7\xe6\xfd\xe7\xc9s\xe7*vu\xa6\x86E\xc2S]\x997N\x04\x8f/\x03\xe2(\xae\v\x95\x9a\xdb\r`\xcc\xcd1O\x1aGfzm\x84\xa6\xb19\v\xdb^R\xafn\xb6\u027a \xd5N\x83\xc5\xf1.\u072c\xb5\x9d\xc3~\x10\xac\xb6\x8e\xe8<\xe7\x95\xe42\xa6G.\xc0\xc2$\xe1z$/]\xa0\xcf;\x9e\x11\xa3\xfd\xdeEU\xbc\xac\u0796\xd7>\xb5U\xaf\xef\x9d\xf5\xba\x17\xb6>\xee\xbc\xe8x\x1d[\x9evZ\u01f6h\xf7\xba\xddN{\xc6\xf2\xce[\xed\x0e]\xc9\xf8\xa7\\ZUos4\x99\x12\x89\xd0bl~a>\x1dH\x19\xd3\u007f\x81\x8aT\x9b'\x8a%\xb3\xc5`\xa2\xb9yg\xca4 \x8d<\x16\xa1\xe1\x00\xf8\x10I\x05W RXR\x83k\xa0~\xf0z\xe7j\x97\xc2\x14\x02b\x126\xd3\x00+S{P)\x91\u016ek\xbeS\xa4d\xe2\x02\x1d\xc8\xcb\t%S\xb8\xd98a\xb2\xa7\xc4\x193\xe5.>\xc0c\xf3\xef\n\xae&\x8f\x12\x16\x8c1\x8b\v\xbe\xf4wA\xe8G\x00\x00\x00\xff\xff;r\x19\f\"\x05\x00\x00")
