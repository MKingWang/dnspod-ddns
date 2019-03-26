package main

import "testing"

func Test_config(t *testing.T) {
	cfg := getconfig()
	t.Log(cfg)
}

func Test_getrecord(t *testing.T) {
	cfg := getconfig()
	ip := getRecord(cfg.Dndpod)
	t.Log(ip)
}
