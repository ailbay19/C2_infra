FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

COPY app.py app.py
COPY bin bin
COPY logs logs 

EXPOSE 18080

CMD ["flask", "--app", "app", "run", "--host", "0.0.0.0", "--port", "18080"]
