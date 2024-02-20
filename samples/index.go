package samples

import (
	"github.com/thuongtruong1009/gouse/samples/api"
)

func Run() {
	apiSample()
	// arraySample()
	// cacheSample()
	// chartSample()
	// configSample()
	consoleSample()
	// cryptoSample()
	dateSample()
	// functionSample()
	// helperSample()
	// iOSample()
	// mathSample()
	// netSample()
	// numberSample()
	// regexSample()
	// stringSample()
	// structSample()
	// toolsSample()
	// typeSample()
}

func apiSample() {
	api.SampleApiPortScanner()
	api.SampleApiPortChecker()
}

// func arraySample() {
// 	arrEqual()
// 	arrIncludes()
// 	arrMost()
// 	arrSum()
// 	arrChunk()
// 	arrDiff()
// 	arrDrop()
// 	arrIndex()
// 	arrMerge()
// 	arrCompact()

// 	arrIntersect()
// 	arrMin()
// 	arrMax()

// 	arrKeyBy()
// 	arrIndexBy()
// 	arrFilterBy()
// 	arrRejectBy()
// 	arrFindBy()
// 	arrForBy()
// 	arrMapBy()
// }

// func cacheSample() {
// 	cacheLocal()
// 	cacheTmp()
// }

// func chartSample() {
// 	chartBar()
// 	chartLine()
// 	chartPie()
// }

// func configSample() {
// 	configJson()
// 	configYaml()
// 	configToml()
// }

func consoleSample() {
	consoleCmd()
	consoleClear()
	consoleWithColor()
	consoleBanner()
	consoleSelect()
	consoleHelp()

	consoleList()
	consolePaper()
	consoleProgress()
	consoleRealtime()
	consoleChoice()
	consoleSpinner()
	consoleSplit()
	consoleStopwatch()
	consoleTable()
	consoleTab()
	consoleCountdown()
	consoleSequence()
	consoleInline()
	consoleParallel()
	consoleDir()
	consoleGlamour()
}

// func cryptoSample() {
// 	cryptoEncode()
// 	cryptoDecode()
// 	cryptoEncryptPassword()
// 	cryptoDecryptPassword()
// 	cryptoEncryptFile()
// 	cryptoDecryptFile()
// }

func dateSample() {
	dateTime()

	dateISO()
	dateShort()
	dateLong()
	dateUTC()

	dateToSecond()
	dateToMinute()
	dateToHour()
	dateSleepSecond()
	dateSleepMinute()
	dateSleepHour()

	dateClock()
}

// func functionSample() {
// 	funcDelay()
// 	funcRetry()
// 	funcTimes()
// 	funcInterval()
// 	funcLock()
// 	funcRunTime()
// }

// func helperSample() {
// 	helperRandomID()
// 	helperUUID()

// 	helperAutoMdDoc()
// }

// func iOSample() {
// 	ioCheckDir()
// 	ioCreateDir()
// 	ioRemoveDir()
// 	ioLsDir()
// 	ioHierarchyDir()
// 	ioCurrentDir()

// 	ioCheckFile()
// 	ioCreateFile()
// 	ioRemoveFile()
// 	ioReadFileByLine()
// 	ioFileInfo()
// 	ioRenameFile()
// 	ioCopyFile()
// 	ioTruncateFile()
// 	ioCleanFile()
// 	ioWriteToFile()
// 	ioAppendToFile()
// 	ioFileObj()

// 	ioCreatePath()
// 	ioReadPath()
// 	ioWritePath()

// 	ioZip()
// 	ioUnzip()
// }

// func mathSample() {
// 	mathIsPrime()
// 	mathIsEven()
// 	mathIsOdd()
// 	mathIsPerfectSquare()

// 	mathAbs()
// 	mathFloor()
// 	mathCeil()
// 	mathRound()
// 	mathMinMax()
// 	mathOperators()
// 	mathPower()
// 	mathFactorial()
// 	mathRoot()

// 	mathLog()
// 	mathPytago()
// 	mathTrigonometry()
// 	mathTransition()

// 	mathRect()
// 	mathCircle()
// 	mathTriangle()
// 	mathSquare()
// 	mathCube()
// 	mathSphere()
// 	mathCylinder()
// 	mathCone()
// 	mathTrapezoid()
// 	mathParallelogram()
// 	mathRhombus()
// 	mathEllipse()
// 	mathPolygon()
// }

// func netSample() {
// 	netOpen()
// 	netEncode()
// 	netDecode()
// 	netCheck()
// 	netCheckWithStatusCode()
// 	netHeader()
// 	netConnectTime()
// }

// func numberSample() {
// 	numRandom()
// 	numClamp()
// 	numInRange()
// }

// func regexSample() {
// 	regexIsMatch()
// 	regexMatch()
// 	regexMatchIndex()
// }

// func stringSample() {
// 	strCapitalize()
// 	strCamelCase()
// 	strSnakeCase()
// 	strKebabCase()

// 	strIsLetter()
// 	strIsDigit()
// 	strIncludes()
// 	strIsLower()
// 	strIsUpper()

// 	strSplit()
// 	strWords()
// 	strReverse()
// 	strLower()
// 	strUpper()
// 	strRepeat()
// 	strTruncate()
// 	strReplace()
// 	strTrim()
// 	strTrimBlank()
// 	strTrimPrefix()
// 	strTrimSuffix()
// 	strJoin()
// 	strConcat()
// 	strSubStr()
// 	strSlice()
// 	strSplice()
// 	strStartsWith()
// 	strEndsWith()
// 	strEscape()
// 	strUnescape()
// 	strPad()

// 	strCount()
// 	strLines()
// 	strIndex()
// 	strRandom()
// 	strAt()
// 	strCodePointAt()
// 	strCodePoint()
// 	strFromCodePointAt()
// 	strFromCodePoint()
// }

// func structSample() {
// 	structMerge()
// 	structRemove()
// 	structAdd()
// 	structSet()
// 	structGet()

// 	structClone()
// 	structHas()
// }

// func toolsSample() {
// 	toolDoc()
// 	toolProfile()
// }

// func typeSample() {
// 	typeStructConvert()
// 	typeStringConvert()
// 	typeCastToString()

// 	typeCheck()
// 	typeCheckUUID()
// 	typeCheckGmail()
// 	typeCheckYahoo()
// 	typeCheckOutlook()
// 	typeCheckEdu()
// 	typeCheckEmail()
// 	typeCheckUsername()
// 	typeCheckPassword()
// 	typeCheckPhone()
// }
