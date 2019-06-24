
mkdir -p artefacts/dist artefacts/build artefacts/osx 

go build -o artefacts/build/diary
cp -r dist/ artefacts/dist
cp -r artefacts/build/ artefacts/dist

# I _think_ this icon is royalty free... 
# it came from the royalty free collection?
#todo - make sure I don't get sued
go run etc/appify.go \
	-assets ./artefacts/dist \
	-bin diary \
	-icon ./icons8-book-64.png \
	-identifier com.github.neetle.diary \
	-name "Diary" \
	-o artefacts/osx/diary
