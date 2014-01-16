golcdsysinfo
============

Go library for Coldtears Electronics LCD Sys Info device (http://coldtearselectronics.wikispaces.com/)

This is a very straight rewrite of the C# LCDSysInfo library. Code could definitely be improved.

## Linux

I have not testing this on linux, but it should work following instructions from [https://github.com/kylelemons/gousb](https://github.com/kylelemons/gousb).

## Windows

This library uses libusb-1.0. 

Follow the driver instructions from [http://www.libusb.org/wiki/windows_backend](http://www.libusb.org/wiki/windows_backend). You want to install the WinUSB driver (I tested with v6.1.7600.16385) with Zadig (I tested with v2.0.1 Build 162) to your LCDSysInfo device.

Here are roughly the steps I took to compile for windows:

- `git clone git://git.libusb.org/libusb.git`
- Ensure you have MinGW64 with MSYS. You can get it at: [http://tdm-gcc.tdragon.net/download](http://tdm-gcc.tdragon.net/download)
- Open MSYS for MinGW64 (open MinGW/msys/1.0/msys.bat)
- CD to your libusb directory (e.g. cd /c/dev/libusb)
- `sh autogen.sh`
- At this point, autogen.sh will configure libusb with debugging forced on. To disable it, run ./configure --disable-debug-log (RECOMMENDED)
- `make`
- Now you will have a libusb directory that has the header files and libraries in libusb/.libs
- `go get https://github.com/kylelemons/gousb`
- You will get an error saying it won't build.
- Find where the gousb library installed, and cd to `github.com/kylelemons/gousb/usb`
- Edit usb.go and make sure you have the following before `import "C"`
   
```
// #cgo LDFLAGS: -LC:/dev/libusb/libusb/.libs -lusb-1.0
// #cgo CFLAGS: -IC:/dev/libusb/libusb/
// #include <libusb-1.0/libusb.h>
```

- `go install github.com/kylelemons/gousb/usb`
- If this fails, make sure you have all your paths correct.
- You should now be able to install golcdsysinfo and run the examples.