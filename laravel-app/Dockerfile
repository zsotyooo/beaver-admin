FROM php:8.1-fpm-alpine

# Arguments defined in docker-compose.yml
ARG user
ARG uid

# Install system dependencies
RUN apk update && apk add --no-cache \
  git \
  curl \
  libpng-dev \
  libxml2-dev \
  oniguruma-dev \
  zip \
  unzip \
  su-exec

# Install PHP extensions
RUN docker-php-ext-install pdo_mysql mbstring exif pcntl bcmath gd sockets

# Get latest Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Create system group 'www'
RUN addgroup -g 1000 -S www

# Create system user to run Composer and Artisan Commands
RUN adduser -D -g 'www' $user

# Set working directory
WORKDIR /var/www/html

# Copy existing application directory contents
COPY ./laravel-app /var/www/html

RUN mkdir -p /var/www/html/storage /var/www/html/bootstrap/cache /var/www/html/vendor

# Change ownership of the storage and bootstrap/cache directories
RUN chown -R $user:www /var/www/html/storage /var/www/html/bootstrap/cache /var/www/html/vendor

# Install Composer dependencies
RUN su-exec $user composer install

# Switch to the non-root user
USER $user