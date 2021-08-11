FROM python:3.8

WORKDIR /ai

COPY common ./common
COPY utils ./utils
COPY web.py ./web.py
COPY pickle_data ./pickle_data

ENV PRODUCTION=1

RUN pip3 install sklearn
RUN pip3 install pandas
RUN pip3 install numpy
RUN pip3 install websockets
RUN pip3 install imbalanced-learn
RUN pip3 install nest-asyncio
RUN pip3 install python-dotenv
RUN pip3 install dataclasses_json

EXPOSE 8083

RUN chmod a+x ./web.py

CMD ["./web.py"]