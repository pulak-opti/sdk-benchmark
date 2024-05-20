# Stage 1: Build
FROM python:3.8-slim-buster as builder

WORKDIR /app

# Copy files from host to container
COPY . .

# Install build dependencies
RUN pip install --upgrade pip && \
    pip install --user --no-warn-script-location -r requirements.txt

# Stage 2: Run
FROM python:3.8-slim-buster

WORKDIR /app

# Copy files from builder
COPY --from=builder /root/.local /root/.local
COPY --from=builder /app .

# Make sure scripts in .local are usable:
ENV PATH=/root/.local/bin:$PATH

# Expose port
EXPOSE 5000

# Run the application
CMD ["python", "app.py"]
