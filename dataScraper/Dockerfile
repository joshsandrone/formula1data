# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /scraper

# Install the MongoDB client
RUN apt-get update

# Copy the current directory contents into the container at /app
COPY . /scraper

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

ENTRYPOINT ["python", "main.py"]
