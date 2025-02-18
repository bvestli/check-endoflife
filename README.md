Simple application to check product end of life dates against the service https://endoflife.date

The products are read from the file products.yaml, and the end of life dates are checked against the service. The results are written to stdout.

products.yaml example:
```yaml
products:
  - name: traefik
    version: 2.11.3
  - name: grafana
    version: 10.4.14
  - name: prometheus
    version: 3.0.1
  - name: argo-cd
    version: v2.5.21+f627b62
```
