package sfgo

import (
	"encoding/json"
)

func getIPStrArray(ips *[]int64) []string {
        strs := make([]string, len(*ips))
        for i, ip := range *ips {
                strs[i] = GetIPStr(int32(ip))
        }
        return strs
}

func (s *Service) MarshalJSON() ([]byte, error) {
        return json.Marshal(&struct {
                ClusterIP  []string `json:"clusterip"`
                Name       string   `json:"name"`
		Id         string   `json:"id"`
		Namespace  string   `json:"namespace"`
                PortList   []*Port  `json:"portList"`
        }{
                ClusterIP: getIPStrArray(&s.ClusterIP),
		Name: s.Name,
		Id: s.Id,
		Namespace: s.Namespace,
		PortList: s.PortList,
        })
}
