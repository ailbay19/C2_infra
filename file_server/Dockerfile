FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

COPY app.py app.py
COPY file_manager.py file_manager.py

CMD ["flask", "--app", "app", "run", "--port", "18081"]
