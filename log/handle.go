package log

import (
	"log"
	"os"
	"time"
)

func Handle(logID, logMessage, output string) {
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Failed to open log file: ", err)
	}

	if err := (os.Chmod(output, 0600)); err != nil {
		log.Printf("Error setting chmod log file handle permissions: %v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Print("Failed to close log file: ", err)
		}
	}()

	// Create a new logger that writes to the log file
	logger := log.New(file, logID, log.Ldate|log.Ltime|log.Lshortfile)

	// Write some log messages
	logger.Println(logMessage)
}

func RotateTruncate(output string, maxFileSize ...int64) {
	maxSize := int64(1 * 1024 * 1024) // 1MB

	if len(maxFileSize) > 0 {
		maxSize = int64(maxFileSize[0])
	}

	logFile, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
	}

	if err := (os.Chmod(output, 0600)); err != nil {
		log.Printf("Error setting chmod rotate log file permissions: %v", err)
	}

	defer func() {
		if err := logFile.Close(); err != nil {
			log.Print("Failed to close log file: ", err)
		}
	}()

	log.SetOutput(logFile)

	for {
		info, err := logFile.Stat()
		if err != nil {
			log.Printf("Error getting log file stats: %v", err)
		}

		if info.Size() >= maxSize {
			err = truncateLogFile(logFile)
			if err != nil {
				log.Printf("Error truncating log file: %v", err)
			}
		}

		time.Sleep(time.Minute)
	}
}

func truncateLogFile(file *os.File) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}
