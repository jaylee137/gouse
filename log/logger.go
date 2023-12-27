package log

import "log"

func Print(msg ...interface{}) {
	log.Print(msg)
}

func Println(msg ...interface{}) {
	log.Println(msg)
}

func Printf(format string, msg ...interface{}) {
	log.Printf(format, msg)
}

func Fatal(err ...interface{}) {
	log.Fatal(err)
}

func Fatalln(err ...interface{}) {
	log.Fatalln(err)
}

func Fatalf(format string, err ...interface{}) {
	log.Fatalf(format, err)
}

func Panic(err ...interface{}) {
	log.Panic(err)
}

func Panicln(err ...interface{}) {
	log.Panicln(err)
}

func Panicf(format string, err ...interface{}) {
	log.Panicf(format, err)
}
