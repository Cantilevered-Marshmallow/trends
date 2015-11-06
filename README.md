# Marshmallow - Trends Service

This is the Trends Service used by [Marshmallow](https://github.com/Cantilevered-Marshmallow/marshmallow).
Refer to the README on the [main repo](https://github.com/Cantilevered-Marshmallow/marshmallow) for more information.

## Team

  - __Product Owner__: Daniel O'Leary
  - __Scrum Master__: Brian Leung
  - __Development Team Members__: Daniel O'Leary, Brian Leung, Brandon Borders

## Deployment

Clone this repo down into your Go workspace and refer to its path in `server/docker-compose.yml` in the [Marshmallow repo](https://github.com/Cantilevered-Marshmallow/marshmallow).

Deployment with the parts of the app is fully handled by docker. Refer to the README on the [main repo](https://github.com/Cantilevered-Marshmallow/marshmallow) for more information.

## Development

To start the service and have it listen to HTTP requests:

```sh
go install
trends
```

## Requirements

- Go

### Dependencies

`go get github.com/asaskevich/govalidator`

## Testing

Simply `go test` from within the trends directory to run `trends_test.go`
