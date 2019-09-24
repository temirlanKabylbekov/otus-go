package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassingEmptyText(t *testing.T) {
	got := frequentWords("", 3)
	require.Equal(t, []string{}, got)
}

func TestPassingZeroCount(t *testing.T) {
	got := frequentWords("hey how hey", 0)
	require.Equal(t, []string{}, got)
}

func TestPassingCountMoreThanWordsInText(t *testing.T) {
	got := frequentWords("hey how hey", 10)
	require.Equal(t, []string{"hey", "how"}, got)
}

func TestGetTheMostFrequentWord(t *testing.T) {
	got := frequentWords("hey how hey are you bro", 1)
	require.Equal(t, []string{"hey"}, got)
}

func TestTextWithTheSameWordFrequency(t *testing.T) {
	got := frequentWords("hey how are you bro", 3)
	require.Equal(t, []string{"you", "how", "hey"}, got)
}

func TestTextWithDiverseWordFrequency(t *testing.T) {
	got := frequentWords("hey how hey are you are bro are", 2)
	require.Equal(t, []string{"are", "hey"}, got)
}

func TestWorkingWithCyrillicText(t *testing.T) {
	got := frequentWords("привет как дела привет как", 2)
	require.Equal(t, []string{"привет", "как"}, got)
}

func TestTextWithTooManyWhitespaces(t *testing.T) {
	got := frequentWords(" привет   как          дела привет  как      ", 2)
	require.Equal(t, []string{"привет", "как"}, got)
}
