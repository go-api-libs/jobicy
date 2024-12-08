# 🌍 Remote Jobs API
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/jobicy.svg)](https://pkg.go.dev/github.com/go-api-libs/jobicy/pkg/jobicy)
[![Official Documentation](https://img.shields.io/badge/docs-API-blue)](https://jobicy.com/jobs-rss-feed)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/jobicy)](https://goreportcard.com/report/github.com/go-api-libs/jobicy)
![Code Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
![API Health](https://img.shields.io/badge/API_health-65%25-yellowgreen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Discover and integrate a diverse range of remote job listings with Jobicy's public API. This API offers the latest remote job opportunities across various industries and regions, making it a valuable resource for developers and businesses looking to enhance their job feed platforms. Supports filters for job region, industry, and keywords to help you target specific job markets.

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/jobicy/pkg/jobicy
```

## Usage

### Example: 

```go
package main

import (
	"context"

	"github.com/go-api-libs/jobicy/pkg/jobicy"
)

func main() {
	c, err := jobicy.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	jobsList, err := c.ListRemoteJobs(ctx, &jobicy.ListRemoteJobsParams{
		Count:    20,
		Geo:      "usa",
		Industry: "marketing",
		Tag:      "seo",
	})
	if err != nil {
		panic(err)
	}

	// Use jobsList object
}

```

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/go-api-libs/jobicy/pkg/jobicy): The Go reference documentation for the client package.
- [**Official Documentation**](https://jobicy.com/jobs-rss-feed): The official API documentation.
- [**OpenAPI Specification**](./api/openapi.json): The OpenAPI 3.1.0 specification.
- [**Go Report Card**](https://goreportcard.com/report/github.com/go-api-libs/jobicy): Check the code quality report.

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/go-api-libs/jobicy).

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
