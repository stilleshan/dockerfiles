# =====================
# STAGE 1: Build React frontend
# =====================
FROM alpine:3.21.3 AS frontend-build

RUN apk add --no-cache nodejs npm python3 build-base git

WORKDIR /app/frontend

# Better caching: copy lockfile/package.json first
COPY frontend/package*.json ./
# Use npm ci for reproducible, faster, clean installs
RUN npm ci

COPY frontend/ ./
RUN npm run build

# =====================
# STAGE 2: Set up Node backend with built frontend
# =====================
FROM alpine:3.21.3

# Runtime deps only
RUN apk add --no-cache nodejs npm vnstat

WORKDIR /app/backend

# Better caching for backend too
COPY backend/package*.json ./
RUN npm ci --omit=dev

# Copy the rest of backend and the built frontend
COPY backend/ ./
COPY --from=frontend-build /app/frontend/build /app/frontend-build

ENV FRONTEND_DIR=frontend-build
ENV PORT=8050

EXPOSE 8050

CMD ["node", "server.js"]
