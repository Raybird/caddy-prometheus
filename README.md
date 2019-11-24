# caddy-prometheus

build caddy 

## Usage

1. docker build -t homemade-caddy/prometheus .
2. docker run -d -v $(pwd)/Caddyfile:/etc/Caddyfile -p 8080:80 -p 9180:9180 homemade-caddy/prometheus
3. http://localhost:9180/metrics
4. http://localhost:8080/

## Reference

1. [Plugging in Plugins Yourself](https://github.com/caddyserver/caddy/wiki/Plugging-in-Plugins-Yourself)
2. [abiosoft/caddy-docker](https://github.com/abiosoft/caddy-docker)