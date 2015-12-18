package passwordmanager

/**
 * Copyright 2015 Chungseok Baek
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 **/

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

const (
	APIKEY  = ""
	ERROR_M = "<center>error!</center>"
)

// init
func init() {
	http.HandleFunc("/", handler)
}

// api handler
func handler(rw http.ResponseWriter, req *http.Request) {

	params := req.URL.Path[1:]

	notyet := &NotYetEncryptedPassword{
		key:   "",
		value: "",
	}

	notyet.SimpleEncrypt()

	fmt.Fprintf(rw, "hi there %s", params)
}

type EncryptedPassword struct {
	key   string
	value string
}
