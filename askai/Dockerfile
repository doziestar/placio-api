FROM tiangolo/uvicorn-gunicorn:python3.11

COPY ./requirements/base.txt .
RUN pip install -r base.txt

COPY ./askai /app
WORKDIR /app

# run the app server with gunicorn
ENTRYPOINT ["gunicorn", "-b", "0.0.0.0:8282", "main:app"]
