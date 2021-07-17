package main

import "testing"

func TestGetNtpTime(t *testing.T) {
	host := "0.centos.pool.ntp.org"

	res := GetNtpTime(host)

	// if res == nil {
	// 	t.Fatalf("Couldn't get data for host %s", host)
	// }
}
