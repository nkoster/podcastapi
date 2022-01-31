CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
ssh master systemctl --user stop podcastapi
sleep 1
scp podcastapi master:
ssh master systemctl --user start podcastapi
