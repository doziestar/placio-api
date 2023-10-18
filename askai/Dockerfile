FROM tiangolo/uvicorn-gunicorn:python3.11

COPY ./requirements/base.txt .
RUN pip install -r base.txt

COPY ./scripts /app
COPY ./askai /app
WORKDIR /app
