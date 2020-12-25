# APK extractor

![Go](https://github.com/Android-Guards/apk-extractor/workflows/Go/badge.svg) ![CodeQL](https://github.com/Android-Guards/apk-extractor/workflows/CodeQL/badge.svg)

This tool just lookup an application by name (or its part), extract it from the device, and saves it along the provided path.
Nothing more.

### Building

Building for the your current OS
```shell script
$ go build .
```

Building for multiple OS
```shell script
$ make build_all
```

### Usage

```shell script
Usage:
  -app string
        Target application name
  -out string
        Output directory for apk file (default "<your current directory>")
```

Example
```shell script
$ ./apk-extractor -app youtube
/system/app/YouTube/YouTube.apk: 1 file pulled, 0 skipped. 67.2 MB/s (45652650 bytes in 0.648s)

APK file is saved as /home/user/Downloads/com.google.android.youtube.apk%     
```

### Feedback

In case you have faced any bugs or have any useful suggestions for improvement of this tool, feel free to create an [issue](https://github.com/Android-Guards/apk-extractor/issues/new).