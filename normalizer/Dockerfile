FROM python:3.8

WORKDIR /normalizer

COPY . .

ENV PRODUCTION 1
ENV SCALER_FILE ./scaler_info.pickle
ENV ADDRESS normalizer
ENV PORT 8082
ENV WS_URL ws://normalizer:8082

RUN pip3 install sklearn numpy websockets python-dotenv dataclasses_json

EXPOSE 8082

RUN chmod a+x ./main.py

CMD ["./main.py"]