test:
	go test ./...
lib:
	go get -u github.com/nak3/note/lib

# create go program from template.
# make init new=$YOUR_NEW_PROGRAM
check-env:
	@ if test "$(new)" = ""; then \
		echo "usag: make init new=YOUR_NEW_PROGRAM"; \
		exit 1; \
	fi

init: check-env
	cp -r .template $(new)
	mv $(new)/template.go $(new)/$(new).go
	mv $(new)/template_test.go $(new)/$(new)_test.go
#	sed -i '' 's/REPLACE/$(new)' $(new)/$(new).go
#	sed -i '' 's/REPLACE/$(new)' $(new)/$(new)_test.go
