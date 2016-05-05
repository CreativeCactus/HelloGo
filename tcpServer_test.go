package main 
import (
  "testing" 
  "net"
)
    
func TestNothing(t *testing.T) {
    if testing.Short() {t.Skip("skipping test in short mode.")}//-test.short
}
func TestEcho(t *testing.T) {
    
}
