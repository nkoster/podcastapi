# Usage

```
cat builder
CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
ssh master systemctl --user stop podcastapi
sleep 1
scp podcastapi master:
ssh master systemctl --user start podcastapi
```
"master" is the destination host where the binary will run.

Tp prepare the master, the first time you have to:

```
ssh master mkdir -p .config/systemd/user
scp systemd/user/podcastapi.service master:.config/systemd/user/
ssh master systemctl --user daemon-relaod
```
