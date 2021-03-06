package simulator

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/treeverse/lakefs/pkg/logging"
)

const EtagExtension = "etag"

type externalRecordDownloader struct {
	downloader *s3manager.Downloader
}

func NewExternalRecordDownloader(region string) *externalRecordDownloader {
	// The session the S3 Downloader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.AnonymousCredentials,
	}))

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	return &externalRecordDownloader{downloader}
}

func getEtagFileName(path string) string {
	return path + "." + EtagExtension
}

func getLocalEtag(path string) (string, error) {
	// if etag exists return
	etagFileName := getEtagFileName(path)
	etag, err := os.ReadFile(etagFileName)
	if err == nil {
		return string(etag), nil
	}
	if os.IsNotExist(err) {
		return "", nil
	}
	return "", err
}

func (d *externalRecordDownloader) DownloadRecording(bucket, key, destination string) error {
	etag, err := getLocalEtag(destination)
	if err != nil {
		return err
	}

	headObject, err := d.downloader.S3.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	s3Etag := aws.StringValue(headObject.ETag)
	if s3Etag == etag {
		return nil
	}

	logging.Default().WithFields(logging.Fields{"bucket": bucket, "key": key, "destination": destination}).Info("download Recording")
	// make sure target folder exists
	dir := filepath.Dir(destination)
	_ = os.MkdirAll(dir, os.ModePerm)
	// create file
	f, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	// Write the contents of S3 Object to the file
	n, err := d.downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	logging.Default().WithFields(logging.Fields{"file": destination, "bytes": n}).Info("file downloaded")

	// write the etag file
	etagFileName := getEtagFileName(destination)
	err = os.WriteFile(etagFileName, []byte(s3Etag), 0644) //nolint:gosec
	return err
}
