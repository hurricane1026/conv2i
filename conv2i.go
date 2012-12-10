package conv2i

import (
    "errors"

)

const (
    UINT64MAX = ^uint64(0)
    INT64MAX = int64(UINT64MAX >> 1)
    UINT32MAX = ^uint32(0)
    INT32MAX = int32(UINT32MAX >> 1)
    UINT16MAX = ^uint16(0)
    INT16MAX = int16(UINT16MAX >> 1)
    UINT8MAX = ^uint8(0)
    INT8MAX = int8(UINT8MAX >> 1)
    UINTMAX = ^uint(0)
    INTMAX = int(UINTMAX >> 1)
)

func Conv2uint (in interface{}) (out uint, err error) {
    if i, ok := in.(uint64); ok {
        if i == uint64(uint(i)) {
            out = uint(i)
        } else {
            err = errors.New("input is bigger than range of uint")
        }
    } else if i, ok := in.(uint32); ok {
        out = uint(i)

    } else if i, ok := in.(uint16); ok {
        out = uint(i)

    } else if i, ok := in.(uint8); ok {
        out = uint(i)

    } else if i, ok := in.(uint); ok {
        out = i

    } else if i, ok := in.(int64); ok {
        if i >= 0 && i == int64(uint(i)) {
            out = uint(i)
        } else {
            err = errors.New("input is a negative or it is bigger than range of uint, can not convert to uint")
        }
    } else if i, ok := in.(int32); ok {
        if i >= 0 {
            out = uint(i)
        } else {
            err = errors.New("input is a negative , can not convert to uint")
        }
    } else if i, ok := in.(int16); ok {
        if i >= 0 {
            out = uint(i)
        } else {
            err = errors.New("input is a negative , can not convert to uint")
        }
    } else if i, ok := in.(int8); ok {
        if i >= 0 {
            out = uint(i)
        } else {
            err = errors.New("input is a negative , can not convert to uint")
        }
    } else if i, ok := in.(int); ok {
        if i >= 0 {
            out = uint(i)
        } else {
            err = errors.New("input is a negative , can not convert to uint")
        }
    } else {
        err = errors.New("input is not a integer, can not convert to uint")
    }
    return

}

func Conv2int (in interface{}) (out int, err error) {
    if i, ok := in.(uint64); ok {
        if i == uint64(int(i)) {
            out = int(i)
        } else {
            err = errors.New("input is bigger than range of int")
        }
    } else if i, ok := in.(uint32); ok {
        if i == uint32(int(i)) {
            out = int(i)
        } else {
            err = errors.New("input is bigger than range of int")
        }

    } else if i, ok := in.(uint16); ok {
        out = int(i)

    } else if i, ok := in.(uint8); ok {
        out = int(i)

    } else if i, ok := in.(uint); ok {
        if i == uint(int(i)) {
            out = int(i)
        } else {
            err = errors.New("input is bigger than range of int")
        }

    } else if i, ok := in.(int64); ok {
        if i == int64(int(i)) {
            out = int(i)
        } else {
            err = errors.New("input is bigger than range of int")
        }

    } else if i, ok := in.(int32); ok {
        out = int(i)

    } else if i, ok := in.(int16); ok {
        out = int(i)

    } else if i, ok := in.(int8); ok {
        out = int(i)

    } else if i, ok := in.(int); ok {
        out = i
    } else {
        err = errors.New("input is not a integer, can not convert to int")
    }
    return
}

func Conv2uint64 (in interface{}) (out uint64, err error) {
    if i, ok := in.(uint64); ok {
        out = i

    } else if i, ok := in.(uint32); ok {
        out = uint64(i)

    } else if i, ok := in.(uint16); ok {
        out = uint64(i)

    } else if i, ok := in.(uint8); ok {
        out = uint64(i)

    } else if i, ok := in.(uint); ok {
        out = uint64(i)

    } else if i, ok := in.(int64); ok {
        if i >= 0 {
            out = uint64(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint64")
        }

    } else if i, ok := in.(int32); ok {
        if i >= 0 {
            out = uint64(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint64")
        }

    } else if i, ok := in.(int16); ok {
        if i >= 0 {
            out = uint64(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint64")
        }

    } else if i, ok := in.(int8); ok {
        if i >= 0 {
            out = uint64(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint64")
        }
    } else if i, ok := in.(int); ok {
        if i >= 0 {
            out = uint64(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint64")
        }
    } else {
        err = errors.New("input is not a integer, can not convert to uint64")
    }
    return
}

func Conv2uint32 (in interface{}) (out uint32, err error) {
    if i, ok := in.(uint); ok {
        temp := uint(uint32(i))
        if temp == i {
            out = uint32(i)
        } else {
            err = errors.New("input is a uint that bigger than range of uint32")
        }

    } else if i, ok := in.(uint64); ok {
        temp := uint64(uint32(i))
        if temp == i {
            out = uint32(i)
        } else {
            err = errors.New("input is a uint64 that bigger than range of uint32")
        }

    } else if i, ok := in.(uint32); ok {
        out = i

    } else if i, ok := in.(uint16); ok {
        out = uint32(i)

    } else if i, ok := in.(uint8); ok {
        out = uint32(i)

    } else if i, ok := in.(int); ok {
        if i >= 0 {
            temp := int(uint32(i))
            if temp == i {
                out = uint32(i)
            } else {
                err = errors.New("input is a int that bigger than range of uint32")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint32")
        }

    } else if i, ok := in.(int64); ok {
        if i >= 0 {
            temp := int64(uint32(i))
            if temp == i {
                out = uint32(i)
            } else {
                err = errors.New("input is a int64 that bigger than range of uint32")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint32")
        }

    } else if i, ok := in.(int32); ok {
        if i >= 0 {
            out = uint32(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint32")
        }

    } else if i, ok := in.(int16); ok {
        if i >= 0 {
            out = uint32(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint32")
        }

    } else if i, ok := in.(int8); ok {
        if i >= 0 {
            out = uint32(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint32")
        }

    } else {
        err = errors.New("input is not a integer, can not convert to uint32")
    }

    return
}

func Conv2uint16 (in interface{}) (out uint16, err error) {
    if i, ok := in.(uint); ok {
        temp := uint(uint16(i))
        if temp == i {
            out = uint16(i)
        } else {
            err = errors.New("input is a uint that bigger than range of uint16")
        }

    } else if i, ok := in.(uint64); ok {
        temp := uint64(uint16(i))
        if temp == i {
            out = uint16(i)
        } else {
            err = errors.New("input is a uint64 that bigger than range of uint16")
        }

    } else if i, ok := in.(uint32); ok {
        temp := uint32(uint16(i))
        if temp == i {
            out = uint16(i)
        } else {
            err = errors.New("input is a uint64 that bigger than range of uint16")
        }

    } else if i, ok := in.(uint16); ok {
        out = i

    } else if i, ok := in.(uint8); ok {
        out = uint16(i)

    } else if i, ok := in.(int); ok {
        if i >= 0 {
            temp := int(uint16(i))
            if temp == i {
                out = uint16(i)
            } else {
                err = errors.New("input is a int that bigger than range of uint16")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint16")
        }

    } else if i, ok := in.(int64); ok {
        if i >= 0 {
            temp := int64(uint16(i))
            if temp == i {
                out = uint16(i)
            } else {
                err = errors.New("input is a int64 that bigger than range of uint16")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint16")
        }

    } else if i, ok := in.(int32); ok {
        if i >= 0 {
            temp := int32(uint16(i))
            if temp == i {
                out = uint16(i)
            } else {
                err = errors.New("input is a int32 that bigger than range of uint16")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint16")
        }

    } else if i, ok := in.(int16); ok {
        if i >= 0 {
            out = uint16(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint16")
        }

    } else if i, ok := in.(int8); ok {
        if i >= 0 {
            out = uint16(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint16")
        }

    } else {
        err = errors.New("input is not a integer, can not convert to uint16")
    }

    return
}


func Conv2uint8 (in interface{}) (out uint8, err error) {
    if i, ok := in.(uint); ok {
        temp := uint(uint8(i))
        if temp == i {
            out = uint8(i)
        } else {
            err = errors.New("input is a uint that bigger than range of uint8")
        }

    } else if i, ok := in.(uint64); ok {
        temp := uint64(uint8(i))
        if temp == i {
            out = uint8(i)
        } else {
            err = errors.New("input is a uint64 that bigger than range of uint8")
        }

    } else if i, ok := in.(uint32); ok {
        temp := uint32(uint8(i))
        if temp == i {
            out = uint8(i)
        } else {
            err = errors.New("input is a uint32 that bigger than range of uint8")
        }

    } else if i, ok := in.(uint16); ok {
        temp := uint16(uint8(i))
        if temp == i {
            out = uint8(i)
        } else {
            err = errors.New("input is a uint16 that bigger than range of uint8")
        }

    } else if i, ok := in.(uint8); ok {
        out = i

    } else if i, ok := in.(int); ok {
        if i >= 0 {
            temp := int(uint8(i))
            if temp == i {
                out = uint8(i)
            } else {
                err = errors.New("input is a int that bigger than range of uint8")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint8")
        }

    } else if i, ok := in.(int64); ok {
        if i >= 0 {
            temp := int64(uint8(i))
            if temp == i {
                out = uint8(i)
            } else {
                err = errors.New("input is a int64 that bigger than range of uint8")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint8")
        }

    } else if i, ok := in.(int32); ok {
        if i >= 0 {
            temp := int32(uint8(i))
            if temp == i {
                out = uint8(i)
            } else {
                err = errors.New("input is a int32 that bigger than range of uint8")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint8")
        }

    } else if i, ok := in.(int16); ok {
        if i >= 0 {
            temp := int16(uint8(i))
            if temp == i {
                out = uint8(i)
            } else {
                err = errors.New("input is a int16 that bigger than range of uint8")
            }
        } else {
            err = errors.New("input is a negative, can not convert to uint8")
        }

    } else if i, ok := in.(int8); ok {
        if i >= 0 {
            out = uint8(i)
        } else {
            err = errors.New("input is a negative, can not convert to uint8")
        }

    } else {
        err = errors.New("input is not a integer, can not convert to uint8")
    }

    return
}


func Conv2int64 (in interface{}) (out int64, err error) {
    if i, ok := in.(int64); ok {
        out = i

    } else if i, ok := in.(int32); ok {
        out = int64(i)

    } else if i, ok := in.(int16); ok {
        out = int64(i)

    } else if i, ok := in.(int8); ok {
        out = int64(i)

    } else if i, ok := in.(int); ok {
        out = int64(i)

    } else if i, ok := in.(uint64); ok {
        if i <= uint64(INT64MAX) {
            out = int64(i)
        } else {
            err = errors.New("input is bigger than INT64MAX, can not convert into int64")
        }

    } else if i, ok := in.(uint32); ok {
        out = int64(i)

    } else if i, ok := in.(uint16); ok {
        out = int64(i)

    } else if i, ok := in.(uint8); ok {
        out = int64(i)

    } else if i, ok := in.(uint); ok {
        if uint64(UINTMAX) == UINT64MAX && uint64(i) > uint64(INT64MAX) {
            err = errors.New("input is bigger than INT64MAX, can not convert into int64")
        } else {
            out = int64(i)
        }

    } else {
        err = errors.New("input is not a integer, can not convert to int64")
    }
    return
}


func Conv2int32 (in interface{}) (out int32, err error) {
    if i, ok := in.(int64); ok {
        if i <= int64(INT32MAX) {
            out = int32(i)
        } else {
            err = errors.New("input is bigger than INT32MAX, can not convert into int32")
        }

    } else if i, ok := in.(int32); ok {
        out = i

    } else if i, ok := in.(int16); ok {
        out = int32(i)

    } else if i, ok := in.(int8); ok {
        out = int32(i)

    } else if i, ok := in.(int); ok {
        if i <= int(INT32MAX) {
            out = int32(i)
        } else {
            err = errors.New("input is bigger than INT32MAX, can not convert into int32")
        }

    } else if i, ok := in.(uint64); ok {
        if i <= uint64(INT32MAX) {
            out = int32(i)
        } else {
            err = errors.New("input is bigger than INT32MAX, can not convert into int32")
        }

    } else if i, ok := in.(uint32); ok {
        if i <= uint32(INT32MAX) {
            out = int32(i)
        } else {
            err = errors.New("input is bigger than INT32MAX, can not convert into int32")
        }

    } else if i, ok := in.(uint16); ok {
        out = int32(i)

    } else if i, ok := in.(uint8); ok {
        out = int32(i)

    } else if i, ok := in.(uint); ok {
        if i <= uint(INT32MAX) {
            out = int32(i)
        } else {
            err = errors.New("input is bigger than INT32MAX, can not convert into int32")
        }

    } else {
        err = errors.New("input is not a integer, can not convert to int32")
    }
    return
}

func Conv2int16 (in interface{}) (out int16, err error) {
    if i, ok := in.(int64); ok {
        if i <= int64(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else if i, ok := in.(int32); ok {
        if i <= int32(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else if i, ok := in.(int16); ok {
        out = i

    } else if i, ok := in.(int8); ok {
        out = int16(i)

    } else if i, ok := in.(int); ok {
        if i <= int(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else if i, ok := in.(uint64); ok {
        if i <= uint64(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else if i, ok := in.(uint32); ok {
        if i <= uint32(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else if i, ok := in.(uint16); ok {
        if i <= uint16(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else if i, ok := in.(uint8); ok {
        out = int16(i)

    } else if i, ok := in.(uint); ok {
        if i <= uint(INT16MAX) {
            out = int16(i)
        } else {
            err = errors.New("input is bigger than INT16MAX, can not convert into int16")
        }

    } else {
        err = errors.New("input is not a integer, can not convert to int16")
    }
    return
}

func Conv2int8 (in interface{}) (out int8, err error) {
    if i, ok := in.(int64); ok {
        if i <= int64(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(int32); ok {
        if i <= int32(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(int16); ok {
        if i <= int16(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(int8); ok {
        out = i

    } else if i, ok := in.(int); ok {
        if i <= int(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(uint64); ok {
        if i <= uint64(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(uint32); ok {
        if i <= uint32(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(uint16); ok {
        if i <= uint16(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(uint8); ok {
        if i <= uint8(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else if i, ok := in.(uint); ok {
        if i <= uint(INT8MAX) {
            out = int8(i)
        } else {
            err = errors.New("input is bigger than INT8MAX, can not convert into int8")
        }

    } else {
        err = errors.New("input is not a integer, can not convert to int8")
    }
    return
}


