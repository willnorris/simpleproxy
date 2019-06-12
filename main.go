// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	target, err := url.Parse(os.Getenv("TARGET_URL"))
	if err != nil {
		log.Fatalf("error parsing TARGET_URL: %v", err)
	}

	rp := httputil.NewSingleHostReverseProxy(target)
	proxy := func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host
		rp.ServeHTTP(w, r)
	}

	log.Printf("Listening on port 8080; Proxying to %v", target.String())
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(proxy)))
}
