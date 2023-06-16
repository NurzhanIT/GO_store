FROM python:3.9-alpine3.16

RUN mkdir /gateway

COPY /API_gateway /gateway

WORKDIR /gateway

#RUN apk add postgresql-client build-base postgresql-dev

RUN pip install -r /gateway/requirements.txt
