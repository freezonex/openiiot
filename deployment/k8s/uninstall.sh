#!/bin/bash

# Ask for confirmation
read -p "Are you sure you want to uninstall all resources in the 'openiiot' namespace? (y/n) " answer
if [ "$answer" = "Y" ] || [ "$answer" = "y" ]; then
    # Proceed with uninstallation
    kubectl delete deployments,services --all -n openiiot
    echo "Uninstallation completed successfully. Please use 'kubectl delete namespace openiiot' to manually delete the namespace if you do not want to keep it."
else
    echo "Uninstallation cancelled."
fi
