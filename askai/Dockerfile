FROM tiangolo/uvicorn-gunicorn:python3.11

COPY ./requirements/base.txt .
RUN pip install -r base.txt

COPY ./askai /app
WORKDIR /app

ENTRYPOINT ["python", "main.py"]
