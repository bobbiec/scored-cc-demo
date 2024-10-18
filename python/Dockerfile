FROM python:3-slim

WORKDIR /app

COPY requirements.txt /app

# Install any needed packages
RUN pip install --no-cache-dir -r requirements.txt

COPY app.py /app

# Run app.py when the container launches
CMD [ "python", "-m", "flask", "run", "--host=0.0.0.0", "--port=8000"]
