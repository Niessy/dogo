package droplet

import (
	"fmt"
	"testing"

	"github.com/domluna/dogo/digitalocean"
)

func TestGetDroplets(t *testing.T) {
	token, err := digitalocean.EnvAuth()
	if err != nil {
		t.Fatal(err)
	}

	cli := NewClient(token)
	droplets, err := cli.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range droplets {
		fmt.Printf("%+v\n", d)
	}
}

// func TestCreateDroplet(t *testing.T) {
// 	token, err := digitalocean.EnvAuth()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cli := NewClient(token)
// 	droplet, err := cli.Create(map[string]interface{}{
// 		"size": "512mb",
// 		"region": "nyc2",
// 		"name": "tester",
// 		"image": 4296335,
// 		"ssh_keys": []int{136188},
// 		"ipv6": true,
// 		"private_networking": true,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("%v\n", droplet)
// }

// func TestDeleteDroplet(t *testing.T) {
// 	token, err := digitalocean.EnvAuth()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	cli := NewClient(token)
// 	err = cli.Destroy(2018368)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }