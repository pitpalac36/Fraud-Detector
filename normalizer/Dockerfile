FROM python:3.8

WORKDIR /normalizer

COPY . .

ENV PRODUCTION=1

RUN pip3 install sklearn numpy websockets python-dotenv dataclasses_json

EXPOSE 8082

RUN chmod a+x ./main.py

CMD ["./main.py"]