# i3lock-honeypot
Capture people trying to hijack your laptop. Just for fun

## Introduction
This objective of this repo is to "extend" i3lock functionality in order to take a picture from any unauthorized person trying to use your laptop.

It's not a security tool, but a trolling tool! When you lock the laptop, it will look like it is unlocked for other people. If anyone tries to hijack it, it will take a picture and show it through the screen!

## Instalation
There is some steps in order to install this. First, be sure to accomplish some requirements:
- [i3lock](https://github.com/i3/i3lock)
- [scrot](https://en.wikipedia.org/wiki/Scrot)
- [ffmpeg](https://github.com/FFmpeg/FFmpeg)
- [mogrify](https://linux.die.net/man/1/mogrify)
- Golang
- Root permissions on the machine

### Compiling golang binary
The logger.go file contains an small keylogger. This app will exit succesfully or not, depending on wether a certain key has been clicked. This is what it's used to guess if an unnauthorized access is happening.
```
#> go build logger.go
#> sudo cp logger /usr/local/bin/logger
```

### Ensure logger can be called with sudo without password
The keylogger requires root permissions to be executed. It's required that it can be called with sudo without password. Just add on /etc/sudoers file:
```
<youruser>    ALL=(ALL) NOPASSWD:/usr/local/bin/logger
```

### Take an screenshot of your screen
For this, you can use the command scrot, like:
```
#> scrot /tmp/screen.png
```
Save the image in a known location.

### Configure i3lock-honeypot
Copy the i3lock-honeypot into a location under $PATH. Then, edit the file and configure it.

Ensure you configure KEYCAPTURE_ARGS correctly. You can find the mouse and keyboard devices on your laptop running:
```
#> xinput
```

### Profit!
