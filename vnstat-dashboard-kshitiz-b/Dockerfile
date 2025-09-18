# =====================
# STAGE 1: Build React frontend
# =====================
FROM node:18 AS frontend-build

WORKDIR /app/frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend/ ./
RUN npm run build

# =====================
# STAGE 2: Set up Node backend with built frontend
# =====================
FROM node:18

WORKDIR /app

# Copy backend files
COPY backend/ ./backend

# Copy backend's package.json
COPY backend/package*.json ./backend/
WORKDIR /app/backend
RUN npm install

# Copy frontend build
COPY --from=frontend-build /app/frontend/build /app/frontend-build

# Install vnstat
RUN apt-get update && \
    apt-get install -y vnstat && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

ENV FRONTEND_DIR=frontend-build
ENV PORT=8050

EXPOSE 8050

CMD ["node", "server.js"]
