package conv2i

import (
    "testing"
)

func Test_Conv2uint64(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2uint64(in1)
    if err != nil && out1!= in1 {
        t.Logf("convert a uint64 to a uint64 failed")
        t.Fail()
    }
}
