FROM python:3.8

WORKDIR /denormalizer

COPY . .

ENV PRODUCTION 1
ENV SCALER_FILE ./scaler_info.pickle
ENV ADDRESS denormalizer
ENV PORT 8085

RUN pip3 install sklearn numpy websockets python-dotenv dataclasses_json

EXPOSE 8085

RUN chmod a+x ./main.py

CMD ["./main.py"]