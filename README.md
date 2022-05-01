# stock-ticker

Stock Ticker is a simple webservice that is written in Go to return Stock price history for a company over a specific range of dates.
It calls the [Alpha Vantage](https://www.alphavantage.co/) API under the hood to retrieve the stock history data.

The latest docker image can be fetched with the following command:
```bash
docker pull docker.io/tariq181290/stock-ticker:v0.0.2
```

# Getting started

## Running in Kubernetes
To get started with `stock-ticker`, you will need to have a Kubernetes Cluster. For illustrative purposes, we will use 
the [kind](https://kind.sigs.k8s.io) cluster, but any Kubernetes cluster is ok to use.
Please ensure that have you have a compatible version of `kubectl` in your local machine.

1. Stand up your Kubernetes cluster. If you're using a kind cluster, you can go [here](https://kind.sigs.k8s.io/docs/user/quick-start)
2. Run `kubectl apply -f ./examples/manifests`
3. You will now need to expose the Kubernetes service, run `kubectl port-forward service/stock-ticker 8090:8090`
4. Run `curl localhost:8090/api/stockticker`. You should get an output that looks like this:
```json
{
   "Meta Data": {
      "1. Information": "Daily Prices (open, high, low, close) and Volumes",
      "2. Symbol": "MSFT",
      "3. Last Refreshed": "2022-04-29",
      "4. Output Size": "Compact",
      "5. Time Zone": "US/Eastern"
   },
   "Time Series (Daily)": {
      "2022-04-25": {
         "1. open": "273.2900",
         "2. high": "281.1100",
         "3. low": "270.7700",
         "4. close": "280.7200",
         "5. volume": "35678852"
      },
      "2022-04-26": {
         "1. open": "277.5000",
         "2. high": "278.3599",
         "3. low": "270.0000",
         "4. close": "270.2200",
         "5. volume": "42047008"
      },
      "2022-04-27": {
         "1. open": "282.1000",
         "2. high": "290.9700",
         "3. low": "279.1600",
         "4. close": "283.2200",
         "5. volume": "63477694"
      },
      "2022-04-28": {
         "1. open": "285.1850",
         "2. high": "290.9800",
         "3. low": "281.4562",
         "4. close": "289.6300",
         "5. volume": "33646570"
      }
   },
   "Output Data": {
      "N Days": 5,
      "Close Average": "280.947502"
   }
}
```
