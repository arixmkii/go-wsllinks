# go-wsllinks
Like a symlink for a binary inside WSL

## How it works

Proxy application calls selected binary inside specified WSL distro.

## Configuration

Application loads 2 ini files relative to itself:

* `wsllinks.ini` - default configuration
* `<binary-name-no-ext>.ini` - overrides

Supported properties:

* `mode` - optional, operation mode, currently supported "wsl" (default) or "direct"
* `distro` - mandatory, wsl mode only, name of the WSL distro, can't be empty
* `user` - optinonal, wsl mode only, user name to use, if skipped then will use default user configured by wsl
* `binary` - optional, should be equal to `<binary-name-no-ext>` or absolute path of the form `/some/path/<binary-name-no-ext>` in WSL mode
or absolute or relative path in Windows format in direct mode.

## Example 1

To output release information from distro named `Ubuntu` and user named `user`:

Build the app:
```bat
go build
```
Rename it to macth `cat` command:
```bat
move go-wsllinks.exe cat.exe
```
Prepare config:
```bat
echo distro = Ubuntu > cat.ini
echo user = user >> cat.ini
```
Call the command:
```bat
cat /etc/os-release
```

## Example 2

To output information from default msys2 installation:

Build the app:
```bat
go build
```
Rename it to macth `uname` command:
```bat
move go-wsllinks.exe uname.exe
```
Prepare config:
```bat
echo mode = direct > uname.ini
echo binary = C:\msys64\usr\bin\uname.exe >> uname.ini
```
Call the command:
```bat
uname -a
```
