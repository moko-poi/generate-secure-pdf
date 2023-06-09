module example.com/watermark

go 1.20

replace example.com/modules => ../modules

require (
	example.com/modules v0.0.0-00010101000000-000000000000
	github.com/aws/aws-lambda-go v1.41.0
	github.com/pdfcpu/pdfcpu v0.4.0
)

require (
	github.com/hhrutter/lzw v0.0.0-20190829144645-6f07a24e8650 // indirect
	github.com/hhrutter/tiff v0.0.0-20190829141212-736cae8d0bc7 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/thanhpk/randstr v1.0.5 // indirect
	golang.org/x/image v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
