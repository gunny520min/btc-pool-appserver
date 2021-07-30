package btcpoolclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Sign(params map[string]string) map[string]string {
	res := make(map[string]string)
	res["nonce"] = RandStringBytes(10)
	res["app_a"] = "6b2llbmFuMmUxb2lpaG9hMm4za2RodWZzZGRmYXMK"
	res["app_b"] = "Y2hhaW5hcHAK"
	res["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	ks := make([]string, 0)
	for k, v := range params {
		res[k] = v
	}
	for k, _ := range res {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	args := make([]string, 0)
	for _, v := range ks {
		args = append(args, res[v])
	}
	secretKey := "7f5a58ff3739aaaf45025599ae5a993269573dd2a215acfccd7cc30eec86efe1"

	hash := hmac.New(sha256.New, []byte(secretKey))
	msg := []byte(strings.Join(args, "|"))
	hash.Write(msg)

	// to base64
	sign := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	res["sign"] = sign
	return res
}

type PoolRankList []PoolRank

type PoolRank struct {
	PoolName               string `json:"pool_name"`
	IconLink               string `json:"icon_link"`
	RealtimeHashrate       string `json:"realtime_hashrate"`
	EstimateHashrate       string `json:"estimate_hashrate"`
	EstimateCur2max        string `json:"estimate_cur2max"`
	Cur2maxPercent         string `json:"cur2max_percent"`
	HashSuffix             string `json:"hashrate_suffix"`
	RealtimeDiff24hPercent string `json:"realtime_diff_24h_percent"`
	HashrateUnit           string `json:"hashrate_unit"`
	EstimateHashrateUnit   string `json:"estimate_hashrate_unit"`
	HashrateSuffix         string `json:"hashrate_suffix"`
	Lucky                  string `json:"lucky"`
	Lucky3d                string `json:"lucky_3d"`
}

type RealtimeData struct {
	List []PoolRank `json:"list"`
}
type PoolRankData struct {
	Realtime RealtimeData `json:"realtime"`
}

// 获取矿池排名
func GetPoolRank(c *gin.Context, params map[string]string) (map[string]PoolRankData, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]PoolRankData `json:"data"`
	}{}
	_, err := doRequest(c, "explorer.poolRank", Sign(params), &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error GetPoolRank: %v", err)
	}
	return dest.Data, nil
}

type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"relayed_by_text"`
	Hash      string `json:"hash"`
	Size      string `json:"size"`
}
type LatestBlockList []LatestBlock

func GetLatestBlockList(c *gin.Context, params map[string]string) (map[string]LatestBlockList, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "explorer.blockList", Sign(params), &dest)
	if err != nil {
		return nil, err //fmt.Errorf("error GetLatestBlockList: %v", err)
	}

	res := make(map[string]LatestBlockList)
	for k, v := range dest.Data {
		resByre, resByteErr := json.Marshal(v["list"])
		if resByteErr != nil {
			return nil, resByteErr
		}
		var newData LatestBlockList
		jsonRes := json.Unmarshal(resByre, &newData)
		if jsonRes != nil {
			return nil, jsonRes
		}
		res[k] = newData
	}
	return res, nil
}
