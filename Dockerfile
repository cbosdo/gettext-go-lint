# SPDX-FileCopyrightText: 2024 SUSE LLC
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.20 as builder

WORKDIR /app
COPY . /app

# Statically compile our app for use in a distroless container
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v .

FROM gcr.io/distroless/static

COPY --from=builder /app/gettext-go-lint /gettext-go-lint

ENTRYPOINT ["/gettext-go-lint"]
