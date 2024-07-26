package controller

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"log"
	"net"
	"os"
	"strings"
)

func Iplist(vip string) bool {
	interface_list, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	var byName *net.Interface
	var addrList []net.Addr
	var oneAddrs []string
	for _, i := range interface_list {
		byName, err = net.InterfaceByName(i.Name)
		if err != nil {
			log.Fatal(err)
		}
		addrList = nil
		addrList, err = byName.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, oneAddr := range addrList {
			oneAddrs = strings.SplitN(oneAddr.String(), "/", 2)
			log.Println(oneAddrs)
			if oneAddrs[0] == vip {
				return true
			}
		}
	}
	return false
}

func Updatelabel(clientset *kubernetes.Clientset, name, label string, ) {
	node, err := clientset.CoreV1().Nodes().Get(name, metav1.GetOptions{})
	if err != nil {
		log.Fatal("get node[%s] err is %s", name+"\n", err)
		return
	}
	labels := node.Labels
	log.Println("old label:", node.Labels)
	if labels["vipbind1"] == label {
		return
	} else {
		labels["vipbind1"] = label
		patchData := map[string]interface{}{"metadata": map[string]map[string]string{"labels": labels}}
		playLoadBytes, _ := json.Marshal(patchData)

		newNode, err := clientset.CoreV1().Nodes().Patch(name, types.StrategicMergePatchType, playLoadBytes)
		if err != nil {
			log.Printf("[Up//datePodByPodSn] %v pod Patch fail %v\n", name, err)
			return
		}
		log.Println("new label:", newNode.Labels)
	}
}

func GetHostname()string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error:", err)
	}
	return hostname
}