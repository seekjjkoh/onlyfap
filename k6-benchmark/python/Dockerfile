FROM python:3.9-slim

WORKDIR /app

COPY app.py ./

RUN pip install --upgrade pip setuptools wheel
RUN pip install Flask gunicorn

EXPOSE 1802

CMD exec gunicorn --bind :1802 --workers 1 --threads 10 --timeout 0 app:app
