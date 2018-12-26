package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
)

func TagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`["ceph-mon","ceph-node","new"]`))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.WriteHeader(http.StatusOK)

	return
}

func TagsNodesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if vars["tag"] == "ceph-node" {
		_, err := w.Write([]byte(`[
            {
                "sku": "5e2de23c-5624-4814-966b-95f9c34435be",
                "name": "6c:92:bf:48:70:b2",
                "identifiers": ["6c:92:bf:48:70:b2"],
                "type": "compute", "autoDiscover": false,
                "relations": [{"relationType": "enclosedBy",
                               "targets": ["5bd9b5d23003d9b40fe97ede"]}],
                "tags": ["ceph-node"], "createdAt": "2018-10-31T13:58:24.152Z",
                "updatedAt": "2018-11-15T16:37:44.121Z", "id": "5bd9b5003003d9b40fe97ec1"
            },
            {
                "sku": "5e2de23c-5624-4814-966b-95f9c34435be", "name": "6c:92:bf:48:70:e5",
                "identifiers": ["6c:92:bf:48:70:e5"], "type": "compute", "autoDiscover": false,
                "relations":
                    [
                        {
                            "relationType": "enclosedBy",
                            "targets": ["5bee958a9d1714ee233d4433"]
                        }
                    ],
                "tags": ["ceph-node", "new"],
                "createdAt": "2018-11-16T09:50:46.695Z",
                "updatedAt": "2018-11-16T10:01:47.166Z",
                "id": "5bee92f69d1714ee233d4416 "
            }
        ]`))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)

		return
	}

	if vars["tag"] == "ceph-mon" {
		_, err := w.Write([]byte(`[
            {
                "sku": "5e2ff-5624-4814-966b-95s9c34435be",
                "name": "6c:92:bf:48:70:b2",
                "identifiers": ["6c:92:bf:48:70:b2"],
                "type": "compute", "autoDiscover": false,
                "relations": [{"relationType": "enclosedBy",
                               "targets": ["5bd9b5d23003d9b40fe97ede"]}],
                "tags": ["ceph-mon"], "createdAt": "2018-10-31T13:58:24.152Z",
                "updatedAt": "2018-11-15T16:37:44.121Z", "id": "5bd9b5003003d9b40fe97ec1"
            },
            {
                "sku": "5e2de23c-5624-4814-966b-95f9cxzh35be", "name": "6c:92:bf:48:70:e5",
                "identifiers": ["6c:92:bf:48:70:e5"], "type": "compute", "autoDiscover": false,
                "relations":
                    [
                        {
                            "relationType": "enclosedBy",
                            "targets": ["5bee958a9d1714ee233d4433"]
                        }
                    ],
                "tags": ["ceph-mon"],
                "createdAt": "2018-11-16T09:50:46.695Z",
                "updatedAt": "2018-11-16T10:01:47.166Z",
                "id": "5bee92f69d1714ee233d4416"
            }
        ]`))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)

		return
	}

	if vars["tag"] == "new" {
		_, err := w.Write([]byte(`[
            {
                "sku": "5e2de23c-5624-4814-966b-95f9c34435be", "name": "6c:92:bf:48:70:e5",
                "identifiers": ["6c:92:bf:48:70:e5"], "type": "compute", "autoDiscover": false,
                "relations":
                    [
                        {
                            "relationType": "enclosedBy",
                            "targets": ["5bee958a9d1714ee233d4433"]
                        }
                    ],
                "tags": ["ceph-node", "new"],
                "createdAt": "2018-11-16T09:50:46.695Z",
                "updatedAt": "2018-11-16T10:01:47.166Z",
                "id": "5bee92f69d1714ee233d4416"
            }
        ]`))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)

		return
	}

	_, err := fmt.Fprintf(w, `[]`)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.WriteHeader(http.StatusOK)

	return
}

func LookupsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`[
        {
            "macAddress": "a0:36:9f:a4:44:75",
            "ipAddress": "192.168.1.1",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-10-25T10:25:41.712Z",
            "id": "5bd19a251f4fccfd679b9fb2"
        },
        {
            "macAddress": "7c:d3:0a:df:10:85",
            "ipAddress": "192.168.1.123",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-10-25T10:25:41.712Z",
            "id": "5bd19a251f4fccfd679b9fb3"
        },
        {
            "node": "",
            "macAddress": "a0:36:9f:a7:80:14",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-10-25T10:30:12.780Z",
            "ipAddress": "192.168.1.120",
            "id": "5bd19b341f4fccfd679b9fb4"
        },
        {
            "macAddress": "7c:d3:0a:df:10:86",
            "ipAddress": "169.254.95.118",
            "updatedAt": "2018-10-29T21:08:10.739Z",
            "createdAt": "2018-10-25T12:09:50.853Z",
            "id": "5bd1b28e971ea3a21eb84096"
        },
        {
            "node": "",
            "macAddress": "50:6b:4b:b4:a1:66",
            "updatedAt": "2018-11-16T09:49:31.936Z",
            "id": "5bd1cd81971ea3a21eb8409b"
        },
        {
            "node": "",
            "macAddress": "50:6b:4b:b4:a1:67",
            "updatedAt": "2018-11-16T09:49:31.936Z",
            "id": "5bd1cd81971ea3a21eb8409c"
        },
        {
            "node": "",
            "macAddress": "50:6b:4b:b4:ac:6e",
            "updatedAt": "2018-11-16T09:49:31.936Z",
            "id": "5bd1cd81971ea3a21eb8409d"
        },
        {
            "node": "",
            "macAddress": "50:6b:4b:b4:ac:6f",
            "updatedAt": "2018-11-16T09:49:31.936Z",
            "id": "5bd1cd81971ea3a21eb8409e"
        },
        {
            "macAddress": "4c:d9:8f:03:54:5a",
            "ipAddress": "192.168.1.122",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-10-26T00:51:45.247Z",
            "id": "5bd26521971ea3a21eb840ac"
        },
        {
            "macAddress": "4c:d9:8f:03:52:32",
            "ipAddress": "192.168.1.124",
            "updatedAt": "2018-11-09T14:55:50.599Z",
            "createdAt": "2018-10-29T23:23:54.520Z",
            "id": "5bd7968a971ea3a21eb840e1"
        },
        {
            "macAddress": "02:42:ac:11:00:02",
            "ipAddress": "172.17.0.2",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-10-30T12:33:15.117Z",
            "id": "5bd84f8b971ea3a21eb840ea"
        },
        {
            "macAddress": "54:bf:64:d9:34:40",
            "ipAddress": "192.168.1.125",
            "updatedAt": "2018-11-09T14:55:50.596Z",
            "createdAt": "2018-10-31T10:48:48.653Z",
            "id": "5bd98890971ea3a21eb840f7"
        },
        {
            "macAddress": "6c:92:bf:b2:2e:dd",
            "ipAddress": "192.168.1.126",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-10-31T10:54:09.117Z",
            "id": "5bd989d1971ea3a21eb840f8"
        },
        {
            "macAddress": "54:bf:64:d9:34:41",
            "ipAddress": "192.168.2.1",
            "updatedAt": "2018-11-09T14:55:50.601Z",
            "createdAt": "2018-10-31T13:45:05.268Z",
            "id": "5bd9b1e13d950437ba9c621c"
        },
        {
            "node": "5bd9b5003003d9b40fe97ec1",
            "macAddress": "6c:92:bf:48:70:b2",
            "ipAddress": "192.168.1.127",
            "updatedAt": "2018-11-09T14:55:50.599Z",
            "createdAt": "2018-10-31T13:45:05.268Z",
            "id": "5bd9b1e13d950437ba9c621d"
        },
        {
            "node": "5bd9b5003003d9b40fe97ec1",
            "macAddress": "98:03:9b:15:9a:e2",
            "updatedAt": "2018-10-31T14:01:18.569Z",
            "id": "5bd9b4053d950437ba9c6223"
        },
        {
            "node": "5bd9b5003003d9b40fe97ec1",
            "macAddress": "98:03:9b:15:9a:e3",
            "updatedAt": "2018-10-31T14:01:18.569Z",
            "id": "5bd9b4053d950437ba9c6224"
        },
        {
            "macAddress": "e4:43:4b:0d:e6:1e",
            "ipAddress": "192.168.1.128",
            "updatedAt": "2018-11-09T14:55:50.599Z",
            "createdAt": "2018-11-06T14:53:36.515Z",
            "id": "5be1aaf03d950437ba9c623a"
        },
        {
            "macAddress": "6c:92:bf:b2:2e:b9",
            "ipAddress": "192.168.1.129",
            "updatedAt": "2018-11-26T11:55:32.655Z",
            "createdAt": "2018-11-06T16:29:28.994Z",
            "id": "5be1c1683d950437ba9c623b"
        },
        {
            "node": "5bee92f69d1714ee233d4416",
            "macAddress": "6c:92:bf:48:70:e4",
            "ipAddress": "192.168.1.130",
            "updatedAt": "2018-11-16T10:01:07.783Z",
            "createdAt": "2018-11-06T19:17:53.700Z",
            "id": "5be1e8e13d950437ba9c623c"
        },
        {
            "node": "5bee92f69d1714ee233d4416",
            "macAddress": "98:03:9b:15:9a:fa",
            "updatedAt": "2018-11-16T10:01:07.783Z",
            "id": "5be1e90b3d950437ba9c623d"
        },
        {
            "node": "5bee92f69d1714ee233d4416",
            "macAddress": "98:03:9b:15:9a:fb",
            "updatedAt": "2018-11-16T10:01:07.783Z",
            "id": "5be1e90b3d950437ba9c623e"
        },
        {
            "macAddress": "7c:d3:0a:df:10:80",
            "updatedAt": "2018-11-25T16:25:23.896Z",
            "createdAt": "2018-11-12T13:01:59.718Z",
            "ipAddress": "192.168.1.3",
            "id": "5be979c72f0f6e9315b6ae44"
        },
        {
            "node": "5bd9b5003003d9b40fe97ec1",
            "macAddress": "6c:92:bf:48:70:b3",
            "updatedAt": "2018-11-28T12:19:22.718Z",
            "ipAddress": "192.168.1.133",
            "id": "5bd9b4053d950437ba9c6225"
        },
        {
            "node": "",
            "macAddress": "a0:36:9f:a7:80:15",
            "updatedAt": "2018-11-16T09:49:31.936Z",
            "id": "5bd1cd81971ea3a21eb8409a"
        },
        {
            "macAddress": "3c:2c:30:21:9c:80",
            "ipAddress": "192.168.1.134",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-11-13T13:13:01.681Z",
            "id": "5beacddd2f0f6e9315b6ae56"
        },
        {
            "macAddress": "3c:2c:30:21:96:80",
            "ipAddress": "192.168.1.135",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-11-13T13:50:00.769Z",
            "id": "5bead6882f0f6e9315b6ae58"
        },
        {
            "macAddress": "54:bf:64:f3:de:c1",
            "ipAddress": "192.168.1.136",
            "updatedAt": "2018-11-25T16:25:23.897Z",
            "createdAt": "2018-11-13T13:50:46.970Z",
            "id": "5bead6b62f0f6e9315b6ae5a"
        },
        {
            "node": "5bee92f69d1714ee233d4416",
            "macAddress": "6c:92:bf:48:70:e5",
            "updatedAt": "2018-11-29T12:40:09.177Z",
            "ipAddress": "192.168.1.137",
            "id": "5be1e90b3d950437ba9c623f"
        },
        {
            "macAddress": "00:e0:4c:00:0d:41",
            "ipAddress": "192.168.1.138",
            "updatedAt": "2018-12-03T13:21:33.546Z",
            "createdAt": "2018-12-03T12:52:35.007Z",
            "id": "5c0527132f0f6e9315b6aef1"
        }
    ]`))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.WriteHeader(http.StatusOK)

	return
}

func RackhdStubServerWithAllEndpoints() *httptest.Server {
	r := mux.NewRouter()
	r.HandleFunc("/api/current/tags/", TagsHandler).Methods("GET")
	r.HandleFunc("/api/current/tags/{tag}/nodes", TagsNodesHandler).Methods("GET")
	r.HandleFunc("/api/2.0/lookups", LookupsHandler).Methods("GET")

	return httptest.NewServer(r)
}
