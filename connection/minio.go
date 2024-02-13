package connection

// import (
// 	"context"
// 	"log"

// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// )

// type MinioClient = *minio.Client

// type MinioConf struct {
// 	Endpoint  string
// 	AccessKey string
// 	SecretKey string
// 	UseSSL    bool
// 	Bucket    string
// 	Location  string
// }

// func Minio(conf MinioConf) (MinioClient, error) {
// 	ctx := context.Background()

// 	minioClient, errInit := minio.New(conf.Endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
// 		Secure: conf.UseSSL,
// 	})
// 	if errInit != nil {
// 		log.Fatalln(errInit)
// 	}

// 	exists, errBucketExists := minioClient.BucketExists(ctx, conf.Bucket)
// 	if errBucketExists == nil && exists {
// 		log.Printf("We already own %s\n", conf.Bucket)
// 		return minioClient, errInit
// 	} else {
// 		err := minioClient.MakeBucket(ctx, conf.Bucket, minio.MakeBucketOptions{Region: conf.Location})
// 		if err != nil {
// 			return minioClient, errInit
// 		} else {
// 			log.Printf("Successfully created %s\n", conf.Bucket)
// 			return minioClient, errInit
// 		}
// 	}
// }
