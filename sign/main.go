package main

import (
	"context"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/signintech/gopdf"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/moko-poi/example.com/modules"
	"github.com/moko-poi/example.com/modules/downloader"
	"github.com/moko-poi/example.com/modules/types/event"
	"github.com/moko-poi/example.com/modules/uploader"
)

func handleRequest(ctx context.Context, evnt event.Event) (event.Event, error) {
	item := evnt.Path
	newItem := modules.GenRndPDFName()
	bucket := os.Getenv("BUCKET_NAME")

	newItemLocalPath := "/tmp/" + newItem
	newItemObjectKey := "tmp/" + newItem
	signFilePath := "/tmp/" + modules.GenRndPDFName()

	itemPath, err := downloader.S3Download(bucket, item)
	if err != nil {
		return evnt, nil
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA5})
	pdf.AddPage()
	if err = pdf.AddTTFFont("Tanuki", "./TanukiMagic.ttf"); err != nil {
		return evnt, err
	}
	if err = pdf.SetFont("Tanuki", "", 12); err != nil {
		return evnt, err
	}
	pdf.Cell(nil, "次の利用者によってダウンロードされました")
	pdf.Br(20)
	pdf.Cell(nil, evnt.Email)
	pdf.WritePdf(signFilePath)

	inFiles := []string{itemPath, signFilePath}
	api.MergeFile(inFiles, newItemLocalPath, nil)
	if err = uploader.S3Upload(newItemLocalPath, bucket, newItemObjectKey); err != nil {
		return evnt, err
	}

	resp := event.Event{
		Email: evnt.Email,
		Path:  newItemObjectKey,
	}

	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}
