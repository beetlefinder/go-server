// Copyright 2018 Vladimir Poliakov and The Contributors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

func main() {
	start(
		"80",
		"postgres",
		"postgresql://postgres:postgres@localhost:5432/beetlefinder?sslmode=disable",
	)
}
