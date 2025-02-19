Simple application to check product end of life dates against the service [endoflife.date](https://endoflife.date)

The products are read from the file `products.yaml`, and the end of life dates are checked against the service. The results are written to stdout.

`products.yaml` example:
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


Output when running the application:
```json
[
  {
    "name": "traefik",
    "releaseDate": "2024-02-12",
    "eol": "29.04.2025",
    "latest": "2.11.20",
    "support": "2025-04-29",
    "myversion": "2.11.3",
    "latestcycle": "3.3"
  },
  {
    "name": "grafana",
    "releaseDate": "2024-03-05",
    "eol": "30.06.2025",
    "latest": "10.4.15",
    "support": "2025-06-30",
    "myversion": "10.4.14",
    "latestcycle": "11.5"
  },
  {
    "name": "prometheus",
    "releaseDate": "2024-11-14",
    "eol": "26.12.2024",
    "latest": "3.0.1",
    "support": null,
    "myversion": "3.0.1",
    "latestcycle": "3.1"
  },
  {
    "name": "argo-cd",
    "releaseDate": "2022-10-25",
    "eol": "07.08.2023",
    "latest": "2.5.22",
    "support": null,
    "myversion": "v2.5.21+f627b62",
    "latestcycle": "2.14"
  }
]
```
