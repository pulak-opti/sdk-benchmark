# Start with a base image containing Java runtime
FROM openjdk:11-jdk-slim

# Install maven
RUN apt-get update
RUN apt-get install -y maven

# Make port 8000 available to the world outside this container
EXPOSE 8000

# Change to the app directory
WORKDIR /app

# Copy the project's pom.xml and resolve all dependencies
COPY pom.xml .

# Copy the project source and build the project
COPY src /app/src
RUN mvn clean package

# Run the jar file 
ENTRYPOINT ["java","-jar","target/java-api-server-1.0-SNAPSHOT.jar"]
