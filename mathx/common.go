package mathx

import (
    "math"
)

// 计算某个值的平方根结果
func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

// 四舍五入
func Round(val float64, places int) float64 {
    var t float64
    f := math.Pow10(places)
    x := val * f
    if math.IsInf(x, 0) || math.IsNaN(x) {
        return val
    }
    if x >= 0.0 {
        t = math.Ceil(x)
        if (t - x) > 0.50000000001 {
            t -= 1.0
        }
    } else {
        t = math.Ceil(-x)
        if (t + x) > 0.50000000001 {
            t -= 1.0
        }
        t = -t
    }
    x = t / f

    if !math.IsInf(x, 0) {
        return x
    }

    return t
}
