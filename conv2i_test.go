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

    in2 := uint32(222)
    out2, err := Conv2uint64(in2)
    if err != nil && out2!= uint64(in2) {
        t.Logf("convert a uint32 to a uint64 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2uint64(in3)
    if err != nil && out3!= uint64(in3) {
        t.Logf("convert a uint16 to a uint64 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2uint64(in4)
    if err != nil && out4 != uint64(in4) {
        t.Logf("convert a uint8 to a uint64 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2uint64(in5)
    if err != nil && out5 != uint64(in5) {
        t.Logf("convert a uint to a uint64 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2uint64(in6)
    if err != nil && out6 != uint64(in6) {
        t.Logf("convert a int64 to a uint64 failed")
        t.Fail()
    }

    in7 := int32(11111)
    out7, err := Conv2uint64(in7)
    if err != nil && out7 != uint64(in7) {
        t.Logf("convert a int32 to a uint64 failed")
        t.Fail()
    }

    in8 := int16(1611)
    out8, err := Conv2uint64(in8)
    if err != nil && out8 != uint64(in8) {
        t.Logf("convert a int16 to a uint64 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2uint64(in9)
    if err != nil && out9 != uint64(in9) {
        t.Logf("convert a int8 to a uint64 failed")
        t.Fail()
    }

    in10 := int(1234223)
    out10, err := Conv2uint64(in10)
    if err != nil && out10 != uint64(in10) {
        t.Logf("convert a int to a uint64 failed")
        t.Fail()
    }

    // boundary test
    in_maxint64 := INT64MAX
    out_maxint64, err := Conv2uint64(in_maxint64)
    if err != nil && out_maxint64 != uint64(in_maxint64) {
        t.Logf("convert a max int64 to a uint64 failed")
        t.Fail()
    }

    _, err = Conv2uint64(-1)
    if err == nil {
        t.Logf("convert a -1 to unint, failed")
        t.Fail()
    }
}

func Test_Conv2uint32(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2uint32(in1)
    if err != nil && out1!= uint32(in1) {
        t.Logf("convert a uint64 to a uint32 failed")
        t.Fail()
    }

    in2 := uint32(222)
    out2, err := Conv2uint32(in2)
    if err != nil && out2!= in2 {
        t.Logf("convert a uint32 to a uint32 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2uint32(in3)
    if err != nil && out3!= uint32(in3) {
        t.Logf("convert a uint16 to a uint32 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2uint32(in4)
    if err != nil && out4 != uint32(in4) {
        t.Logf("convert a uint8 to a uint32 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2uint32(in5)
    if err != nil && out5 != uint32(in5) {
        t.Logf("convert a uint to a uint32 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2uint32(in6)
    if err != nil && out6 != uint32(in6) {
        t.Logf("convert a int64 to a uint32 failed")
        t.Fail()
    }

    in7 := int32(11111)
    out7, err := Conv2uint32(in7)
    if err != nil && out7 != uint32(in7) {
        t.Logf("convert a int32 to a uint32 failed")
        t.Fail()
    }

    in8 := int16(1611)
    out8, err := Conv2uint32(in8)
    if err != nil && out8 != uint32(in8) {
        t.Logf("convert a int16 to a uint32 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2uint32(in9)
    if err != nil && out9 != uint32(in9) {
        t.Logf("convert a int8 to a uint32 failed")
        t.Fail()
    }

    in10 := int(1234223)
    out10, err := Conv2uint32(in10)
    if err != nil && out10 != uint32(in10) {
        t.Logf("convert a int to a uint32 failed")
        t.Fail()
    }

    // boundary test
    in_maxint32 := INT32MAX
    out_maxint32, err := Conv2uint32(in_maxint32)
    if err != nil && out_maxint32 != uint32(in_maxint32) {
        t.Logf("convert a max int32 to a uint32 failed")
        t.Fail()
    }

    _, err = Conv2uint32(-1)
    if err == nil {
        t.Logf("convert a -1 to unint, failed")
        t.Fail()
    }

    _, err = Conv2uint32(INT64MAX)
    if err == nil {
        t.Logf("convert a man int64 to unint, failed")
        t.Fail()
    }
}

func Test_Conv2uint16(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2uint16(in1)
    if err != nil && out1!= uint16(in1) {
        t.Logf("convert a uint64 to a uint16 failed")
        t.Fail()
    }

    in2 := uint32(222)
    out2, err := Conv2uint16(in2)
    if err != nil && out2!= uint16(in2) {
        t.Logf("convert a uint16 to a uint16 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2uint16(in3)
    if err != nil && out3!= in3 {
        t.Logf("convert a uint16 to a uint16 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2uint16(in4)
    if err != nil && out4 != uint16(in4) {
        t.Logf("convert a uint8 to a uint16 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2uint16(in5)
    if err != nil && out5 != uint16(in5) {
        t.Logf("convert a uint to a uint16 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2uint16(in6)
    if err != nil && out6 != uint16(in6) {
        t.Logf("convert a int64 to a uint16 failed")
        t.Fail()
    }

    in7 := int32(11111)
    out7, err := Conv2uint16(in7)
    if err != nil && out7 != uint16(in7) {
        t.Logf("convert a int32 to a uint16 failed")
        t.Fail()
    }

    in8 := int16(1611)
    out8, err := Conv2uint16(in8)
    if err != nil && out8 != uint16(in8) {
        t.Logf("convert a int16 to a uint16 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2uint16(in9)
    if err != nil && out9 != uint16(in9) {
        t.Logf("convert a int8 to a uint16 failed")
        t.Fail()
    }

    in10 := int(1234)
    out10, err := Conv2uint16(in10)
    if err != nil && out10 != uint16(in10) {
        t.Logf("convert a int to a uint16 failed")
        t.Fail()
    }

    // boundary test
    in_maxint16 := INT16MAX
    out_maxint16 , err := Conv2uint16(in_maxint16)
    if err != nil && out_maxint16 != uint16(in_maxint16) {
        t.Logf("convert a max int16 to a uint16 failed")
        t.Fail()
    }

    _, err = Conv2uint16(-1)
    if err == nil {
        t.Logf("convert a -1 to unint, failed")
        t.Fail()
    }

    _, err = Conv2uint16(uint32(UINT16MAX) +1)
    if err == nil {
        t.Logf("convert a uint16 max +1 to unint, failed")
        t.Fail()
    }
}

func Test_Conv2uint8(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2uint8(in1)
    if err != nil && out1!= uint8(in1) {
        t.Logf("convert a uint64 to a uint8 failed")
        t.Fail()
    }

    in2 := uint32(222)
    out2, err := Conv2uint8(in2)
    if err != nil && out2!= uint8(in2) {
        t.Logf("convert a uint16 to a uint8 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2uint8(in3)
    if err != nil && out3!= uint8(in3) {
        t.Logf("convert a uint16 to a uint8 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2uint8(in4)
    if err != nil && out4 != in4 {
        t.Logf("convert a uint8 to a uint8 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2uint8(in5)
    if err != nil && out5 != uint8(in5) {
        t.Logf("convert a uint to a uint8 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2uint8(in6)
    if err != nil && out6 != uint8(in6) {
        t.Logf("convert a int64 to a uint8 failed")
        t.Fail()
    }

    in7 := int32(111)
    out7, err := Conv2uint8(in7)
    if err != nil && out7 != uint8(in7) {
        t.Logf("convert a int32 to a uint8 failed")
        t.Fail()
    }

    in8 := int16(161)
    out8, err := Conv2uint8(in8)
    if err != nil && out8 != uint8(in8) {
        t.Logf("convert a int16 to a uint8 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2uint8(in9)
    if err != nil && out9 != uint8(in9) {
        t.Logf("convert a int8 to a uint8 failed")
        t.Fail()
    }

    in10 := int(123)
    out10, err := Conv2uint8(in10)
    if err != nil && out10 != uint8(in10) {
        t.Logf("convert a int to a uint8 failed")
        t.Fail()
    }

    // boundary test
    in_maxint8:= INT8MAX
    out_maxint8, err := Conv2uint8(in_maxint8)
    if err != nil && out_maxint8!= uint8(in_maxint8) {
        t.Logf("convert a max int8 to a uint8 failed")
        t.Fail()
    }

    _, err = Conv2uint8(-1)
    if err == nil {
        t.Logf("convert a -1 to uint8, failed")
        t.Fail()
    }

    _, err = Conv2uint8(uint32(UINT16MAX) +1)
    if err == nil {
        t.Logf("convert a uint16 max +1 to unint, failed")
        t.Fail()
    }
}

func Test_Conv2int64(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2int64(in1)
    if err != nil && out1!= int64(in1) {
        t.Logf("convert a uint64 to a int64 failed")
        t.Fail()
    }

    in2 := uint32(222)
    out2, err := Conv2int64(in2)
    if err != nil && out2!= int64(in2) {
        t.Logf("convert a uint32 to a int64 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2int64(in3)
    if err != nil && out3!= int64(in3) {
        t.Logf("convert a uint16 to a int64 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2int64(in4)
    if err != nil && out4 != int64(in4) {
        t.Logf("convert a uint8 to a int64 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2int64(in5)
    if err != nil && out5 != int64(in5) {
        t.Logf("convert a uint to a int64 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2int64(in6)
    if err != nil && out6 != int64(in6) {
        t.Logf("convert a int64 to a int64 failed")
        t.Fail()
    }

    in7 := int32(-11111)
    out7, err := Conv2int64(in7)
    if err != nil && out7 != int64(in7) {
        t.Logf("convert a int32 to a int64 failed")
        t.Fail()
    }

    in8 := int16(1611)
    out8, err := Conv2int64(in8)
    if err != nil && out8 != int64(in8) {
        t.Logf("convert a int16 to a int64 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2int64(in9)
    if err != nil && out9 != int64(in9) {
        t.Logf("convert a int8 to a int64 failed")
        t.Fail()
    }

    in10 := int(1234223)
    out10, err := Conv2int64(in10)
    if err != nil && out10 != int64(in10) {
        t.Logf("convert a int to a int64 failed")
        t.Fail()
    }

    // boundary test
    in_maxint64 := INT64MAX
    out_maxint64, err := Conv2int64(uint64(in_maxint64))
    if err != nil && out_maxint64 != int64(uint64(in_maxint64)) {
        t.Logf("convert a max int64 to a uint64 failed")
        t.Fail()
    }
}

func Test_Conv2int32(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2int32(in1)
    if err != nil && out1!= int32(in1) {
        t.Logf("convert a uint64 to a int32 failed")
        t.Fail()
    }

    in2 := uint32(222)
    out2, err := Conv2int32(in2)
    if err != nil && out2!= int32(in2) {
        t.Logf("convert a uint32 to a int32 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2int32(in3)
    if err != nil && out3!= int32(in3) {
        t.Logf("convert a uint16 to a int32 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2int32(in4)
    if err != nil && out4 != int32(in4) {
        t.Logf("convert a uint8 to a int32 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2int32(in5)
    if err != nil && out5 != int32(in5) {
        t.Logf("convert a uint to a int32 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2int32(in6)
    if err != nil && out6 != int32(in6) {
        t.Logf("convert a int64 to a int32 failed")
        t.Fail()
    }

    in7 := int32(11111)
    out7, err := Conv2int32(in7)
    if err != nil && out7 != int32(in7) {
        t.Logf("convert a int32 to a int32 failed")
        t.Fail()
    }

    in8 := int16(-1611)
    out8, err := Conv2int32(in8)
    if err != nil && out8 != int32(in8) {
        t.Logf("convert a int16 to a int32 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2int32(in9)
    if err != nil && out9 != int32(in9) {
        t.Logf("convert a int8 to a int32 failed")
        t.Fail()
    }

    in10 := int(1234223)
    out10, err := Conv2int32(in10)
    if err != nil && out10 != int32(in10) {
        t.Logf("convert a int to a int32 failed")
        t.Fail()
    }

    // boundary test
    in_maxint32 := INT32MAX
    out_maxint32, err := Conv2int32(uint32(in_maxint32))
    if err != nil && out_maxint32 != in_maxint32 {
        t.Logf("convert a max int32 to a uint32 failed")
        t.Fail()
    }


    _, err = Conv2int32(INT64MAX)
    if err == nil {
        t.Logf("convert a man int64 to int32, failed")
        t.Fail()
    }
}


func Test_Conv2int16(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2int16(in1)
    if err != nil && out1!= int16(in1) {
        t.Logf("convert a uint64 to a int16 failed")
        t.Fail()
    }

    in2 := uint32(222)
    out2, err := Conv2int16(in2)
    if err != nil && out2!= int16(in2) {
        t.Logf("convert a uint16 to a int16 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2int16(in3)
    if err != nil && out3!= int16(in3) {
        t.Logf("convert a uint16 to a int16 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2int16(in4)
    if err != nil && out4 != int16(in4) {
        t.Logf("convert a uint8 to a int16 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2int16(in5)
    if err != nil && out5 != int16(in5) {
        t.Logf("convert a uint to a int16 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2int16(in6)
    if err != nil && out6 != int16(in6) {
        t.Logf("convert a int64 to a int16 failed")
        t.Fail()
    }

    in7 := int32(11111)
    out7, err := Conv2int16(in7)
    if err != nil && out7 != int16(in7) {
        t.Logf("convert a int32 to a int16 failed")
        t.Fail()
    }

    in8 := int16(1611)
    out8, err := Conv2int16(in8)
    if err != nil && out8 != in8 {
        t.Logf("convert a int16 to a int16 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2int16(in9)
    if err != nil && out9 != int16(in9) {
        t.Logf("convert a int8 to a int16 failed")
        t.Fail()
    }

    in10 := int(1234)
    out10, err := Conv2int16(in10)
    if err != nil && out10 != int16(in10) {
        t.Logf("convert a int to a int16 failed")
        t.Fail()
    }

    // boundary test
    in_maxint16 := INT16MAX
    out_maxint16 , err := Conv2int16(uint16(in_maxint16))
    if err != nil && out_maxint16 != in_maxint16 {
        t.Logf("convert a max int16 to a uint16 failed")
        t.Fail()
    }

    _, err = Conv2int16(uint32(INT16MAX) +1)
    if err == nil {
        t.Logf("convert a int16 max +1 to unint, failed")
        t.Fail()
    }
}

func Test_Conv2int8(t *testing.T) {
    in1 := uint64(22)
    out1, err := Conv2int8(in1)
    if err != nil && out1!= int8(in1) {
        t.Logf("convert a uint64 to a int8 failed")
        t.Fail()
    }

    in2 := uint32(22)
    out2, err := Conv2int8(in2)
    if err != nil && out2!= int8(in2) {
        t.Logf("convert a uint16 to a int8 failed")
        t.Fail()
    }

    in3 := uint16(11)
    out3, err := Conv2int8(in3)
    if err != nil && out3!= int8(in3) {
        t.Logf("convert a uint16 to a int8 failed")
        t.Fail()
    }

    in4 := uint8(11)
    out4, err := Conv2int8(in4)
    if err != nil && out4 != int8(in4) {
        t.Logf("convert a uint8 to a int8 failed")
        t.Fail()
    }

    in5 := uint(11)
    out5, err := Conv2int8(in5)
    if err != nil && out5 != int8(in5) {
        t.Logf("convert a uint to a int8 failed")
        t.Fail()
    }

    in6 := int64(11)
    out6, err := Conv2int8(in6)
    if err != nil && out6 != int8(in6) {
        t.Logf("convert a int64 to a int8 failed")
        t.Fail()
    }

    in7 := int32(111)
    out7, err := Conv2int8(in7)
    if err != nil && out7 != int8(in7) {
        t.Logf("convert a int32 to a int8 failed")
        t.Fail()
    }

    in8 := int16(61)
    out8, err := Conv2int8(in8)
    if err != nil && out8 != int8(in8) {
        t.Logf("convert a int16 to a int8 failed")
        t.Fail()
    }

    in9 := int8(16)
    out9, err := Conv2int8(in9)
    if err != nil && out9 != in9 {
        t.Logf("convert a int8 to a int8 failed")
        t.Fail()
    }

    in10 := int(123)
    out10, err := Conv2int8(in10)
    if err != nil && out10 != int8(in10) {
        t.Logf("convert a int to a int8 failed")
        t.Fail()
    }

    // boundary test
    in_maxint8:= INT8MAX
    out_maxint8, err := Conv2uint8(uint64(in_maxint8))
    if err != nil && out_maxint8!= uint8(in_maxint8) {
        t.Logf("convert a max int8 to a int8 failed")
        t.Fail()
    }

    _, err = Conv2uint8(uint32(UINT8MAX) +1)
    if err == nil {
        t.Logf("convert a uint16 max +1 to unint, failed")
        t.Fail()
    }
}


