#!/bin/bash

# Ask for confirmation
read -p "Are you sure you want to uninstall all resources in the 'openiiot' namespace? (y/n) " answer
if [ "$answer" = "Y" ] || [ "$answer" = "y" ]; then
    # Proceed with uninstallation
    echo "Delete all deployments, services, jobs, pvc in namespace openiiot."
    kubectl delete deployments,services,jobs,pvc --all -n openiiot

    echo "Delete all pv starting with 'openiiot'..."
    # List all PVs, filter those starting with 'openiiot', and delete them
    kubectl get pv --no-headers | grep '^openiiot' | awk '{print $1}' | while read pv; do
        echo "Deleting PV: $pv"
        kubectl delete pv "$pv" || echo "Failed to delete PV: $pv"
    done

    echo "Uninstallation completed successfully. Please use 'kubectl delete namespace openiiot' to manually delete the namespace if you do not want to keep it."
else
    echo "Uninstallation cancelled."
fi
