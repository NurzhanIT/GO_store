FROM python:3.9-alpine3.16

RUN mkdir /news_paper

COPY /news_paper /news_paper

WORKDIR /news_paper

#RUN apk add postgresql-client build-base postgresql-dev

RUN pip install -r /news_paper/requirements.txt
