package digitalocean

import (
	"testing"
)

func Test_ListRegions(t *testing.T) {
	setup(t)
	defer teardown()

	mux.HandleFunc("/regions", testServer(200, listRegionsExample))

        want := &Region{
                Slug: "ams1",
                Name: "Amsterdam",
                Sizes: []string{
                        "1gb",
                        "512mb",
                },
                Available: true,
                Features: []string{
                        "virtio",
                        "backups",
                },

        }
	regions, err := client.ListRegions()
        assertEqual(t, err, nil)
        assertEqual(t, regions[2].Slug, want.Slug)
        assertEqual(t, regions[2].Name, want.Name)
        assertEqual(t, regions[2].Sizes, want.Sizes)
        assertEqual(t, regions[2].Available, want.Available)
        assertEqual(t, regions[2].Features, want.Features)

}

var listRegionsExample = `{
  "regions": [
    {
      "slug": "nyc1",
      "name": "New York",
      "sizes": [

      ],
      "available": false,
      "features": [
        "virtio",
        "private_networking",
        "backups",
        "ipv6"
      ]
    },
    {
      "slug": "sfo1",
      "name": "San Francisco",
      "sizes": [
        "1gb",
        "512mb"
      ],
      "available": true,
      "features": [
        "virtio",
        "backups"
      ]
    },
    {
      "slug": "ams1",
      "name": "Amsterdam",
      "sizes": [
        "1gb",
        "512mb"
      ],
      "available": true,
      "features": [
        "virtio",
        "backups"
      ]
    }
  ],
  "meta": {
    "total": 3
  }
}`

