package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pack-sizes/pkg/split"
	"strconv"
)

type PackageSplitInfo struct {
	PackageSize int `json:"bucketSize"`
	Count       int `json:"count"`
}

func packageSplitsInfoFromMap(m map[int]int) []PackageSplitInfo {
	ret := make([]PackageSplitInfo, 0)
	for k, v := range m {
		ret = append(ret, PackageSplitInfo{
			PackageSize: k,
			Count:       v,
		})
	}
	return ret
}

func MakeSplitEndpoint(buckets []int, f split.Split, opt split.Options) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ordersParamName := "orders"
		orderStr := r.URL.Query().Get(ordersParamName)
		if len(orderStr) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(w, "`%s` parameter required and not sent - received: %v", ordersParamName, r.URL.Query())
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			return
		}
		order, err := strconv.Atoi(orderStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err := fmt.Fprintf(w, "unable to parse `orders` as number: %s", err)
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			return
		}

		ret := f(order, buckets, opt)
		packagesInfo := packageSplitsInfoFromMap(ret)
		jsonData, err := json.Marshal(packagesInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "unable to marshal response: %s", err)
			return
		}

		_, err = fmt.Fprintf(w, "%s\n", jsonData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "unable to parse `orders` as number: %s", err)
			return
		}
	}
}
