# Stage 1: Build the application
FROM node:20 AS build-stage
WORKDIR /app

# Copy the package.json and package-lock.json to install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the entire project to the container
COPY . .

# Build the application
RUN npm run build

# Stage 2: Serve the application using Nginx
FROM nginx:alpine AS production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Expose port 80 to make the app accessible
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]

