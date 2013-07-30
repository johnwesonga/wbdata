wbdata
======

World Bank open data API for Go

Find and request information from the
World Bank's various databases

Usage
=====

```go
import "github.com/johnwesonga/wbdata"
```

Construct a new Wbdata client, then use the various services on the client to
access different parts of the World Bank Open Data API.  For example, to list all
countries:

```go
client := wbdata.NewClient(nil)
orgs, err := client.Countries.GetCountries()
```
