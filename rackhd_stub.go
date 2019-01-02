package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RackhdPathHandlers() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc(tagsPath, TagsHandler).Methods("GET")
	r.HandleFunc(tagsPathTemplate, TagsNodesHandler).Methods("GET")
	r.HandleFunc(lookupPath, LookupsHandler).Methods("GET")

	return r
}

func TagsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(TagsHandlerOutput))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

}

func TagsNodesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if vars["tag"] == TagCephNode {
		_, err := w.Write([]byte(TagsCephNodeOutput))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		return
	}

	if vars["tag"] == TagCephMon {
		_, err := w.Write([]byte(TagsCephMonOutput))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		return
	}

	if vars["tag"] == TagNew {
		_, err := w.Write([]byte(TagsNewOutput))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		return
	}

	_, err := fmt.Fprintf(w, EmptyBoxBrackets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

}

func LookupsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(lookupsHandlerOutput))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	return
}

const TagsHandlerOutput = `["ceph-mon","ceph-node","new"]`

const TagsCephNodeOutput = `[
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
        ]`

const TagsCephMonOutput = `[
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
        ]`

const TagsNewOutput = `[
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
        ]`

const lookupsHandlerOutput = `[
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
            "node": "5bee92f69d1714ee233d4416",
            "macAddress": "6c:92:bf:48:70:e5",
            "updatedAt": "2018-11-29T12:40:09.177Z",
            "ipAddress": "192.168.1.137",
            "id": "5be1e90b3d950437ba9c623f"
        }
    ]`
