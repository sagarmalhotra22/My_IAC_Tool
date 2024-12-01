# MY IAC TOOL

My Iac Tool is a simple CLI-based Go IaC application that allows users to manage (create and delete) Google Cloud Platform (GCP) virtual machine (VM) instances. The project demonstrates:

1. **Creating a VM instance on GCP**: The `apply` command creates a VM instance and stores its details (instance name, project ID, and zone) in a state file.
2. **Deleting a VM instance from GCP**: The `destroy` command reads the instance details from the state file and deletes the VM instance from GCP.

The project uses **Cobra** for building a command-line interface and manages instance states using a JSON-based state file.


## Features

1. **State Management**: Tracks instance details (name, project ID, zone) in a state.json file for managing VM lifecycle.
2. **CLI Commands**: Provides intuitive commands for creating and destroying VMs using GCP's Compute Engine API.
3. **Modular Design**: Built using the Cobra package, making it extensible for additional features.


---

## Project Structure

```plaintext
gcp-vm-manager/
├── main.go                 # Entry point of the application
├── config/
│   └── config.go           # Loads configuration of the VM to be created
├── state/
│   └── state.go            # Handles state management (save/load instance details)
└── cmd/
    ├── apply.go            # Command for creating a VM
    └── destroy.go          # Command for deleting a VM

