# Index

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/samples/api"
	"github.com/thuongtruong1009/gouse/samples/array"
	"github.com/thuongtruong1009/gouse/samples/cache"
	"github.com/thuongtruong1009/gouse/samples/chart"
	"github.com/thuongtruong1009/gouse/samples/config"
	"github.com/thuongtruong1009/gouse/samples/console"
	"github.com/thuongtruong1009/gouse/samples/crypto"
	"github.com/thuongtruong1009/gouse/samples/date"
	"github.com/thuongtruong1009/gouse/samples/function"
	"github.com/thuongtruong1009/gouse/samples/helper"
	"github.com/thuongtruong1009/gouse/samples/io"
	"github.com/thuongtruong1009/gouse/samples/io/dir"
	"github.com/thuongtruong1009/gouse/samples/io/file"
	"github.com/thuongtruong1009/gouse/samples/io/path"
	math_check "github.com/thuongtruong1009/gouse/samples/math/check"
	math_fomular "github.com/thuongtruong1009/gouse/samples/math/fomular"
	math_geometry "github.com/thuongtruong1009/gouse/samples/math/geometry"
	math_operator "github.com/thuongtruong1009/gouse/samples/math/operator"
	"github.com/thuongtruong1009/gouse/samples/net"
	"github.com/thuongtruong1009/gouse/samples/number"
	"github.com/thuongtruong1009/gouse/samples/regex"
	"github.com/thuongtruong1009/gouse/samples/strings"
	"github.com/thuongtruong1009/gouse/samples/structs"
	"github.com/thuongtruong1009/gouse/samples/tools"
	types_cast "github.com/thuongtruong1009/gouse/samples/types/cast"
	types_check "github.com/thuongtruong1009/gouse/samples/types/check"
)
```
## Functions


### Run

```go
func Run() {
	apiSample()
	arraySample()
	cacheSample()
	chartSample()
	configSample()
	consoleSample()
	cryptoSample()
	dateSample()
	functionSample()
	helperSample()
	ioSample()
	mathSample()
	netSample()
	numberSample()
	regexSample()
	stringSample()
	structSample()
	toolsSample()
	typeSample()
}
```

### apiSample

```go
func apiSample() {
	api.SampleApiPortScanner()
	api.SampleApiPortChecker()
}
```

### arraySample

```go
func arraySample() {
	array.SampleArrayEqual()
	array.SampleArrayIncludes()
	array.SampleArrayMost()
	array.SampleArraySum()
	array.SampleArrayChunk()
	array.SampleArrayDiff()
	array.SampleArrayDrop()
	array.SampleArrayIndex()
	array.SampleArrayMerge()
	array.SampleArrayCompact()

	array.SampleArrayIntersect()
	array.SampleArrayMin()
	array.SampleArrayMax()

	array.SampleArrayKeyBy()
	array.SampleArrayIndexBy()
	array.SampleArrayFilterBy()
	array.SampleArrayRejectBy()
	array.SampleArrayFindBy()
	array.SampleArrayForBy()
	array.SampleArrayMapBy()
}
```

### cacheSample

```go
func cacheSample() {
	cache.SampleCacheLocal()
	cache.SampleCacheTmp()
}
```

### chartSample

```go
func chartSample() {
	chart.SampleChartBar()
	chart.SampleChartLine()
	chart.SampleChartPie()
}
```

### configSample

```go
func configSample() {
	config.SampleConfigJson()
	config.SampleConfigYaml()
	config.SampleConfigToml()
}
```

### consoleSample

```go
func consoleSample() {
	console.SampleConsoleCmd()
	console.SampleConsoleClear()
	console.SampleConsoleWithColor()
	console.SampleConsoleBanner()
	console.SampleConsoleSelect()
	console.SampleConsoleHelp()

	console.SampleConsoleList()
	console.SampleConsolePaper()
	console.SampleConsoleProgress()
	console.SampleConsoleRealtime()
	console.SampleConsoleChoice()
	console.SampleConsoleSpinner()
	console.SampleConsoleSplit()
	console.SampleConsoleStopwatch()
	console.SampleConsoleTable()
	console.SampleConsoleTab()
	console.SampleConsoleCountdown()
	console.SampleConsoleSequence()
	console.SampleConsoleInline()
	console.SampleConsoleParallel()
	console.SampleConsoleDir()
	console.SampleConsoleGlamour()
}
```

### cryptoSample

```go
func cryptoSample() {
	crypto.SampleCryptoEncode()
	crypto.SampleCryptoDecode()
	crypto.SampleCryptoEncryptPassword()
	crypto.SampleCryptoDecryptPassword()
	crypto.SampleCryptoEncryptFile()
	crypto.SampleCryptoDecryptFile()
}
```

### dateSample

```go
func dateSample() {
	date.SampleDateTime()

	date.SampleDateISO()
	date.SampleDateShort()
	date.SampleDateLong()
	date.SampleDateUTC()

	date.SampleDateToSecond()
	date.SampleDateToMinute()
	date.SampleDateToHour()
	date.SampleDateSleepSecond()
	date.SampleDateSleepMinute()
	date.SampleDateSleepHour()

	date.SampleDateClock()
}
```

### functionSample

```go
func functionSample() {
	function.SampleFuncDelay()
	function.SampleFuncRetry()
	function.SampleFuncTimes()
	function.SampleFuncInterval()
	function.SampleFuncLock()
	function.SampleFuncRunTime()
}
```

### helperSample

```go
func helperSample() {
	helper.SampleHelperRandomID()
	helper.SampleHelperUUID()

	helper.SampleHelperAutoMdDoc()
}
```

### ioSample

```go
func ioSample() {
	dir.SampleIoCheckDir()
	dir.SampleIoCreateDir()
	dir.SampleIoRemoveDir()
	dir.SampleIoLsDir()
	dir.SampleIoHierarchyDir()
	dir.SampleIoCurrentDir()

	file.SampleIoCheckFile()
	file.SampleIoCreateFile()
	file.SampleIoRemoveFile()
	file.SampleIoReadFileByLine()
	file.SampleIoFileInfo()
	file.SampleIoRenameFile()
	file.SampleIoCopyFile()
	file.SampleIoTruncateFile()
	file.SampleIoCleanFile()
	file.SampleIoWriteToFile()
	file.SampleIoAppendToFile()
	file.SampleIoFileObj()

	path.SampleIoCreatePath()
	path.SampleIoReadPath()
	path.SampleIoWritePath()

	io.SampleIoZip()
	io.SampleIoUnzip()
}
```

### mathSample

```go
func mathSample() {
	math_check.SampleMathIsPrime()
	math_check.SampleMathIsEven()
	math_check.SampleMathIsOdd()
	math_check.SampleMathIsPerfectSquare()

	math_operator.SampleMathAbs()
	math_operator.SampleMathFloor()
	math_operator.SampleMathCeil()
	math_operator.SampleMathRound()
	math_operator.SampleMathMin()
	math_operator.SampleMathMax()
	math_operator.SampleMathSum()
	math_operator.SampleMathMean()
	math_operator.SampleMathOperators()
	math_operator.SampleMathPower()
	math_operator.SampleMathFactorial()
	math_operator.SampleMathRoot()

	math_fomular.SampleMathLog()
	math_fomular.SampleMathPytago()
	math_fomular.SampleMathTrigonometry()
	math_fomular.SampleMathTransition()

	math_geometry.SampleMathRect()
	math_geometry.SampleMathCircle()
	math_geometry.SampleMathTriangle()
	math_geometry.SampleMathSquare()
	math_geometry.SampleMathCube()
	math_geometry.SampleMathSphere()
	math_geometry.SampleMathCylinder()
	math_geometry.SampleMathCone()
	math_geometry.SampleMathTrapezoid()
	math_geometry.SampleMathParallelogram()
	math_geometry.SampleMathRhombus()
	math_geometry.SampleMathEllipse()
	math_geometry.SampleMathPolygon()
}
```

### netSample

```go
func netSample() {
	net.SampleNetOpen()
	net.SampleNetEncode()
	net.SampleNetDecode()
	net.SampleNetCheck()
	net.SampleNetCheckWithStatusCode()
	net.SampleNetHeader()
	net.SampleNetConnectTime()
}
```

### numberSample

```go
func numberSample() {
	number.SampleNumRandom()
	number.SampleNumClamp()
	number.SampleNumInRange()
}
```

### regexSample

```go
func regexSample() {
	regex.SampleRegexIsMatch()
	regex.SampleRegexMatch()
	regex.SampleRegexMatchIndex()
}
```

### stringSample

```go
func stringSample() {
	strings.SampleStringCapitalize()
	strings.SampleStringCamelCase()
	strings.SampleStringSnakeCase()
	strings.SampleStringKebabCase()

	strings.SampleStringIsLetter()
	strings.SampleStringIsDigit()
	strings.SampleStringIncludes()
	strings.SampleStringIsLower()
	strings.SampleStringIsUpper()

	strings.SampleStringSplit()
	strings.SampleStringWords()
	strings.SampleStringReverse()
	strings.SampleStringLower()
	strings.SampleStringUpper()
	strings.SampleStringRepeat()
	strings.SampleStringTruncate()
	strings.SampleStringReplace()
	strings.SampleStringTrim()
	strings.SampleStringTrimBlank()
	strings.SampleStringTrimPrefix()
	strings.SampleStringTrimSuffix()
	strings.SampleStringJoin()
	strings.SampleStringConcat()
	strings.SampleStringSubStr()
	strings.SampleStringSlice()
	strings.SampleStringSplice()
	strings.SampleStringStartsWith()
	strings.SampleStringEndsWith()
	strings.SampleStringEscape()
	strings.SampleStringUnescape()
	strings.SampleStringPad()

	strings.SampleStringCount()
	strings.SampleStringLines()
	strings.SampleStringIndex()
	strings.SampleStringRandom()
	strings.SampleStringAt()
	strings.SampleStringCodePointAt()
	strings.SampleStringCodePoint()
	strings.SampleStringFromCodePointAt()
	strings.SampleStringFromCodePoint()
}
```

### structSample

```go
func structSample() {
	structs.SampleStructMerge()
	structs.SampleStructRemove()
	structs.SampleStructAdd()
	structs.SampleStructSet()
	structs.SampleStructGet()

	structs.SampleStructClone()
	structs.SampleStructHas()
}
```

### toolsSample

```go
func toolsSample() {
	tools.SampleToolDoc()
	tools.SampleToolProfile()
}
```

### typeSample

```go
func typeSample() {
	types_cast.SampleTypeStringConvert()
	types_cast.SampleTypeStructConvert()
	types_cast.SampleTypeCastToString()

	types_check.SampleTypeCheck()
	types_check.SampleTypeCheckUUID()
	types_check.SampleTypeCheckGmail()
	types_check.SampleTypeCheckYahoo()
	types_check.SampleTypeCheckOutlook()
	types_check.SampleTypeCheckEdu()
	types_check.SampleTypeCheckEmail()
	types_check.SampleTypeCheckUsername()
	types_check.SampleTypeCheckPassword()
	types_check.SampleTypeCheckPhone()
}
```
