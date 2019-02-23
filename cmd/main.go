package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nicolas-martin/wox"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Upload a directory to a given bucket
//
// Usage:
// sync <params>
//	-region <region> // required
//	-bucket <bucket> // required
//	-path  <path> // required
func main() {
	bucketPtr := flag.String("bucket", "", "bucket to upload to")
	regionPtr := flag.String("region", "", "region to be used when making requests")
	pathPtr := flag.String("path", "", "path of directory to be synced")
	flag.Parse()

	sess := session.New(&aws.Config{
		Region: regionPtr,
	})
	uploader := s3manager.NewUploader(sess)

	iter := wox.NewSyncFolderIterator(*pathPtr, *bucketPtr)
	if err := uploader.UploadWithIterator(aws.BackgroundContext(), iter); err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error has occurred: %v", err)
	}

	if err := iter.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error occurred during file walking: %v", err)
	}

	fmt.Println("Success")
}
