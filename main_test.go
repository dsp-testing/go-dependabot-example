package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assert.Equal(t, 5, result, "2 + 3 should equal 5")
}
