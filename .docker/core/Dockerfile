# Build Step
# docker build -f .docker/core/Dockerfile -t hackyeah .
FROM golang:1.24 AS build

ARG vite_api_url="http://localhost"
ENV VITE_API_URL=$vite_api_url

ARG vite_vapid_key=""
ENV VITE_VAPID_KEY=$vite_vapid_key

RUN apt-get update && apt-get install -y --no-install-recommends \
    nodejs \
    npm \
    make \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY . .
RUN rm -fr apps/webapp/node_modules
RUN rm -f apps/webapp/package-lock.json
RUN make core_build

# Runtime Step
FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build /app/dist/core /app/core
WORKDIR /app
CMD ["/app/core"]