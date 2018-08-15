// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package addrmgr

import (
	"math"
	"testing"
	"time"

	"github.com/Legenddigital/lddld/wire"
)

func newKnownAddress(na *wire.NetAddress, attempts int, lastattempt, lastsuccess time.Time, tried bool, refs int) *KnownAddress {
	return &KnownAddress{na: na, attempts: attempts, lastattempt: lastattempt,
		lastsuccess: lastsuccess, tried: tried, refs: refs}
}

func TestChance(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	var tests = []struct {
		addr     *KnownAddress
		expected float64
	}{
		{
			// Test normal case
			newKnownAddress(&wire.NetAddress{Timestamp: now.Add(-35 * time.Second)},
				0, now.Add(-30*time.Minute), now, false, 0),
			1.0,
		}, {
			// Test case in which lastseen < 0
			newKnownAddress(&wire.NetAddress{Timestamp: now.Add(20 * time.Second)},
				0, now.Add(-30*time.Minute), now, false, 0),
			1.0,
		}, {
			// Test case in which lastattempt < 0
			newKnownAddress(&wire.NetAddress{Timestamp: now.Add(-35 * time.Second)},
				0, now.Add(30*time.Minute), now, false, 0),
			1.0 * .01,
		}, {
			// Test case in which lastattempt < ten minutes
			newKnownAddress(&wire.NetAddress{Timestamp: now.Add(-35 * time.Second)},
				0, now.Add(-5*time.Minute), now, false, 0),
			1.0 * .01,
		}, {
			// Test case with several failed attempts.
			newKnownAddress(&wire.NetAddress{Timestamp: now.Add(-35 * time.Second)},
				2, now.Add(-30*time.Minute), now, false, 0),
			1 / 1.5 / 1.5,
		},
	}

	err := .0001
	for i, test := range tests {
		chance := test.addr.chance()
		if math.Abs(test.expected-chance) >= err {
			t.Errorf("case %d: got %f, expected %f", i, chance, test.expected)
		}
	}
}

func TestIsBad(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	future := now.Add(35 * time.Minute)
	monthOld := now.Add(-43 * time.Hour * 24)
	secondsOld := now.Add(-2 * time.Second)
	minutesOld := now.Add(-27 * time.Minute)
	hoursOld := now.Add(-5 * time.Hour)
	zeroTime := time.Time{}

	futureNa := &wire.NetAddress{Timestamp: future}
	minutesOldNa := &wire.NetAddress{Timestamp: minutesOld}
	monthOldNa := &wire.NetAddress{Timestamp: monthOld}
	currentNa := &wire.NetAddress{Timestamp: secondsOld}

	// Test addresses that have been tried in the last minute.
	if newKnownAddress(futureNa, 3, secondsOld, zeroTime, false, 0).isBad() {
		t.Errorf("test case 1: addresses that have been tried in the last minute are not bad.")
	}
	if newKnownAddress(monthOldNa, 3, secondsOld, zeroTime, false, 0).isBad() {
		t.Errorf("test case 2: addresses that have been tried in the last minute are not bad.")
	}
	if newKnownAddress(currentNa, 3, secondsOld, zeroTime, false, 0).isBad() {
		t.Errorf("test case 3: addresses that have been tried in the last minute are not bad.")
	}
	if newKnownAddress(currentNa, 3, secondsOld, monthOld, true, 0).isBad() {
		t.Errorf("test case 4: addresses that have been tried in the last minute are not bad.")
	}
	if newKnownAddress(currentNa, 2, secondsOld, secondsOld, true, 0).isBad() {
		t.Errorf("test case 5: addresses that have been tried in the last minute are not bad.")
	}

	// Test address that claims to be from the future.
	if !newKnownAddress(futureNa, 0, minutesOld, hoursOld, true, 0).isBad() {
		t.Errorf("test case 6: addresses that claim to be from the future are bad.")
	}

	// Test address that has not been seen in over a month.
	if !newKnownAddress(monthOldNa, 0, minutesOld, hoursOld, true, 0).isBad() {
		t.Errorf("test case 7: addresses more than a month old are bad.")
	}

	// It has failed at least three times and never succeeded.
	if !newKnownAddress(minutesOldNa, 3, minutesOld, zeroTime, true, 0).isBad() {
		t.Errorf("test case 8: addresses that have never succeeded are bad.")
	}

	// It has failed ten times in the last week
	if !newKnownAddress(minutesOldNa, 10, minutesOld, monthOld, true, 0).isBad() {
		t.Errorf("test case 9: addresses that have not succeeded in too long are bad.")
	}

	// Test an address that should work.
	if newKnownAddress(minutesOldNa, 2, minutesOld, hoursOld, true, 0).isBad() {
		t.Errorf("test case 10: This should be a valid address.")
	}
}
