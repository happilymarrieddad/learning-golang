# client.example.com

> A Vue.js project

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).


## Notes

```go
package orders

import (
	v3_db "../../../../pkg/v3_db"
)

type Orders []Order

type Order v3_db.Orders

func (u *Order) TableName() string {
	return "orders"
}
```
```go
package orders

import (
	. "../../../../system/errors"
	Orders "../../models/orders"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//userId := r.Header.Get("user_id")
	customerId := r.Header.Get("customer_id")
	var offset int
	limit := int(25)

	if keys, ok := r.URL.Query()["limit"]; ok {
		limit, _ = strconv.Atoi(keys[0])
	}
	if keys, ok := r.URL.Query()["page"]; ok {
		page, _ := strconv.Atoi(keys[0])
		offset = (page - 1) * limit
	}

	customer_id, _ := strconv.Atoi(customerId)

	orders, err := Orders.FindBy(db, int64(customer_id), limit, offset)

	packet, err := json.Marshal(orders)
	if err != nil {
		log.Println(err)
		http.Error(w, UNABLE_TO_PARSE_JSON.String(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}

```
