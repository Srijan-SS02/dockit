# Use an official Python runtime as a parent image
FROM python:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Install any dependencies using pip (if needed)
# RUN pip install -r requirements.txt

# Expose a port for the application (change to your application's port)
EXPOSE 5000

# Run the Flask application when the container starts
CMD ["python", "app.py"]
