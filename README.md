# go-wsllinks
Like a symlink for a binary inside WSL

## How it works

Proxy application calls selected binary inside specified WSL distro.

## Configuration

Application loads 2 ini files relative to itself:

* `wsllinks.ini` - default configuration
* `<binary-name-no-ext>.ini` - overrides

Supported properties:

* `distro` - mandatory, name of the WSL distro, can't be empty
* `user` - mandatory, user name to use, can't be empty
* `binary` - optional, should be equal to `<binary-name-no-ext>` or absolute path of the form `/some/path/<binary-name-no-ext>`

## Example

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