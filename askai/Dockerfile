FROM tiangolo/uvicorn-gunicorn:python3.11

# Copy the requirements and install
COPY ./requirements/base.txt .
RUN pip install -r base.txt

# Copy your FastAPI application
COPY ./askai /app
WORKDIR /app
