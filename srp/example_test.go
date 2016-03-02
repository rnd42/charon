// Copyright 2013 Tad Glines
// Copyright 2016 Alex Mayfield <alexmax2742@gmail.com>
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package srp

import (
	"crypto/sha256"
	"log"
)

func ExampleNewSRP() {
	username := []byte("example")
	password := []byte("3x@mp1e")

	srp, err := NewSRP("openssl.1024", sha256.New, nil)
	if err != nil {
		log.Fatal(err)
	}

	cs := srp.NewClientSession(username, password)
	salt, v, err := srp.ComputeVerifier(username, password)
	if err != nil {
		log.Fatal(err)
	}

	ss := srp.NewServerSession(username, salt, v)

	ckey, err := cs.ComputeKey(salt, ss.GetB())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The Client's computed session key is: %v\n", ckey)

	skey, err := ss.ComputeKey(cs.GetA())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The Server's computed session key is: %v\n", skey)

	cauth := cs.ComputeAuthenticator()
	if !ss.VerifyClientAuthenticator(cauth) {
		log.Print("Client Authenticator is not valid")
	}
	sauth := ss.ComputeAuthenticator(cauth)
	if !cs.VerifyServerAuthenticator(sauth) {
		log.Print("Server Authenticator is not valid")
	}
}