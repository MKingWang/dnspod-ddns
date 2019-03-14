package main

import "testing"

func Test_config(t *testing.T) {
	cfg := getconfig()
	t.Log(cfg)
}
