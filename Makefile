DISTDIR = ./dist

copy_assets:
	@echo "Copying assets"
	@echo "Copying config.json"
	mkdir -p $(DISTDIR)
	cp config/config.json $(DISTDIR)/config.json

build_dataCollector:
	@echo "Building dataCollector..."
	go build -o $(DISTDIR)/data_collector cmd/dataCollector/main.go

build_web:
	@echo "Building web server..."
	go build -o $(DISTDIR)/htrade_web web/main.go web/routing.go

clean:
	@echo "Cleaning dist dir: $(DISTDIR)"
	rm -v -f -r $(DISTDIR)/
	rm -v -f -r *.log

build: clean copy_assets build_dataCollector build_web
