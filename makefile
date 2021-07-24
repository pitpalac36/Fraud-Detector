run-all:
	python ./normalizer/main.py &
	go run ./ms1/{main.go, handler.go, models.go} &
	go run ./gateway/{main.go, handler.go, models.go} &