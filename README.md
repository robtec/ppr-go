# Property Price Register CLI (unoffical)

This cli allows you to query the sales listed on `https://www.propertypriceregister.ie`

## Installation

Using go

```
$ go get github.com/robtec/ppr-go/cmd/...
```

## Usage

`ppr` supports the following options

* `-c` - county to search, default Galway
* `-yf` - year from to search, default 2010
* `-yt` - year to search to, default 2020
* `-a` - address to search
* `-o` - change output, options are `json`, default `table`

### Example

```
$ ppr -c mayo -yf 2019 -a "castlebar"

+------------+--------------------------------+----------------+
|    DATE    |            ADDRESS             |     PRICE      |
+------------+--------------------------------+----------------+
| 20/12/2019 | DUMMY, CASTLEBAR, MAYO         | €45,500.00     |
+------------+--------------------------------+----------------+
| 19/12/2019 | LANE, CASTLEBAR, MAYO          | €321,000.00    |
+------------+--------------------------------+----------------+
| 17/12/2019 | SPRINGFIELD, CASTLEBAR,        | €543,000.00    |
|            | MAYO                           |                |
+------------+--------------------------------+----------------+
| 17/12/2019 | MICKY MOUSE DRIVE, CASTLEBAR   | €99,000.00     |
+------------+--------------------------------+----------------+
| 13/12/2019 | 123 FAKE ROAD, CASTLEBAR, CO   | €75,000.00     |
|            | MAYO                           |                |
+------------+--------------------------------+----------------+
# real addresses/prices omitted for privacy
...
```

JSON output

```
$ ppr -o json -c mayo -yf 2019 -a "castlebar"

{
    "sales": [
        {
            "date": "20/12/2019",
            "price": "€45,500.00",
            "address": "DUMMY, CASTLEBAR, MAYO "
        },
        {
            "date": "19/12/2019",
            "price": "€321,000.00",
            "address": "LANE, CASTLEBAR, MAYO"
        },
        {
            "date": "17/12/2019",
            "price": "€543,000.00",
            "address": "SPRINGFIELD, CASTLEBAR, MAYO"
        },
        {
            "date": "17/12/2019",
            "price": "€99,000.00",
            "address": "MICKY MOUSE DRIVE, CASTLEBAR"
        },
        {
            "date": "13/12/2019",
            "price": "€75,000.00",
            "address": "123 FAKE ROAD, CASTLEBAR, CO MAYO"
        }]
}
```