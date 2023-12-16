# Step 1: Build the application
# Use an official Node runtime as a parent image
FROM node:alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Bundle app source inside the docker image
COPY frontend/. .

# Generate Prisma client
RUN npx prisma generate

# Build the Next.js application
RUN npm run build

# Step 2: Run the application
# Use an official Node runtime as a parent image
FROM node:alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY frontend/package*.json ./

# Install only production dependencies
RUN npm install --only=production

# Copy the build artifacts from the builder stage
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/src ./src
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/prisma ./prisma
COPY --from=builder /app/jsconfig.json ./jsconfig.json
COPY --from=builder /app/next.config.js ./next.config.js
COPY --from=builder /app/next.config.js ./next.config.js
COPY --from=builder /app/node_modules ./node_modules
# COPY --from=builder /app/public ./public

# Your app binds to port 81 so you'll use the EXPOSE instruction to have it mapped by the docker daemon
EXPOSE 81

# Define the command to run your app using CMD which defines your runtime
CMD ["npm", "start"]
