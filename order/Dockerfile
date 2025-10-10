# syntax=docker/dockerfile:1

# ---- Base image ----
FROM python:3.11-slim AS base

# Prevent Python from writing .pyc files and using stdout buffering
ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1 \
    PATH="/usr/local/bin:$PATH"

# Create a non-root user (for security)
RUN adduser --disabled-password --gecos "" appuser

WORKDIR /app

# ---- Install dependencies ----
COPY requirements.txt .

# Install only system dependencies you actually need
RUN apt-get update && \
    apt-get install -y --no-install-recommends gcc && \
    pip install --no-cache-dir -r requirements.txt && \
    apt-get remove -y gcc && apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/*

# ---- Copy app code ----
COPY . .

# Change ownership to non-root
RUN chown -R appuser /app
USER appuser

EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
