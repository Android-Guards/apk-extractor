package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

type Application struct {
	apkPath     string
	packageName string
}

func main() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	app := flag.String("app", "", "Target application name")
	out := flag.String("out", workingDirectory, "Output directory for apk file")

	flag.Parse()

	if *app == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		flag.PrintDefaults()
		return
	}

	if _, err := exec.LookPath("adb"); err != nil {
		log.Fatal(errors.New(string(err.(*exec.ExitError).Stderr)))
	}

	application, err := findApplication(*app)
	if err != nil {
		log.Fatal(err)
	}

	apkPath, err := saveApk(application, *out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("APK file is saved as %s", apkPath)
}

func findApplication(appName string) (application *Application, err error) {
	packages, err := exec.Command("adb", "shell", "pm", "list", "packages", "-f").Output()
	if err != nil {
		return nil, errors.New(string(err.(*exec.ExitError).Stderr))
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?im):(?P<path>.*?%s.*?\.apk)=(?P<package>.*?)$`, appName))

	match := re.FindStringSubmatch(string(packages))
	if match == nil {
		return nil, errors.New(fmt.Sprintf("`%s` not found", appName))
	}

	pathIndex := re.SubexpIndex("path")
	packageIndex := re.SubexpIndex("package")

	re = regexp.MustCompile(`\r?`)

	return &Application{
		apkPath:     match[pathIndex],
		packageName: re.ReplaceAllString(match[packageIndex], ""),
	}, nil
}

func saveApk(application *Application, outputDir string) (path string, err error) {
	targetPath := filepath.Join(outputDir, fmt.Sprintf("%s.apk", application.packageName))
	output, err := exec.Command("adb", "pull", application.apkPath, targetPath).Output()

	if err != nil {
		return "", errors.New(string(err.(*exec.ExitError).Stderr))
	}

	fmt.Println(string(output))

	return targetPath, nil
}
