# Builder image installs tools needed to build aro-hcp-backend
FROM --platform=linux/amd64 mcr.microsoft.com/oss/go/microsoft/golang:1.23-fips-cbl-mariner2.0@sha256:97d76864911dfbaf6b3387b6e0583abe5af74d7b744773c4facbaf02389e654f as builder
WORKDIR /app
ADD archive.tar.gz .
# https://github.com/microsoft/go/tree/microsoft/main/eng/doc/fips#build-option-to-require-fips-mode
ENV CGO_ENABLED=1 GOFLAGS='-tags=requirefips'
RUN cd backend && make backend


# Deployment image copies aro-hcp-backend from builder image
FROM --platform=linux/amd64 mcr.microsoft.com/cbl-mariner/distroless/base:2.0-nonroot@sha256:ce44fc29db88c9aba8041a50c1abcd19a54f997c2b99a8c513e8ec113261374a
WORKDIR /
COPY --from=builder /app/backend/aro-hcp-backend .
ENTRYPOINT ["/aro-hcp-backend"]
