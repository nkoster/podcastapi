#### Usage with local systemd

Clone repo and `cd` into `podcastapi`

Have a look at `builder`:

```
cat builder
CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
ssh master systemctl --user stop podcastapi
sleep 1
scp podcastapi master:
ssh master systemctl --user start podcastapi
```
"master" is the destination host where the binary will run.

Adjust "master" to your needs.

To prepare "master", the first time you have to:

```
ssh master mkdir -p .config/systemd/user
scp systemd/user/podcastapi.service master:.config/systemd/user/
ssh master systemctl --user daemon-relaod
```

Run `builder` to build and deploy to "master":

```
sh builder
```
