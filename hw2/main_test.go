package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassingEmptyString(t *testing.T) {
	got, err := unpack("")
	require.Nil(t, err)
	require.Equal(t, "", got)
}

func TestPassingNotLetters(t *testing.T) {
	_, err := unpack("[2a3B")
	require.EqualError(t, err, "Passed string \"[2a3B\" should contain only letters and digits")
}

func TestStringStartingWithNumber(t *testing.T) {
	_, err := unpack("2a")
	require.EqualError(t, err, "Passed string \"2a\" starts with digit or have two digits next to each other")
}

func TestPassingNumber(t *testing.T) {
	_, err := unpack("45")
	require.EqualError(t, err, "Passed string \"45\" starts with digit or have two digits next to each other")
}

func TestPassingStringContainsDigitsNextToEachOther(t *testing.T) {
	_, err := unpack("ab45c1")
	require.EqualError(t, err, "Passed string \"ab45c1\" starts with digit or have two digits next to each other")
}

func TestUnpacking(t *testing.T) {
	got, err := unpack("a4bc2d5e")
	require.Nil(t, err)
	require.Equal(t, "aaaabccddddde", got)
}

func TestUnpackingWithZeroCount(t *testing.T) {
	got, err := unpack("g0b")
	require.Nil(t, err)
	require.Equal(t, "b", got)
}

func TestUnpackingWithCountEqualOne(t *testing.T) {
	got, err := unpack("g1")
	require.Nil(t, err)
	require.Equal(t, "g", got)
}

func TestUnpackingWithoutCounters(t *testing.T) {
	got, err := unpack("abcd")
	require.Nil(t, err)
	require.Equal(t, "abcd", got)
}

func TestUnpackingCyrillicSymbols(t *testing.T) {
	got, err := unpack("Щ2я3")
	require.Nil(t, err)
	require.Equal(t, "ЩЩяяя", got)
}
