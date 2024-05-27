FROM php:8.1-cli

# Install necessary tools
RUN apt-get update && apt-get install -y \
    git \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Install composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Set working directory
WORKDIR /app

# Copy composer.json and composer.lock files to the working directory
COPY composer.json composer.lock ./

# Install dependencies
RUN composer install --no-plugins --no-scripts

# Copy the rest of the application
COPY . .

# Expose port 8000 for the application
EXPOSE 8000

# Start the PHP built-in server
CMD ["php", "-S", "0.0.0.0:8000", "app.php"]
