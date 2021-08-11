run-all:
	python ./normalizer/main.py &
	python ./denormalizer/main.py &
	go run github.com/pitpalac36/Fraud-Detector/aggregator &
	npm start ./dashboard &
	python ./algorithm/web.py &
	go run github.com/pitpalac36/Fraud-Detector/processor &
	go run github.com/pitpalac36/Fraud-Detector/gateway &
	go run github.com/pitpalac36/Fraud-Detector/load_test2