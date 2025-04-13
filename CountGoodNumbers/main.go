

func countGoodNumbers(n int64) int {
    mod := int64(1000000007)
    return int(binExpo(5, (n+1)/2, mod) * binExpo(4, n / 2, mod) % mod)
}

func binExpo(x, y, mod int64) int64 {
    ret := int64(1)
    mul := x

    for y > 0 {
        if y % 2 == 1 {
            ret = ret * mul % mod
        }
        mul = mul * mul % mod
        y /= 2
    }
    return ret
}