package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/therecipe/qt/internal/binding/templater"
	"github.com/therecipe/qt/internal/utils"
)

func main() {
	var (
		buildTarget = "desktop"
		tagFlags    string
		env         map[string]string
	)
	if len(os.Args) > 1 {
		buildTarget = os.Args[1]
	}

	if buildTarget != "desktop" || (runtime.GOOS == "windows" && buildTarget == "desktop") {

		switch buildTarget {
		case
			"android":
			{
				tagFlags = fmt.Sprintf("-tags=\"%v\"", buildTarget)

				switch runtime.GOOS {
				case "darwin", "linux":
					{
						env = map[string]string{
							"PATH":   os.Getenv("PATH"),
							"GOPATH": os.Getenv("GOPATH"),
							"GOROOT": runtime.GOROOT(),

							"GOOS":   "android",
							"GOARCH": "arm",
							"GOARM":  "7",

							"CGO_ENABLED":  "1",
							"CC":           filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
							"CXX":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
							"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm", "usr", "include")),
							"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm")),
						}
					}

				case "windows":
					{
						env = map[string]string{
							"PATH":   os.Getenv("PATH"),
							"GOPATH": os.Getenv("GOPATH"),
							"GOROOT": runtime.GOROOT(),

							"TMP":  os.Getenv("TMP"),
							"TEMP": os.Getenv("TEMP"),

							"GOOS":   "android",
							"GOARCH": "arm",
							"GOARM":  "7",

							"CGO_ENABLED":  "1",
							"CC":           filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-gcc"),
							"CXX":          filepath.Join(utils.ANDROID_NDK_DIR(), "toolchains", "arm-linux-androideabi-4.9", "prebuilt", runtime.GOOS+"-x86_64", "bin", "arm-linux-androideabi-g++"),
							"CGO_CPPFLAGS": fmt.Sprintf("-isystem %v", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm", "usr", "include")),
							"CGO_LDFLAGS":  fmt.Sprintf("--sysroot=%v -llog", filepath.Join(utils.ANDROID_NDK_DIR(), "platforms", "android-9", "arch-arm")),
						}
					}
				}
			}

		case
			"ios", "ios-simulator":
			{
				tagFlags = "-tags=\"ios\""

				var (
					GOARCH = func() string {
						if buildTarget == "ios" {
							return "arm64"
						}
						return "amd64"
					}()

					CLANGARCH, CLANGDIR, CLANGFLAG = func() (string, string, string) {
						if buildTarget == "ios" {
							return "arm64", "iPhoneOS", "iphoneos"
						}
						return "x86_64", "iPhoneSimulator", "ios-simulator"
					}()
				)

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": os.Getenv("GOPATH"),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   runtime.GOOS,
					"GOARCH": GOARCH,

					"CGO_ENABLED":  "1",
					"CGO_CPPFLAGS": fmt.Sprintf("-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v.sdk -m%v-version-min=7.0 -arch %v", CLANGDIR, CLANGDIR, CLANGFLAG, CLANGARCH),
					"CGO_LDFLAGS":  fmt.Sprintf("-isysroot /Applications/Xcode.app/Contents/Developer/Platforms/%v.platform/Developer/SDKs/%v.sdk -m%v-version-min=7.0 -arch %v", CLANGDIR, CLANGDIR, CLANGFLAG, CLANGARCH),
				}
			}

		case
			"desktop":
			{
				if runtime.GOOS == "windows" {
					env = map[string]string{
						"PATH":   os.Getenv("PATH"),
						"GOPATH": os.Getenv("GOPATH"),
						"GOROOT": runtime.GOROOT(),

						"TMP":  os.Getenv("TMP"),
						"TEMP": os.Getenv("TEMP"),

						"GOOS":   runtime.GOOS,
						"GOARCH": "386",

						"CGO_ENABLED": "1",
					}
				}
			}

		case "sailfish", "sailfish-emulator":
			{
				tagFlags = fmt.Sprintf("-tags=\"%v\"", strings.Replace(buildTarget, "-", "_", -1))

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": os.Getenv("GOPATH"),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   "linux",
					"GOARCH": "386",
				}

				if runtime.GOOS == "windows" {
					env["TMP"] = os.Getenv("TMP")
					env["TEMP"] = os.Getenv("TEMP")
				}

				//build linux,386 tools
				var _, err = ioutil.ReadDir(filepath.Join(runtime.GOROOT(), "bin", "linux_386"))
				if err != nil {
					var build = exec.Command(filepath.Join(runtime.GOROOT(), "src", func() string {
						if runtime.GOOS == "windows" {
							return "run.bat"
						}
						return "run.bash"
					}()))
					for key, value := range env {
						build.Env = append(build.Env, fmt.Sprintf("%v=%v", key, value))
					}
					build.Run()
				}

				if buildTarget == "sailfish" {
					env["GOARCH"] = "arm"
					env["GOARM"] = "7"
				}
			}

		case "rpi1", "rpi2", "rpi3":
			{
				tagFlags = fmt.Sprintf("-tags=\"%v\"", buildTarget)

				env = map[string]string{
					"PATH":   os.Getenv("PATH"),
					"GOPATH": os.Getenv("GOPATH"),
					"GOROOT": runtime.GOROOT(),

					"GOOS":   "linux",
					"GOARCH": "arm",
					"GOARM":  "7",

					"CGO_ENABLED": "1",
					"CC":          fmt.Sprintf("%v/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf-gcc", utils.RPI_TOOLS_DIR()),
					"CXX":         fmt.Sprintf("%v/arm-bcm2708/arm-rpi-4.9.3-linux-gnueabihf/bin/arm-linux-gnueabihf-g++", utils.RPI_TOOLS_DIR()),
				}

				if buildTarget == "rpi1" {
					env["GOARM"] = "6"
				}
			}
		}

		var cmd = exec.Command("go", "install")
		if tagFlags != "" {
			cmd.Args = append(cmd.Args, tagFlags)
		}
		if buildTarget != "desktop" {
			cmd.Args = append(cmd.Args, fmt.Sprintf("-installsuffix=%v", strings.Replace(buildTarget, "-", "_", -1)))
		}
		cmd.Args = append(cmd.Args, "std")
		for key, value := range env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
		}
		runCmd(cmd, "install.std_1")

		//armv7
		if buildTarget == "ios" && (strings.HasPrefix(runtime.Version(), "go1.7") || strings.HasPrefix(runtime.Version(), "devel")) {
			var cmdiOS = exec.Command("go", "install")
			if tagFlags != "" {
				cmdiOS.Args = append(cmdiOS.Args, tagFlags)
			}
			if buildTarget != "desktop" {
				cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("-installsuffix=%v", buildTarget))
			}
			cmdiOS.Args = append(cmdiOS.Args, "std")
			var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
			tmp = strings.Replace(tmp, "arm64", "arm", -1)
			cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
			runCmd(cmdiOS, "install.std_2")
		}
	}

	if buildTarget != "desktop" || strings.ToLower(os.Getenv("CI")) == "true" {
		return
	}

	fmt.Println("------------------------install-------------------------")

	for _, m := range templater.GetLibs() {

		if !(buildTarget == "android" && (m == "DBus" || m == "WebEngine" || m == "Designer") || strings.HasSuffix(m, "Extras")) &&
			!(strings.HasPrefix(buildTarget, "ios") && (m == "SerialPort" || m == "SerialBus" || m == "WebEngine" || m == "PrintSupport" || m == "Designer") || strings.HasSuffix(m, "Extras")) && //TODO: support for PrintSupport
			!(strings.HasPrefix(buildTarget, "rpi") && (m == "WebEngine" || m == "Designer") || strings.HasSuffix(m, "Extras")) { //TODO: support for WebEngine (rpi2 + rpi3)

			var before = time.Now()

			fmt.Print(strings.ToLower(m))

			if templater.ShouldBuild(m) {
				var cmd = exec.Command("go", "install")
				if tagFlags != "" {
					cmd.Args = append(cmd.Args, tagFlags)
				}
				if buildTarget != "desktop" {
					cmd.Args = append(cmd.Args, fmt.Sprintf("-installsuffix=%v", strings.Replace(buildTarget, "-", "_", -1)))
				}
				cmd.Args = append(cmd.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))

				if buildTarget != "desktop" || (runtime.GOOS == "windows" && buildTarget == "desktop") {
					for key, value := range env {
						cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", key, value))
					}
				}
				runCmd(cmd, fmt.Sprintf("install.%v_1", m))

				//armv7
				if buildTarget == "ios" && (strings.HasPrefix(runtime.Version(), "go1.7") || strings.HasPrefix(runtime.Version(), "devel")) {
					var cmdiOS = exec.Command("go", "install")
					if tagFlags != "" {
						cmdiOS.Args = append(cmdiOS.Args, tagFlags)
					}
					if buildTarget != "desktop" {
						cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("-installsuffix=%v", buildTarget))
					}
					cmdiOS.Args = append(cmdiOS.Args, fmt.Sprintf("github.com/therecipe/qt/%v", strings.ToLower(m)))
					var tmp = strings.Replace(strings.Join(cmd.Env, "|"), "-arch arm64", "-arch armv7", -1)
					tmp = strings.Replace(tmp, "arm64", "arm", -1)
					cmdiOS.Env = append(strings.Split(tmp, "|"), "GOARM=7")
					runCmd(cmdiOS, fmt.Sprintf("install.%v_2", m))
				}
			}

			fmt.Println(strings.Repeat(" ", 45-len(m)), time.Since(before)/time.Second*time.Second)
		}
	}
}

func runCmd(cmd *exec.Cmd, name string) {
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("\n\n%v\noutput:%s\nerror:%s\n\n", name, out, err)
		os.Exit(1)
	}
}
