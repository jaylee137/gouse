package math

import "math"

func Log(number, base int) int {
	return int(math.Log(float64(number)) / math.Log(float64(base)))
}

func LogF(number, base float64) float64 {
	return math.Log(number) / math.Log(base)
}

func Log2(number int) int {
	return Log(number, 2)
}

func Log2F(number float64) float64 {
	return LogF(number, 2)
}

func Log10(number int) int {
	return Log(number, 10)
}

func Log10F(number float64) float64 {
	return LogF(number, 10)
}

func Pytago(side1, side2 int) float64 {
	return SqrtF(float64(Pow(side1, 2) + Pow(side2, 2)))
}

func PytagoF(side1, side2 float64) float64 {
	return SqrtF(float64(PowF(side1, 2) + PowF(side2, 2)))
}

func Sin(angle int) float64 {
	return math.Sin(float64(angle))
}

func SinF(angle float64) float64 {
	return math.Sin(angle)
}

func Cos(angle int) float64 {
	return math.Cos(float64(angle))
}

func CosF(angle float64) float64 {
	return math.Cos(angle)
}

func Tan(angle int) float64 {
	return math.Tan(float64(angle))
}

func TanF(angle float64) float64 {
	return math.Tan(angle)
}

func Speed(distance, time float64) float64 {
	return DivideF(distance, time)
}

func Distance(speed, time float64) float64 {
	return MultiF(speed, time)
}

func Time(distance, speed float64) float64 {
	return DivideF(distance, speed)
}

// func Density(mass, volume int) int {
// 	return Divide(mass, volume)
// }

// func DensityF(mass, volume float64) float64 {
// 	return DivideF(mass, volume)
// }

// func Pressure(force, area int) int {
// 	return Divide(force, area)
// }

// func PressureF(force, area float64) float64 {
// 	return DivideF(force, area)
// }

// func Work(force, distance int) int {
// 	return Multi(force, distance)
// }

// func WorkF(force, distance float64) float64 {
// 	return MultiF(force, distance)
// }

// func Power(work, time int) int {
// 	return Divide(work, time)
// }

// func PowerF(work, time float64) float64 {
// 	return DivideF(work, time)
// }

// func Momentum(mass, velocity int) int {
// 	return Multi(mass, velocity)
// }

// func MomentumF(mass, velocity float64) float64 {
// 	return MultiF(mass, velocity)
// }

// func Acceleration(velocity, time int) int {
// 	return Divide(velocity, time)
// }

// func AccelerationF(velocity, time float64) float64 {
// 	return DivideF(velocity, time)
// }

// func Force(mass, acceleration int) int {
// 	return Multi(mass, acceleration)
// }

// func ForceF(mass, acceleration float64) float64 {
// 	return MultiF(mass, acceleration)
// }

// func Weight(mass, gravity int) int {
// 	return Multi(mass, gravity)
// }

// func WeightF(mass, gravity float64) float64 {
// 	return MultiF(mass, gravity)
// }

// func MomentumForce(momentum, time int) int {
// 	return Divide(momentum, time)
// }

// func MomentumForceF(momentum, time float64) float64 {
// 	return DivideF(momentum, time)
// }

// func MomentumVelocity(momentum, mass int) int {
// 	return Divide(momentum, mass)
// }

// func MomentumVelocityF(momentum, mass float64) float64 {
// 	return DivideF(momentum, mass)
// }

// func MomentumMass(momentum, velocity int) int {
// 	return Divide(momentum, velocity)
// }

// func MomentumMassF(momentum, velocity float64) float64 {
// 	return DivideF(momentum, velocity)
// }

// func MomentumTime(momentum, force int) int {
// 	return Divide(momentum, force)
// }

// func MomentumTimeF(momentum, force float64) float64 {
// 	return DivideF(momentum, force)
// }
