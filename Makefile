LDFLAGS = -ldflags "-X main.BuildVersion=$$VERSION -X cmd.gitCommit=$$COMMIT_ID"
GOARCH = amd64



build: clean
	go build $(LDFLAGS) -o aws-backups

.PHONY: clean
clean:
	rm -rf bin