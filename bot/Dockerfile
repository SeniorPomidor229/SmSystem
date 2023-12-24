FROM python:3.11.6-slim

RUN pip install aiogram requests pandas

COPY main.py /app/main.py

WORKDIR /app

CMD ["python", "main.py"]
