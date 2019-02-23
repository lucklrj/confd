package apollo

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Client struct {
	Nodes     []string
	AppId     string
	Cluster   string
	Namespace string
}

func NewApolloClient(nodes []string, appId string, cluster string, namespace string) (*Client, error) {
	return &Client{Nodes: nodes, AppId: appId, Cluster: cluster, Namespace: namespace}, nil
}

func (c *Client) GetValues(keys []string) (map[string]string, error) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nodeIndex := r.Intn(len(c.Nodes))

	node := c.Nodes[nodeIndex] + "/configs/" + c.AppId + "/" + c.Cluster + "/" + c.Namespace

	res, _ := http.Get(node)

	defer res.Body.Close()
	html, _ := ioutil.ReadAll(res.Body)

	var dat map[string]interface{}
	vars := make(map[string]string)

	json.Unmarshal(html, &dat)
	if configs, ok := dat["configurations"]; ok {
		for key, val := range configs.(map[string]interface{}) {
			vars[key] = val.(string)
		}
	}
	return vars, nil
}
func (c *Client) WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	<-stopChan
	return 0, nil
}
