# Stage 1: Build
FROM ruby:3.3.1 as builder

WORKDIR /app

# Copy Gemfile and Gemfile.lock
COPY Gemfile Gemfile.lock ./

# Update Bundler and install gems
RUN gem install bundler:2.5.10 && bundle update --bundler && bundle update && bundle install

# Copy the rest of the application
COPY . .

# Expose port
EXPOSE 4567

# Run the application
CMD ["bundle", "exec", "ruby", "./app.rb"]
