
package main

import (
	"fmt"
	"time"
)

// Network represents a network entity with performance metrics.
type Network struct {
	Name       string
	Latency    float64 // In milliseconds
	PacketLoss float64 // In percentage
	Bandwidth  float64 // In Mbps
}

// NetworkAnalyzer provides methods to analyze network performance.
type NetworkAnalyzer struct {
	ThresholdLatency    float64
	ThresholdPacketLoss float64
}

// Analyze identifies sluggish network behavior.
func (analyzer *NetworkAnalyzer) Analyze(network *Network) bool {
	if network.Latency > analyzer.ThresholdLatency || network.PacketLoss > analyzer.ThresholdPacketLoss {
		return true
	}
	return false
}

// Optimizer improves network performance using an algorithm.
type Optimizer struct {
	ImprovementFactor float64
}

// Optimize attempts to reduce latency and packet loss for a network.
func (optimizer *Optimizer) Optimize(network *Network) {
	network.Latency -= network.Latency * optimizer.ImprovementFactor
	network.PacketLoss -= network.PacketLoss * optimizer.ImprovementFactor
	if network.Latency < 0 {
		network.Latency = 0
	}
	if network.PacketLoss < 0 {
		network.PacketLoss = 0
	}
	fmt.Printf("Optimized Network: %s | Latency: %.2fms | Packet Loss: %.2f%%\n",
		network.Name, network.Latency, network.PacketLoss)
}

// NetworkMonitor monitors and optimizes networks.
type NetworkMonitor struct {
	Networks   []*Network
	Analyzer   *NetworkAnalyzer
	Optimizer  *Optimizer
	CheckCycle time.Duration
}

// Monitor continuously analyzes and optimizes networks.
func (monitor *NetworkMonitor) Monitor() {
	for {
		fmt.Println("Monitoring networks...")
		for _, network := range monitor.Networks {
			if monitor.Analyzer.Analyze(network) {
				fmt.Printf("Network %s is sluggish. Optimizing...\n", network.Name)
				monitor.Optimizer.Optimize(network)
			} else {
				fmt.Printf("Network %s is performing well.\n", network.Name)
			}
		}
		time.Sleep(monitor.CheckCycle)
	}
}

func main() {
	// Create sample networks
	network1 := &Network{Name: "Office LAN", Latency: 120, PacketLoss: 2, Bandwidth: 100}
	network2 := &Network{Name: "Home WiFi", Latency: 300, PacketLoss: 5, Bandwidth: 50}

	// Set up the analyzer and optimizer
	analyzer := &NetworkAnalyzer{ThresholdLatency: 100, ThresholdPacketLoss: 1}
	optimizer := &Optimizer{ImprovementFactor: 0.2}

	// Initialize the network monitor
	monitor := &NetworkMonitor{
		Networks:   []*Network{network1, network2},
		Analyzer:   analyzer,
		Optimizer:  optimizer,
		CheckCycle: 5 * time.Second,
	}

	// Start monitoring
	monitor.Monitor()
}
