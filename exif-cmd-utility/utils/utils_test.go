package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExtractGPSInfo tests the ExtractGPSInfo function.
func TestExtractGPSInfo(t *testing.T) {
	// Add test cases here to test the ExtractGPSInfo function.
	// For example:
	expectedLatitude := "49.254508"
	expectedLongitude := "-123.100225"
	latitude, longitude, err := ExtractGPSInfo("testdata/test.jpg")
	assert.Equal(t, expectedLatitude, latitude)
	assert.Equal(t, expectedLongitude, longitude)
	assert.NoError(t, err)
}

// TestIsDirectory tests the IsDirectory function.
func TestIsDirectory(t *testing.T) {
	isDir, err := IsDirectory("../utils")
	assert.True(t, isDir)
	assert.NoError(t, err)
}

// TestIsHiddenFile tests the IsHiddenFile function.
func TestIsHiddenFile(t *testing.T) {
	assert.True(t, IsHiddenFile(".hidden_file"))
	assert.False(t, IsHiddenFile("not_hidden.txt"))
}

// Note: You can add more test cases to each test function as needed.
