# APK extractor

This tool just lookup an application by name (or its part), extract it from the device, and saves it along the provided path.
Nothing more.

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