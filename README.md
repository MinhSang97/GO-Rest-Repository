# GO Rest Repository

## Frameworks Used:
- Gin
- Gorm

## How to Run the Code:
1. Navigate to the `cmd` directory.
2. Run the following command:
   ```
   go run main.go
   ```

## Application of GO Rest Repository:

The purpose of this repository is to build a REST API called `/get_histories` to pull market prices.

### Input Parameters:

The API accepts the following parameters:

- `start_date`: The beginning date of the data (start of the day).
- `end_date`: The ending date of the data (end of the day).
- `period`: The period of 1 tick, which can be specified as `30M` (per 30 minutes), `1H` (hourly), or `1D` (daily).
- `symbol`: The symbol for which the data is requested.

Example Request Body:
```json
{
    "start_date": "07-03-2024 12:00:00",
    "end_date": "07-03-2024 19:00:00",
    "period": "1H",
    "symbol": "BTC"
}
```

### Returned Data:

The API returns data in the following format:

```json
{
    "high": double, // Highest price in the given period
    "low": double, // Lowest price in the given period
    "open": double, // The first price in the given period
    "close": double, // The last price in the given period
    "time": int, // The timestamp of the beginning period 
    "change": double // The percentage of change compared to the last API call (per user)
}
```

Example Response Body:
```json
{
    "data": [
        {
            "timestamp": 1709787600,
            "open": 66165.8,
            "high": 67269.8,
            "low": 65726,
            "close": 66691.3,
            "change": -0.24500748633986472
        }
    ]
}
```
